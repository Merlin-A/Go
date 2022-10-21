package main

import (
	"fmt"
)

func main() {
	var day int = 900

	// ----------------------------------//
	switch day {
	case 1, 52, 100:
		fmt.Println("One")
	case day:
		fmt.Println("Ninety")
	default:
		fmt.Println("Something")
	}
	// ------------------------------------------//
	i := 10
	switch {
	case i <= 10:
		fmt.Println("One")
		// Break is self implied, unlike Python and C
		// fallthrough // Logic Less, will execute even if condition doesn't pass
	case i >= 9:
		fmt.Println("Fo")
	default:
		fmt.Println("Something")
	}

	//---------------------------------------//
	// ------------ Type Switching ---------------------//

	var test interface{} = [...]int{9, 3}

	switch test.(type) {
	case int:
		fmt.Println("Integer")

	case string:
		fmt.Println("String")

	case []int:
		fmt.Println("This is a Slice")

	default:
		fmt.Println("Something")
	}

	//-----------------------------------------//
}
