package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // To Sync Go Routines

func main() {
	ch := make(chan int) // Declaring Channels of Type int
	wg.Add(2)

	// 2. This will receive 42
	go func() {
		i := <-ch // Receiving Data from th channel/ Channel Sending Data to variable

		fmt.Println(i)
		wg.Done()

	}()

	// 1. This will Send 42
	go func() {
		// var i *int
		// *i = 42
		// ch <- *i
		ch <- 42 // Sending Data to the channel/ Channel receiving external Data
		wg.Done()
	}()

	wg.Wait()
}
