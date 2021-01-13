package main

// Your Name

import (
	"fmt"

	"github.com/pschlump/myPkg/test1" // import package you created
)

func main() {
	out := test1.DoubleValue(8)   // Call function in your package
	fmt.Printf("out = %d\n", out) // should print "out = 16"
	// add call to TripleValue at this point
}
