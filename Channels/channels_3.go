package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // Buffer channels, to get some time to proceess more data
	wg.Add(2)

	// go func(ch <-chan int) { // Recieve Only Channel (Only Recieves data from the channel )
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	i = <-ch
	// 	fmt.Println(i)

	// 	// ch <- 27
	// 	wg.Done()
	// }(ch)

	// ----------------------------------------------------------
	// FOR RANGE LOOP

	// --------------------------------------------------------------
	// go func(ch <-chan int) { // Recieve Only Channel (Only Recieves data from the channel )

	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()

	// 	// ch <- 27
	// }(ch)

	//---------------------------------------------------------------------
	// i OK SYNTAX LOOP

	go func(ch <-chan int) { // Recieve Only Channel (Only Recieves data from the channel)
		for {
			i, ok := <-ch // i has value, and ok has boolen if channel is open or not
			if ok == true {
				fmt.Println(i)
			} else {
				break
			}

			// Can also be written as

			// if i, ok := <-ch; ok {
			// 	fmt.Println(i)
			// } else {
			// 	break
			// }
		}
		wg.Done()
	}(ch)

	// --------------------------------------------------

	go func(ch chan<- int) { // Send Only Channel(Only Sends)
		ch <- 42
		ch <- 32 // Two senders
		// fmt.Println(<-ch)
		close(ch) // To let the receiver know that there are no more messages to send and to close the channel, otherwise it would result in a receiver side deadlock
		wg.Done()
	}(ch)

	wg.Wait()

}
