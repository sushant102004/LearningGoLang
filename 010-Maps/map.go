package main

import "fmt"

func main() {
	cars := map[string]int{
		"BMW":      2009,
		"Mercedes": 2021,
		"Jeep":     2015,
	}

	for key, value := range cars {
		fmt.Println(key, " ", value)
	}

	fmt.Println()

	students := make(map[int]string)

	// Here 11212531 is key and "Sushant" is value
	students[11212531] = "Sushant"
	students[11212534] = "Kunal"

	delete(students, 11212534)

	for key, value := range students {
		fmt.Println(key, value)
	}

	_, found := students[11212534]
	if !found {
		fmt.Println("No student found")
	} else {
		fmt.Println("Student found")
	}

}
