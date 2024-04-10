package main

import "fmt"

func main() {

	fmt.Println("Struct in Golang")

	Harsh := User{"Harsh", "harsh@gmail", true, 22}
	fmt.Println(Harsh)
	fmt.Printf(" My Name is : %v My email is : %v", Harsh.Name, Harsh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
