package pool

import (
	"log"

	work "github.com/shibe23/example-worker_pool/work"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan Work
	End           chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.WorkerChannel:
				work.DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping\n", w.ID)
}
