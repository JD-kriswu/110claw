package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	DBDriver      string
	DBDSN         string
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	// Feishu sync (optional — sync is disabled when AppID is empty)
	FeishuAppID       string
	FeishuAppSecret   string
	FeishuWikiToken   string
	FeishuSyncInterval time.Duration
}

func Load() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_DSN", "root:password@tcp(127.0.0.1:3306)/110claw?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	viper.SetDefault("FEISHU_APP_ID", "")
	viper.SetDefault("FEISHU_APP_SECRET", "")
	viper.SetDefault("FEISHU_WIKI_TOKEN", "OCGUw7De8i1zFUkpWPQc6oi8nOb")
	viper.SetDefault("FEISHU_SYNC_INTERVAL", "1h")

	// Optionally load .env file if present
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig()

	interval, err := time.ParseDuration(viper.GetString("FEISHU_SYNC_INTERVAL"))
	if err != nil || interval < time.Minute {
		interval = time.Hour
	}

	return &Config{
		ServerPort:    viper.GetString("SERVER_PORT"),
		DBDriver:      viper.GetString("DB_DRIVER"),
		DBDSN:         viper.GetString("DB_DSN"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),

		FeishuAppID:        viper.GetString("FEISHU_APP_ID"),
		FeishuAppSecret:    viper.GetString("FEISHU_APP_SECRET"),
		FeishuWikiToken:    viper.GetString("FEISHU_WIKI_TOKEN"),
		FeishuSyncInterval: interval,
	}
}
