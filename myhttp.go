package main

import (
	"flag"
	"net/http"

	"github.com/http-api/httpworker"
)

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func parseCli() (int, []string) {
	var requestLimit int
	flag.IntVar(&requestLimit, "parallel", 10, "specify the parallel request limit to use. Defaults to 10.")
	flag.Parse()

	return requestLimit, flag.Args()
}

func main() {
	parallel, addrs := parseCli()
	numberOfJobs := len(addrs)

	requestQueue := make(chan string, numberOfJobs)
	resultQueue := make(chan int, numberOfJobs)

	wm := &httpworker.WorkerManager{
		Client:       &http.Client{},
		RequestQueue: requestQueue,
		ResultQueue:  resultQueue,
	}
	wm.CreateWorkerPool(min(parallel, numberOfJobs))
	wm.PopulateJobs(addrs)

	for results := 1; results <= numberOfJobs; results++ {
		<-resultQueue
	}
}
