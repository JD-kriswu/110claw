package main

import (
	"log"

	"github.com/110claw/backend/config"
	"github.com/110claw/backend/crawler"
	"github.com/110claw/backend/model"
	"github.com/110claw/backend/router"
	"github.com/110claw/backend/store"
)

func main() {
	cfg := config.Load()

	store.InitDB(cfg.DBDSN)
	if err := model.Migrate(store.DB); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	store.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Start news crawler scheduler if enabled
	var crawlerScheduler *crawler.Scheduler
	if cfg.NewsCrawlEnabled {
		crawlerScheduler = crawler.NewScheduler(cfg.NewsCrawlInitial, cfg.NewsCrawlDaily)
		crawlerScheduler.Start()
		defer crawlerScheduler.Stop()
	}

	r := router.Setup()

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}