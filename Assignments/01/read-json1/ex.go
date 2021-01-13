package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
Use the following structure to read a JSON file.  Create a program that will read
JSON in, print it, write it out in a new JSON file.   Get the file name for the input
and output from the command line.

Add a map/dictionary field in the JSON input that is

```
	{
		"TxHash": "Your actual Name",
		"TxIn": 22,
		"TxOut": 44
	}
```

Example Run:

```SH
	./read-json1 --input in.json --output out.json
```

The struct to include in your code.

```Go
	type TransactionType struct {
		TxHash	string
		TxIn	int
		TxOut	int
	}
```

You will need to use:

1. `ioutil.WriteFile` [ioutil package](https://golang.org/pkg/io/ioutil/)
2. `IndentJSON` from above.

Submit:

1. Your program.
2. Your JSON input file.
3. Your JSON output file.

### References

1. [ioutil package](https://golang.org/pkg/io/ioutil/)
*/

type TransactionType struct {
	TxHash string
	TxIn   int
	TxOut  int
}

func main() {

	data, err := ioutil.ReadFile("data.json")
	_ = err

	var tt TransactionType

	err = json.Unmarshal(data, &tt)
	_ = err

	fmt.Printf("%+v\n", tt)

	tt.TxHash = "Colin"
	// tt.TxIn = 1000

	buf, err := json.Marshal(tt)
	_ = err

	fmt.Printf("Data: %s\n", buf)

}
