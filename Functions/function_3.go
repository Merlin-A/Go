package main

import (
	"fmt"
)

func main() {
	var divide func(float64, float64) (float64, error)

	divide = func(num1, num2 float64) (float64, error) {

		if num2 == 0.0 {
			return 0.0, fmt.Errorf("Can't Divide By Zero")
		}

		return num1 / num2, nil

	}

	var a, b float64
	fmt.Print("Enter Input - ")
	fmt.Scan(&a, &b)

	d, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}
