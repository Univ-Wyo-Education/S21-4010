package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// DoGet performs a HTTP GET request and returns a status and a body if successful.
// Any errors will be printed out on standard error.  You can pass multiple values
// to DoGet with:
//
// DoGet("http://localhost:3000/status","name","value","name1","value1")
//
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
		fmt.Fprintf(os.Stderr, "Error on GET of %s with parameters %s - Error:%s\n", uri, args, err)
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

func main() {
	status, body := DoGet("http://localhost:3000/status")

	// TODO: Test also with 'http://localhost:3000/double?value=100'
	// TODO: Test also with an invalid URL - so you don't get a status of 200.

	if status != 200 {
		fmt.Printf("Failed: status = %d\n", status)
	} else {
		fmt.Printf("Body ->%s<-\n", body)
	}
}
