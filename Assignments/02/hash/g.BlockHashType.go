package hash

import (
	"encoding/hex"
	"encoding/json"
)

type BlockHashType []byte // Hash of data in block -- Eventually a Merkle hash

func (ww BlockHashType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(ww) + `"`), nil
}

func (ww *BlockHashType) UnmarshalJSON(b []byte) error {
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
