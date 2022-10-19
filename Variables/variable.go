package main

import "fmt"

func main() {
	// var i int i = 40
	i := 40
	var j int = 90
	k := 99
	var z int
	z = i + j + k
	name := "Dina"
	fmt.Println("The variable is", z)
	fmt.Printf("The type of this variable is %T", z)
	fmt.Println("\n", name)

	var (
		roll_no int    = 20
		name1   string = "S"
		marks   int    = 90
		subject string = "Maths"
	)

	fmt.Print(name1, "  ", marks, subject, roll_no)

}

// LIfe span small, vairable mame slmall
// Long life, Long Name
