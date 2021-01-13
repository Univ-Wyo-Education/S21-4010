package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/config"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/hash"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/index"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/mine"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/transactions"
	"github.com/pschlump/MiscLib"
	"github.com/pschlump/godebug"
)

type CLI struct {
	GCfg       config.GlobalConfigData
	AllBlocks  []*block.BlockType
	BlockIndex index.BlockIndex
}

// NewCLI  returns a new command line config.
func NewCLI(c config.GlobalConfigData) *CLI {
	return &CLI{
		GCfg: c,
	}
}

// BuildIndexFileName returns the name of the index.json file
// with the correct path from the configuration.
func (cc *CLI) BuildIndexFileName() (fnIndexPath string) {
	fnIndexPath = filepath.Join(cc.GCfg.DataDir, "index.json") //
	return
}

// BuildBlockFileName takes a hashStr that is the name of the JSON
// file withouth the path and `.json` and combines to make a full
// file name.
func (cc *CLI) BuildBlockFileName(hashStr string) (fnBlockPath string) {
	fnBlockPath = filepath.Join(cc.GCfg.DataDir, hashStr+".json") //
	return
}

// CreateGenesis creates and writes out the genesis block and the
// initial index.json files.  Theis is the ""genesis"" of the
// blockchain.
func (cc *CLI) CreateGenesis(args []string) {
	gb := block.InitGenesisBlock()
	os.MkdirAll(cc.GCfg.DataDir, 0755)

	fnIndexPath := cc.BuildIndexFileName()
	if lib.Exists(fnIndexPath) {
		fmt.Fprintf(os.Stderr,
			"Error: %s already exists - you will need to remove it if you"+
				" want to re-create a new chain.\n", fnIndexPath)
		os.Exit(1)
		return
	}

	cc.BlockIndex = index.BuildIndex(cc.AllBlocks) // Build an initial index.

	for _, act := range cc.GCfg.InitialAccounts {
		out := transactions.TxOutputType{
			Account:     act.Acct,
			Amount:      act.Value,
			TxOutputPos: 0, //  Position of the output in this block. In the
			// 					block[this].Tx[TxOffset].Output[TxOutptuPos]
			// TxOffset  - will be set by AppendTxToBlock
		}
		tx := &transactions.TransactionType{
			Output:  []transactions.TxOutputType{out},
			Account: cc.GCfg.AcctCoinbase,
			Comment: "Initial Balance",
		}
		cc.AppendTxToBlock(gb, tx)
	}

	gb.ThisBlockHash = hash.HashOf(block.SerializeBlock(gb))
	fn := fmt.Sprintf("%x", gb.ThisBlockHash)
	fnPath := cc.BuildBlockFileName(fn)
	if lib.Exists(fnPath) {
		fmt.Fprintf(os.Stderr, "Error: %s already exists - you will need to remove it "+
			"if you want to re-create a new chain.\n", fnPath)
		os.Exit(1)
		return
	}

	cc.AppendBlock(gb) // Append block to list of blocks, set the index postion.  Write block and index.

}

// TestReadBlock Test code for command line.
func (cc *CLI) TestReadBlock(args []string) {
	fnPath := cc.BuildBlockFileName("29e68e530d0718dd01759e8c9a5276d20687bc5ec23e60dce063c2b97ba6b04f")
	if !lib.Exists(fnPath) {
		fmt.Printf("You must run \"create-genesis\" first before this test.\n")
		os.Exit(1)
	}
	_, err := block.ReadBlock(fnPath)
	if err != nil {
		fmt.Printf("FAIL\n")
		os.Exit(1)
	}
	fmt.Printf("PASS\n")
}

// TestWriteBlock Test code for command line.
func (cc *CLI) TestWriteBlock(args []string) {
	// Initialize a block
	bk := block.InitBlock(12, "Good Morning 4010/5010 class", []byte{1, 2, 3, 4})
	// Write out that block
	err := block.WriteBlock(cc.BuildBlockFileName("TestWriteBlock"), bk)
	if err != nil {
		fmt.Printf("FAIL\n")
		os.Exit(1)
	}
	fmt.Printf("PASS\n")
}

