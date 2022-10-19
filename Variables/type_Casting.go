package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 30
	fmt.Println(i)
	var j string
	j = strconv.Itoa(i) // Type Casting
	fmt.Println(j)
}

// Interger - Float -> Type casting int(i)
// Float - Integer -> Data Loss
// Integer - String -> returns unicode charachter of that integer
