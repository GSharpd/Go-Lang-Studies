package main

import "fmt"

func calculate(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	fmt.Println("Named returns")

	sum, sub := calculate(3, 6)
	fmt.Println(sum, sub)

}
