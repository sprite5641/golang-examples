package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printCount(prefix string, count int) {
	defer wg.Done()
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	wg.Add(2)
	go printCount("Goroutine A", 5)
	go printCount("Goroutine B", 5)
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
