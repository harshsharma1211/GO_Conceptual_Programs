package main

import "fmt"

func main() {

	fmt.Println("Struct in Golang")

	Harsh := User{"Harsh", "harsh@gmail", true, 22}
	fmt.Println(Harsh)
	fmt.Printf(" My Name is : %v My email is : %v", Harsh.Name, Harsh.Email)
	Harsh.GetStatus()
}

func (u User) GetStatus() { //methods
	fmt.Println(" Is user active :", u.Status)
}

func (u User) GetEmail() {

	fmt.Println("Email of User :", u.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
