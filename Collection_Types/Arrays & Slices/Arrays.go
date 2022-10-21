package main

import "fmt"

// --------------------------------//
// Initial Arrays

// func main() {
// 	grades := [...]int{90}
// 	// grades[4] = 78
// 	fmt.Println(grades[0])
// 	// var arr [4] int {1, 2, 3, 4}
// 	var arr [4]string
// 	arr[0] = "1"
// 	arr[3] = "Test"
// 	arr[3] = "Ts"

// 	fmt.Println(arr)
// 	fmt.Println(len(arr))

// }

// --------------------------------//
// 2D Matrix

// func main() {
// 	var i_matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 0, 1}, [3]int{0, 1, 0}} // Redundant

// 	var matrix [3][3]int

// 	matrix[0] = [3]int{1, 0, 2}

// 	fmt.Println(i_matrix)
// 	fmt.Println(matrix)

// }

// --------------------------------//
// TAKING INPUT FROM USER

// func main() {
// 	var t [4]int
// 	for i := 0; i < 4; i++ {
// 		fmt.Printf("Enter %v Letter - ", i+1)
// 		fmt.Scanln(&t[i])
// 	}

// 	fmt.Println(t)
// }

// --------------------------------//

// --------------------------------//
// COPYING AN ARRAY

func main() {
	a := []int{1, 2, 3, 4}
	b := a //Copies the Entire Thing
	// b := &a // Points to the location of the Array Elements
	b[1] = 6
	fmt.Println(a)
	// fmt.Println(b)
}
