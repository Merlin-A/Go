package main

import (
	"fmt"
	"math"
)

// Immutable but can be shadowed by Masking

const (
	xo = iota
	x1 = iota // IOTA acts as a counter
	x2 = iota
	x3 // Compiler Recognizes pattern, scoped to a constant block
)

const (
	xx = iota
)

func main() {
	fmt.Println("Hello")

	// var test int = 90
	// fmt.Println(test)
	// test = 89
	// fmt.Println(test)

	const constat int = 90
	fmt.Println(constat)
	sin := math.Sin(1.57)
	fmt.Println(sin + float64(constat))
	// const sin float64 = math.Sin(1.57) // Can't declare executables as constant
	// constat = 3 // Constants can't be redeclared

	const test_const = 4
	var b int16 = 28
	fmt.Print(test_const + b) // Implicit Conversions are possible in COnstant
	fmt.Println()
	fmt.Println("THis is package declaration", xo, x1, x2, x3, xx)

}

const (
	_ = iota // Disregards the first value
	cat
	dog
	horse
)

func main() {
	var special int = cat // Constant can be accessed in main package
	var speical int       // initial int value is 0
	fmt.Println(dog, special)
	// fmt.Println(special == cat)
	fmt.Print(speical == cat)
}

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Println(KB, MB, GB, TB)
	fileSize := 40000000000000
	fmt.Printf("%.2f GB", float64(fileSize/GB))
}

const (
	isAdmin = 1 << iota
	canSeeA
	canSeeD

	canSeeB
)

func main() {
	var roles byte = isAdmin | canSeeA | canSeeB
	fmt.Printf("%b", roles)
	fmt.Println()

	// And OPeration on BITS, bitmasking
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}
