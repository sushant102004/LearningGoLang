package main

import (
	"fmt"
	"time"
	"unicode"
)

func printSomething(msg string) {
	fmt.Println(msg)
}

func main() {
	go printSomething("Hello")
	go printSomething("I'm")
	go printSomething("Sushant")
	go printSomething("From")
	go printSomething("Panipat")

	var capitalized []rune

	data := []rune{'a', 'b', 'c', 'd'}

	capitalize := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
	}

	for i := 0; i < len(data); i++ {
		go capitalize(data[i])
	}

	fmt.Println(string(capitalized)) // Result is not ready yet.

	time.Sleep(100 * time.Millisecond)

	fmt.Println(string(capitalized)) // Code is ready after waiting 100 ms.
}
