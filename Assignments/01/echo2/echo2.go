package main

// Your Name - it is important if you want to get credit for your assignment.
// Assignment 1.4 echo command line arguments and parse arguments.

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigData struct {
	Name  string
	Value string
}

func main() {
	var Cfg = flag.String("cfg", "cfg.json", "config file for this call")

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Usage: ./echo2 [-cfg cfg.json] [arg1 ...]\n")
		os.Exit(1)
	}

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	gCfg, err := ReadConfig(*Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	fmt.Printf("Congiguration: %+v\n", gCfg)

	for ii, ag := range fns {
		if ii < len(fns) {
			fmt.Printf("%s ", ag)
		} else {
			fmt.Printf("%s", ag)
		}
	}
	fmt.Printf("\n")
}

func ReadConfig(filename string) (rv ConfigData, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
	}
	err = json.Unmarshal(buf, &rv)
	if err != nil {
	}
	return
}
