package main

import (
	"coloriAI/internal/bot"
	"coloriAI/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}

	cfgStruct := cfg.GetConfig()

	tgBot, err := bot.New(&cfgStruct)
	if err != nil {
		panic(err.Error())
	}

	tgBot.Start()
}
