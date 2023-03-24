package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == "linux" {
		fmt.Println("This code will run on Linux")
	} else if runtime.GOOS == "windows" {
		fmt.Println("This section will run on Windows.")
	} else {
		fmt.Println("This program will exit now.")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Running main function.")
}
