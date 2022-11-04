package main

import (
	"fmt"
)

func main() {

	thing := "ti"
	chicken := "de"
	ed := 34

	hello("Adia", "Wd", ed)
	hello_u(thing, chicken)

	hello_u_pointer(&thing, &chicken)

}

func hello(name string, caste string, number int) {
	fmt.Printf("\n%v  %v %v", name, caste, number)
}

func hello_u(name, caste string) {
	fmt.Printf("\n%v  %v", name, caste)
}

func hello_u_pointer(name, caste *string) {
	*name = "Ted"
	fmt.Println(*name, &caste)
}
