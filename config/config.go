package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres
	Server
}

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Postgres struct {
	PgHost     string
	PgPort     string
	PgUsername string
	PgDBName   string
	PgSSLMode  string
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
