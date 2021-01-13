package block

import (
	"bytes"
	"encoding/binary"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/hash"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/merkle"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/transactions"
)

// GenesisDesc is the constant that marks that a block is a genesis block.
const GenesisDesc = "GenesisBlock Fri Aug 17 23:16:12 MDT 2018"

// BlockType is a single block in the block chain.
type BlockType struct {
	Index         int                             // position of this block in the chain, 0, 1, 2, ...
	Desc          string                          // if "genesis" string then this is a genesis block.
	ThisBlockHash hash.BlockHashType              //
	PrevBlockHash hash.BlockHashType              // This is 0 length if this is a "genesis" block
	Nonce         uint64                          //
	Seal          hash.SealType                   //
	MerkleHash    hash.MerkleHashType             // AS-03
	Tx            []*transactions.TransactionType // Add Transactions to Block Later, (AS-04 will do this)
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
		MerkleHash:    []byte{},
		Tx:            make([]*transactions.TransactionType, 0, 1), // Add Transactions to Block Later, (AS-04 will do this)
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
		MerkleHash:    []byte{},
		Tx:            make([]*transactions.TransactionType, 0, 1), // Add Transactions to Block Later, (AS-04 will do this)
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
	buf.Write([]byte(bk.PrevBlockHash))
	if len(bk.Tx) > 0 {
		mData := make([][]byte, 0, len(bk.Tx))
		for _, tx := range bk.Tx {
			it := transactions.SerializeForSeal(tx)
			buf.Write(it)
			mData = append(mData, it)
		}
		bk.MerkleHash = merkle.MerkleHash(mData)
		buf.Write([]byte(bk.MerkleHash))
	}
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
	if len(bk.Tx) > 0 {
		mData := make([][]byte, 0, len(bk.Tx))
		for _, tx := range bk.Tx {
			it := transactions.SerializeForSeal(tx)
			buf.Write(it)
			mData = append(mData, it)
		}
		bk.MerkleHash = merkle.MerkleHash(mData)
		buf.Write([]byte(bk.MerkleHash))
	}
	return buf.Bytes()
}
