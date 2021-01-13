package cli

import (
	"fmt"
	"os"
)

func (cc *CLI) Usage() {
	fmt.Printf("got ->%s<-\n", os.Args[1])
	fmt.Printf("Usage: bc02 [ --cfg cfg.json ] [ --create-genesis ] xyzzy\n")
}
