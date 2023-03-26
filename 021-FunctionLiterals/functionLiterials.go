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

// Closures
// When we take value from outside then a copy of that value is created. It is closure same as JS

func calculatePrice(price float64, discountFunc func(price float64) float64) float64 {
	return price - (price * discountFunc(price))
}

func calculateDiscount(price float64) float64 {
	discount := 0.0
	if price < 100 {
		discount = 0.03
	} else {
		discount = 0.07
	}
	fmt.Println("Discount: ", discount)
	return discount
}

func main() {
	hello()

	// Passing printError() as argument. It will have access to errorCode
	customError(printError(), " 504")
	finalPrice := calculatePrice(20760.0, calculateDiscount)
	fmt.Println(finalPrice)
}
