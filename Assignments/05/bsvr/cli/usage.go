package cli

import (
	"fmt"
	"os"
)

func (cc *CLI) Usage() {
	if len(os.Args) > 1 {
		fmt.Printf("got ->%s<-\n", os.Args[1])
	}
	fmt.Printf(`Usage: %s [ --cfg cfg.json ] [ --create-genesis ] [ --test-read-block ] [ --test-send-funds ] [ --list-accounts ] [ --show-balance AccoutNumber ] [ --server http://Host:Port ]

		--cfg cfg.json          Configuration file
		--list-accounts			List all of the current accounts in the system
		--show-balance	Acct	Show the current value in a specifed account
		--server URL			Run as a server on the specified host/port.

`, os.Args[0])
}
