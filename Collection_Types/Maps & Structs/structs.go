package main

import (
	"fmt"
)

type Doctor struct { // New Data type struct declared
	number  int
	name    string
	company []string
}

func main() {

	doctor := Doctor{ // Instantiating the struct type
		number: 90,

		company: []string{
			"Hi",
			"N",
			"E",
			"e",
		},
		name: "A",
	}

	another_Doctor := struct { // For Short Lived Structs
		name   string
		number int
	}{name: "Hi", number: 9}

	// another_Doctor := Doctor{  // Works the same but not recommended, misbehaves if order changed
	// 	91,
	// 	"Adi",
	// 	[]string{
	// 		"Hello",
	// 	},
	// }
	fmt.Println(doctor.name)
	fmt.Println(doctor)

	fmt.Println("Another Doctor - ", another_Doctor)

	a_doctor := another_Doctor // Value based, main value doesn't change based on reference
	a_doctor.name = "Tom"

	fmt.Println(a_doctor)

	a2_doctor := &a_doctor // pointer to that variable
	a2_doctor.name = "Hideous"
	fmt.Println("A_DOCTOR - ", a_doctor)
	fmt.Println("A2_DOCTOR - ", a2_doctor)

}
