package main

import (
	"SIVTGAdmin/common/logger"
	"SIVTGAdmin/internal/bot"
	"SIVTGAdmin/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.InitConfig()

	if err != nil {
		panic(fmt.Sprintf("App with down with err parsing config: %v", err))
	}

	l := logger.NewLogger(cfg.App.DebugMode)

	botServise := bot.InitBotService(l)
	botServise.Start()
}
