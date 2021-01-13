package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			fmt.Printf("a=%d\n", a)
		}(12 + i)
	}
	wg.Wait()
}
