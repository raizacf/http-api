package main

import (
	"./httpclient"
	"flag"
)

func worker(requestQueue chan string, resultQueue chan int) {
	for url := range requestQueue {
		HandleURL(url)
		resultQueue <- 1
	}
}

func createWorkerPool(workerNum int, requestQueue chan string, resultQueue chan int) {
	for i := 0; i < workerNum; i++ {
		go worker(requestQueue, resultQueue)
	}
}

func populateJobs(addrs []string, requestQueue chan string) {
	for _, addr := range addrs {
		requestQueue <- addr
	}
	close(requestQueue)
}

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

	createWorkerPool(min(parallel, numberOfJobs), requestQueue, resultQueue)
	populateJobs(addrs, requestQueue)

	for a := 1; a <= numberOfJobs; a++ {
		<-resultQueue
	}
}
