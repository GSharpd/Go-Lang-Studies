package main

import "fmt"

type person struct {
	name string
	age  int8
}

type student struct {
	person
	course string
	school string
}

func main() {

	prsn1 := person{name: "cleber", age: 20}

	fmt.Println(prsn1)

	s1 := student{prsn1, "prog", "fiap"}

	fmt.Println(s1)

	fmt.Println(s1.name)

	s2 := student{person{"cleito", 21}, "eng", "usp"}

	fmt.Println(s2.age, s2.course)
}
