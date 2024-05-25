package main

import (
	"agentYL/internal/app"
	"agentYL/internal/config"
	"agentYL/internal/service"
)

func main() {
	const (
		COMPUTING_POWER = 5
		PORT            = 8080
	)

	logger := config.LoadLogger()
	// Загрузка конфигурации
	cfg, err := config.LoadConfig(PORT, COMPUTING_POWER)
	if err != nil {
		logger.Fatalf("Could not load config: %s\n", err.Error())
		return
	}
	// новый service
	srv := service.New(cfg, logger)
	app.Run(srv)
}
