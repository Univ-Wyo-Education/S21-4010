package lib

import (
	"encoding/json"
	"errors"
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

// Assert that 'assertion' is true.  If not exit program.
func Assert(assertion bool) {
	if !assertion {
		fmt.Fprintf(os.Stderr, "%sFatal: Failed Assertion AT: %s%s\n", MiscLib.ColorRed, godebug.LF(2), MiscLib.ColorReset)
		os.Exit(2)
	}
}

// Exists returns true if file exists. It will return true for directories also.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

var invalidMode = errors.New("Invalid Mode")

// Fopen opens the file name with mode having the same meaing as the C/C++ stdio.h Fopen.
func Fopen(fn string, mode string) (file *os.File, err error) {
	file = nil
	if mode == "r" {
		file, err = os.Open(fn) // For read access.
	} else if mode == "w" {
		file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	} else if mode == "a" {
		file, err = os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			file, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		}
	} else {
		err = invalidMode
	}
	return
}
