package lib

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pschlump/MiscLib"
	"github.com/pschlump/godebug"
)

// SVarI  marshals and indents a data structure into JSON.
func SVarI(v interface{}) string {
	// s, err := json.Marshal ( v )
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	}
	return string(s)
}

// Exists returns true if a file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Assert(assertion bool) {
	if !assertion {
		fmt.Fprintf(os.Stderr, "%sFatal: Failed Assertion AT: %s%s\n", MiscLib.ColorRed, godebug.LF(2), MiscLib.ColorReset)
		os.Exit(2)
	}
}
