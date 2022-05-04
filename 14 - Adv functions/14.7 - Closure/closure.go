package main

import "fmt"

func closure() func() {
	text := "Within closure()"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	fmt.Println("Closure")

	text := "Within main()"

	fmt.Println(text)

	// Functions that reference variables outside their scope
	// have a function return a function and declare the variables within the context of the first function, that way it always references that variable

	newFunc := closure()

	newFunc()
}
