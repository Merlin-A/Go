package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) { // Recieve Only Channel (Only Recieves data from the channel )
		i := <-ch
		fmt.Println(i)
		// ch <- 27
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // Send Only Channel(Only Sends)
		ch <- 42
		// fmt.Println(<-ch)
		wg.Done()
	}(ch)

	wg.Wait()

}
