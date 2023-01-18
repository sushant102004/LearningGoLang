package main

import "fmt"

type Counter struct {
	hits int
}

func increaseCounter(counter *Counter) {
	counter.hits = counter.hits + 1
	fmt.Println(counter.hits)
}

func main() {
	// var a int = 10
	// var ptr *int = &a

	// fmt.Println(a)
	// fmt.Println(&a)
	// fmt.Println(&ptr)

	// *ptr = *ptr + 1

	// fmt.Println(a)

	var counterOne = Counter{}
	increaseCounter(&counterOne)
	fmt.Println(counterOne.hits)
}
