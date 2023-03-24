package main

import "fmt"

func main() {
	name := "Sushant"
	age := 19

	fmt.Printf("%s is %d years old. \n", name, age)

	msg := "Hi, I'm learning goLang."

	result := fmt.Sprintln(msg)
	fmt.Println(result)

	// Format Indexing
	n1 := 1
	n2 := 2
	n3 := 3

	fmt.Printf("There are %[2]d oranges and %[1]d apple.\n", n1, n2, n3)

	// Format Conversion
	fmt.Printf("%f\n", 542.3)
	fmt.Printf("%e\n", 542.3)
	fmt.Printf("%c, %c, %c, %c \n", 'J', 'h', 'o', 'n')
}
