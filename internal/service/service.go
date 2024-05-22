package service

import (
	"agentYL/internal/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Service struct {
	Cfg    *config.Config
	Logger *logrus.Logger
}

func New(cfg *config.Config, logger *logrus.Logger) *Service {
	return &Service{
		Cfg:    cfg,
		Logger: logger,
	}
}

func (s *Service) evaluatePostfix(postfix []string) (float64, error) {
	stack := make([]float64, 0)
	for _, token := range postfix {
		if operand, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, operand)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64

			switch token {
			case "+":
				result = a + b
				time.Sleep(s.Cfg.Addition)
			case "-":
				result = a - b
				time.Sleep(s.Cfg.Subtraction)
			case "*":
				result = a * b
				time.Sleep(s.Cfg.Multiplication)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
				time.Sleep(s.Cfg.Division)
			default:
				return 0, fmt.Errorf("invalid operator: %s", token)
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}
