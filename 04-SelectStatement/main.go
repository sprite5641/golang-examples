package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received on channel 1:", v)
		case v := <-ch2:
			fmt.Println("Received on channel 2:", v)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go producer(ch1, time.Second)
	go producer(ch2, 2*time.Second)
	go reader(ch1, ch2)
	time.Sleep(10 * time.Second)
}
