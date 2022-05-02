package main

import "fmt"

func main() {
	separator := "-----------"
	// Arithmetics ( + - / * % )
	// Different types of int cannot be operated without being cast to the same type first

	var nmb1 int16 = 10
	var nmb2 int32 = 10

	// sum := nmb1 + nmb2 -> throws an error

	sum := int32(nmb1) + int32(nmb2)

	fmt.Println(sum)
	fmt.Println(separator)
	// End

	// Attibution ( = -> type declared | := -> type inferred )

	// Truthy operators
	// Bigger than
	fmt.Println(5 > 3)
	fmt.Println(5 >= 5)
	fmt.Println(5 < 3)
	fmt.Println(5 != 5)
	fmt.Println(5 == 5)
	fmt.Println(separator)
	// End

	// Logical Operators
	// AND && - OR || - Diff !

	// Unary operators
	// Add one ++ | Add specific value += | Works for subtraction aswell
	nmb3 := 10 // nmb3 = nmb3 + 10
	nmb3++
	fmt.Println(nmb3)
	nmb3 *= 2 // nmb3 = nmb3 * 2 - works for division and modulo
	fmt.Println(nmb3)
	fmt.Println(separator)
	// End

	// Ternary operators
	// Does not exist in golang
	// var something == true ? "It's true" : "It's false"
	// Do regular if checks instead

	var truth bool = true

	var text string
	if truth == true {
		text = "It's true"
	} else {
		text = "It's false"
	}

	fmt.Println(text)
	fmt.Println(separator)
}
