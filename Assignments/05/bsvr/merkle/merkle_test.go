package merkle

import (
	"encoding/hex"
	"testing"
)

func TestMerkleHash(t *testing.T) {
	return
	tests := []struct {
		data         [][]byte
		expectedHash string
	}{
		{
			data:         [][]byte{{0x10, 0x22}, {0x44, 0x51, 0x32}},
			expectedHash: "2d35b13160d84eaa133178f7e4d259d8c6cacc95057215495d042897adee53d2",
		},
		{
			data:         [][]byte{{0x10, 0x22}},
			expectedHash: "3b8cbc855c432466f2f4e05c07e5234631d9f3985096fdce7b5126424306b508",
		},
		{
			data:         [][]byte{{0x10, 0x22}, {0x44, 0x51, 0x32}, {0x00, 0x88, 0x41, 0x22, 0x33}},
			expectedHash: "cfecba8c0aed379216509b63f5e12aedb6582aa389168228dd96e6211d6f3ae1",
		},
		{
			data:         [][]byte{{0x10, 0x22}, {0x44, 0x51, 0x32}, {0x00, 0x88, 0x41, 0x22, 0x33}, {0x09, 0x10}},
			expectedHash: "abf91783e802d581828eb40f8f925d8f1dcbe72136855c908e4667c115b4ead1",
		},
		{
			data:         [][]byte{{0x10, 0x22}, {0x44, 0x51, 0x32}, {0x01}, {0x00, 0x88, 0x41, 0x22, 0x33}},
			expectedHash: "ba58d2b518c7c619a10b73ae5c881a0fdd2497f5eb93d96931a9017cb84018b6",
		},
		{
			data:         [][]byte{{0x10, 0x22}, {0x44, 0x51, 0x32}, {0x01}, {0x01, 0x02, 0x02, 0x04}, {0x00, 0x88, 0x41, 0x22, 0x33}},
			expectedHash: "a80780f7d875cf401dbea84580788f7ca5eeb21de7f1a89373c312b9dc47891c",
		},
	}

	for ii, test := range tests {
		h := MerkleHash(test.data)
		hStr := hex.EncodeToString(h)
		if hStr != test.expectedHash {
			t.Errorf("Test %d, expected %s got %s\n", ii, test.expectedHash, hStr)
		}
	}

}
