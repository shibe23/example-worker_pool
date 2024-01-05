package main

import (
	"log"

	"github.com/shibe23/example-worker_pool/pool"
	work "github.com/shibe23/example-worker_pool/work"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("Starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT)

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <- pool.Work{Job: job, ID: i} // Worker(job+ID)を生成してcollectorにchannelとして送信する
	}
}
