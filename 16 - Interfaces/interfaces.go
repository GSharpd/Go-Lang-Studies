package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	height, width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func writeArea(s shape) {
	fmt.Printf("The shapes area is: %0.2f\n", s.area())
}

func main() {
	fmt.Println("Interfaces")
	rec := rectangle{10, 15}
	writeArea(rec)

	circ := circle{10}
	writeArea(circ)

	// Interface implementation is implicit
	// if a struct implements the method signature defined in the interface, it implicitly implements the interface
}