// TestSendFunds will process command argumetns and walk through the transaction process
// of sending funds once.  This is essentially the transaction process - but driven from
// the command line.
func (cc *CLI) TestSendFunds(args []string) {

	// In Assignment 5: args should be 6 - FromAddress, ToAddress, AmountToSend, Signature, Msg, Memo
	//			Last 4 args can just be 'x' for the moment - placeholders - not checked - not used.
	if len(args) != 7 {
		fmt.Fprintf(os.Stderr, "Should have 6 argumetns after the flag.  FromAddress, ToAddress, AmountToSend, x, x, Memo\n")
		os.Exit(1)
		return
	}

	// -----------------------------------------------------------------------------
	// Read in index and all of the blocks.
	// -----------------------------------------------------------------------------
	cc.ReadGlobalConfig()

	// fmt.Printf("%sSUCCESS #1 - Read in the index and the blocks!%s\n", MiscLib.ColorGreen, MiscLib.ColorReset)
	// fmt.Printf("%s        #1 - AT: %s%s\n", MiscLib.ColorGreen, godebug.LF(), MiscLib.ColorReset)

	// -----------------------------------------------------------------------------
	// Do the send funds stuff
	// -----------------------------------------------------------------------------
	from := addr.MustParseAddr(args[0])
	to := addr.MustParseAddr(args[1])
	amt64, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid parse of amout [%s] error [%s]\n", args[2], err)
		os.Exit(1)
	}
	amount := int(amt64) // type cast from int64 to int
	if amount <= 0 {     // Validate that no negative amount is used.
		fmt.Fprintf(os.Stderr, "Amount is out of range - can not send 0 or negative amounts [%d]\n", amount)
		os.Exit(1)
	}

	bk := cc.NewEmptyBlock()
	lib.Assert(bk.Index == len(cc.AllBlocks))

	tx, err := cc.SendFundsTransaction(from, lib.SignatureType(args[3]), args[4], to, amount, args[5])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to transfer error [%s]\n", err)
		os.Exit(1)
		return
	}
	lib.Assert(tx != nil)
	cc.AppendTxToBlock(bk, tx)

	// -----------------------------------------------------------------------------
	// Write out updated index and new block at end.
	// -----------------------------------------------------------------------------
	cc.AppendBlock(bk)

}

// SendFundsTransaction is the core chunk of moving funds from one account to
// another.
//
// This is your homework.  Finish out this function and test.
//
func (cc *CLI) SendFundsTransaction(
	from addr.AddressType, // account to transfer from
	sig lib.SignatureType, // not used yet - ditital signature - Assignment 5
	message string, //        not used yet - JSON message - Assignment 5
	to addr.AddressType, //   account to send funds to
	amount int, //            Amount of funds to send
	memo string, //           Memo to add to transaction (Comment)
) (
	tx *transactions.TransactionType,
	err error,
) {
	if !lib.ValidSignature(sig, message, from) { // Assignment 5 implements, just true for now.
		return nil, fmt.Errorf("Signature not valid")
	}

	// --- Homework Section for Assignment 4 ----------------------------
	// Replace the line below with code that performs a transaction
	return cc.InstructorSendFundsTransaction(from, sig, message, to, amount, memo)
	// return cc.StudentSendFundsTransaction(from, sig, message, to, amount, memo)

	//
	// Pseudo Code:
	// 1. Calcualte the total value of the account 'from'.  Call this 'tot'.
	//    You can do this by calling `cc.GetTotalValueForAccount(from)`.
	// 2. If the total, `tot` is less than the amount that is to be transfered,
	//	  `amount` then fail.  Return an error "Insufficient funds".  The person
	//    is trying to bounce a check.
	// 3. Get the list of output tranactions ( ../transactions/tx.go TxOutputType ).
	//    Call this 'oldOutputs'.
	// 4. Find the set of (may be empty - check for that) values that are pointed
	//    to in the index - from the 'from' account.  Delete this from the
	//    index.
	// 5. Create a new empty transaction.  Call `transctions.NewEmptyTx` to create.
	//	  Pass in the 'memo' and the 'from' for this tranaction.
	// 6. Convert the 'oldOutputs' into a set of new inputs.  The type is
	//    ../transctions/tx.go TxInputType.  Call `transactions.CreateTxInputsFromOldOutputs`
	//	  to do this.
	// 7. Save the new inputs in the tx.Input.
	// 8. Create the new output for the 'to' address.  Call `transactions.CreateTxOutputWithFunds`.
	//    Call this `txOut`.    Take `txOut` and append it to the tranaction by calling
	//    `transactions.AppendTxOutputToTx`.
	// 9. Calcualte the amount of "change" - if it is larger than 0 then we owe 'from'
	//    change.  Create a 2nd tranaction with the change.  Append to the tranaction the
	//    TxOutputType.
	// 10. Return
	//
}

