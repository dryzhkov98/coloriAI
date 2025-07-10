package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

var configPath = ".env"

type Config struct {
	BotToken         string `env:"BOT_TOKEN"`
	BotName          string `env:"BOT_NAME"`
	DebugMode        bool   `env:"DEBUG_MODE"`
	UpdateBotTimeout int    `env:"UPDATE_BOT_TIMEOUT" default:"60"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s is not exist", configPath)
	}

	err := cleanenv.ReadConfig(".env", &s.config)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) GetBotToken() string {
	return s.config.BotToken
}

func (s *Service) GetConfig() Config {
	return s.config
}
