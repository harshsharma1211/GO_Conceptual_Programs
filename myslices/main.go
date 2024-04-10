package main

import (
	"fmt"
)

func main() {

	var fruitList = []string{"apple", "Tomato,", "Peach"}

	fmt.Printf(" Tye of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	var course = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	index := 2
	fmt.Println(course)

	course = append(course[:index], course[index+1:]...)

	fmt.Println(course)
}
