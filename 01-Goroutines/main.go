package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second * 2)
}
