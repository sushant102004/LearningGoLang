package main

import "fmt"

func printRoute(slice []string) {
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Mess", "College", "Lab"}
	printRoute(route)

	fmt.Println()
	route = append(route, "Hostel")
	printRoute(route)

	fmt.Println()
	smallRoute := route[0:2]
	printRoute(smallRoute)
}
