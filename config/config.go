package config

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres     PostgresData
	CloudflareR2 CloudflareR2 // used for image
}

type PostgresData struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

type CloudflareR2 struct {
	R2_ACCESS_KEY_ID     string
	R2_SECRET_ACCESS_KEY string
	R2_ACCOUNT_ID        string
}

func initConfig(path string) error {
	envPath := filepath.Join(path, ".env")
	if err := godotenv.Load(envPath); err != nil {
		return err
	}

	v := viper.New()
	v.AutomaticEnv() // Load from environment

	cfg = Config{
		Postgres: PostgresData{
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetString("POSTGRES_PORT"),
			Username: v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			DB:       v.GetString("POSTGRES_DB"),
		},
		CloudflareR2: CloudflareR2{
			R2_ACCESS_KEY_ID:     v.GetString("R2_ACCESS_KEY_ID"),
			R2_SECRET_ACCESS_KEY: v.GetString("R2_SECRET_ACCESS_KEY"),
			R2_ACCOUNT_ID:        v.GetString("R2_ACCOUNT_ID"),
		},
		// JWTSecretKey: v.GetString("JWT_SECRET_KEY"),
	}

	return nil
}

var (
	cfg  Config
	once sync.Once
)

func LoadConfig(path string) Config {
	once.Do(func() {
		if err := initConfig(path); err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
	})
	return cfg
}
