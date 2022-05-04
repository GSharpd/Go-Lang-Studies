package main

import "fmt"

func function1() {
	fmt.Println("Executing function 1")
}

func function2() {
	fmt.Println("Executing function 2")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("Score calculated, returning result...")
	fmt.Println("Checking if student is approved...")

	median := (n1 + n2) / 2

	if median >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Defer")

	defer function1()
	function2() // will execute function2() before 1

	// defer makes deferred function be executed in the last moment possible

	fmt.Println(studentApproved(8, 4))
}
