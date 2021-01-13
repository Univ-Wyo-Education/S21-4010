package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/config"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/index"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/lib"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/mine"
)

type CLI struct {
	GCfg config.GlobalConfigData
}

func NewCLI(c config.GlobalConfigData) *CLI {
	return &CLI{
		GCfg: c,
	}
}

var AllBlocks []*block.BlockType

func (cc *CLI) CreateGenesis(args []string) {
	gb := block.InitGenesisBlock()
	os.MkdirAll(cc.GCfg.DataDir, 0755)
	fn := fmt.Sprintf("%x.json", gb.ThisBlockHash)
	fnPath := filepath.Join(cc.GCfg.DataDir, fn)
	if lib.Exists(fnPath) {
		fmt.Printf("Error: %s already exists - you will need to remove it if you want to re-create a new chain.\n", fnPath)
		return
	}
	mine.MineBlock(gb, cc.GCfg.MiningDifficulty)                // mine the block
	AllBlocks = append(AllBlocks, gb)                           // put in set of global blocks:
	fnIndexPath := filepath.Join(cc.GCfg.DataDir, "index.json") //
	index.WriteIndex(fnIndexPath, AllBlocks)                    // write index
	ioutil.WriteFile(fnPath, []byte(lib.SVarI(gb)), 0644)       //
}

func (cc *CLI) TestReadBlock(args []string) {
	fnPath := "./data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json"
	if !lib.Exists(fnPath) {
		fmt.Printf("You must run \"create-genesis\" first before this test.\n")
		os.Exit(1)
	}
	_, err := block.ReadBlock("./data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json")
	if err != nil {
		fmt.Printf("FAIL\n")
		os.Exit(1)
	}
	fmt.Printf("PASS\n")
}

func (cc *CLI) TestWriteBlock(args []string) {
	// Initialize a block
	bk := block.InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	// Write out that block
	block.WriteBlock("./data/TestWriteBlock.json", bk)
}
