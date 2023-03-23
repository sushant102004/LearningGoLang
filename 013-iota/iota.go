package main

import "fmt"

type Direction byte

func main() {
	const (
		East Direction = iota
		West
		North
		South
	)

	fmt.Println(West)

	east := East
	fmt.Println(east)
}
