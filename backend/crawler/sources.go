package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Article represents a crawled article from any source.
type Article struct {
	Title       string
	Summary     string
	Content     string
	SourceURL   string
	SourceName  string
	Tags        string
	PublishedAt *time.Time
}

// Source defines the interface for news data sources.
type Source interface {
	Fetch(query string, limit int) ([]Article, error)
	Name() string
}

// Translator defines the interface for translation services.
type Translator interface {
	Translate(text, targetLang string) (string, error)
}

// HackerNewsSource fetches articles from Hacker News via Algolia API.
type HackerNewsSource struct {
	client     *http.Client
	translator Translator
}

func NewHackerNewsSource(translator Translator) *HackerNewsSource {
	return &HackerNewsSource{
		client:     &http.Client{Timeout: 30 * time.Second},
		translator: translator,
	}
}

func (s *HackerNewsSource) Name() string {
	return "hackernews"
}

type hnResponse struct {
	Hits []hnHit `json:"hits"`
}

type hnHit struct {
	ObjectID    string `json:"objectID"`
	Title       string `json:"title"`
	StoryText   string `json:"story_text"`
	URL         string `json:"url"`
	Author      string `json:"author"`
	CreatedAt   int64  `json:"created_at_i"`
	Points      int    `json:"points"`
	NumComments int    `json:"num_comments"`
}

func (s *HackerNewsSource) Fetch(query string, limit int) ([]Article, error) {
	encodedQuery := url.QueryEscape(query)
	apiURL := fmt.Sprintf("https://hn.algolia.com/api/v1/search?query=%s&hitsPerPage=%d&tags=story", encodedQuery, limit)

	resp, err := s.client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("hackernews fetch error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hackernews api returned status %d", resp.StatusCode)
	}

	var result hnResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("hackernews decode error: %w", err)
	}

	var articles []Article
	for _, hit := range result.Hits {
		sourceURL := hit.URL
		if sourceURL == "" {
			sourceURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%s", hit.ObjectID)
		}

		publishedAt := time.Unix(hit.CreatedAt, 0)

		// Translate title and summary
		title := hit.Title
		summary := truncateText(hit.StoryText, 500)
		content := hit.StoryText

		if s.translator != nil {
			if translated, err := s.translator.Translate(title, "zh"); err == nil {
				title = translated
			}
			if summary != "" && s.translator != nil {
				if translated, err := s.translator.Translate(summary, "zh"); err == nil {
					summary = translated
				}
			}
		}

		articles = append(articles, Article{
			Title:       title,
			Summary:     summary,
			Content:     content,
			SourceURL:   sourceURL,
			SourceName:  s.Name(),
			Tags:        "hackernews,tech",
			PublishedAt: &publishedAt,
		})
	}

	return articles, nil
}

// RedditSource fetches posts from Reddit API.
type RedditSource struct {
	client     *http.Client
	translator Translator
}

func NewRedditSource(translator Translator) *RedditSource {
	return &RedditSource{
		client:     &http.Client{Timeout: 30 * time.Second},
		translator: translator,
	}
}

func (s *RedditSource) Name() string {
	return "reddit"
}

type redditResponse struct {
	Data struct {
		Children []struct {
			Data redditPost `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type redditPost struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	SelfText    string  `json:"selftext"`
	URL         string  `json:"url"`
	Permalink   string  `json:"permalink"`
	Subreddit   string  `json:"subreddit"`
	Author      string  `json:"author"`
	Score       int     `json:"score"`
	NumComments int     `json:"num_comments"`
	CreatedUTC  float64 `json:"created_utc"`
}

func (s *RedditSource) Fetch(query string, limit int) ([]Article, error) {
	encodedQuery := url.QueryEscape(query)
	apiURL := fmt.Sprintf("https://www.reddit.com/search.json?q=%s&sort=relevance&limit=%d&type=link", encodedQuery, limit)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("reddit request error: %w", err)
	}
	req.Header.Set("User-Agent", "110claw-crawler/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("reddit fetch error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reddit api returned status %d", resp.StatusCode)
	}

	var result redditResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("reddit decode error: %w", err)
	}

	var articles []Article
	for _, child := range result.Data.Children {
		post := child.Data
		sourceURL := post.URL
		if post.Permalink != "" && !strings.HasPrefix(post.Permalink, "http") {
			sourceURL = "https://www.reddit.com" + post.Permalink
		}

		publishedAt := time.Unix(int64(post.CreatedUTC), 0)

		// Translate title and summary
		title := post.Title
		summary := truncateText(post.SelfText, 500)
		content := post.SelfText

		if s.translator != nil {
			if translated, err := s.translator.Translate(title, "zh"); err == nil {
				title = translated
			}
			if summary != "" && s.translator != nil {
				if translated, err := s.translator.Translate(summary, "zh"); err == nil {
					summary = translated
				}
			}
		}

		tags := "reddit," + post.Subreddit
		articles = append(articles, Article{
			Title:       title,
			Summary:     summary,
			Content:     content,
			SourceURL:   sourceURL,
			SourceName:  s.Name(),
			Tags:        tags,
			PublishedAt: &publishedAt,
		})
	}

	return articles, nil
}

// MyMemoryTranslator uses MyMemory free translation API with fallback to Google Translate.
type MyMemoryTranslator struct {
	client *http.Client
}

func NewMyMemoryTranslator() *MyMemoryTranslator {
	return &MyMemoryTranslator{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

type myMemoryResponse struct {
	ResponseStatus int    `json:"responseStatus"`
	ResponseData   struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
	QuotaFinished bool `json:"quotaFinished"`
}

func (t *MyMemoryTranslator) Translate(text, targetLang string) (string, error) {
	if text == "" {
		return "", nil
	}

	// First try MyMemory API
	translated, err := t.translateMyMemory(text, targetLang)
	if err == nil && translated != "" && !strings.Contains(translated, "MYMEMORY WARNING") {
		return translated, nil
	}

	// Fallback to Google Translate unofficial API
	return t.translateGoogle(text, targetLang)
}

func (t *MyMemoryTranslator) translateMyMemory(text, targetLang string) (string, error) {
	apiURL := fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=en|%s",
		url.QueryEscape(text), targetLang)

	resp, err := t.client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("mymemory fetch error: %w", err)
	}
	defer resp.Body.Close()

	var result myMemoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("mymemory decode error: %w", err)
	}

	if result.ResponseStatus != 200 {
		return "", fmt.Errorf("mymemory api returned status %d", result.ResponseStatus)
	}

	return result.ResponseData.TranslatedText, nil
}

func (t *MyMemoryTranslator) translateGoogle(text, targetLang string) (string, error) {
	// Google Translate unofficial API
	apiURL := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=%s&dt=t&q=%s",
		targetLang, url.QueryEscape(text))

	resp, err := t.client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("google translate fetch error: %w", err)
	}
	defer resp.Body.Close()

	// Response format: [["translated text","original text",null,null,10],null,"en",...]
	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("google translate decode error: %w", err)
	}

	if len(result) == 0 {
		return "", fmt.Errorf("google translate empty response")
	}

	// Extract translated text from the response
	if sentences, ok := result[0].([]interface{}); ok && len(sentences) > 0 {
		if first, ok := sentences[0].([]interface{}); ok && len(first) > 0 {
			if translated, ok := first[0].(string); ok {
				return translated, nil
			}
		}
	}

	return "", fmt.Errorf("google translate unexpected response format")
}

// truncateText truncates text to maxLen characters.
func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	return text[:maxLen-3] + "..."
}