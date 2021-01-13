package merkle

import (
	"fmt"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/hash"
)

// TODO - uncomment the imports when you build your code.
/*
import (
	"fmt"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/hash"
)
*/

func MerkleHash(data [][]byte) []byte {
	// return InstructorMerkleHash(data) // TODO: Replace this line with your code.

	// Build a place to put the hashes for the leaves
	hLeaf := make([][]byte, 0, len(data))
	// Calculate Leaf Hashes
	for ii := range data {
		aHash := hash.HashOf(data[ii])
		hLeaf = append(hLeaf, aHash)
	}
	// fmt.Printf("Leaf Hashes : %s, AT: %s\n", dumpSSB(hLeaf), godebug.LF())
	hMid := make([][]byte, 0, (len(hLeaf)/2)+1)
	ln := len(hLeaf)/2 + 1
	// fmt.Printf("ln=%d AT:%s\n", ln, godebug.LF())
	for ln >= 1 {
		// fmt.Printf("\n%s-------- TOP -------- AT:%s %s\n", MiscLib.ColorGreen, godebug.LF(), MiscLib.ColorReset)
		for ii := 0; ii < len(hLeaf); ii += 2 {
			// fmt.Printf("ii+1 = %d len(hLeaf) = %d AT:%s\n", ii+1, len(hLeaf), godebug.LF())
			if ii+1 < len(hLeaf) {
				hT := hash.Keccak256(hLeaf[ii], hLeaf[ii+1])
				hMid = append(hMid, hT)
				// fmt.Printf("AT:%s\n", godebug.LF())
			} else {
				hT := hash.Keccak256(hLeaf[ii])
				hMid = append(hMid, hT)
				// fmt.Printf("AT:%s\n", godebug.LF())
			}
		}
		hLeaf = hMid
		ln = len(hLeaf) / 2
		hMid = make([][]byte, 0, ln)
		// fmt.Printf("ln = %d Mid Hashes : %s, AT: %s\n", ln, dumpSSB(hLeaf), godebug.LF())
	}
	// fmt.Printf("\nFinal Hash : %s, AT: %s\n", dumpSSB(hLeaf), godebug.LF())
	return hLeaf[0]
}

func dumpSSB(x [][]byte) (s string) {
	s = "["
	com := ""
	for ii := range x {
		st := fmt.Sprintf("%x", x[ii])
		s += com + st
		com = ", "
	}
	s += "]"
	return
}
