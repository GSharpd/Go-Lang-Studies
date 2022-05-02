package main

import "fmt"

// Structs are GoLangs respective of classes
// Important to Object Oriented Programming

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number int16
}

func main() {
	separator := "------"
	fmt.Println("Structs file...")
	fmt.Println(separator)

	var usr user

	fmt.Println(usr)

	usr.name = "Mad"
	usr.age = 30

	fmt.Println(usr)

	// Type inference for structs
	usr2 := user{"Mad", 31, address{"Rua", 458}}
	fmt.Println(usr2)

	// You can't inform only some properties of a struct, otherwise compiler throws and error
	// eg. usr3 := user { 31 } -> Throws "too few values in struct literal"
	// Works if you inform the specific property with its property name -> user {age: 29}

	usr3 := user{age: 29}
	fmt.Println(usr3)

	addr1 := address{"rua random", 100}
	usr4 := user{"Nome", 40, addr1}

	fmt.Println(usr4)
}
