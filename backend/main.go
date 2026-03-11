package main

import (
	"log"

	"github.com/110claw/backend/config"
	"github.com/110claw/backend/feishu"
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

	// Start Feishu sync scheduler if credentials are configured.
	if cfg.FeishuAppID != "" && cfg.FeishuAppSecret != "" {
		syncer := feishu.New(cfg.FeishuAppID, cfg.FeishuAppSecret, cfg.FeishuWikiToken)
		feishu.StartScheduler(syncer, cfg.FeishuSyncInterval)
		log.Printf("Feishu sync enabled (wiki=%s interval=%s)",
			cfg.FeishuWikiToken, cfg.FeishuSyncInterval)
	} else {
		log.Println("Feishu sync disabled (FEISHU_APP_ID not set)")
	}

	r := router.Setup()

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
