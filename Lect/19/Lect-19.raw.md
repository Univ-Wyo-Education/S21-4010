
m4_include(../../setup.m4)

# Lecture 19 - Smart Contract Pitfalls

Smart contracts have a unique set of limitings and pitfalls.

First put them in context.  An analogy...  Cryptocurency is to government-money, what smart-contracts are to law.

## General Limitations on Solidity

1. no way to iterate over maps
2. data stored in arrays has to be returned 1 at a time
3. friction between 256bit ints and JavaScript
4. multiple-inheritance object oriented is just generally bad
5. solidity keeps changing
5. weird data types byres32 is same storage as int256/uint256, strings are 1 byte in a uint256 etc.
5. "address" type may or may not be "payable"
5. lots of "type-casting" required.

## Limitations to the code.

A smart contract is a piece of code that is tied to a blockchain, triggered by transactions and data passed in the transactions and which reads and writes data in that blockchain’s history.
The code behaves as if it is an embed software system.   To maintain consistency between nodes in the chain - all nodes must see the same data.
This means that the smart contract can not go pull the price of a stock itself - if node1 sees a different value than node2 they will get different output results.
The data must be fixed and written to the chain before the smart contact is called, or it must be a parameter in the transaction and shared across all the nodes.
This means that everything that takes place on a blockchain must be deterministic.

## Smart contracts are forever.

Defects are forever.  You can't go back and change the contact and get a consistent new state for the block.
Microsoft did a book called "Code Complete" that found that shipping code from Microsoft had 15-20 defects
per 1000 lines of code.  Some systems like NASA's control system for the Shuttle had much fewer defects, close
to 0 per 1000 lines of code.   The cost for developing systems like that was 100,000x as much.   So...
For smart contracts which development system are you using?  

## Available forever.  

Contracts don't just go away.  Version 1 of your contract is still there when you
really want to be on version 3.   If you want to stop people from using version one you have to have built
into the contract a flag to say stop running.   You have to protect the flag so that only the contract
owner or some authorized list of people can set/unset the flag.

## Upgrade-ability. 

Since the contract and it's address are forever - you have to plan a way to upgrade
the contract.  There is a "Proxy" system that has been used for this but upgrades could also be
handled by forcing users to use a new-address and a new contract.

The "Proxy" model for upgrades is really nasty and complex.

```
m4_include(ProxyTo.sol)
```

```
m4_include(Proxy.sol)
```

## Security is up to the contract writer.  An example

```
m4_include(s1.sol)
```

## You have to plan for the future - hard to do. 
What happens when the author of the contract dies?
What happens if the business goes under?

## Integer underflow has to be considered in ALL calculations!

```
m4_include(s2.sol)
```


## Contract to contact calls must be carefully handled.  (Or not used at all).

What happens if the "call" in the contract leads back to the same withdraw
call?

```
m4_include(s3.sol)
```


## You need other "things" in the contract that are usually not obvious.

```
m4_include(InsLogEvent.sol)
```


Code is MIT Licensed.<br>
Copyright (C) University of Wyoming, 2018-2021.