func (cc *CLI) NewEmptyBlock() (bk *block.BlockType) {
	lastBlock := len(cc.AllBlocks) - 1
	prev := cc.AllBlocks[lastBlock].ThisBlockHash
	bk = block.InitBlock(len(cc.AllBlocks), "" /*block-comment-TODO*/, prev)
	return
}

// ListAccounts will walk through the index and find all the accounts, construct a non-dup
// list the accounts and print it out.
//
// Improvement - this could be split into a library function to get the accoutns and
// then just print.
func (cc *CLI) ListAccounts(args []string) {
	cc.ReadGlobalConfig()
	// Go through index - and list out the accounts.
	accts := make(map[string]bool)
	for key := range cc.BlockIndex.FindValue.AddrIndex {
		// fmt.Printf("Search Tx for Addr: %s\n", key)
		accts[key] = true
	}
	for key := range cc.BlockIndex.AddrData.AddrIndex {
		// fmt.Printf("Search SC for Addr: %s\n", key)
		accts[key] = true
	}
	fmt.Printf("\nList of Addresses\n")
	for key := range accts { // TODO : Should this be sorted?  If so why?
		fmt.Printf("\t%s\n", key)
	}
}

// ShowBalance will use the index to find an account, then walk through all the
// unused outputs (the balance) and add that up.  Then it will print out the
// balance for that account.
func (cc *CLI) ShowBalance(args []string) {
	cc.ReadGlobalConfig()

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Should have 1 argumetn after the flag.  AcctToList\n")
		os.Exit(1)
		return
	}

	acct := addr.MustParseAddr(args[0])

	fmt.Printf("Acct: %s Value: %d\n", acct, cc.GetTotalValueForAccount(acct))
}

// GetTotalValueForAccount walks the index finding all the non-zero tranactions for an
// account and adds up the total value for the account.
func (cc *CLI) GetTotalValueForAccount(acct addr.AddressType) (sum int) {

	unusedOutput := cc.BlockIndex.FindValue.AddrIndex[fmt.Sprintf("%s", acct)]
	// fmt.Fprintf(os.Stderr, "Acct: [%s] cc=%s AT:%s\n", acct, lib.SVarI(cc), godebug.LF())

	sum = 0
	for _, blockLoc := range unusedOutput.Value {
		if db4 {
			fmt.Fprintf(os.Stderr, "blocLoc: [%s] acct[%s] AT:%s\n", lib.SVarI(blockLoc), acct, godebug.LF())
		}

		lib.Assert(blockLoc.BlockIndex >= 0 && blockLoc.BlockIndex < len(cc.AllBlocks))
		bk := cc.AllBlocks[blockLoc.BlockIndex]

		if bk.Tx != nil {

			lib.Assert(blockLoc.TxOffset >= 0 && blockLoc.TxOffset < len(bk.Tx))
			tx := bk.Tx[blockLoc.TxOffset]

			if tx.Output != nil {

				lib.Assert(blockLoc.TxOutputPos >= 0 && blockLoc.TxOutputPos < len(tx.Output))
				out := tx.Output[blockLoc.TxOutputPos]

				lib.Assert(out.Amount >= 0)
				sum += out.Amount

				// fmt.Printf("bottom of loop: sum=[%d] AT:%s\n", sum, godebug.LF())
			}
		}
	}

	lib.Assert(sum >= 0)
	return
}

