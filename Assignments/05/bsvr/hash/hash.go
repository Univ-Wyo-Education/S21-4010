package hash

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Keccak256 use the Ethereum Keccak hasing fucntions to return a hash from a list of values.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// HashOfBlock calcualtes the hash of the 'data' and returns it.
func HashOf(data []byte) (h []byte) {
	h = Keccak256(data)
	return
}

// HashStringOf calcualtes the hash of the 'data' and returns it.
func HashStrngOf(data string) (h []byte) {
	h = Keccak256([]byte(data))
	return
}

// HashStringOfReturnHex calcualtes the hash of the 'data' and returns it.
func HashStrngOfReturnHex(data string) (s string) {
	h := Keccak256([]byte(data))
	s = hex.EncodeToString(h)
	return
}
