package main

import "fmt"

func main() {
	// Simple For Loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// While Loop
	j := 0

	for j < 5 {
		fmt.Print(j, " ")
		j++
	}

	// Infinite Loop

	// for {
	// This will be infinite

	// if somethingHappend {
	// break
	// }
	// }

	fmt.Println()

	for k := 0; k < 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Print(k, " ")
	}
}
