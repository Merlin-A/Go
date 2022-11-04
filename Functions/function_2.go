package main

import "fmt"

func main() {

	s, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The Answer is - ", s)

}
func divide(num1, num2 float32) (float32, error) {

	if num2 == 0.0 {
		return 0.0, fmt.Errorf("Can't Divide By Zero")
	}

	div := num1 / num2
	return div, nil

}
