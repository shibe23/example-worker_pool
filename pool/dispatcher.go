package pool

import "log"

var WorkerChannel = make(chan Work)

type Collector struct {
	Work chan Work
	End  chan bool
}

func StartDispatcher(workerCount int) Collector {
	workers := activateWorker(workerCount)
	return createCollector(workers)
}

// 定数で指定した数だけWorkerのgoroutineを生成する
// workerは生成後workersに追加される
func activateWorker(workerCount int) []Worker {
	var i int
	var workers []Worker

	for i < workerCount {
		i++
		log.Println("Startning worker: ", i)

		// Workerを生成
		worker := Worker{
			ID:            i,
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker)
	}
	return workers
}

// Collectorを生成する
// CorellatorはJobを受け取り、Workerに割り当てる役割
// chan inputでWork(job+ID)構造体を受け取り、chan WorkerChannelに詰め直している
func createCollector(workers []Worker) Collector {
	input := make(chan Work)
	end := make(chan bool)
	collector := Collector{Work: input, End: end}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				WorkerChannel <- work // dispatch work to worker
			}
		}
	}()
	return collector
}
