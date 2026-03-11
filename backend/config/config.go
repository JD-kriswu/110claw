package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	DBDriver      string
	DBDSN         string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func Load() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_DSN", "root:password@tcp(127.0.0.1:3306)/110claw?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	// Optionally load .env file if present
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig()

	return &Config{
		ServerPort:    viper.GetString("SERVER_PORT"),
		DBDriver:      viper.GetString("DB_DRIVER"),
		DBDSN:         viper.GetString("DB_DSN"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),
	}
}