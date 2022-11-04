package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5)
	s := sum(2, 4)
	fmt.Println("The Sum is", *s)
}

// func sum(values ...int) float32 { // Take all the arguments given and wrap them in a slice named values
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values { // _ is the index & v is value
// 		result += v
// 	}
// 	return float32(result)

//		// fmt.Println("The sum is ", result)
//	}

// func sum(values ...int) (result int) { // Take all the arguments given and wrap them in a slice named values
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values { // _ is the index & v is value
// 		result += v
// 	}
// 	return

//		// fmt.Println("The sum is ", result)
//	}

func sum(values ...int) *int { // Take all the arguments given and wrap them in a slice named values
	fmt.Println(values)
	result := 0
	for _, v := range values { // _ is the index & v is value
		result += v
	}
	return &result

	// fmt.Println("The sum is ", result)
}
