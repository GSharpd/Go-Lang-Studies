package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interfaces as generic types")

	generic("string")
	generic(20)
	generic(true)

	// println is a generic type, it works with many interfaces

	mapa := map[interface{}]interface{}{
		1:                      "string",
		float64(299321.299321): true,
		"coisa":                "outra coisa",
		true:                   "false",
	}

	fmt.Println(mapa)

	// Using like the example above is not recommended, ignoring type can result in confusing code
}
