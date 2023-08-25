# Go Example Code

This repository contains example code written in Go (Golang) to demonstrate various concepts and techniques. Each example is self-contained and can be used as a starting point for your own Go projects.

## Table of Contents

1. [Installation](#installation)
2. [Examples](#examples)
   - [Example 1: Goroutines](#example-1-Goroutines)
   - [Example 2: Channels](#example-2-Channels)
   - [Example 3: WaitGroup](#example-3-WaitGroup)
   - [Example 4: SelectStatement](#example-4-SelectStatement)
   - [Example 5: BufferedChannels](#example-5-BufferedChannels)
   - [Example 6: Context](#example-6-Context)
   - [Example 7: Errors](#example-7-Errors)
   - [Example 8: WorkerPool](#example-8-WorkerPool)
   - [Example 9: Mutex and Shared State](#example-9-Mutex-and-Shared-State)
   - [Example 10: Atomic Operations](#example-10-Atomic-Operations)
   - [Example 11: Rate Limiting with time.Ticker](#example-11-Rate-Limiting-with-time.Ticker)
   - [Example 12: Defer](#example-12-Defer)
   - [Example 13: Panic and Recover](#example-13-Panic-and-Recover)
   - [Example 14: Custom Types and Methods](#example-14-Custom-Types-and-Methods)
   - [Example 15: JSON Marshaling and Unmarshaling](#example-15-JSON-Marshaling-and-Unmarshaling)
   - [Example 16: File I/O](#example-16-File-I/O)
   - [Example 17: Generics](#example-17-Generics)
3. [Contributing](#contributing)
4. [License](#license)

## Installation

To run the example code, make sure you have Go installed on your system. You can download and install Go from the official website: https://golang.org/

Once Go is installed, clone this repository to your local machine using the following command:

```
git clone https://github.com/sprite5641/golang-examples.git
```

## Examples

This repository contains the following examples:

### Example 1: Goroutines

This example demonstrates how to use goroutines to perform concurrent operations. It includes examples of creating goroutines, passing arguments to goroutines, and waiting for goroutines to finish.

### Example 2: Channels

This example demonstrates how to use channels to communicate between goroutines. It includes examples of creating channels, sending values to channels, and receiving values from channels.

### Example 3: WaitGroup

This example demonstrates how to use a WaitGroup to wait for goroutines to finish. It includes examples of creating a WaitGroup, adding goroutines to a WaitGroup, and waiting for goroutines to finish.

### Example 4: SelectStatement

This example demonstrates how to use a select statement to receive values from multiple channels. It includes examples of creating channels, sending values to channels, and receiving values from channels.

### Example 5: BufferedChannels

This example demonstrates how to use buffered channels to send and receive multiple values. It includes examples of creating buffered channels, sending values to buffered channels, and receiving values from buffered channels.

### Example 6: Context

This example demonstrates how to use a context to cancel goroutines. It includes examples of creating a context, passing a context to goroutines, and canceling goroutines.

### Example 7: Errors

This example demonstrates how to use errors to handle unexpected conditions. It includes examples of creating errors, checking for errors, and handling errors.

### Example 8: WorkerPool

This example demonstrates how to use a worker pool to perform concurrent operations. It includes examples of creating a worker pool, adding tasks to a worker pool, and waiting for tasks to finish.


### Example 9: Mutex and Shared State

This example demonstrates how to use a mutex to protect shared resources from concurrent access. It includes examples of creating a mutex, using the mutex to protect critical sections of code, and avoiding race conditions.

### Example 10: Atomic Operations

This example demonstrates how to use atomic operations to perform safe and efficient updates to shared variables without the need for locks.

### Example 11: Rate Limiting with time.Ticker

This example demonstrates how to use a ticker to limit the rate of execution of certain operations.

### Example 12: Defer

This example demonstrates how to use the defer statement to ensure that a function call is executed when the surrounding function exits, regardless of the control flow.

### Example 13: Panic and Recover

This example demonstrates how to use the panic and recover mechanisms to handle exceptional situations gracefully.

### Example 14: Custom Types and Methods

This example demonstrates how to define custom types and methods on them.

### Example 15: JSON Marshaling and Unmarshaling

This example demonstrates how to marshal Go structs into JSON and unmarshal JSON into Go structs.

### Example 16: File I/O

This example demonstrates how to read and write files in Go.

### Example 17: Generics

This example demonstrates how to use generic in Go.

## Contributing

If you would like to contribute to this project, you can follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use the code in this repository as a reference or starting point for your own Go projects.