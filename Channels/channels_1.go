package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // To Sync Go Routines

func main() {
	ch := make(chan int) // Declaring Channels of Type int
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch // Receiving Data from th channel/ Channel Sending Data to variable

			fmt.Println(i)
			wg.Done()

		}()

		go func() {

			ch <- 42 // Sending Data to the channel/ Channel receiving external Data
			wg.Done()
		}()

	}

	wg.Wait()
}
