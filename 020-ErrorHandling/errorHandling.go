package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("temp.go")

	if err != nil {
		fmt.Println("Error Reading File: ", err)
	} else {
		fmt.Println(string(file))
	}
}
