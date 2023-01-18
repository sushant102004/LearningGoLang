package main

import "fmt"

type Student struct {
	rollNo     int
	name       string
	isVerified bool
}

func main() {
	var sushant = Student{11212531, "Sushant Dhiman", true}
	fmt.Println(sushant.name)
}
