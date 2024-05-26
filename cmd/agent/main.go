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
		DEBUG_LEVEL     = true
	)

	logger := config.LoadLogger(DEBUG_LEVEL)             // Загрузка логгера
	cfg, err := config.LoadConfig(PORT, COMPUTING_POWER) // Загрузка конфигурации
	if err != nil {
		logger.Fatalf("Could not load config: %s\n", err.Error())
		return
	}
	srv := service.New(cfg, logger) // новый service
	app.Run(srv)                    // запуск приложения
}
