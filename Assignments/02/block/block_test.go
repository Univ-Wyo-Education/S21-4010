package block

import (
	"fmt"
	"testing"
)

func Test_GenesisBlock(t *testing.T) {
	gb := InitGenesisBlock()
	if !IsGenesisBlock(gb) {
		t.Errorf("Should be genesis block")
	}
}

func Test_InitBlock(t *testing.T) {
	bk := InitBlock(12, "Hello World", []byte{1, 2, 3})
	if IsGenesisBlock(bk) {
		t.Errorf("Should not be genesis block")
	}
	if bk.Index != 12 {
		t.Errorf("Expected index of 12")
	}
	exp := "592fa743889fc7f92ac2a37bb1f5ba1daf2a5c84741ca0e0061d243a2e6707ba"
	got := fmt.Sprintf("%x", bk.ThisBlockHash)
	if exp != got {
		t.Errorf("Block hash incorrect, expected ->%s<- got ->%s<-", exp, got)
	}
}

//
func Test_SerializeBlock(t *testing.T) {
	bk := InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	data := SerializeBlock(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "476f6f64204d6f726e696e6720343031302f3530313020636c617373"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block, expected %s got %s", testDataStr, dataStr)
	}
}

//
func Test_SerializeForSeal(t *testing.T) {
	bk := InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	data := SerializeForSeal(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "476f6f64204d6f726e696e6720343031302f3530313020636c617373855dacd594fe8549500a64210b373222f9c06ecc2a25b3530334b404081166bc010203040000000000000000"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block, expected %s got %s", testDataStr, dataStr)
	}
}