// GetNonZeroForAccount returns a slice of tranactions that have a positive (Non-Zero) balance.
// This is the set of output tranactions that will need to be turned into input tranactions
// to make a funds transfer occure.
func (cc *CLI) GetNonZeroForAccount(acct addr.AddressType) (rv []*transactions.TxOutputType) {

	unusedOutput := cc.BlockIndex.FindValue.AddrIndex[fmt.Sprintf("%s", acct)]

	for _, blockLoc := range unusedOutput.Value {
		if db6 {
			fmt.Fprintf(os.Stderr, "blocLoc: [%s] acct[%s] AT:%s\n", lib.SVarI(blockLoc), acct, godebug.LF())
		}

		lib.Assert(blockLoc.BlockIndex >= 0 && blockLoc.BlockIndex < len(cc.AllBlocks))
		bk := cc.AllBlocks[blockLoc.BlockIndex]

		if bk.Tx != nil {

			lib.Assert(blockLoc.TxOffset >= 0 && blockLoc.TxOffset < len(bk.Tx))
			tx := bk.Tx[blockLoc.TxOffset]

			if tx.Output != nil {

				lib.Assert(blockLoc.TxOutputPos >= 0 && blockLoc.TxOutputPos < len(tx.Output))
				out := tx.Output[blockLoc.TxOutputPos]

				if out.Amount > 0 {
					rv = append(rv, &out)
				}
			}
		}
	}

	return
}

// AppendTxToBlock takes a transaction and appends it to the set of transactions in the
// block.  TxOffset values for transaction outputs have to be set.  The BlockNo is
// also set to the current blocks.  If the list of addresses,
// cc.BLockIndex.FindValue.AddrIndex is nil then the map is allocated.
func (cc *CLI) AppendTxToBlock(bk *block.BlockType, tx *transactions.TransactionType) {
	Offset := len(bk.Tx)
	tx.TxOffset = Offset
	for ii := range tx.Output {
		tx.Output[ii].BlockNo = bk.Index
		tx.Output[ii].TxOffset = Offset
	}
	bk.Tx = append(bk.Tx, tx)

	if cc.BlockIndex.FindValue.AddrIndex == nil {
		// fmt.Printf("Allocate space for: cc.BlockIndex.FindValue.AddrIndex, AT: %s\n", godebug.LF())
		cc.BlockIndex.FindValue.AddrIndex = make(map[string]index.TxWithValue)
	}

	// --------------------------------------------------------------------------------
	// Check to see if existing value exists in account.
	// --------------------------------------------------------------------------------
	hashKey := fmt.Sprintf("%s", tx.Account)
	_, hasValueNow := cc.BlockIndex.FindValue.AddrIndex[hashKey]
	if db2 {
		fmt.Printf("Append: hasValueNow = %v, AT:%s\n", hasValueNow, godebug.LF())
	}

	// Find all outputs - add them to index
	for ii, out := range tx.Output {

		aTxWithValue := index.TxWithValue{
			Addr: out.Account, // Address of destination account for output
			Value: []index.TxWithAValue{{
				BlockIndex:  bk.Index,
				TxOffset:    Offset, // position of this Tx in the array of Tx in the block, this is in block.Tx[TxOffset]
				TxOutputPos: ii,     // positon of the output with a positive value in the transaction, block.Tx[TxOffset].Output[TxOutputPos]
			}}, // List of Values in a set of blocks, may have more than one value per block.
		}

		hashKey := fmt.Sprintf("%s", out.Account)
		if val, has := cc.BlockIndex.FindValue.AddrIndex[hashKey]; has {
			for _, av := range aTxWithValue.Value {
				val.Value = append(val.Value, av)
				cc.BlockIndex.FindValue.AddrIndex[hashKey] = val
			}
		} else {
			cc.BlockIndex.FindValue.AddrIndex[hashKey] = aTxWithValue
		}
	}
}

// ReadGlobalConfig reads in the index and all of the blocks.
func (cc *CLI) ReadGlobalConfig() {

	// -------------------------------------------------------------------------------
	// Read in index and blocks.
	// -------------------------------------------------------------------------------

	// Read in index so that we know what all the hashs for the blocks are.
	fnIndexPath := cc.BuildIndexFileName()          //
	BlockIndex, err := index.ReadIndex(fnIndexPath) //
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading index [%s] error [%s]\n", fnIndexPath, err)
		os.Exit(1)
	}
	cc.BlockIndex = *BlockIndex

	if db5 {
		fmt.Fprintf(os.Stderr, "dbg: AT: %s ->%s<-\n", godebug.LF(), cc.BlockIndex.Index)
	}
	for ii, key := range cc.BlockIndex.Index {
		// fmt.Printf("dbg: AT: %s ->%s<-\n", godebug.LF(), key)
		fnBlock := cc.BuildBlockFileName(key) // take the key, hash of block as a string, and generate file name.
		aBk, err := block.ReadBlock(fnBlock)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading block [%s] error [%s]\n", fnBlock, err)
			os.Exit(1)
		}
		lib.Assert(aBk.Index == ii)
		cc.AllBlocks = append(cc.AllBlocks, aBk)
		if db5 {
			fmt.Fprintf(os.Stderr, "Read In ii=%d Block[%d] Hash[%x]\n", ii, aBk.Index, aBk.ThisBlockHash)
		}
	}
}

