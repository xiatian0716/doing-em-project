package scheduler

import "go-crawler-test/engine"

type QueuedScheduler struct {
	// requestQ    <- s.requestChan     <- request
	// workerChanQ <- s.workerChanChan  <- workerChan
	requestChan    chan engine.Request
	workerChanChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) WorkerReady(workerChan chan engine.Request) {
	s.workerChanChan <- workerChan
}

func (s *QueuedScheduler) Run() {
	// 因为要生成它们，我们改变了s的内容，所以都要改成
	// 指针(*)接受者，指针(*)接受者才能改变里面的内容
	s.requestChan = make(chan engine.Request)
	s.workerChanChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerChanQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorkerChan chan engine.Request
			if len(requestQ) > 0 &&
				len(workerChanQ) > 0 {
				activeRequest = requestQ[0]
				activeWorkerChan = workerChanQ[0]
			}

			select {
			case request := <-s.requestChan:
				requestQ = append(requestQ, request)
			case activeWorkerChan <- activeRequest:
				requestQ = requestQ[1:]
				workerChanQ = workerChanQ[1:]
			case workerChan := <-s.workerChanChan:
				workerChanQ = append(workerChanQ, workerChan)
			}
		}
	}()
}
