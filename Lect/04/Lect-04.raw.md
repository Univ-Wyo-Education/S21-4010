m4_include(../../setup.m4)

More On Go 
====================

1. Janet Yellen offense stance.



What is a Hash
====================================================

A hash is a mapping from an input set to a unique number.
This is so that we have a high probability of a unique number for each different input.
The output will have a fixed length.


There are a number of different kind of hashes.

1. md5 a deprecated hash - it has weaknesses and should not be used.
2. sha1 some weaknesses.
3. sha256 
3. sha512 
3. sha3

hashes are used to validate passwords.  pbkdf and Ashley Madison.

Some hashes are specificly designed to be slow for passwords.

Some are build to be in hardware (sha1, sha256) or to not be in
hardware (Keccak used in Etherum).
 
Bitcoin uses sha256.

Example:

| Input             | Output                                                            |
|-------------------|-------------------------------------------------------------------|
| Hi                | c01a4cfa25cb895cdd0bb25181ba9c1622e93895a6de6f533a7299f70d6b0cfb  |
| Bitcoin           | deb10ca6fd85a5eba792ea8561da390635242f0c37c376f8eb7d7859adbffca9  |
| war-and-peace.txt | 6cbcc5ca5e590fb9ace161a5b93e4ecf280d7118104be0d63b686c004cfa70ae  |


Hashes for our purposes are unidirectional.  You can take the same input and get the same output, but you can't derive the input from the output.


```
// HashStrings hash a set of strings and return in hex-strings form
func HashStrings(a ...string) string {
	h := sha256.New()
	for _, z := range a {
		h.Write([]byte(z))
	}
	return fmt.Sprintf("%x", (h.Sum(nil)))
}
```

One of your first Assignmetns was to build a 'ksum' that works like sha256sum or md5sum and reads a file and calculates the hash for that file.

Once you have a hash you can prove that the original file has not changed.

This is the basis of lots of Blockchain metadata applications.



Blockchain and Mining
====================================================


What is Mining and How is it implemented.
-----

1. More on Go

Maps do not synchronize automatically.
So... Synchronization Primitives:

```Go

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}


```

### A Go Core/Panic 

First the Code

```Go
m4_include(Lect-04_00.go)
```

Then the bad output.


```
m4_include(Lect_04_02.txt)
```

### Pseudo Code for Mining (Homework 02)


```Go
m4_include(Lect-04_01.go)
```


