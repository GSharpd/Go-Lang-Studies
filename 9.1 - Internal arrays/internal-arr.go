package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Internal arrays")

	slc1 := make([]float32, 10, 11)
	fmt.Println(slc1)

	fmt.Println(reflect.TypeOf(slc1)) // type
	fmt.Println(len(slc1))            // length
	fmt.Println(cap(slc1))            // capacity

	slc1 = append(slc1, 1)
	fmt.Println(slc1)
	fmt.Println(len(slc1))
	fmt.Println(cap(slc1))

	slc1 = append(slc1, 2)
	fmt.Println(len(slc1)) // exceeds previous capacity
	fmt.Println(cap(slc1)) // capacity is now twice the size of the length

	slc2 := make([]float32, 5)
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

	slc2 = append(slc2, 1)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

	// Slices reference an array created internally
}
