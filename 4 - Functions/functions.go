package main

import "fmt"

func somar(n1 int8, n2 int8) int16 {
	return int16(n1 + n2)
}

var func1 = func(txt string) string {
	fmt.Println(txt)
	return txt
}

func multipleReturn(n1, n2 int8) (int, int) {
	sum := n1 + n2
	multiply := n1 * n2

	return int(sum), int(multiply)
}

func main() {
	soma := somar(10, 9)

	fmt.Println(soma)

	fmt.Println(somar(5, 10))

	result := func1("Funcs are types")
	fmt.Println(result)

	fmt.Println(multipleReturn(2, 5))

	result2, result3 := multipleReturn(3, 7)

	fmt.Println(result2, result3)

	result4, _ := multipleReturn(4, 3)
	fmt.Println(result4)
}
