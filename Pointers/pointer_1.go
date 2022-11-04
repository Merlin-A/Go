package main

import "fmt"

type mystruct struct {
	foo int
}

func main() {
	a := [3]int{2, 3, 4}

	b := &a[0] // stores address of a[0]
	c := &a[1]
	e := &a[2]

	fmt.Println(a, *b, c, e)

	var ms *mystruct
	ms = &mystruct{
		foo: 90,
	}
	fmt.Printf("%T %T %T %T", ms, ms.foo, (*ms).foo, *ms)
}
