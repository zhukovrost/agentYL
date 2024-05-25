package service

import (
	"agentYL/internal/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	NoTaskError = errors.New("no task")
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

type Response struct {
	Task *TaskResponse `json:"task"`
}

type TaskResponse struct {
	Id            uint    `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime uint    `json:"operation_time"`
}

func (s *Service) GetTask() (*TaskResponse, error) {
	resp, err := http.Get(s.Cfg.GetURL())
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var response Response
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}
		return response.Task, nil
	case http.StatusNotFound:
		return nil, NoTaskError
	case http.StatusInternalServerError:
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code %d, and failed to read response body: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil, fmt.Errorf("failed to get response: %w", err)
}

func (s *Service) Compute(task *TaskResponse) error {
	defer s.Cfg.Wg.Done()
	s.Cfg.Sem <- struct{}{}        // Захватываем место в семафоре
	defer func() { <-s.Cfg.Sem }() // Освобождаем место в семафоре после завершения

	var result float64
	exp := fmt.Sprintf("%f %s %f", task.Arg1, task.Operation, task.Arg2)

	s.Logger.Debugf("calculating task (ID: %d): %s...", task.Id, exp)

	switch task.Operation {
	case "+":
		result = task.Arg1 + task.Arg2
	case "-":
		result = task.Arg1 - task.Arg2
	case "*":
		result = task.Arg1 * task.Arg2
	case "/":
		result = task.Arg1 / task.Arg2
	}
	calculationTime := time.Duration(task.OperationTime) * time.Millisecond
	time.Sleep(calculationTime)
	s.Logger.Infof("calculation done (ID: %d): %s = %f", task.Id, exp, result)

	return s.Response(task.Id, result)
}

type ResultResp struct {
	Id  uint    `json:"id"`
	Var float64 `json:"result"`
}

func (s *Service) Response(id uint, result float64) error {
	res := &ResultResp{Id: id, Var: result}

	// Преобразуем структуру в JSON
	jsonData, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Создаем HTTP POST-запрос
	resp, err := http.Post(s.Cfg.GetURL(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
