package app

import (
	"agent/internal/service"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(srv *service.Service) {
	// Создаем контекст, который может быть отменён
	ctx, cancel := context.WithCancel(context.Background())

	// Настраиваем канал для прослушивания сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для прослушивания сигналов ОС и отмены контекста
	go func() {
		<-sigChan
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			srv.Logger.Info("shutting down gracefully...")
			srv.Cfg.Wg.Wait() // Ожидаем завершения всех задач
			srv.Logger.Info("all tasks completed, exiting")
			return
		default:
			task, err := srv.GetTask()
			if err != nil {
				if !errors.Is(err, service.NoTaskError) {
					srv.Logger.Errorf("failed to get task: %v", err.Error())
				}
				time.Sleep(1 * time.Second) // Подождем перед следующей попыткой
				continue
			}

			srv.Logger.Infof("got new task: %d", task.Id)
			srv.Cfg.Wg.Add(1)
			go func() {
				if err := srv.Compute(task); err != nil {
					srv.Logger.Fatalf("failed to compute task: %v", err.Error())
				}
			}()
			time.Sleep(1 * time.Second)
		}
	}
}
