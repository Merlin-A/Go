package main

import (
	"fmt"
	"reflect"
)

// type Animal struct {
// 	Name string
// 	Type string
// }

// type Bird struct {
// 	canFly string
// 	Speed  int
// }

// // Bird has animal like charachteristics
// type Bird_emd struct {
// 	Animal // Embedding (Inheritance)
// 	canFly string
// 	Speed  int
// }

// func main() {
// 	fmt.Println("HI")
// 	A := Animal{
// 		"Tiger",
// 		"Cat",
// 	}

// 	B := Bird{
// 		"NO",
// 		43,
// 	}

// 	B_emd := Bird_emd{}
// 	B_emd.Name = "Hello"
// 	B_emd.Type = "Land"
// 	B_emd.canFly = "No"
// 	B_emd.Speed = 50

// 	B_emd_1 := Bird_emd{ // Another Way
// 		Animal: Animal{
// 			Name: "Emu",
// 			Type: "Land",
// 		},

// 		canFly: "No",
// 		Speed:  90,
// 	}
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(B_emd)
// 	fmt.Println(B_emd_1)

// }

// ----------------------------------------------------------//
// type Vehicle struct {
// 	reg_no int

// }

// type Car struct {
// 	Vehicle
// 	model  string
// 	car_no int
// }

// func main() {
// 	Brezza := Car{
// 		Vehicle: Vehicle{
// 			23,
// 			"Friday",
// 		},

// 		model:  "Brezza",
// 		car_no: 345,
// 	}

// 	fmt.Println(Brezza)
// }

// -----------------------------------------------------------------//

type Vehicle struct {
	reg_no int
	date   string `required max: 10` // Tags
}

type Car struct {
	Vehicle
	model  string
	car_no int
}

func main() {
	t := reflect.TypeOf(Vehicle{})
	field, _ := t.FieldByName("date")
	fmt.Println(field.Tag)

	// Brezza := Car{
	// 	Vehicle: Vehicle{
	// 		23,
	// 		"Friday",
	// 	},

	// 	model:  "Brezza",
	// 	car_no: 345,
	// }

	// fmt.Println(Brezza)
}
