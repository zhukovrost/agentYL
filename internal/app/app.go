package app

import (
	"agentYL/internal/service"
	"errors"
	"time"
)

func Run(srv *service.Service) {
	for {
		task, err := srv.GetTask()
		if err != nil {
			if errors.Is(err, service.NoTaskError) {
				srv.Logger.Debug("no task")
				time.Sleep(1 * time.Second) // Подождем перед следующей попыткой
				continue
			} else {
				srv.Logger.Errorf("failed to get task: %v", err.Error())
				time.Sleep(1 * time.Second) // Подождем перед следующей попыткой
				continue
			}
		}

		srv.Logger.Infof("got new task: %d", task.Id)
		srv.Cfg.Wg.Add(1)
		go srv.Compute(task)
		time.Sleep(1 * time.Second)
	}
	srv.Cfg.Wg.Wait()
}
