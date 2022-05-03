package main

import "fmt"

func main() {
	fmt.Println("Control structures")

	nmbr := 10

	if nmbr > 15 {
		fmt.Println("Bigger than 15")
	} else {
		fmt.Println("Less or equal to 15")
	}

	// No if-else statements without brackets

	if nmbr2 := nmbr; nmbr2 > 0 && nmbr2 < 5 {
		fmt.Println("Bigger than zero")
	} else if nmbr2 > 5 {
		fmt.Println("Also bigger than five")
	}

	// nmbr2 only exists within condition scope
	// not accessible outside conditional statement

}
