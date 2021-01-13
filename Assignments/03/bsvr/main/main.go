package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/03/bsvr/cli"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/03/bsvr/config"
)

func main() {
	//
	// Flags
	//	--cfg cfg.json
	//  --create-gensisi
	//  --test-read-block
	//  --test-write-block
	//  --test-send-funds From To Amount
	//
	var Cfg = flag.String("cfg", "cfg.json", "config file for this call")
	var CreateGenesisFlag = flag.Bool("create-genesis", false, "init command")
	var TestReadBlockFlag = flag.Bool("test-read-block", false, "test read a block")
	var TestWriteBlockFlag = flag.Bool("test-write-block", false, "test write a block")
	var TestSendFunds = flag.Bool("test-send-funds", false, "test sending funds from one account to another")
	var ShowBalance = flag.Bool("show-balance", false, "Show the balance on an account")
	var ListAccounts = flag.Bool("list-accounts", false, "List the addresses of known accounts")

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

	gCfg := config.GetGlobalConfig()
	cc := cli.NewCLI(gCfg)

	// xyzzy - TODO - read index
	// xyzzy - TODO - read blocks

	if *CreateGenesisFlag {
		cc.CreateGenesis(fns)
	} else if *TestReadBlockFlag {
		cc.TestReadBlock(fns)
	} else if *TestWriteBlockFlag {
		cc.TestWriteBlock(fns)
	} else if *TestSendFunds {
		cc.TestSendFunds(fns)
	} else if *ShowBalance {
		cc.ShowBalance(fns)
	} else if *ListAccounts {
		cc.ListAccounts(fns)
	} else {
		cc.Usage()
		os.Exit(1)
	}

}

/*
Tx - initialization - create initial values in accounts.

	func IsAccount ( acct ) -> bool
	func AccountInitialValue ( bk, acct, value ) // create Tx with value in genesis block
		if IsAcct(acct)
		if bk IsGenesis
			CreateFunds(bk,acct,value)

	Verify testing with this....

*/
