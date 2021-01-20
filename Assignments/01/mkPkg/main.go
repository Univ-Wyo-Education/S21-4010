package main

// Your Name

import (
	"fmt"

	"github.com/Univ-Wyo-Education/S21-4010/Assignments/01/mkPkg/test1"
)

func main() {
	out := test1.DoubleValue(8)   // Call function in your package
	fmt.Printf("out = %d\n", out) // should print "out = 16"
	// add call to TripleValue at this point
}
