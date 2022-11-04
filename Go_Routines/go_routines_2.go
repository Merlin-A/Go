// Using Something Else instead of SLeep, as it restricts the actual working speed of the program

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // 1.Syncs Go routine together

// 2.Here we two go routines, one func & one main
// We need two sync these two go routines

func main() {
	wg.Add(1) // 3.We tell the Wait Group that we intend to add one go routine in this program
	msg_1 := "String"
	go func(some string) {
		fmt.Println(some)
		wg.Done() // 5. This Decrements the number of go routines by 1

	}(msg_1)

	msg_1 = "Another"
	wg.Wait() // 4. Wait Group will wait till the number of go routines is zero
}
