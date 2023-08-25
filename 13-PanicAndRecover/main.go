package main

import "fmt"

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	defer recoverFromPanic()
	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // This won't be executed
}
