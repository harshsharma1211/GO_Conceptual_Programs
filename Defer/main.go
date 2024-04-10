package main

import "fmt"

func main() {

	defer fmt.Println("Hello again")
	defer fmt.Println("\n HS")

	fmt.Println("Hey")
	MyDefer()

}
func MyDefer() {

	for i := 0; i < 5; i++ {

		fmt.Print(i)
	}

}
