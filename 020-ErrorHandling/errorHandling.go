package main

import (
	// "errors"
	"fmt"
	"os"
	"strconv"
)

type CustomError struct {
	message string
	code    int
}

func (c CustomError) Error() string {
	return c.message + " " + strconv.Itoa(c.code)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// return float64(0), errors.New("can't divide by zero")
		return 0.0, CustomError{"can't divide by zero", -1}
	} else {
		return float64(a) / float64(b), nil
	}
}

func main() {
	file, err := os.ReadFile("/mnt/07df2e6f-7d6b-4909-9ed8-b51af0bfa52c/GoLang/LearningGoLang/020-ErrorHandling/data.txt")

	if err != nil {
		fmt.Println("Error Reading File: ", err)
	} else {
		fmt.Println(string(file))
	}

	ans, error := divide(3, 0)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(ans)
	}
}
