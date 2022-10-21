package main

import "fmt"

// func main() {

// 	a := []int{1, 2, 3, 4}
// 	b := a
// 	b[1] = 3
// 	fmt.Println(a, b)

// 	fmt.Println("Length: ", len(a[:2]), "\nCapacity: ", cap(a)) // Slicing

// }

// func main() {

// 	a := make([]int, 10, 100) // To Predefine extra space
// 	a[0] = 4
// 	a = append(a, 20)
// 	fmt.Printf("%v %T", a, a)
// 	fmt.Println()
// 	fmt.Printf("Length: %v\n", len(a))   // Current Occupied Length
// 	fmt.Printf("Capacity: %v\n", cap(a)) // Total Capacity of the thing

// }

// func main() {
// 	a := []int{}
// 	a = append(a, 2, 3, 5) // Variadic Function
// 	a = append(a, 90)
// 	b := len(a)
// 	c := cap(a)

// 	fmt.Println(a, "---", b, c)

// }

func main() {
	a := []int{}
	a = append(a, 9, 8, 21, 24, 90, 22)

	b := append(a[:3], a[4:]...) // To Slice out an element
	// fmt.Println(a)
	fmt.Println(b)
}
