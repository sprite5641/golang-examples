package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan<- string, message string) {
	time.Sleep(time.Second)
	ch <- message
}

func main() {
	messages := make(chan string)
	go sendMessage(messages, "Hello")
	go sendMessage(messages, "World!")
	msg1 := <-messages
	msg2 := <-messages
	fmt.Println(msg1, msg2)
}
