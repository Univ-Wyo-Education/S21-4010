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
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CheckPassword if ignore is false will use remote site to validate passwrod.
// ignore should be true for things like the old password when chaning
// passwrods (so you can upgrade a password).
//
// Documented: https://haveibeenpwned.com/API/v2#PwnedPasswords
// GET https://api.pwnedpasswords.com/range/{first 5 hash chars}
func CheckPassword(pw string, ignore bool) bool {
	// If an "old" password on a change - then should ignore this check!  FIXME -- TODO --
	if true || DbMap["ignore-password-check"] {
		return true
	}

	hash := fmt.Sprintf("%X", Sha1String(pw))
	if DbMap["db05"] {
		fmt.Printf("hash of ->%s<- [%s]\n", pw, hash)
	}
	status, data := DoGet("https://api.pwnedpasswords.com/range/" + hash[0:5])
	if DbMap["db05"] {
		fmt.Printf("status = %d, results...\n%s\n", status, data)
	}
	rest := strings.Split(data, "\n")
	for _, line := range rest {
		if strings.HasPrefix(line, hash[5:]) {
			if DbMap["db05"] {
				fmt.Printf("Match %s to %s\n", hash[5:], line)
			}
			return false
		}
	}
	return true // just return true, password is ok for now.
}

// Sha1String takes the sha1 has of a string and returns it.
func Sha1String(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	return hash
}

// DoGet performs a get request.
func DoGet(uri string, args ...string) (status int, rv string) {

	sep := "?"
	var qq bytes.Buffer
	qq.WriteString(uri)
	for ii := 0; ii < len(args); ii += 2 {
		// q = q + sep + name + "=" + value;
		qq.WriteString(sep)
		qq.WriteString(url.QueryEscape(args[ii]))
		qq.WriteString("=")
		if ii < len(args) {
			qq.WriteString(url.QueryEscape(args[ii+1]))
		}
		sep = "&"
	}
	url_q := qq.String()

	res, err := http.Get(url_q)
	if err != nil {
		return 500, ""
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 500, ""
		}
		status = res.StatusCode
		if status == 200 {
			rv = string(body)
		}
		return
	}
}
