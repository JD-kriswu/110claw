// Package feishu provides a client for the Feishu (Lark) Open Platform API
// and a scheduler that periodically syncs wiki content into learn_steps.
//
// Required env vars (in backend/.env):
//   FEISHU_APP_ID      - Feishu app ID
//   FEISHU_APP_SECRET  - Feishu app secret
//   FEISHU_WIKI_TOKEN  - wiki node token (last path segment of the wiki URL)
//   FEISHU_SYNC_INTERVAL - e.g. "1h", "30m"  (default: 1h)
package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"gorm.io/gorm"
)

const baseURL = "https://open.feishu.cn/open-apis"

// ─── client ───────────────────────────────────────────────

// Syncer holds credentials and syncs a Feishu wiki into learn_steps.
type Syncer struct {
	appID     string
	appSecret string
	wikiToken string

	mu          sync.Mutex
	accessToken string
	tokenExpiry time.Time
}

// New creates a Syncer. wikiToken is the last path segment of the wiki URL.
func New(appID, appSecret, wikiToken string) *Syncer {
	return &Syncer{appID: appID, appSecret: appSecret, wikiToken: wikiToken}
}

// token returns a valid tenant_access_token, refreshing when expired.
func (s *Syncer) token() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Now().Before(s.tokenExpiry) {
		return s.accessToken, nil
	}

	payload, _ := json.Marshal(map[string]string{
		"app_id":     s.appID,
		"app_secret": s.appSecret,
	})
	resp, err := http.Post(
		baseURL+"/auth/v3/tenant_access_token/internal",
		"application/json",
		bytes.NewReader(payload),
	)
	if err != nil {
		return "", fmt.Errorf("auth request: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Code              int    `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("auth decode: %w", err)
	}
	if result.Code != 0 {
		return "", fmt.Errorf("auth error %d: %s", result.Code, result.Msg)
	}

	s.accessToken = result.TenantAccessToken
	// Refresh 2 minutes before expiry
	s.tokenExpiry = time.Now().Add(time.Duration(result.Expire-120) * time.Second)
	return s.accessToken, nil
}

// get performs an authenticated GET and returns the response body.
func (s *Syncer) get(path string) ([]byte, error) {
	tok, err := s.token()
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("GET", baseURL+path, nil)
	req.Header.Set("Authorization", "Bearer "+tok)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// ─── wiki API types ───────────────────────────────────────

type wikiNode struct {
	SpaceID   string `json:"space_id"`
	NodeToken string `json:"node_token"`
	ObjType   string `json:"obj_type"`  // "docx" | "doc" | "sheet" …
	ObjToken  string `json:"obj_token"` // underlying document ID
	Title     string `json:"title"`
	HasChild  bool   `json:"has_child"`
}

// getNode fetches info about a wiki node by its token.
func (s *Syncer) getNode(token string) (*wikiNode, error) {
	data, err := s.get("/wiki/v2/spaces/get_node?token=" + token)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Code int      `json:"code"`
		Msg  string   `json:"msg"`
		Data struct { Node wikiNode `json:"node"` } `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("get_node %d: %s (raw: %s)", resp.Code, resp.Msg, data)
	}
	return &resp.Data.Node, nil
}

// getChildren lists immediate children of a wiki node.
func (s *Syncer) getChildren(spaceID, parentToken string) ([]wikiNode, error) {
	path := fmt.Sprintf("/wiki/v2/spaces/%s/nodes?parent_node_token=%s&page_size=50",
		spaceID, parentToken)
	data, err := s.get(path)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Items []wikiNode `json:"items"`
		} `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("get_children %d: %s", resp.Code, resp.Msg)
	}
	return resp.Data.Items, nil
}

// getRawContent returns the plain-text content of a docx document.
func (s *Syncer) getRawContent(docToken string) (string, error) {
	data, err := s.get("/docx/v1/documents/" + docToken + "/raw_content?lang=0")
	if err != nil {
		return "", err
	}
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	if resp.Code != 0 {
		return "", fmt.Errorf("raw_content %d: %s (token=%s)", resp.Code, resp.Msg, docToken)
	}
	return resp.Data.Content, nil
}

// ─── content parsing ──────────────────────────────────────

var (
	// Matches "Day 1", "第1天", "第一天", "1.", "1、", "D1:" etc.
	dayRe = regexp.MustCompile(`(?i)(?:day|第|d)\s*([1-7])|^([1-7])\s*[.、：:]`)
	cnNum = map[rune]int{'一': 1, '二': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7}
)

