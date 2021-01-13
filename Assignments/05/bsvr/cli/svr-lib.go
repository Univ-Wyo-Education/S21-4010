package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
)

// ReadGlobalData is an intialization function to read the locally stored
// set of blocks and index from ./data.
func (cc *CLI) ReadGlobalData(args []string) {
	cc.ReadGlobalConfig()
}

// ShowBalanceJSON will use the index to find an account, then walk through all the
// unused outputs (the balance) and add that up.  Then it will convert this to a
// JSON responce and return it.
func (cc *CLI) ShowBalanceJSON(acctStr string) string {
	acct, err := addr.ParseAddr(acctStr)
	if err != nil {
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}\n", err)
	}

	return fmt.Sprintf("{ \"status\":\"success\", \"acct\": %q\n  \"value\": %d }\n", acct, cc.GetTotalValueForAccount(acct))
}

// ListAccounts will walk through the index and find all the accounts, construct a non-dup
// list the accounts and print it out.
//
// Improvement - this could be split into a library function to get the accoutns and
// then just print.
func (cc *CLI) ListAccountsJSON() string {
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
	acctList := make([]string, 0, len(accts))
	for key := range accts {
		acctList = append(acctList, key)
	}
	return lib.SVarI(acctList)
}

// SendFundsJSON will process command argumetns and walk through the transaction process
// of sending funds once.  This is essentially the transaction process - but driven from
// the command line.
func (cc *CLI) SendFundsJSON(fromStr, toStr, amountStr, signature, msg, memo string) (jsonData string) {

	// In Assignment 5: args should be 6 - FromAddress, ToAddress, AmountToSend, Signature, MsgHash, Msg, Memo
	isValid, err := cc.ValidateSignature(fromStr, signature, msg)
	if err != nil {
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}\n", err)
	}
	if !isValid {
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}\n", "Signature is not valid for this account.")
	}

	// -----------------------------------------------------------------------------
	// Do the send funds stuff
	// -----------------------------------------------------------------------------
	from, err := addr.ParseAddr(fromStr)
	if err != nil {
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}\n", err)
	}
	to, err := addr.ParseAddr(toStr)
	if err != nil {
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}\n", err)
	}
	amt64, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid parse of amout [%s] error [%s]\n", toStr, err)
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}", fmt.Sprintf("Invalid parse of amout [%s] error [%s]\n", toStr, err))
	}
	amount := int(amt64) // type cast from int64 to int
	if amount <= 0 {     // Validate that no negative amount is used.
		fmt.Fprintf(os.Stderr, "Amount is out of range - can not send 0 or negative amounts [%d]\n", amount)
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}", fmt.Sprintf("Amount is out of range - can not send 0 or negative amounts [%d]\n", amount))
	}

	bk := cc.NewEmptyBlock()
	lib.Assert(bk.Index == len(cc.AllBlocks))

	tx, err := cc.SendFundsTransaction(from, lib.SignatureType(signature), msg, to, amount, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to transfer error [%s]\n", err)
		return fmt.Sprintf("{\"status\":\"error\",\"msg\":%q}", fmt.Sprintf("Unable to transfer error [%s]\n", err))
	}
	cc.AppendTxToBlock(bk, tx)

	// -----------------------------------------------------------------------------
	// Write out updated index and new block at end.
	// -----------------------------------------------------------------------------
	cc.AppendBlock(bk)

	return fmt.Sprintf("{\"status\":\"success\", \"blockNo\":%d }", bk.Index) // xyzzy other items in block, Hash, Tranaction Hash, etc.
}

func (cc *CLI) ValidateSignature(addrStr, signature, msg string) (isValid bool, err error) {
	isValid = true
	addr, err := addr.ParseAddr(addrStr)
	if err != nil {
		return
	}
	_ = addr

	// Test Validation of Signature at this point.
	return cc.InstructorValidateSignature(addrStr, signature, msg)
}
