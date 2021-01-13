# handout-2.md


# Data types in Code

from block/block.go:

```
    type BlockType struct {
        Index         int                             // position of this block in the chain, 0, 1, 2, ...
        Desc          string                          // if "genesis" string then this is a genesis block.
        ThisBlockHash hash.BlockHashType              //
        PrevBlockHash hash.BlockHashType              // This is 0 length if this is a "genesis" block
        Nonce         uint64                          //
        Seal          hash.SealType                   //
        MerkleHash    hash.MerkleHashType             // AS-03
        Tx            []*transactions.TransactionType // Add Transactions to Block Later, (AS-04 will do this)
    }
```

from transactions/tx.go:

```
    type TransactionType struct {
        TxOffset       int               // The position of this in the block.
        Input          []TxInputType     // Set of inputs to a transaction
        Output         []TxOutputType    // Set of outputs to a tranaction
        ...
    }

    type TxInputType struct {
        BlockNo     int // Which block is this from
        TxOffset    int // The transaction in the block. In the block[BlockHash].Tx[TxOffset]
        TxOutputPos int // Position of the output in the transaction. In the  block[BlockHash].Tx[TxOffset].Output[TxOutptuPos]
        Amount      int // Value
    }

    type TxOutputType struct {
        BlockNo     int              // Which block is this in
        TxOffset    int              // Which transaction in this block.  block[this].Tx[TxOffset]
        TxOutputPos int              // Position of the output in this block. In the  block[this].Tx[TxOffset].Output[TxOutptuPos]
        Account     addr.AddressType // Acctount funds go to (If this is ""chagne"" then this is the same as TransactionType.Account
        Amount      int              // Amoutn to go to accoutn
    }
```
