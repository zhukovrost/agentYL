package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"time"

	"fmt"
	"os"
	"strings"
)

type Config struct {
	Power          uint16        `json:"power"`
	Addition       time.Duration `json:"time_addition_ms"`
	Subtraction    time.Duration `json:"time_subtraction_ms"`
	Multiplication time.Duration `json:"time_multiplications_ms"`
	Division       time.Duration `json:"time_divisions_ms"`
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
func LoadConfig(power, addition, subtraction, multiplication, division int) (*Config, error) {
	if addition <= 0 || subtraction <= 0 || multiplication <= 0 || division <= 0 || power <= 0 {
		return nil, errors.New("invalid duration, change configuration")
	}

	return &Config{
		Power:          uint16(power),
		Addition:       time.Millisecond * time.Duration(addition),
		Subtraction:    time.Millisecond * time.Duration(subtraction),
		Multiplication: time.Millisecond * time.Duration(multiplication),
		Division:       time.Millisecond * time.Duration(division),
	}, nil
}

func LoadLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	return log
}
