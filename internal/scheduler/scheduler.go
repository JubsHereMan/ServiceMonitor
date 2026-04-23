package scheduler

import (
	"monitoramento/internal/services"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()

	services.UpdateStatus()

	c.AddFunc("@every 10s",func() {
		services.UpdateStatus()
	})

	c.Start()
}