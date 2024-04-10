package main

import "fmt"

func main() {

	defer fmt.Println("Hello again")
	defer fmt.Println("HS")

	fmt.Println("Hey")

}
func MyDefer() {

	for i := 0; i < 5; i++ {

		fmt.Println("Hi")
	}

}
