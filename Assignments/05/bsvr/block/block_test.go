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
	exp := "a5939c6b3f2c404798db8d78bc9de6d2839e8129a9c7ce5ea0c8abe4a35b2eee"
	got := fmt.Sprintf("%x", bk.ThisBlockHash)
	if exp != got {
		t.Errorf("Block hash incorrect\nexpected: ->%s<-\nactual:   ->%s<-\n", exp, got)
	}
}

//
func Test_SerializeBlock(t *testing.T) {
	bk := InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	data := SerializeBlock(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "476f6f64204d6f726e696e6720343031302f3530313020636c61737301020304"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block\nexpected: ->%s<-\nactual:   ->%s<-\n", testDataStr, dataStr)
	}
}

//
func Test_SerializeForSeal(t *testing.T) {
	bk := InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	data := SerializeForSeal(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "476f6f64204d6f726e696e6720343031302f3530313020636c61737330647b743e55e60afab61e6d2e63da424bc11a1ccdd73295b9ea00825f4deda4010203040000000000000000"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block\nexpected: ->%s<-\nactual:   ->%s<-\n", testDataStr, dataStr)
	}
}
