package main

import (
	"fmt"
)

func main() {

	n := 1 == 1
	m := 2 == 4
	var z bool // Initial Value is False
	fmt.Printf("%v, %T", n, n)
	fmt.Printf("\n%v, %T", m, m)
	fmt.Printf("%v", z)
}
