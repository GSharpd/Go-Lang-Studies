package main

import "fmt"

func main() {
	fmt.Println("Anonymous Functions")

	// Functions without names and executed right after their declaration

	func() {
		fmt.Println("This is an anonymous function")
	}()

	// Can have params and return objects

	func(text string) {
		fmt.Println(text)
	}("test")

	fmt.Println(func(texts ...string) string {
		var finalText string

		for _, text := range texts {
			finalText = finalText + text
		}

		return finalText
	}("this", " is a ", "test"))

	returnObj := func(text string, number int) string {
		return fmt.Sprintf("Received %s %d", text, number)
	}("parameter", 15)

	fmt.Println(returnObj)
}
