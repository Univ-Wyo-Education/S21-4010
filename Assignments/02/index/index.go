package index

import (
	"fmt"
	"io/ioutil"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/lib"
)

/*

Example Data (sort of with stuff ...ed out)
	{
		"Index": [ "126c..", .... ]
		"TxHashIndex": {
			"1212121...": {
				"Index": 0,
				"BlockHash": "1212222",
			},
			"2212121...": {
				"Index": 1,
				"BlockHash": "1212222",
			}
		}
	}

*/

type TxHashIndexType struct {
	Index     int
	BlockHash string
}

type BlockIndex struct {
	Index       []string
	TxHashIndex map[string]TxHashIndexType
}

func ReadIndex() {
	// xyzzy
}

func WriteIndex(fn string, bkslice []*block.BlockType) {
	indexForBlocks := BuildIndex(bkslice)
	ioutil.WriteFile(fn, []byte(lib.SVarI(indexForBlocks)), 0644)
}

func BuildIndex(bkslice []*block.BlockType) (idx BlockIndex) {
	for _, bk := range bkslice {
		idx.Index = append(idx.Index, fmt.Sprintf("%x", bk.ThisBlockHash))
	}
	idx.TxHashIndex = make(map[string]TxHashIndexType)
	for ii, bk := range bkslice {
		idx.TxHashIndex[fmt.Sprintf("%x", bk.ThisBlockHash)] = TxHashIndexType{
			Index:     ii,
			BlockHash: fmt.Sprintf("%x", bk.ThisBlockHash),
		}
	}
	return
}
