package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	flag.Parse()

	fns := flag.Args()
	if len(fns) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: ./ksum [file ...]\n")
		os.Exit(1)
	}

	for _, fn := range fns {
		data, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read %s, error: %s\n", fn, err)
			continue
		}
		fmt.Printf("%s %x\n", fn, Keccak256(data))
	}
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
