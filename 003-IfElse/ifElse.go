package main

import "fmt"

func calculateAge(yearOfBirth int) int {
	return 2023 - yearOfBirth
}

func main() {
	var userRole string = "admin"

	if userRole == "admin" {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	if age := calculateAge(2004); age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}
}
