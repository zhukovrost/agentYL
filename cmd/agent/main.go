package main

import (
	"agent/internal/app"
	"agent/internal/config"
	"agent/internal/service"
	"flag"
	"os"
	"strconv"
)

func main() {
	const (
		PORT = 8080
	)

	// Определение флагов (подробнее go run cmd/agent/main.go -h)
	debug := flag.Bool("debug", false, "enable debug level logging")
	flag.Parse() // Парсинг флагов

	powerStr := os.Getenv("COMPUTING_POWER")
	var power uint
	power = 3
	if powerStr != "" {
		if p, err := strconv.Atoi(powerStr); err == nil {
			power = uint(p)
		}
	}

	logger := config.LoadLogger(*debug)        // Загрузка логгера
	cfg, err := config.LoadConfig(PORT, power) // Загрузка конфигурации
	if err != nil {
		logger.Fatalf("Could not load config: %s\n", err.Error())
		return
	}
	srv := service.New(cfg, logger) // новый service
	app.Run(srv)                    // запуск приложения
}
