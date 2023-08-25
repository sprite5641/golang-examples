package main

import (
	"fmt"
	"sync"
)

var (
	counter     int
	counterLock sync.Mutex
	wg          sync.WaitGroup
)

func incrementCounter() {
	defer wg.Done()
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
}

func main() {
	const numGoroutines = 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter()
	}
	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter)
}
