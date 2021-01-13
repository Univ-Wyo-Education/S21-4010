package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pschlump/godebug"
)

// ReadBlock reads in from a file a block and returns it.
func ReadBlock(fn string) (bk *BlockType, err error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	bk = &BlockType{}
	err = json.Unmarshal(data, bk)
	if !IsGenesisBlock(bk) {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	return bk, nil
}

// WriteBlock read i;n a block from a file.
func WriteBlock(fn string, bk *BlockType) (err error) {

	data := IndentJSON(bk)

	err = ioutil.WriteFile(fn, []byte(data), 0644)

	return
}

// IndentJSON converts data to JSON format and returns it as a string.
func IndentJSON(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}
