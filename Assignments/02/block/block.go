package block

import (
	"bytes"
	"encoding/binary"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/hash"
)

// GenesisDesc is the constant that marks that a block is a genesis block.
const GenesisDesc = "GenesisBlock Fri Aug 17 23:16:12 MDT 2018"

// BlockType is a single block in the block chain.
type BlockType struct {
	Index         int                // position of this block in the chain, 0, 1, 2, ...
	Desc          string             // if "genesis" string then this is a genesis block.
	ThisBlockHash hash.BlockHashType //
	PrevBlockHash hash.BlockHashType // This is 0 length if this is a "genesis" block
	Nonce         uint64             //
	Seal          hash.SealType      //

	// Add Merkle Hash to Block Later, (AS-03 will do this)
	// MerkleHash hash.BlockHashType

	// Add Transactions to Block Later, (AS-04 will do this)
}

// InitGenesisBlock create and return a genesis block.   This is a special
// block at the beginning of a chain.
func InitGenesisBlock() (gb *BlockType) {
	gb = &BlockType{
		Index:         0,
		Desc:          GenesisDesc,
		ThisBlockHash: []byte{},
		PrevBlockHash: []byte{},
		Nonce:         0,
		Seal:          []byte{},

		// Add fields in AS-03
		// MerkleHash:          []byte{},

		// Add fields in AS-04
	}
	gb.ThisBlockHash = hash.HashOf(SerializeBlock(gb))
	return
}

// InitBlock creates and returns a block.
func InitBlock(ii int, dd string, prev hash.BlockHashType) (bk *BlockType) {
	bk = &BlockType{
		Index:         ii,
		Desc:          dd,
		ThisBlockHash: []byte{},
		PrevBlockHash: prev,
		Nonce:         0,
		Seal:          []byte{},

		// Add fields in AS-03
		// MerkleHash:          []byte{},

		// Add fields in AS-04
	}
	bk.ThisBlockHash = hash.HashOf(SerializeBlock(bk))
	return
}

// IsGenesisBlock returns true if this is a genisis block.
func IsGenesisBlock(bk *BlockType) bool {
	if bk.Desc == GenesisDesc && len(bk.PrevBlockHash) == 0 {
		return true
	}
	return false
}

// SearalizeBlock searializes into bytes the fields that will be hashed the hash of the block.
// This is the hash that the next block will use to point to this block and the hash that
// this block will be saved as.
func SerializeBlock(bk *BlockType) []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, bk.Index)

	buf.Write([]byte(bk.Desc))

	// Additional fields will need to be added at this point. (AS-03)
	// MerkleHash:          []byte{},

	// Additional fields will need to be added at this point. (AS-04)

	return buf.Bytes()
}

// SearalizeForSeal searializes into bytes the fields that will be hashed for the mining seal.
func SerializeForSeal(bk *BlockType) []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, bk.Index)

	buf.Write([]byte(bk.Desc))

	buf.Write([]byte(bk.ThisBlockHash))

	buf.Write([]byte(bk.PrevBlockHash))

	binary.Write(&buf, binary.BigEndian, bk.Nonce)

	// Additional fields will need to be added at this point. (AS-03)
	// MerkleHash:          []byte{},

	// Additional fields will need to be added at this point. (AS-04)

	return buf.Bytes()
}
