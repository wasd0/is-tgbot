package config

import (
	"is-tgbot/internal/keys"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	EnvDev   = "dev"
	EnvProd  = "prod"
	EnvStage = "stage"
)

type Config struct {
	Env     string `yaml:"env" env-default:"prod" env-required:"true"`
	LogPath string `yaml:"log_path"`
	Server  server `yaml:"server" env-required:"true"`
}

type server struct {
	Port        string        `yaml:"port" env-default:"8100"`
	Host        string        `yaml:"host" env-default:""`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"69s"`
}

var isLoaded = false

func MustRead() *Config {

	if !isLoaded {
		isLoaded = true
		err := godotenv.Load()

		if err != nil {
			panic("Error loading .env file")
		}
	}

	configPath := os.Getenv(keys.EnvConfig)

	if len(configPath) == 0 {
		panic("Config file path not found in environment")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exists by path: " + configPath)
	}

	config := Config{}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("Error while Config read")
	}

	return &config
}