func parseDayNumber(s string) int {
	m := dayRe.FindStringSubmatch(s)
	for _, g := range m[1:] {
		if g != "" {
			n, _ := strconv.Atoi(g)
			return n
		}
	}
	for _, r := range s {
		if n, ok := cnNum[r]; ok {
			return n
		}
	}
	return 0
}

func extractSummary(content string) string {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		runes := []rune(line)
		if len(runes) > 120 {
			return string(runes[:120]) + "…"
		}
		return line
	}
	return ""
}

type dayEntry struct {
	Day     int
	Title   string
	Content string
}

// parseContent splits a raw document into per-day entries.
// It looks for section headings like "第1天: …" or "Day 1 …".
func parseContent(raw string) []dayEntry {
	var entries []dayEntry
	var cur *dayEntry

	for _, line := range strings.Split(raw, "\n") {
		trimmed := strings.TrimSpace(line)
		if day := parseDayNumber(trimmed); day >= 1 && day <= 7 {
			if cur != nil {
				cur.Content = strings.TrimSpace(cur.Content)
				entries = append(entries, *cur)
			}
			// Extract title after separator if present
			title := trimmed
			if idx := strings.IndexAny(trimmed, ":："); idx != -1 {
				if rest := strings.TrimSpace(trimmed[idx+1:]); rest != "" {
					title = rest
				} else {
					title = strings.TrimSpace(trimmed[:idx])
				}
			}
			cur = &dayEntry{Day: day, Title: title}
		} else if cur != nil {
			cur.Content += line + "\n"
		}
	}
	if cur != nil {
		cur.Content = strings.TrimSpace(cur.Content)
		entries = append(entries, *cur)
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Day < entries[j].Day })
	return entries
}

// ─── main sync ────────────────────────────────────────────

// Sync fetches the configured wiki node and upserts learn_steps rows.
func (s *Syncer) Sync() error {
	log.Println("[feishu] sync start")

	root, err := s.getNode(s.wikiToken)
	if err != nil {
		return fmt.Errorf("get root node: %w", err)
	}
	log.Printf("[feishu] root: title=%q type=%s hasChild=%v", root.Title, root.ObjType, root.HasChild)

	var steps []model.LearnStep

	if root.HasChild {
		// Each child wiki node represents one learning day.
		children, err := s.getChildren(root.SpaceID, root.NodeToken)
		if err != nil {
			return fmt.Errorf("get children: %w", err)
		}

		for i, child := range children {
			day := parseDayNumber(child.Title)
			if day == 0 {
				day = i + 1 // fall back to insertion order
			}
			if day < 1 || day > 7 {
				continue
			}

			var content string
			if child.ObjType == "docx" {
				content, err = s.getRawContent(child.ObjToken)
				if err != nil {
					log.Printf("[feishu] warn: content for day %d: %v", day, err)
				}
			}

			steps = append(steps, model.LearnStep{
				Day:         day,
				Title:       child.Title,
				Description: extractSummary(content),
				Content:     content,
				Status:      1,
			})
		}
	} else {
		// Single document – parse section headings for each day.
		if root.ObjType != "docx" {
			return fmt.Errorf("unsupported obj_type %q (only docx is supported)", root.ObjType)
		}
		raw, err := s.getRawContent(root.ObjToken)
		if err != nil {
			return fmt.Errorf("get doc content: %w", err)
		}
		for _, e := range parseContent(raw) {
			steps = append(steps, model.LearnStep{
				Day:         e.Day,
				Title:       e.Title,
				Description: extractSummary(e.Content),
				Content:     e.Content,
				Status:      1,
			})
		}
	}

	if len(steps) == 0 {
		return errors.New("sync produced 0 learn steps — check document structure")
	}

	// Upsert: update existing row or create new one.
	for i := range steps {
		step := &steps[i]
		var existing model.LearnStep
		err := store.DB.Where("day = ?", step.Day).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := store.DB.Create(step).Error; err != nil {
				log.Printf("[feishu] create day %d: %v", step.Day, err)
			}
		} else if err == nil {
			store.DB.Model(&existing).Updates(map[string]any{
				"title":       step.Title,
				"description": step.Description,
				"content":     step.Content,
				"status":      step.Status,
			})
		} else {
			log.Printf("[feishu] query day %d: %v", step.Day, err)
		}
	}

	log.Printf("[feishu] sync done — %d steps written", len(steps))
	return nil
}
