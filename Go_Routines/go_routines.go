package main

import (
	"fmt"
	"time"
)

func main() {
	// go sayHello()                      // Forms a Go Routine
	// time.Sleep(100 * time.Millisecond) // Go Routine is generated and function ends before it can read SayHello  Function is held open by sleep to let the print be executed

	//---------------------------------------------------//

	msg := "Test"

	go func() {
		fmt.Println(msg)
	}()

	msg = "Night" // As the function print is called after go routine is executed & because of the Sleep time we provide, the msg value is overwritten
	time.Sleep(100 * time.Millisecond)

	// -----------------------------------------------//

	// This can be avoided by just passing the arguments in function

	// msg_1 := "String"
	// go func(some string) {
	// 	fmt.Println(some)

	// }(msg_1)

	// msg_1 = "Another"

}

func sayHello() {
	fmt.Println("Hello")
}
