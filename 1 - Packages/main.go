package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main.go")
	auxiliar.Write()

	result := checkmail.ValidateFormat("gustavo.maddog@gmail.com")
	result2 := checkmail.ValidateFormat("temqueerrar")

	fmt.Println(result)
	fmt.Println(result2)
}
