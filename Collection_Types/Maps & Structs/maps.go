package main

import (
	"fmt"
)

func main() {
	// Maps are unordered

	// state := make(map[string]int)  // Make Syntax
	state := map[string]int{ // LiteraL Syntax
		"Cal": 34,
		"Te":  32,

		"ui": 90,
		"hi": 300,
	}

	s := state // References to the location of the map elements
	fmt.Println(state["Cal"])
	delete(s, "ui")
	// pop, ok := state["ui"]

	// fmt.Println(pop, ok)
	state["Te"] = 890
	fmt.Println(s)
	fmt.Println(state)

}
