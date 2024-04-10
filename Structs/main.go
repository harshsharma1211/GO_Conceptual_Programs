package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
	contactinfo
}

type contactinfo struct {
	email   string
	pincode int
}

func main() {

	jim := person{
		Firstname: "Harsh",
		Lastname:  "Sharma",
		contactinfo: contactinfo{
			email:   "jim@gmail.com",
			pincode: 20134,
		}}

	jim.print()

}
func (p person) updateName(newFirstName string) {
	p.Firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
