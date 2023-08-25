package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data := []byte("Hello, File I/O in Go!")

	// Writing to a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Reading from the file
	readData := make([]byte, len(data))
	_, err = file.Read(readData)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Read from file:", string(readData))

	// Removing the file and handling the potential error
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error removing file:", err)
	}
}
