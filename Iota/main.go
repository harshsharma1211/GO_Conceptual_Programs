package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * 1)
	MB
	GB
	TB
	PB
	EB
)

func main() {

	fmt.Printf("%d and %b \n", KB, KB)
	fmt.Printf("%d and %b \n", MB, MB)
	fmt.Printf("%d and %b \n", GB, GB)
	fmt.Printf("%d and %b \n", TB, TB)
	fmt.Printf("%d and %b \n", PB, PB)
	fmt.Printf("%d and %b \n", EB, EB)
}
