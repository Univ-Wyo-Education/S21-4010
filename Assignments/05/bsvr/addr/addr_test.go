package addr

import (
	"fmt"
	"testing"
)

/*
// String is the insterface for fmt.String so that you can print an address as a "%s" and get human output.
func (aa AddressType) String() string {
	return "0x" + hex.EncodeToString(aa)
}
*/

func Test_FormatAsString(t *testing.T) {
	var aa AddressType
	aa = []byte{4, 1, 3, 1, 8, 5}
	ss := fmt.Sprintf("%s", aa)
	// fmt.Printf("%s\n", ss)
	if ss != "0x040103010805" {
		t.Errorf("Should have formatted, did not.")
	}
}
