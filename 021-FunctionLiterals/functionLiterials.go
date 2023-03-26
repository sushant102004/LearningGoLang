// Function literals accept other functions as parameters and can also return function as parameter.
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

// customError will trim the code and set errorCode appropriately.
// It have fn function which we are calling at last to give errorCode result to another function.

func customError(fn func(code string), code string) {
	errorCode := strings.Trim(code, " ")
	errorMessage := ""

	if errorCode[0] == '4' {
		errorMessage = "Not Found"
	} else {
		errorMessage = "Internal Server Error"
	}

	// Giving control to other function.
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

	// Passing printError() as argument. It will have access to errorCode
	customError(printError(), " 504")
}
