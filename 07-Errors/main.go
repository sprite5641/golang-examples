package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, result chan<- int, errCh chan<- error) {
	defer wg.Done()

	if true {
		errCh <- fmt.Errorf("an example error")
		return
	}

	result <- 42
}

func main() {
	var wg sync.WaitGroup
	result := make(chan int, 1)
	errCh := make(chan error, 1)
	wg.Add(1)
	go worker(&wg, result, errCh)
	wg.Wait()
	select {
	case res := <-result:
		fmt.Println("Result:", res)
	case err := <-errCh:
		fmt.Println("Error:", err)
	}
}
