package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Printf("a=%d\n", a)
		}(12 + i)
	}
}
