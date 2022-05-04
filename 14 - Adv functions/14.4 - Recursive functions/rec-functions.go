package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position // stopping condition
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Recursive Functions")

	// Basically functions that call themselves

	// 1 1 2 3 5 8 13

	position := uint(10)
	fmt.Println(fibonacci(position))

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}
