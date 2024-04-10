package main

import (
	"fmt"
	"time"
)

func hello(name string, c chan string) {

	for i := 0; i < 5; i++ {

		msg := fmt.Sprintf("Hello  %s", name)

		c <- msg
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {

	c := make(chan string)

	go hello("Harsh", c)

	for msg := range c {

		fmt.Println(msg)

	}

}
