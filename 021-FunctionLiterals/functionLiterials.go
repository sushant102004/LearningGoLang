// These are also called closures in JS

package main

import (
	"fmt"
	"strings"
)

func hello() {
	fmt.Print("Hello,")
	world := func() {
		fmt.Println("World")
	}
	world()
}

func customError(fn func(code string), code string) {
	errorCode := strings.Trim(code, " ")
	errorMessage := ""

	if errorCode[0] == '4' {
		errorMessage = "Not Found"
	} else {
		errorMessage = "Internal Server Error"
	}
	fn(errorMessage)
}

func printError() func(errorCode string) {
	return func(errorCode string) {
		fmt.Printf("%.*s\n", len(errorCode), "---------------")
		fmt.Println(errorCode)
		fmt.Printf("%.*s\n", len(errorCode), "---------------")
	}
}

func main() {
	hello()
	customError(printError(), " 504")
}
