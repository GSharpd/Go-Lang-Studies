package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and slices")

	// Array is a list of values

	var arr1 [5]int
	fmt.Println(arr1) // All items with 0 value

	// Acces via index
	fmt.Println(arr1[0])

	arr2 := [5]string{"pos1", "pos2"}
	fmt.Println(arr2)

	// arr2[5] = "pos6" -> out of bounds error

	arr2[4] = "pos5"

	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5} // instances array with inferred number of positions, but can't go over positions after being created (arr3[5] = 6 -> out of bounds)
	fmt.Println(arr3)

	// slice is something of a more flexible array
	// no fixed size

	slc1 := []int{1, 2, 3}

	fmt.Println(slc1)

	fmt.Println(reflect.TypeOf(slc1))
	fmt.Println(reflect.TypeOf(arr3))

	// to add item to slice use append()

	slc1 = append(slc1, 4) // essentially recreates the slice with the new item

	fmt.Println(slc1)

	slc2 := arr3[1:3] // inclusive:exclusive gets the first of the range and excludes the last
	fmt.Println(slc2)

	arr3[1] = 3
	fmt.Println(slc2)
}
