package main

import "fmt"

func main() {
	var variable1 int = 10

	var variable2 int = variable1

	fmt.Println(variable2, variable1)

	variable1 += 7

	fmt.Println(variable2, variable1)

	// Pointers are memory references

	var variable3 int
	var pointer *int

	fmt.Println(variable3, pointer)

	variable3 = 100
	// pointer = variable3 -> throws an error since pointer is a pointer of int (*int), not an int

	pointer = &variable3

	fmt.Println(variable3, pointer)

	// Use & marker to reference the memory... use * before pointer var to dereference it

	fmt.Println(variable3, *pointer) // dereferenced pointer

	variable3 = 150

	fmt.Println(variable3, pointer) // Reference doesn't change even when value changed
}
