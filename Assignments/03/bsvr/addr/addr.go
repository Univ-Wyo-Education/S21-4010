package addr

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type AddressType []byte

func (ww AddressType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(ww) + `"`), nil
}

func (ww *AddressType) UnmarshalJSON(b []byte) error {
	// First, deserialize everything into a local data type that matches with data.
	var objMap string
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	// fmt.Printf("string ->%s<-, %s\n", objMap, godebug.LF())

	xx, err := hex.DecodeString(objMap)
	if err != nil {
		return err
	}
	// fmt.Printf("hex ->%x<-, %s\n", xx, godebug.LF())

	*ww = append(*ww, xx...)

	// fmt.Printf("got ->%x<-, %s\n", *ww, godebug.LF())
	return nil
}

// String is the insterface for fmt.String so that you can print an address as a "%s" and get human output.
func (aa AddressType) String() string {
	return "0x" + hex.EncodeToString(aa)
}

func ParseAddr(s string) (addr AddressType, err error) {
	if s[0:2] == "0x" {
		s = s[2:]
	}
	addr, err = hex.DecodeString(s)
	if err != nil {
		return []byte{}, err
	}
	return
}

func MustParseAddr(s string) AddressType {
	addr, err := ParseAddr(s)
	if err != nil {
		fmt.Printf("Error parsing address [%s] error [%s]\n", s, err)
		os.Exit(1)
	}
	return addr
}
