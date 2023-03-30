// Channels are used to communicate between goroutines.

package main

import (
	"fmt"
	"sync"
)

func main() {
	myChan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	/*
		We can only send data to a channel if someone is listening to it.
		Otherwise we will get a deadlock error.

		That's why we are first creating a function that is listening to
		channel and printing data from it.

		Then we are creating a function to send data to the channel.
	*/

	// This is for recieving data into channel.
	go func(myChan chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChan)
		wg.Done()
	}(myChan, wg)

	// This is for sending data into channel.
	go func(myChan chan int, wg *sync.WaitGroup) {
		myChan <- 2
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
