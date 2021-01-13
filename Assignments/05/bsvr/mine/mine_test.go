package mine

import (
	"encoding/hex"
	"testing"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/hash"
)

func TestMineBlock(t *testing.T) {

	tests := []struct {
		bk               block.BlockType
		expectedSealHash string
		expectedNonce    uint64
	}{
		{
			bk: block.BlockType{
				Index:         0,
				Desc:          block.GenesisDesc,
				ThisBlockHash: []byte{},
				PrevBlockHash: []byte{},
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "0000ae2cab130b4836988969f731c4f884ac4675790e5575a5161e5b96ab13d7",
			expectedNonce:    54586,
		},
		{
			bk: block.BlockType{
				Index:         1,
				Desc:          "My First Block",
				ThisBlockHash: []byte{},
				PrevBlockHash: []byte{},
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "0000adc29a80f1f0df08c8687c013d179050f5d1b449599e4d1437e4fad23525",
			expectedNonce:    46734,
		},
		{
			bk: block.BlockType{
				Index:         1,
				Desc:          "My First Block",
				ThisBlockHash: []byte{},
				PrevBlockHash: MustDecodeString("136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e"),
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "000026ac02359874bdfa606769a8f2a4e9d3d3de781214774128d604c67ffc46",
			expectedNonce:    143620,
		},
	}

	for ii, test := range tests {
		test.bk.ThisBlockHash = hash.HashOf(block.SerializeBlock(&test.bk))
		tests[ii] = test
	}

	for ii, test := range tests {
		bk := &test.bk
		MineBlock(bk, "0000")
		SealString := hex.EncodeToString(bk.Seal)
		if SealString != test.expectedSealHash {
			t.Errorf("Test %d, expected %s got %s\n", ii, test.expectedSealHash, SealString)
		}
		if bk.Nonce != test.expectedNonce {
			t.Errorf("Test %d, expected %d got %d\n", ii, test.expectedNonce, bk.Nonce)
		}
	}

}

func MustDecodeString(s string) []byte {
	rv, err := hex.DecodeString("136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e")
	if err != nil {
		panic(err)
	}
	return rv
}
