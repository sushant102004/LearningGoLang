package main

import "fmt"

func main() {
	var fName string = "Sushant"
	var lName = "Dhiman"

	age := 19

	fmt.Println("Hello ", fName, " ", lName, ". You are ", age, " years old")

	// Ok Error Idiom

	partOne, other := 10, 20
	fmt.Println("Part One ", partOne, " ", " Other ", other)

	partTwo, other := 30, 40
	fmt.Println("Part Two ", partTwo, " ", " Other ", other)

	// Multiple Variables

	var (
		userName = "sushant102004"
		userId   = "11212531"
	)

	fmt.Println("Username ", userName)
	fmt.Println("UserID ", userId)

	// Ignoring values

	a, b, _ := 'A', 'B', 'C'

	fmt.Println(string(a))
	fmt.Println(string(b))

}
