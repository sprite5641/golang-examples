package main

import "fmt"

func main() {
	defer fmt.Println("Deferred statement executed")
	fmt.Println("This will be printed first")
}
