package pool

import "log"

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work chan Work
	End  chan bool
}

func StartDispatcher(workerCount int) Collector {
	workers := activateWorker(workerCount)
	return createCollectaor(workers)
}

func activateWorker(workerCount int) []Worker {
	var i int
	var workers []Worker

	for i < workerCount {
		i++
		log.Println("Startning worker: ", i)

		// Workerを生成
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker)
	}
	return workers
}

// input channelでWorkを受け取り、WorkerChannelに詰め直す
func createCollectaor(workers []Worker) Collector {
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
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}
		}
	}()
	return collector
}
