package main

import "fmt"

// GenericSum adds two values of any type that supports addition.
func GenericSum[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	intSum := GenericSum(5, 10)
	fmt.Println("Integer Sum:", intSum)

	floatSum := GenericSum(3.14, 2.71)
	fmt.Println("Float Sum:", floatSum)

	stringConcat := GenericSum("Hello, ", "World!")
	fmt.Println("String Concatenation:", stringConcat)
}
