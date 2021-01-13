package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(from string) {
	wg.Add(1)
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %v\n", from, i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			fmt.Printf("%s\n", msg)
		}(fmt.Sprintf(" I am %d ", i))
	}

	wg.Wait()
	fmt.Printf("All Done\n")
}
