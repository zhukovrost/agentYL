package main

import (
	"agent/internal/app"
	"agent/internal/config"
	"agent/internal/service"
	"flag"
)

func main() {
	const (
		PORT = 8080
	)

	// Определение флагов (подробнее go run cmd/agent/main.go -h)
	debug := flag.Bool("debug", false, "enable debug level logging")
	COMPUTING_POWER := flag.Uint("power", 3, "configure computing power")

	flag.Parse() // Парсинг флагов

	logger := config.LoadLogger(*debug)                   // Загрузка логгера
	cfg, err := config.LoadConfig(PORT, *COMPUTING_POWER) // Загрузка конфигурации
	if err != nil {
		logger.Fatalf("Could not load config: %s\n", err.Error())
		return
	}
	srv := service.New(cfg, logger) // новый service
	app.Run(srv)                    // запуск приложения
}
