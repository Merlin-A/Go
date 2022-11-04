package main

import (
	"fmt"
)

func main() {
	a := 42
	var b *int = &a // B is pointer to a

	// b is now pointer type & &a is pointer type

	// fmt.Printf("%T %T %T %T", &b, *b, &a, a)
	// fmt.Printf("%v %v %v %v", &b, *b, &a, a)

	/*
		{
			& - Reference Pointer to Integer
			* - Deferenc Pointer from a pointer to INteger (Deferencing Operator)
			* - Form a pointer to integer while declaring

		}

	*/

	*b = 23 // B refers to the address of a, so if we change b we change a

	fmt.Println(a + *b)

}
