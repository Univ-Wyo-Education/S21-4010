package main

// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var errorOutput = io.MultiWriter(os.Stdout, os.Stderr)

func init() {
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		errorOutput = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			errorOutput = os.Stderr
		}
	}
}

// Fatalf formats a message to standard error and exits the program.
// On Linux/Mac the message is also printed to standard output if standard error
// is redirected to a different file.
func Fatalf(rc int, format string, args ...interface{}) {
	fmt.Fprintf(errorOutput, "Fatal: "+format+"\n", args...)
	fmt.Fprintf(errorOutput, " From: %s\n", LF(2))
	os.Exit(rc)
}
