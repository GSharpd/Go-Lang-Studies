package main

import "fmt"

func invertNumber(number int) (returnNumber int) {
	returnNumber = number * -1
	return
}

func invertNumberByPointer(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Functions and pointers")

	number := 20
	invertedNumber := invertNumber(number)

	fmt.Println(invertedNumber)
	fmt.Println(number)

	number2 := 20

	fmt.Println(number2)
	invertNumberByPointer(&number2)

	fmt.Println(number2)
}
