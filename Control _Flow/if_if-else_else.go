package main

import (
	"fmt"
	"math"
)

// JUST IF ELSE STATEMENTS

/*
if {
}else if {

}else{}
*/
func main() {
	state := map[string]int{ // LiteraL Syntax
		"Cal": 34,
		"Te":  32,

		"ui": 90,
		"hi": 300,
	}

	var test string
	fmt.Scan(&test)

	if pop, ok := state[test]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println("Not Present")
	}

	// -------------------------------------------------------//

	num := 9.9
	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("Different")
	}

	// -------------------------------------- //

}
