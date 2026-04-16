package scheduler

import (
	"monitoramento/internal/services"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()

	go services.UpdateStatus()

	c.AddFunc("@every 1m",func() {
		services.UpdateStatus()
	})

	c.Start()
}