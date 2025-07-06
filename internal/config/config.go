package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"localhost"`
	DatabaseURL string `yaml:"database_url" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"PASSWORD_HTTP"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH not found")
	}
	
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file not exist: %s", configPath)
	}
	
	var cfg Config
	
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	
	return &cfg
}