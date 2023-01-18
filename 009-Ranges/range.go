package main

import "fmt"

func main() {
	slice := []string{"Hello", "I", "am", "Sushant"}

	for i, element := range slice {
		fmt.Println("I: ", i, " Element: ", element)
		for _, ch := range element {
			fmt.Printf("%q", ch)
		}
	}
}
