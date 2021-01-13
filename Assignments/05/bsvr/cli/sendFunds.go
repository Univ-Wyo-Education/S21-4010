package cli

import (
	"fmt"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/transactions"
	"github.com/pschlump/MiscLib"
	"github.com/pschlump/godebug"
)

func (cc *CLI) InstructorSendFundsTransaction(
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
	tot := cc.GetTotalValueForAccount(from)
	if tot < amount {
		return nil, fmt.Errorf("Insufficient funds")
	}
	oldOutputs := cc.GetNonZeroForAccount(from)
	if db3 {
		fmt.Printf("%sOld Outputs (Step 1): %s, AT:%s%s\n", MiscLib.ColorYellow, lib.SVarI(oldOutputs), godebug.LF(), MiscLib.ColorReset)
	}
	// remvoe inputs from index.
	fromHashKey := fmt.Sprintf("%s", from)
	if _, ok := cc.BlockIndex.FindValue.AddrIndex[fromHashKey]; ok {
		delete(cc.BlockIndex.FindValue.AddrIndex, fromHashKey)
	}
	tx = transactions.NewEmptyTx(memo, from)
	// create inputs into tranaction from "oldOutputs"
	txIn, err := transactions.CreateTxInputsFromOldOutputs(oldOutputs)
	if err != nil {
		return nil, err
	}
	if db3 {
		fmt.Printf("%sNew Inputs (Step 2): %s, AT:%s%s\n", MiscLib.ColorYellow, lib.SVarI(txIn), godebug.LF(), MiscLib.ColorReset)
	}
	tx.Input = txIn
	txOut, err := transactions.CreateTxOutputWithFunds(to, amount)
	if err != nil {
		return nil, err
	}
	transactions.AppendTxOutputToTx(tx, txOut)
	change := tot - amount
	if change > 0 {
		txOut, err := transactions.CreateTxOutputWithFunds(from, change)
		if err != nil {
			return nil, err
		}
		transactions.AppendTxOutputToTx(tx, txOut)
	}
	return
}
