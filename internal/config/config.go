package config

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Sem  chan struct{}
	Wg   sync.WaitGroup
	port uint16
}

func (c *Config) GetURL() string {
	return fmt.Sprintf("http://localhost:%d/internal/task", c.port)
}

// CustomFormatter определяет свой собственный формат вывода для логгера
type CustomFormatter struct{}

// Format форматирует запись лога с заданным форматом времени
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] [%s] %s\n",
		entry.Time.Format("15:04:05.0000000"), // Формат времени: часы:минуты:секунды.микросекунды
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)), nil
}

// LoadConfig принимает порт для сервера и длительность математических операций и возвращает конфиг
func LoadConfig(port uint16, power uint) (*Config, error) {
	if power <= 0 {
		return nil, errors.New("invalid power, change configuration")
	}

	return &Config{
		Sem:  make(chan struct{}, power),
		port: port,
	}, nil
}

func LoadLogger(debugLevel bool) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(os.Stdout)
	if debugLevel {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	return log
}
