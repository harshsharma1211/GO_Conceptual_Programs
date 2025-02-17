package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("Oranges")
		wg.Done()
	}()

	go func() {

		count("Apples")
		wg.Done()
	}()

	wg.Wait()

}
func count(thing string) {

	for i := 0; i < 4; i++ {

		fmt.Printf("Counting %s \n", thing)
		time.Sleep(time.Millisecond * 500)
	}

}
