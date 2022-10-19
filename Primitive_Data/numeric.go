package main

import (
	"fmt"
)

func main() {

	k := -42
	c := 2
	var j uint
	j = uint(k)
	fmt.Printf("%T", j)

	fmt.Println(k + int(j))
	fmt.Println(k / c)

	a := 10 //1010
	b := 2  //0010

	fmt.Println(a & b) // And of Binaries
	fmt.Println(a | b) // Or of Binaries

	fmt.Println(a ^ b)  // Ex-Or of Binaries
	fmt.Println(a &^ b) // NOR of Binaries

	s := 4              // 2 ^ 2
	fmt.Println(s << 3) // 2 ^ 2 * 2 ^ 3 = 32
	fmt.Println()
	fmt.Println(float64(s >> 3))

	var m complex128 = 1 + 2i //Complex Numbers, 128 -> 64 floating, 64 -> 32 floating point
	n := 2i

	fmt.Println(m + n)
	fmt.Println(m - n)
	fmt.Println(m * n)
	fmt.Println(m / n)
	fmt.Println(imag(n), ' ', real(n))
	fmt.Printf("%v,  %T", imag(n), real(n))

}
