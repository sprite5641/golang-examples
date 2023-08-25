package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func incrementCounter() {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
}

func main() {
	const numGoroutines = 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter()
	}
	wg.Wait()
	fmt.Printf("Counter value: %d\n", atomic.LoadInt64(&counter))
}
