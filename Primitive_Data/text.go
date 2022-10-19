package main

import (
	"fmt"
)

func main() {

	// STRING TYPE
	// Strings are immutable
	s := "This is a string"
	s2 := " This is also a string"
	fmt.Printf("%v,  %T", s[0], s[0])                 // Strings are aliases of bytes
	fmt.Printf("%v,  %T", string(s[0]), string(s[0])) // By typecasting
	fmt.Println("\n", s+s2)
	fmt.Println(string(s[0:])) // Slicing

	// Runes

	r := 'a' // Runes are just int32 bytes
	fmt.Println(rune(r))

}
