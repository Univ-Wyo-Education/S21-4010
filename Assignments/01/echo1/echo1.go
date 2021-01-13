package main

// Your Name - it is important if you want to get credit for your assignment.
// Assignment 1.3 echo command line arguments.

import (
	"fmt"
	"os"
)

func main() {
	ags := os.Args[1:]
	for ii, ag := range ags {
		if ii < len(ags)-1 {
			fmt.Printf("%s ", ag)
		} else {
			fmt.Printf("%s", ag)
		}
	}
	fmt.Printf("\n")
}
