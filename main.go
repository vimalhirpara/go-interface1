package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	width  float32
	length float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func (r rect) perim() float32 {
	return 2 * r.width * 2 * r.length
}

func (r rect) area() float32 {
	return r.width * r.length
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	myCircle := circle{radius: 5}
	myRect := rect{width: 30, length: 20}

	// fmt.Println(myCircle.area())
	// fmt.Println(myRect.area())

	measure(myCircle)
	measure(myRect)
}
