package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {

	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {

	return 0.5 * t.height * t.base
}

func printArea(s shape) {

	fmt.Println(s.getArea())
}

func main() {

	t := triangle{34, 56}
	s := square{3}

	printArea(t)
	printArea(s)

}
