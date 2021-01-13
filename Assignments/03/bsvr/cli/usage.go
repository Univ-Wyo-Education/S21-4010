package cli

import (
	"fmt"
	"os"
)

func (cc *CLI) Usage() {
	if len(os.Args) > 1 {
		fmt.Printf("got ->%s<-\n", os.Args[1])
	}
	fmt.Printf("Usage: bc02 [ --cfg cfg.json ] [ --create-genesis ] -- should have more help at this point -- TODO \n")
}
