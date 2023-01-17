package main

import "fmt"

func main() {

	switch age := 3; {
	case age == 0:
		fmt.Println("New Born")
	case age <= 3:
		fmt.Println("Toodler")
	case age <= 10:
		fmt.Println("Child")
	}
}
