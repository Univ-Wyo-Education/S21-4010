## Install truffle


Sites of Interest:

1. Truffle Site: 
[https://www.trufflesuite.com/docs/truffle/getting-started/installation](https://www.trufflesuite.com/docs/truffle/getting-started/installation)
2. Ganache - the command line runnable and front end for a stand-alone Ethereum system: 
[https://www.trufflesuite.com/docs/ganache/quickstart](https://www.trufflesuite.com/docs/ganache/quickstart)
3. Ethereum github - this is the actual Ethereum client/server:
[https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)
4. Solidity Compiler - the language that the contracts are written in: [https://github.com/ethereum/solidity](https://github.com/ethereum/solidity)


With node.js and npm installed it really should be this easy:
The `-g` is to make it a global command (that means that you
should be able to run it from any directory).

```
npm install -g truffle
```


Special note for windows users:  
[https://www.trufflesuite.com/docs/truffle/reference/configuration#resolving-naming-conflicts-on-windows](https://www.trufflesuite.com/docs/truffle/reference/configuration#resolving-naming-conflicts-on-windows)
As far as I know this has been resolved.  Note that it specifically directs anybody on Widnows to use
PowerShell or the Bash Shell that is installed when you install the command line GIT on your system.
Since I almost always use the GIT (MinGW) bash on my Windows system - I may not have seen the problem.

Lecture 16 has the source for a simple contract - but before we
do that let's get the sample 'MetaCoin' contract to work.

As soon as you get "truffle" installed the steps to follow are at: [https://www.trufflesuite.com/docs/truffle/quickstart](https://www.trufflesuite.com/docs/truffle/quickstart)
Build the "MetaCoin" example and get the test to run.

Then create an new empty project and get a the sample from Lecture 16 to work (The .sol file is in this directory also).


