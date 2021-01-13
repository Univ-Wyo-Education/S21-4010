package main

import (
	"encoding/json"
	"fmt"
)

var IVar int
var SVar string
var I64Var int64
var UIVar uint64

type Example17 struct {
	A int
	B string
}

var E17 Example17

// Add int64 and uint64 types
var SliceOfString []string
var MapOfString map[string]string
var MapOfBool map[string]bool

// init will initilize data before main() runs.  You can have more than one init() function.
func init() {
	SliceOfString = make([]string, 0, 10)
	MapOfString = make(map[string]string)
	MapOfBool = make(map[string]bool)
}

func main() {
	SliceOfString = append(SliceOfString, "AAA", "BBB")
	MapOfString["mark"] = "first"
	MapOfString["twain"] = "last"
	MapOfBool["mark"] = true
	MapOfBool["twain"] = false

	fmt.Printf("IVar = %d, type of IVar %T\n", IVar, IVar)
	fmt.Printf("IVar = %v, type of IVar %T\n", IVar, IVar)
	// add prints for your int64 and uint64 types
	fmt.Printf("SVar = %s, type of SVar %T\n", SVar, SVar)
	fmt.Printf("Address of SVar = %s, type of SVar %T\n", &SVar, &SVar)
	fmt.Printf("E17 = %s, type of E17 %T\n", &E17, &E17)
	fmt.Printf("    E17 = %+v, E17 as JSON: %s\n", &E17, IndentJSON(E17))
	// add prints for the other types above - so you can see them printed out.
	// use a %s and a %T for SliceOfString
	// use a %s and a %T for MapOfString
	// use a %#v and a %T for MapOfBool
	// Print out each of them with the IndentJSON function.
	fmt.Printf("SliceOfString = %s, type of SliceOfString %T\n", SliceOfString, SliceOfString) // xyzzy - answer
	fmt.Printf("MapOfString = %s, type of MapOfString %T\n", MapOfString, MapOfString)         // xyzzy - answer
	fmt.Printf("MapOfBool = %#v, type of MapOfBool %T\n", MapOfBool, MapOfBool)                // xyzzy - answer
	fmt.Printf("    SliceOfString: %s\n", IndentJSON(SliceOfString))                           // xyzzy - answer
	fmt.Printf("    MapOfString: %s\n", IndentJSON(MapOfString))                               // xyzzy - answer
	fmt.Printf("    MapOfBool: %s\n", IndentJSON(MapOfBool))                                   // xyzzy - answer
}

func IndentJSON(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}
