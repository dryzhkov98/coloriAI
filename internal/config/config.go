package config

import (
	"coloriAI/internal/utils/dictionary"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

var configPath = ".env"

type (
	Config struct {
		AppConfig AppConfig
		BotConfig BotConfig
		DBConfig  DbConfig
	}

	AppConfig struct {
		LogLevel    string `env:"LOG_LEVEL" default:"info"`
		Environment string `env:"ENVIRONMENT"`
	}

	BotConfig struct {
		BotToken         string `env:"BOT_TOKEN"`
		BotName          string `env:"BOT_NAME"`
		DebugMode        bool   `env:"DEBUG_MODE"`
		UpdateBotTimeout int    `env:"UPDATE_BOT_TIMEOUT" default:"60"`
		Dictionary       *dictionary.Dictionary
	}

	DbConfig struct {
		DbUrl       string        `env:"DB_URL"`
		MaxConns    int32         `env:"MAX_CONNS"`
		MinConns    int32         `env:"MIN_CONNS"`
		MaxConnIdle time.Duration `env:"MAX_CONN_IDLE"`
	}
)

func New() (*Config, error) {
	conf := &Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s is not exist", configPath)
	}

	err := cleanenv.ReadConfig(".env", conf)
	if err != nil {
		return nil, err
	}

	err = conf.loadDictionary()
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (cfg *Config) loadDictionary() error {
	d, err := dictionary.MustLoad("ru")
	if err != nil {
		return err
	}

	cfg.BotConfig.Dictionary = d

	return nil
}
