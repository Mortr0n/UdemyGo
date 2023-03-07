package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func main() {
	s := square{10}
	tr := triangle{2, 5}
	printArea(s)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("The area of the object is : ", s.getArea())
}

func (sq square) getArea() float64 {
	a := sq.sideLength * sq.sideLength
	return a
}

func (t triangle) getArea() float64 {
	a := .5 * t.base * t.height
	return a
}
