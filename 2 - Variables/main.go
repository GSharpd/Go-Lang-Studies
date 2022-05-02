package main

import "fmt"

func main() {
	var string1 string = "1st string"
	string2 := "2nd string"

	fmt.Println(string1)
	fmt.Println(string2)

	var (
		var1 string = "sumthing"
		var2 string = "sumthing else"
	)

	fmt.Println(var1, var2)

	var3, var4 := "variable 3", "variable 4"

	fmt.Println(var3, var4)

	const const1 string = "first constant"
	fmt.Println(const1)
}
