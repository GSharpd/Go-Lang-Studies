package main

import (
	"fmt"
)

func sum(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}

	return
}

func writeSum(text string, numbers ...int) {
	var total int
	for _, number := range numbers {
		total += number
	}

	fmt.Println(text, total)
}

func main() {
	fmt.Println("Variatic (is this a word) functions")

	fmt.Println(sum(11, 13, 15, 20, 87, 65))

	writeSum("Total is:", 5, 7, 10, 87, 54, 32) // only 1 variatic param per func, must be last parameter
}
