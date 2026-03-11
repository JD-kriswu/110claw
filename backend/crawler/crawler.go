package crawler

import (
	"log"
	"sync"
	"time"

	"github.com/110claw/backend/model"
	"github.com/110claw/backend/store"
	"gorm.io/gorm/clause"
)

// Crawler orchestrates fetching from multiple sources.
type Crawler struct {
	sources []Source
}

// NewCrawler creates a new Crawler with all available sources.
func NewCrawler(translator Translator) *Crawler {
	return &Crawler{
		sources: []Source{
			NewHackerNewsSource(translator),
			NewRedditSource(translator),
		},
	}
}

// FetchAll fetches articles from all sources concurrently.
func (c *Crawler) FetchAll(query string, limitPerSource int) []Article {
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		articles []Article
	)

	for _, source := range c.sources {
		wg.Add(1)
		go func(s Source) {
			defer wg.Done()
			log.Printf("[crawler] fetching from %s with query: %s", s.Name(), query)
			result, err := s.Fetch(query, limitPerSource)
			if err != nil {
				log.Printf("[crawler] %s fetch error: %v", s.Name(), err)
				return
			}
			mu.Lock()
			articles = append(articles, result...)
			mu.Unlock()
			log.Printf("[crawler] %s returned %d articles", s.Name(), len(result))
		}(source)
	}

	wg.Wait()
	return articles
}

// SaveArticles saves articles to the database, skipping duplicates by SourceURL.
func (c *Crawler) SaveArticles(articles []Article) (int, error) {
	var saved int
	now := time.Now()

	for _, article := range articles {
		// Skip if SourceURL is empty
		if article.SourceURL == "" {
			continue
		}

		// Check if article already exists
		var existing model.News
		result := store.DB.Where("source_url = ?", article.SourceURL).First(&existing)
		if result.Error == nil {
			// Article already exists, skip
			continue
		}

		// Create new News record with truncated fields
		news := model.News{
			Title:       truncateString(article.Title, 200),
			Summary:     truncateString(article.Summary, 500),
			Content:     article.Content,
			SourceURL:   article.SourceURL,
			SourceName:   article.SourceName,
			Tags:        truncateString(article.Tags, 200),
			Status:      1, // published
			PublishedAt: article.PublishedAt,
			ViewCount:   0,
		}

		// Set PublishedAt to now if not provided
		if news.PublishedAt == nil {
			news.PublishedAt = &now
		}

		// Use upsert to handle potential race conditions
		if err := store.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "source_url"}},
			DoNothing: true,
		}).Create(&news).Error; err != nil {
			log.Printf("[crawler] failed to save article %s: %v", article.SourceURL, err)
			continue
		}

		saved++
	}

	return saved, nil
}

// truncateString truncates a string to maxLen characters.
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// Crawl fetches articles from all sources and saves them to the database.
func (c *Crawler) Crawl(query string, limitPerSource int) (int, error) {
	articles := c.FetchAll(query, limitPerSource)
	log.Printf("[crawler] total articles fetched: %d", len(articles))

	saved, err := c.SaveArticles(articles)
	if err != nil {
		return saved, err
	}

	log.Printf("[crawler] saved %d new articles", saved)
	return saved, nil
}

// CrawlMultiQuery runs multiple queries and saves results.
func (c *Crawler) CrawlMultiQuery(queries []string, limitPerSource int) (int, error) {
	var totalSaved int

	for _, query := range queries {
		saved, err := c.Crawl(query, limitPerSource)
		if err != nil {
			log.Printf("[crawler] error crawling query '%s': %v", query, err)
			continue
		}
		totalSaved += saved
	}

	return totalSaved, nil
}