package main

import (
	"fmt"
)

type bot interface {
	getgretting() string
}

type Englishbot struct{}
type Spanishbot struct{}

func main() {

	eb := Englishbot{}
	sb := Spanishbot{}

	printgretting(eb)
	printgretting(sb)

}

func (eb Englishbot) getgretting() string {

	return " Hello There"

}

func (sb Spanishbot) getgretting() string {

	return " Holla"

}

func printgretting(b bot) {
	fmt.Println(b.getgretting())

}
