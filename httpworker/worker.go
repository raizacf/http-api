package httpworker

import (
	"fmt"
	"log"

	"github.com/http-api/apiclient"
)

// WorkerManager handles the queues related to the http workers
type WorkerManager struct {
	Client       apiclient.HTTPClient
	RequestQueue chan string
	ResultQueue  chan int
}

// CreateWorkerPool creates a pool of workers to handle http requests in parallel
func (wm *WorkerManager) CreateWorkerPool(workerNum int) {
	for i := 0; i < workerNum; i++ {
		go wm.worker()
	}
}

// PopulateJobs populate the RequestQueue with all given urls
func (wm *WorkerManager) PopulateJobs(addrs []string) {
	for _, addr := range addrs {
		wm.RequestQueue <- addr
	}

	close(wm.RequestQueue)
}

func (wm *WorkerManager) worker() {
	for addr := range wm.RequestQueue {
		url, hash, err := apiclient.HashResponse(addr, wm.Client)
		if err != nil {
			log.Printf("Impossible to reach given URL: %s, Please refer to the error below:\n%s", url, err)
			wm.ResultQueue <- 1
			continue
		}

		fmt.Println(url, hash)
		wm.ResultQueue <- 1
	}
}
