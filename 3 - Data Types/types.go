package main

import (
	"errors"
	"fmt"
)

func main() {
	//Numbers
	var nmb1 int = 1000000
	// uint exists as an unsigned integer type, meaning it is necessarily a positive number
	// int32 alias = rune
	var nmb2 rune = 9999

	// uint8 alias = byte
	var nmb3 byte = 255

	fmt.Println(nmb1, nmb2, nmb3)

	//Real numbers
	var float1 float32 = 123.45

	var float2 float64 = 100000000.45

	fmt.Println(float1)
	fmt.Println(float2)

	//both inferred int and float with no explicit bits signature use the system bit architecture, meaning: in 64bit OS int is int64
	//e.g
	var float3 = 167598.94
	fmt.Println(float3)

	//String
	//go has no char type
	var str string = "text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	char1 := 'G'
	fmt.Println(char1)

	//Zero value - initial value of any data type | inferred vars cannot be zero value
	var txt1 string
	fmt.Println(txt1)

	//Boolean
	var bool1 bool
	fmt.Println(bool1)

	var bool2 bool = true
	fmt.Println(bool2)

	//Error | error is a type in golang
	var err1 error
	fmt.Println(err1)

	var err2 error = errors.New("Internal error")
	fmt.Println(err2)
}
