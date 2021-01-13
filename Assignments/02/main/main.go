package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/cli"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/config"
)

func main() {

	os.Mkdir("./data", 0766)

	//
	//	TODO: AS-02 -- Add in the flags.
	//
	// Flags
	//	--cfg cfg.json
	//  --create-gensisi
	//  --test-read-block
	//  --test-write-block
	//
	// xyzzy - remove for class --------------------------------------------------------------------------------------------
	var Cfg = flag.String("cfg", "cfg.json", "config file for this call")
	var CreateGenesisFlag = flag.Bool("create-genesis", false, "init command")
	var TestReadBlockFlag = flag.Bool("test-read-block", false, "test read a block")
	var TestWriteBlockFlag = flag.Bool("test-write-block", false, "test write a block")

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	err := config.ReadConfig(*Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	// Commands
	//		create-genesis
	//		test-read-block
	//		test-write-write
	//  More to come

	gCfg := config.GetGlobalConfig()
	cc := cli.NewCLI(gCfg)

	if *CreateGenesisFlag {
		cc.CreateGenesis(fns)
	} else if *TestReadBlockFlag {
		cc.TestReadBlock(fns)
	} else if *TestWriteBlockFlag {
		cc.TestWriteBlock(fns)
	} else {
		cc.Usage()
		os.Exit(1)
	}

}