// AppendBlock appends a new block to the set of blocks in the blockchain.
// The block mining reward will be added as the last transaction in the block.
// Mining will be performed to seal the block.  The block will be written
// out to the file system and the index of blocks is updated.  Verification
// occures that the block has produced a unique hash.  (Hash Collisions
// are possible but very rare - It would be simple to add a fix so that if
// you get a collision it fixes it.  This has not been done).
func (cc *CLI) AppendBlock(bk *block.BlockType) {
	bk.Index = len(cc.AllBlocks)

	// -------------------------------------------------------------------------------
	// add in Tx for mining reward
	// -------------------------------------------------------------------------------
	tx := transactions.NewEmptyTx("Mining Reward", cc.GCfg.AcctCoinbase)
	txOut, err := transactions.CreateTxOutputWithFunds(cc.GCfg.AcctCoinbase, cc.GCfg.MiningReward)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to supply mining reward error [%s]\n", err)
		os.Exit(1)
		return
	}
	transactions.AppendTxOutputToTx(tx, txOut)
	cc.AppendTxToBlock(bk, tx)

	// -------------------------------------------------------------------------------
	// Calculate hash for block now that the transactions are complete.
	// -------------------------------------------------------------------------------
	if db1 {
		hd := block.SerializeBlock(bk)
		fmt.Printf("Searilized Block: %x\n", hd)
	}
	bk.ThisBlockHash = hash.HashOf(block.SerializeBlock(bk))

	if db6 {
		fmt.Printf("%sbk.ThisBlockHash = %x, AT:%s%s\n", MiscLib.ColorCyan, bk.ThisBlockHash,
			godebug.LF(), MiscLib.ColorReset)
	}

	// Verify hash is unique - never seen before.
	if _, ok := cc.BlockIndex.BlockHashToIndex[fmt.Sprintf("%x", bk.ThisBlockHash)]; ok {
		lib.Assert(false)
	}

	// -------------------------------------------------------------------------------
	// Mine the block
	// -------------------------------------------------------------------------------
	mine.MineBlock(bk, cc.GCfg.MiningDifficulty)

	// -------------------------------------------------------------------------------
	// Update the block index - this is the hard part.
	// -------------------------------------------------------------------------------
	cc.BlockIndex.Index = append(cc.BlockIndex.Index, fmt.Sprintf("%x", bk.ThisBlockHash))
	if cc.BlockIndex.BlockHashToIndex == nil {
		cc.BlockIndex.BlockHashToIndex = make(map[string]int)
	}
	cc.BlockIndex.BlockHashToIndex[fmt.Sprintf("%x", bk.ThisBlockHash)] = bk.Index
	cc.AllBlocks = append(cc.AllBlocks, bk) // put in set of global blocks:

	// -------------------------------------------------------------------------------
	// Write out both the block and the updated index.
	// -------------------------------------------------------------------------------
	fnPath := cc.BuildBlockFileName(fmt.Sprintf("%x", bk.ThisBlockHash))
	fnIndexPath := cc.BuildIndexFileName()
	err = index.WriteIndex(fnIndexPath, &cc.BlockIndex) // write index
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write index to [%s] error [%s]\n", fnIndexPath, err)
		os.Exit(1)
		return
	}
	err = block.WriteBlock(fnPath, bk)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write block to [%s] error [%s]\n", fnPath, err)
		os.Exit(1)
		return
	}
}

// Debug flags to turn on output in sections of the code.
const db1 = false
const db2 = false
const db3 = false
const db4 = false
const db5 = false
const db6 = false
