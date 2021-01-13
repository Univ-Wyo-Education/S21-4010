package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
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

//type TxHashIndexType struct {
//	Index     int
//	BlockHash string
//}

type AddrHashIndexType struct {
	Addr addr.AddressType
	Data string
}

type TxWithValue struct {
	Addr  addr.AddressType // Address of Self
	Value []TxWithAValue   // List of Values in a set of blocks, may have more than one value per block.
}

type TxWithAValue struct {
	BlockIndex  int // Index of this block
	TxOffset    int // position of this Tx in the array of Tx in the block, this is in block.Tx[TxOffset]
	TxOutputPos int // positon of the output with a positive value in the transaction, block.Tx[TxOffset].Output[TxOutputPos]
}

type AddressIndex struct {
	AddrIndex map[string]AddrHashIndexType // Map of address to date (for S.C.)
}

type ValueIndex struct {
	AddrIndex map[string]TxWithValue // Map of address to list of Transaction with output value.
}

type BlockIndex struct {
	Index            []string       // List of block-hash
	BlockHashToIndex map[string]int // map from hash back to index
	AddrData         AddressIndex   // Contains map of addresses to data on this address
	FindValue        ValueIndex     // Locaitons of value
	// TxHashIndex      map[string]TxHashIndexType // Map of block-hash to transactions
}

func ReadIndex(fn string) (idx *BlockIndex, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	idx = &BlockIndex{}
	err = json.Unmarshal(buf, idx)
	if err != nil {
		err = fmt.Errorf("invalid initialization - Unable to parse JSON file, %s", err)
		return nil, err
	}
	return
}

//func WriteIndex(fn string, bkslice []*block.BlockType) {
//	indexForBlocks := BuildIndex(bkslice)
//	ioutil.WriteFile(fn, []byte(lib.SVarI(indexForBlocks)), 0644)
//}

func WriteIndex(fn string, indexForBlocks *BlockIndex) error {
	return ioutil.WriteFile(fn, []byte(lib.SVarI(indexForBlocks)), 0644)
}

func BuildIndex(bkslice []*block.BlockType) (idx BlockIndex) {
	for _, bk := range bkslice {
		idx.Index = append(idx.Index, fmt.Sprintf("%x", bk.ThisBlockHash))
	}
	//idx.TxHashIndex = make(map[string]TxHashIndexType)
	//for ii, bk := range bkslice {
	//	idx.TxHashIndex[fmt.Sprintf("%x", bk.ThisBlockHash)] = TxHashIndexType{
	//		Index:     ii,
	//		BlockHash: fmt.Sprintf("%x", bk.ThisBlockHash),
	//	}
	//}
	return
}
