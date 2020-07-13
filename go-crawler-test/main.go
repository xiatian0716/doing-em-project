package main

import (
	"go-crawler-test/config"
	"go-crawler-test/engine"
	"go-crawler-test/parse/start"
	"go-crawler-test/scheduler"
)

func main() {
	Configs := config.ConfigsParse("./config/config.yaml")
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: Configs.CEConfig.WorkerCount,
		ESSave:      Configs.CEConfig.ESSave,
	}
	e.Run(engine.Request{
		Url:       Configs.StartUrl,
		ParseFunc: start.Woke,
	})
}
