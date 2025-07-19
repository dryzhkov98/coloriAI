package main

import (
	"coloriAI/cmd/app"
	"coloriAI/internal/config"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	botApp := app.NewApp(cfg)

	botApp.Run()

}
