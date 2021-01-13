package merkle

import (
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/03/bsvr/hash"
)

func MerkleHash(data [][]byte) []byte {
	// Build a place to put the hashes for the leaves
	hLeaf := make([][]byte, 0, len(data))
	// Calculate Leaf Hashes
	for ii := range data {
		aHash := hash.HashOf(data[ii])
		hLeaf = append(hLeaf, aHash)
	}
	ln := (len(hLeaf) / 2) + (len(hLeaf) % 2)
	hMid := make([][]byte, 0, ln)
	for ln > 0 {
		for ii := 0; ii < len(hLeaf); ii += 2 {
			if ii+1 < len(hLeaf) {
				hT := hash.Keccak256(hLeaf[ii], hLeaf[ii+1])
				hMid = append(hMid, hT)
			} else {
				hT := hash.Keccak256(hLeaf[ii])
				hMid = append(hMid, hT)
			}
		}
		hLeaf = hMid
		ln = (len(hLeaf) / 2) + (len(hLeaf) % 2)
		hMid = make([][]byte, 0, ln)
		if len(hLeaf) == 1 {
			return hLeaf[0]
		}
	}
	return hLeaf[0]
}
