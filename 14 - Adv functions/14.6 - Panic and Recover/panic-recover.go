package main

import (
	"fmt"
)

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Recovering code execution!")
	}
}

func isApproved(n1, n2 float64) bool {
	defer recoverExecution()
	median := (n1 + n2) / 2

	if median > 6 {
		return true
	} else if median < 6 {
		return false
	}

	panic("Median is exactly 6!")
}

func main() {
	fmt.Println("Panic and Recover")

	fmt.Println(isApproved(6, 6))

	fmt.Println("Post execution!")
}
