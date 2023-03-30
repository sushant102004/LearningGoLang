// WaitGroup don't let program to finish until all goroutines have finished.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex

// Will be storing result of request in this.
var result = []string{}

func main() {
	endpoints := []string{
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		// Adding a wait group
		waitGroup.Add(1)
	}

	// Waiting before finishing main function.
	waitGroup.Wait()
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	} else {
		// Locking the operation so that only 1 goroutine can access it.
		mutex.Lock()
		result = append(result, endpoint)
		// Releasing lock over the operation.
		mutex.Unlock()

		fmt.Printf("%s : %d\n", endpoint, res.StatusCode)
	}

	// Determining that work had been done and lock can be released.
	waitGroup.Done()
}
