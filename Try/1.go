// Make an array of 5 strings
// Make a slice of 5 int

// Custom input

package main

import (
	"fmt"
)

func main() {
	arr := [...]string{}

	for i := 0; i < 4; i++ {
		fmt.Print("Enter String ", i+1, " ")
		fmt.Scanln(&arr[i])
		fmt.Println(arr)
	}

}
