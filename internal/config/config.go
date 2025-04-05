package config

import (
	"errors"
	"io/fs"
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
	ReadTimeout time.Duration `yaml:"read_timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Mongo struct {
	Password     string   `yaml:"password"`
	Login        string   `yaml:"login"`
	ConnectCount uint64   `yaml:"connect_count"`
	Hosts        []string `yaml:"hosts"`
	DbName       string   `yaml:"db_name"`
	AuthSource   string   `yaml:"auth_source"`
}

// ParseConfig parses config from yaml to Config
func ParseConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = ".env"
		log.Println("CONFIG_PATH is empty. parsing ENV")
	}

	if _, err := os.Stat(configPath); errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}
