package main

import (
	"fmt"
)

func main() {
	f := greeter{ // First instantiate struct
		greeting: "Hello",
		name:     "There",
	}

	f.greet() // Calls the method
	fmt.Println(f.name)

}

type greeter struct {
	greeting string
	name     string
}

// func  (object datatype) name_method{}
func (g *greeter) greet() {
	// 						Implicit Deferencing in GO
	fmt.Println(g.greeting, (*g).name) // Accesses struct objects via object
	g.name = "This"
}

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name) // Accesses struct objects via object
// 	g.name = "This"
// }
