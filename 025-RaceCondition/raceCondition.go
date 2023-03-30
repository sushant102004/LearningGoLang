/*
Objective: -
	Create 3 IIFE and add functionality to append a number into int silce.
	Use waitgroups with pointer to continue execution until all goroutines
	had done their work.

	Then check for race conditions.

	Remove race condition using Mutex lock and unlock.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initializing wait group with pointer.
	wg := &sync.WaitGroup{}

	// Initializing mutex using pointer to restrict access to resource and prevent race conditions
	mt := &sync.Mutex{}

	data := []int{0}

	// Here 3 means number of goroutines.
	wg.Add(3)

	go func(wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("Goroutine One")
		mt.Lock()
		data = append(data, 1)
		mt.Unlock()
		wg.Done()
	}(wg, mt)

	go func(wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("Goroutine Two")
		mt.Lock()
		data = append(data, 2)
		mt.Unlock()
		wg.Done()
	}(wg, mt)

	go func(waitGroup *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("Goroutine Three")
		mt.Lock()
		data = append(data, 3)
		mt.Unlock()
		waitGroup.Done()
	}(wg, mt)
	wg.Wait()

	fmt.Println(data)
}
