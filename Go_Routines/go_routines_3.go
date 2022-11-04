package main

import (
	"fmt"
	"sync"
)

// 1. Initial Declared
var counter int = 0
var wg = sync.WaitGroup{}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)     // 2. Add 2 Go Routines on every iteration
		go sayHello() // 3. Run the Go routines
		go count()

	}

	wg.Wait()
	// Here our Go Routines are racing against each other to get the execution don
	// without any mutual synchronization

}

func sayHello() {
	fmt.Println("Hello", counter) // 4. Accessed by the go routine call
	wg.Done()
}

func count() {
	counter++
	wg.Done()
}
