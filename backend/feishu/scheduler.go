package feishu

import (
	"log"
	"time"
)

// StartScheduler runs Sync immediately in a goroutine, then repeats
// every interval. It never blocks the caller.
func StartScheduler(s *Syncer, interval time.Duration) {
	go func() {
		if err := s.Sync(); err != nil {
			log.Printf("[feishu] initial sync error: %v", err)
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			if err := s.Sync(); err != nil {
				log.Printf("[feishu] sync error: %v", err)
			}
		}
	}()
}
