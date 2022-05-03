package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Increasing variable...")
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}

	fmt.Println("Final result: ", i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Increasing j...")
		time.Sleep(time.Millisecond)
		fmt.Println("j value: ", j)
	}

	// iterating through lists - range
	// default is index and then the value of the item in the list
	names := [3]string{"mad", "dog", "foo"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	// iterating a string
	for index, char := range "word" {
		fmt.Println(index, string(char)) // use string() otherwise go returns the asc code for each character
	}

	usr := map[string]string{
		"id":   "nick",
		"name": "name",
	}

	// iterate through maps
	for key, value := range usr {
		fmt.Println(key, value)
	}

	// range does not work on structs
	// use for with no condition to loop indefinitely
}
