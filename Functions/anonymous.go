package main

import (
	"fmt"
)

func main() {

	// var f func() = func(){}
	f := func() {
		fmt.Println()
	}

	f()
	f()

	func() { //

		fmt.Println("Hello")

		for i := 0; i < 5; i++ {
			func(i int) {
				fmt.Println(i)
			}(i)
		}
	}()

	fmt.Printf("%T", f)

}
