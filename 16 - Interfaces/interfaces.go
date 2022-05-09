package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func WriteArea(s Shape) {
	fmt.Printf("The shapes area is: %0.2f\n", s.Area())
}

// func main() {
// 	fmt.Println("Interfaces")
// 	rec := Rectangle{10, 15}
// 	WriteArea(rec)

// 	circ := Circle{10}
// 	WriteArea(circ)

// 	Interface implementation is implicit
// 	if a struct implements the method signature defined in the interface, it implicitly implements the interface
// }
