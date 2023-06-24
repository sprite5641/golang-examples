package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(ch chan int, start, end int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		ch <- i
	}
}

func consumer(ch chan int) {
	defer wg.Done()
	for n := range ch {
		fmt.Println("Received", n)
	}
}

func main() {
	ch := make(chan int, 10)
	wg.Add(2)
	go producer(ch, 1, 5)
	go producer(ch, 6, 10)
	wg.Add(1)
	go consumer(ch)
	wg.Wait()
	close(ch)
	fmt.Println("All goroutines have finished.")
}
