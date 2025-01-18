package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env:"WAREHOUSE_ENV" env-default:"local"`
	HTTPServer `yaml:"http_server" env:"WAREHOUSE_SERVER"`
	Mongo      `yaml:"mongo" env:"WAREHOUSE_MONGO"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password" env:"HTTP_SERVER_PASSWORD"`
}

type Mongo struct {
	Password         string        `yaml:"password"`
	Login            string        `yaml:"login"`
	LowExecutionTime time.Duration `yaml:"low_execution_time"`
	ConnectCount     uint64        `yaml:"connect_count"`
	Hosts            []string      `yaml:"hosts"`
	DbName           string        `yaml:"db_name"`
	AuthSource       string        `yaml:"auth_source"`
}

func ParseConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
