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
		waitGroup.Add(1)
	}

	waitGroup.Wait()
}
