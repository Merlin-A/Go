package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 1. Initial Declared
var counter int = 0
var m = sync.RWMutex{}

/* Simple Mutex is locked or unlocked -> Only one can access the data, that too only when the mutex is unlocked
RW (Read WRite) Mutex  -> Many Can Read but only one can write, once the writing is ongoing, mutex is locked so no more reading
*/

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(100) // Assigns number of threads to work with in the program
	for i := 0; i < 10; i++ {
		wg.Add(2)     // 2. Add 2 Go Routines on every iteration
		m.RLock()     // 4. Locks that particular mutex after a single go routin is reading it
		go sayHello() // 3. Run the Go routines
		m.Lock()      // 5. Same as 4
		go count()

	}

	wg.Wait()

}

func sayHello() {
	fmt.Println("Hello", counter)
	m.RUnlock() // 6. Once the statement is exectued Unlocks for other go routines to access
	wg.Done()
}

func count() {
	counter++
	m.Unlock() // 7. Same as 6
	wg.Done()
}
