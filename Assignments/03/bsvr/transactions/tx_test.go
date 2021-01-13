package transactions

import (
	"fmt"
	"testing"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/03/bsvr/addr"
)

func Test_SerializeTransaction(t *testing.T) {
	tx := NewEmptyTx("Good Morning 4010/5010 class", addr.AddressType([]byte{4, 1, 3, 1, 8, 5}))
	data := SerializeTransaction(tx)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "040103010805476f6f64204d6f726e696e6720343031302f3530313020636c617373"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block, expected %s got %s", testDataStr, dataStr)
	}
}
