package main

import "fmt"

var n int

func init() {
	fmt.Println("Executing init() \n setting n to 10")
	n = 10
}

func main() {
	fmt.Println("Executing main()")
	fmt.Println(n)
	// init function runs before main function
	// there can be one init function per file
}
