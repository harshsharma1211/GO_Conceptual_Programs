package main

import (
	"fmt"
)

func main() {

	langauges := make(map[string]string)

	langauges["JS"] = "Javascript"
	langauges["RB"] = "Ruby"
	langauges["PY"] = "Pyton"

	fmt.Println("These are the langauges", langauges)
	fmt.Println(langauges["JS"])

	delete(langauges, "JS")

	fmt.Println("These are the langauges", langauges)

	for key, value := range langauges {

		fmt.Println("value of key is : %v and Value of Language is : %v ", key, value)

	}

}
