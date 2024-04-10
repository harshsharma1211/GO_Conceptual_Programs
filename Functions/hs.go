package main

import "fmt"

func adder(q, w int) int {

	return q + w

}
func main() {

	a := 23
	b := 47

	o := adder(a, b)

	fmt.Printf("Sum of a and b is %d", o)

}
