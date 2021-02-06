


<style>
.pagebreak { page-break-before: always; }
.half { height: 200px; }
</style>








<style>
.pagebreak { page-break-before: always; }
</style>

# L08 - Transactions and Blocks

[4010 L08 pt1 Transactions and Blocks - https://youtu.be/X8uBPt22AZ8](https://youtu.be/X8uBPt22AZ8)<br>
[4010 L08 pt2 Transactions and Blocks - https://youtu.be/o0qkLigh2oE](https://youtu.be/o0qkLigh2oE)<br>

From Amazon S3 - for download (same as youtube videos)

[4010 L08 pt1 Transactions and Blocks](http://uw-s20-2015.s3.amazonaws.com/4010-L08-pt1-tranactions-and-blocks.mp4)<br>
[4010 L08 pt2 Transactions and Blocks](http://uw-s20-2015.s3.amazonaws.com/4010-L08-pt2-tranactions-and-blocks.mp4)<br>

Transactions in Blockchain
======================

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
  MerkleHash    hash.MerkleHashType             // Hw 02 2nd part 
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
  Signature      lib.SignatureType //  Used in HW 4 - Signature 
  Message        string            //  Used in HW 4 - Message
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


Portions of this are code that is MIT Licensed.<br>
Copyright (C) University of Wyoming, 2018-2021.

