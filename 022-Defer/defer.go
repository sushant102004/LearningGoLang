// Defers are used to call a function after a functions successfully complete
// its execution.

package main

import "fmt"

func printData() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("Printing...")
	defer printData()
	fmt.Println("Printing Done!")
}
