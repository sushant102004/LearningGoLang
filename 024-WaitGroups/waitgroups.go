// WaitGroup don't let program to finish until all goroutines have finished.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s : %d\n", endpoint, res.StatusCode)
	}

	// Determining that work had been done and lock can be released.
	waitGroup.Done()
}

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
