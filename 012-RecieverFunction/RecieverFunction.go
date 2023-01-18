/*
A Reciever function is a function that makes it easy to mutate a structure.
It modified function signature.
*/

package main

import "fmt"

type MyStruct struct {
	x int
	y int
}

type Coordinate struct {
	X, Y int
}

// Simple Function
func modifyNormal(x int, y int, myStruct *MyStruct) {
	myStruct.x += x
	myStruct.y += y
}

// Reciever Function (Pointer)

func (myStruct *MyStruct) modifyWithRec(x int, y int) {
	myStruct.x += x
	myStruct.y += y
}

// Reciever Function (Value)

func (start Coordinate) calculateDistance(end Coordinate) Coordinate {
	return Coordinate{start.X - end.X, start.Y - end.Y}
}

func main() {
	myStruct := MyStruct{x: 10, y: 10}

	// modifyNormal(1, 2, &myStruct)

	myStruct.modifyWithRec(3, 4)
	fmt.Println(myStruct)

	startCord := Coordinate{4, 2}
	endCord := Coordinate{5, 1}

	distance := startCord.calculateDistance(endCord)
	fmt.Println(distance)
}
