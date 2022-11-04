package main

import "fmt"

func main() {
	var arr [3]*int
	*arr[0] = 23
	*arr[1] = 34

	fmt.Print(&arr[0])

}
