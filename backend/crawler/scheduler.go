package crawler

import (
	"log"
	"sync"
	"time"
)

// Scheduler manages periodic crawling.
type Scheduler struct {
	crawler       *Crawler
	queries       []string
	initialCount  int
	dailyCount    int
	stopChan      chan struct{}
	wg            sync.WaitGroup
	hasRunInitial bool
}

// Default queries for Claude Code / OpenClaw related news.
var defaultQueries = []string{
	"claude code",
	"claude-code",
	"openclaw",
	"anthropic claude",
	"claude ai",
}

// NewScheduler creates a new crawler scheduler.
func NewScheduler(initialCount, dailyCount int) *Scheduler {
	translator := NewMyMemoryTranslator()
	return &Scheduler{
		crawler:      NewCrawler(translator),
		queries:      defaultQueries,
		initialCount: initialCount,
		dailyCount:   dailyCount,
		stopChan:     make(chan struct{}),
	}
}

// Start begins the scheduled crawling.
// It runs immediately on start, then daily at 6:00 AM.
func (s *Scheduler) Start() {
	// Run initial crawl immediately
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.runInitialCrawl()
	}()

	// Start daily scheduler
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.runScheduler()
	}()

	log.Println("[crawler] scheduler started")
}

// Stop gracefully stops the scheduler.
func (s *Scheduler) Stop() {
	close(s.stopChan)
	s.wg.Wait()
	log.Println("[crawler] scheduler stopped")
}

// runInitialCrawl runs the first crawl with a larger batch.
func (s *Scheduler) runInitialCrawl() {
	if s.hasRunInitial {
		return
	}

	log.Printf("[crawler] running initial crawl, fetching %d articles per source per query", s.initialCount)
	_, err := s.crawler.CrawlMultiQuery(s.queries, s.initialCount)
	if err != nil {
		log.Printf("[crawler] initial crawl error: %v", err)
	}
	s.hasRunInitial = true
}

// runScheduler handles the daily scheduling loop.
func (s *Scheduler) runScheduler() {
	for {
		// Calculate next run time (6:00 AM local time)
		nextRun := nextSixAM()
		waitDuration := time.Until(nextRun)

		log.Printf("[crawler] next scheduled crawl at %s", nextRun.Format("2006-01-02 15:04:05"))

		select {
		case <-s.stopChan:
			return
		case <-time.After(waitDuration):
			s.runDailyCrawl()
		}
	}
}

// runDailyCrawl runs the daily crawl.
func (s *Scheduler) runDailyCrawl() {
	log.Printf("[crawler] running daily crawl, fetching %d articles per source per query", s.dailyCount)
	saved, err := s.crawler.CrawlMultiQuery(s.queries, s.dailyCount)
	if err != nil {
		log.Printf("[crawler] daily crawl error: %v", err)
	} else {
		log.Printf("[crawler] daily crawl completed, saved %d new articles", saved)
	}
}

// nextSixAM returns the next 6:00 AM time.
// If it's before 6 AM today, returns today's 6 AM.
// Otherwise, returns tomorrow's 6 AM.
func nextSixAM() time.Time {
	now := time.Now()
	sixAM := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
	if now.Before(sixAM) {
		return sixAM
	}
	return sixAM.Add(24 * time.Hour)
}

// RunOnce runs a single crawl immediately (for manual triggering).
func (s *Scheduler) RunOnce(limitPerSource int) (int, error) {
	return s.crawler.CrawlMultiQuery(s.queries, limitPerSource)
}