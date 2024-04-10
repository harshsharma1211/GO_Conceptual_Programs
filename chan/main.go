package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.google.co.in/",
		"https://pkg.go.dev/",
		"https://www.linkedin.com/groups/4565934/",
		"https://www.google.com/maps",
	}

	for _, link := range links {
		Checklink(link)

	}

}

func Checklink(link string) {

	_, err := http.Get(link)
	if err != nil {

		fmt.Println(link, "Might be down")

	} //error handling

	fmt.Println(link, " is up")

}
