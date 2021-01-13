<style>
.pagebreak { page-break-before: always; }
</style>

Last Part on Go
======================

Transactions in Blockchain
-------

Data Structure from ./bsvr/block/block.go:

```
// BlockType is a single block in the block chain.
type BlockType struct {
  Index         int                             // position of this
                                                // block in the
                                                // chain, 0, 1, ...
  Desc          string                          // if "genesis" str.
                                                // then this is a
                                                // genesis block.
  ThisBlockHash hash.BlockHashType              //
  PrevBlockHash hash.BlockHashType              // This is 0 len.
                                                // if this is a
                                                // "genesis" block
  Nonce         uint64                          //
  Seal          hash.SealType                   //
  MerkleHash    hash.MerkleHashType             // Hw 03 
  Tx            []*transactions.TransactionType // Tx for Block
}

```


<br><div class="pagebreak"> </div>
Data Structure from ./bsvr/transactions/tx.go:

```
type TransactionType struct {
  TxOffset       int               // The pos. of this in the block.
  Input          []TxInputType     // Set of inputs to a transaction
  Output         []TxOutputType    // Set of outputs to a tranaction
  SCOwnerAccount addr.AddressType  // ... for SmartContracts ... 
  SCAddress      addr.AddressType  // ... for SmartContracts ... 
  SCOutputData   string            // ... for SmartContracts ... 
  Account        addr.AddressType  //
  Signature      lib.SignatureType //  Used in HW 5 - Signature 
  Message        string            //  Used in HW 5 - Message
  Comment        string            //
}

type TxInputType struct {
  BlockNo     int // Which block is this from
  TxOffset    int // The transaction in the block.
                  // In the block[BlockHash].Tx[TxOffset]
  TxOutputPos int // Position of the output in the transaction.
                  // In the block[BlockHash].Tx[TxOffset].
                  // Output[TxOutptuPos]
  Amount      int // Value $$
}

type TxOutputType struct {
  BlockNo     int              // Which block is this in
  TxOffset    int              // Which transaction in this block. 
                               // block[this].Tx[TxOffset]
  TxOutputPos int              // Pos. of the output in this block.
                               // In the  block[this].Tx[TxOffset].
                               // Output[TxOutptuPos]
  Account     addr.AddressType // Acctount funds go to (If this is
                               // ""chagne"" then this is the same
                               // as TransactionType.Account
  Amount      int              // Amoutn to go to accoutn
}
```

