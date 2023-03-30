// Channels are used to communicate between goroutines.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Here 2 is used to create a buffered channel.
	/*
		Buffered channels allow us to send a number of values without having
		reciever for each of that value.
	*/
	myChan := make(chan int, 2)
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
	go func(myChan <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myChan

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChan, wg)

	// This is for sending data into channel.
	go func(myChan chan<- int, wg *sync.WaitGroup) {
		myChan <- 2
		wg.Done()
		close(myChan)
	}(myChan, wg)

	wg.Wait()
}
