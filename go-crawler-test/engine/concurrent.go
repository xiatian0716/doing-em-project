package engine

import (
	"go-crawler-test/persist"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ESSave      bool
}

type Scheduler interface {
	Submit(Request)           // s.requestChan <- request
	WorkerChan() chan Request // make(chan engine.Request)
	ReadyNotifier             // s.workerChan <- workerChan
	Run()                     // (s *QueuedScheduler) Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 启动Scheduler
	e.Scheduler.Run() // (s *QueuedScheduler) Run()

	// 启动Request
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	// 启动Worker
	out := make(chan ParseResult) // make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, e.Scheduler.WorkerChan(), out)
	}

	// 拿到output
	itemCount := 0
	outSaveChan, _ := persist.ItemSaver()
	for {
		result := <-out

		// Items结果
		for _, item := range result.Items {
			if e.ESSave {
				outSaveChan <- item // 存储
			} else {
				log.Printf("Got item#%d: %s\n", itemCount, item) // 打印
			}
			itemCount++
		}

		// 把新的Requesrts送给Scheduler加进去
		for _, request := range result.Requesrts {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
func createWorker(ready ReadyNotifier, workerChan chan Request, out chan<- ParseResult) {
	go func() {
		for {
			ready.WorkerReady(workerChan)
			request := <-workerChan
			// fetcher网页body(Url+ParseFunc)
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
