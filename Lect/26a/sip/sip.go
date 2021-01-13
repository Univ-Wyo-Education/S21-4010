package main

import (
	"fmt"
	"math/big"
)

// From: https://arrow.dit.ie/cgi/viewcontent.cgi?article=1031&context=itbj
// IdProtocalsInCrypto.pdf

func main() {
	// From p40
	p := big.NewInt(88667)
	q := big.NewInt(1031)
	alpha := big.NewInt(70322)
	a := big.NewInt(755)

	// v = (alpha ^ ( q-a )) % p
	t1 := big.NewInt(0)
	t1.Sub(q, a)
	v := big.NewInt(0)
	v.Exp(alpha, t1, p)

	fmt.Printf("Setup Complete: v=%s\n", v)

	// Alice is the Client:

	// ----------------------------------------------------------
	// Message 1 - Client to Server
	// ----------------------------------------------------------

	// Alice Chooses, and send to Bob
	r := big.NewInt(543) // Should be random, but for this example
	x := big.NewInt(0)
	x.Exp(alpha, r, p) // x=(alpha^r) % p

	fmt.Printf("Send To Bob : x=%s\n", x)

	// ------------------------------------------------------------------------
	// Response to Message 1, Server back to client
	// ------------------------------------------------------------------------

	// Bob is the Server:
	// Bob sends the challenge 'e' back to Alice e to do the computation
	e := big.NewInt(1000) // how chose (random?)

	// Alice now computes: y = a*e % q
	t2 := big.NewInt(0)
	t2.Mul(a, e)
	t2.Mod(t2, q) // 45664
	t2.Add(t2, r) // 851 is correct

	y := t2
	fmt.Printf("y=%s\n", y) // Prints 851

	// ------------------------------------------------------------------------
	// Message 2 - Client (Alice) with response to challenge.
	// ------------------------------------------------------------------------

	// Bob (server) verifies: x == z == (a^y) * (v^e) % p
	// Bob (server) verifies: x == z == ((a^y)%p)*((v^e)%p) % p

	t3 := big.NewInt(0)
	t3.Exp(alpha, y, p)
	t4 := big.NewInt(0)
	t4.Exp(v, e, p)
	t5 := big.NewInt(0)
	t5.Mul(t3, t4)
	t5.Mod(t5, p)

	z := t5
	fmt.Printf("z=%s\n", z)

	// ------------------------------------------------------------------------
	// Response 2 - Success/Fail message from server back to client
	// ------------------------------------------------------------------------
	if x.Cmp(z) == 0 {
		fmt.Printf("Authoized! Yea\n")
	} else {
		fmt.Printf("Nope nope nope\n")
	}

}
