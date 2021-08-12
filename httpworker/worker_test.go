package httpworker_test

import (
	"errors"
	"net/http"
	"runtime"
	"testing"

	"github.com/http-api/apiclient/mocks"
	"github.com/http-api/httpworker"
)

func TestCreateWorkerPool(t *testing.T) {
	previouslyRunning := runtime.NumGoroutine()

	requestQueue := make(chan string, 3)
	resultQueue := make(chan int, 3)

	wm := &httpworker.WorkerManager{
		Client:       nil,
		RequestQueue: requestQueue,
		ResultQueue:  resultQueue,
	}
	wm.CreateWorkerPool(3)

	if runtime.NumGoroutine() != previouslyRunning+3 {
		t.Errorf("TestCreateWorkerPool failed, expected 3 created routines, got %v.", runtime.NumGoroutine()-previouslyRunning)
	}
}

func TestCreateWorkerPoolRemoveJobAfterWork(t *testing.T) {
	requestQueue := make(chan string, 3)
	resultQueue := make(chan int, 3)

	wm := &httpworker.WorkerManager{
		Client:       &mocks.MockClient{},
		RequestQueue: requestQueue,
		ResultQueue:  resultQueue,
	}
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}

	wm.PopulateJobs([]string{"sample_url"})
	wm.CreateWorkerPool(1)

	for a := 1; a <= 1; a++ {
		<-resultQueue
	}

	if len(wm.RequestQueue) != 0 {
		t.Errorf("TestCreateWorkerPoolRemoveJobAfterWork failed, expected no job to run, got %v.", len(wm.RequestQueue))
	}

}

func TestPopulateJobs(t *testing.T) {
	requestQueue := make(chan string, 3)
	resultQueue := make(chan int, 3)

	wm := &httpworker.WorkerManager{
		Client:       nil,
		RequestQueue: requestQueue,
		ResultQueue:  resultQueue,
	}

	wm.PopulateJobs([]string{"sample_url_1", "sample_url_2", "sample_url_3"})

	if len(wm.RequestQueue) != 3 {
		t.Errorf("TestPopulateJobs failed, expected 3 created jobs, got %v.", len(wm.RequestQueue))
	}
}
