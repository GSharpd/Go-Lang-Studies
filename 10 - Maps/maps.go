package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Immutable structure to save data
	// key-value pair

	usr := map[string]string{
		"name":     "Mad",
		"lastName": "dog",
		"age":      "30",
	}

	fmt.Println(usr)

	// Does not work referencing individual properties, must acces the key to get the value
	// eg.

	fmt.Println(usr["age"])
	fmt.Println(usr["lastName"])

	// Can be nested

	usr2 := map[string]map[string]string{
		"completeName": {
			"first": "gustavo",
			"last":  "vieira",
		},
	}

	fmt.Println(usr2)

	delete(usr2, "completeName")

	fmt.Println(usr2)

	usr2["course"] = map[string]string{
		"campus": "one",
		"school": "law",
	}

	fmt.Println(usr2)
}
