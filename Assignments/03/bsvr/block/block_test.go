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
	exp := "3ba3cce9afa8549bd38878defc9cccac4b503756c6c6e40c8be78c849ce50a6e"
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
	testDataStr := "000000000000000c476f6f64204d6f726e696e6720343031302f3530313020636c61737301020304"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block, expected %s got %s", testDataStr, dataStr)
	}
}

//
func Test_SerializeForSeal(t *testing.T) {
	bk := InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	data := SerializeForSeal(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "000000000000000c476f6f64204d6f726e696e6720343031302f3530313020636c61737383e435c862836a963f52e06e463f827bcd44578e1b501e1264a8aef5a1b6c2fe010203040000000000000000"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block, expected %s got %s", testDataStr, dataStr)
	}
}
