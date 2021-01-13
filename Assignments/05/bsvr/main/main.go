package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/cli"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/config"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
	"github.com/pschlump/godebug"
)

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

var ServerHostPort = flag.String("server", "", "http://127.0.0.1:9191 is a good example - if set then a server will be started.   Non server options are ignored.")
var Dir = flag.String("dir", "", "Direcotry to serve files from if server is enabled.")
var ServerCfgFlag = flag.String("scf", "{}", "Arbitrary server config data in JSON.")

func main() {
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

	if len(*Dir) > 1 && (*Dir)[0:1] == "." {
		tDir, _ := os.Getwd()
		tDir = filepath.Join(tDir, (*Dir)[:1])
		Dir = &tDir
	} else if *Dir == "." {
		tDir, _ := os.Getwd()
		Dir = &tDir
	}

	gCfg := config.GetGlobalConfig()
	cc := cli.NewCLI(gCfg)

	if *ServerHostPort != "" {
		cc.ReadGlobalData([]string{})
		var wg sync.WaitGroup
		wg.Add(1)
		// fmt.Printf(" Start Server: %s\n", godebug.LF())
		go func() {
			// fmt.Printf("1 Start Server: %s\n", godebug.LF())
			RunServer(cc, *ServerHostPort)
			wg.Done()
		}()
		wg.Wait()
	} else {
		// fmt.Printf("1 Start Server: %s\n", godebug.LF())
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
			fmt.Printf("1 Start Server: %s\n", godebug.LF())
			cc.Usage()
			os.Exit(1)
		}
	}

}

func RunServer(cc *cli.CLI, ServerHostPort string) {

	http.HandleFunc("/api/status", respHandlerStatus)
	http.HandleFunc("/api/shutdown", getRespHandlerShutdown(cc))

	http.HandleFunc("/api/acct-list", getRespHandlerAcctList(cc))
	http.HandleFunc("/api/acct-value", getRespHandlerAcctValue(cc))
	http.HandleFunc("/api/validate-signed-message", getRespHandlerValidateSignedMessage(cc))
	http.HandleFunc("/api/send-funds-to", getRespHandlerSendFundsTo(cc))
	/*
		TODO: Homework 6 ---------------------------------------------------------------------------

		4	/api/load-contract (A-06)
		5	/api/get-contract-source (A-06)
			/api/get-contract-data (A-06)
			/api/get-contract-private-data (A-06) (requires signed message of contract owner)
		6	/api/call-contract-method (A-06)
		7	/api/call-contract-method-with-funds (A-06)
		8	/api/call-contract-view-method (A-06)
		9	/api/get-contract-info (A-06) ( empty | name | name+owner | name+owner+hash(src) )
				//$$name$$ ContractName
				//Follow Go's lead, if name of functions starts with 'CAP' then public, else private.
				//Follow Go's lead, if name of data starts with 'CAP' then public, else private.

		Misc Networking ---------------------------------------------------------------------------

		3	/api/send-block	( bkn= or bkh= )
			/api/receive-block {post}

		TODO: Remember every request requires an Account/Pin to auth it.
			this will be done with ValidAcctPin(res,req) -> Bool
	*/
	http.Handle("/", http.FileServer(http.Dir(*Dir)))
	log.Fatal(http.ListenAndServe(ServerHostPort, nil))

}

type HandlerFunc func(www http.ResponseWriter, req *http.Request)

func respHandlerStatus(www http.ResponseWriter, req *http.Request) {
	q := req.RequestURI

	var rv string
	www.Header().Set("Content-Type", "application/json")
	rv = fmt.Sprintf(`{"status":"success","name":"go-server version 1.0.0","URI":%q,"req":%s, "response_header":%s}`,
		q, lib.SVarI(req), lib.SVarI(www.Header()))

	io.WriteString(www, rv)
}

func getRespHandlerShutdown(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		www.Header().Set("Content-Type", "application/json")
		if !ValidateAcctPin(cc, www, req) {
			return
		}
		signature, fail := GetParam(cc, www, req, "signature")
		if fail {
			return
		}
		msg, fail := GetParam(cc, www, req, "msg")
		if fail {
			return
		}
		gCfg := config.GetGlobalConfig()
		isValid, err := cc.ValidateSignature(gCfg.ControlAcct, signature, msg)
		if err != nil {
			return
		}
		if !isValid {
			return
		}
		fmt.Printf("Shutdown Now\n")
		fmt.Fprintf(os.Stderr, "Shutdown Now\n")
		os.Exit(0)
	}
}

func ValidateAcctPin(cc *cli.CLI, www http.ResponseWriter, req *http.Request) bool {
	gCfg := config.GetGlobalConfig()
	un, pw, havePw := req.BasicAuth()
	if !havePw {
		return false
	}
	if acct, ok := gCfg.AcctPin[pw]; ok {
		if strings.ToUpper(acct) == strings.ToUpper(un) {
			return true
		}
	}
	return false
}

func GetParam(cc *cli.CLI, www http.ResponseWriter, req *http.Request, name string) (string, bool) {
	// fmt.Printf("Request: %s, AT:%s\n", godebug.SVarI(req), godebug.LF())
	param1 := req.URL.Query().Get(name)
	// fmt.Printf("Param: name ->%s<- value ->%s<-, AT: %s\n", name, param1, godebug.LF())
	if param1 != "" {
		return param1, false
	}
	http.Error(www, "Required parameter "+name+" missing.", http.StatusInternalServerError)
	return "", true
}

func getRespHandlerAcctList(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		www.Header().Set("Content-Type", "application/json")
		if !ValidateAcctPin(cc, www, req) {
			return
		}
		rv := cc.ListAccountsJSON()
		www.Header().Set("Content-Type", "application/json")
		io.WriteString(www, rv)
	}
}

func getRespHandlerAcctValue(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		// fmt.Printf("AT: %s\n", godebug.LF())
		if !ValidateAcctPin(cc, www, req) {
			fmt.Printf("AT: %s\n", godebug.LF())
			return
		}
		www.Header().Set("Content-Type", "application/json")
		acct, fail := GetParam(cc, www, req, "acct")
		if fail {
			// fmt.Printf("AT: %s\n", godebug.LF())
			// fmt.Fprintf(www, "Missing Parameter %s\n", "acct")
			return
		}
		// fmt.Printf("AT: %s\n", godebug.LF())
		rv := cc.ShowBalanceJSON(acct)
		// fmt.Printf("rv = ->%s<- AT: %s\n", rv, godebug.LF())
		// fmt.Printf("AT: %s\n", godebug.LF())
		io.WriteString(www, rv)
	}
}

func getRespHandlerValidateSignedMessage(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		www.Header().Set("Content-Type", "application/json")
		if !ValidateAcctPin(cc, www, req) {
			return
		}
		acct, fail := GetParam(cc, www, req, "acct")
		if fail {
			return
		}
		signature, fail := GetParam(cc, www, req, "signature")
		if fail {
			return
		}
		msg, fail := GetParam(cc, www, req, "msg")
		if fail {
			return
		}
		isValid, err := cc.ValidateSignature(acct, signature, msg)
		if err != nil {
			fmt.Fprintf(www, "{\"status\":\"error\",\"msg\":%q}\n", err)
			return
		}
		if isValid {
			fmt.Fprintf(www, "{\"status\":\"success\",\"msg\":%q}\n", "Signature validated")
		} else {
			fmt.Fprintf(www, "{\"status\":\"fail\",\"msg\":%q}\n", "Signature is not valid")
		}
	}
}

func getRespHandlerSendFundsTo(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		www.Header().Set("Content-Type", "application/json")
		if !ValidateAcctPin(cc, www, req) {
			return
		}
		from, fail := GetParam(cc, www, req, "from")
		if fail {
			return
		}
		to, fail := GetParam(cc, www, req, "to")
		if fail {
			return
		}
		amount, fail := GetParam(cc, www, req, "amount")
		if fail {
			return
		}
		signature, fail := GetParam(cc, www, req, "signature")
		if fail {
			return
		}
		msg, fail := GetParam(cc, www, req, "msg")
		if fail {
			return
		}
		memo, fail := GetParam(cc, www, req, "memo")
		if fail {
			return
		}
		rv := cc.SendFundsJSON(from, to, amount, signature, msg, memo)
		io.WriteString(www, rv)
	}
}
