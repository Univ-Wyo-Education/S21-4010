m4_include(../../setup.m4)

# L15 - Ethereum Contracts

[Solidity Overview - https://youtu.be/FXhz-zcH8ws](https://youtu.be/FXhz-zcH8ws)<br>
[Install Node.js and NPM - https://youtu.be/X8n7u7XxKrs](https://youtu.be/X8n7u7XxKrs)<br>
[Install Truffle and Files and Tools - https://youtu.be/oL_wuPsfhBA](https://youtu.be/oL_wuPsfhBA)<br>
[Hello world Eth and Truffle - ]()<br>
[Run a Test - https://youtu.be/G7Ambek_vWs](https://youtu.be/G7Ambek_vWs)<br>
[Vesting Contract - https://youtu.be/QRscb_JVC9M](https://youtu.be/QRscb_JVC9M)<br>

From Amazon S3 - for download (same as youtube videos)

[Solidity Overview](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt1-solidity-eth-overview.mp4)<br>
[Install Node.js and NPM](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt2-install-npm-and-node.mp4)<br>
[Install Truffle and Files and Tools](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt3-install-truffle-and-tools.mp4)<br>
[Hello world Eth and Truffle](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt4-hello-world-truffle-solidity.mp4)<br>
[Run a Test](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt5-run-truffle-test-on-completed-code.mp4)<br>
[Vesting Contract](http://uw-s20-2015.s3.amazonaws.com/4010-L15-pt6-vesting-contract-walkthrough.mp4)<br>

## Overview

1. Tools we are going to work with
2. The Eth Network
3. 95% of dApps on Eth
4. Cardano a new alternative
5. EIP - Ethereum Improvement Proposal
6. ERC - Ethereum Request for Comment ( ERC-20 )
7. Replacement for ERC-20 is ERC-777 - we are going to use that.
8. Homework from the contract down.
8. web3.js - how to use a contract from the front end.

## The network.

Diagram.

Use an outside test system. RMix

Single Node. Truffle.

## Install

Let's go and install truffle and build a simple contract.

Things to install:

```
npm install --save -g ganache-cli
npm install --save -g truffle
npm install --save truffle-assertions
npm install --save @openzeppelin/contracts
npm install --save @openzeppelin/test-helpers
npm install --save @nomiclabs/hardhat-web3 web3
npm install --save truffle-contract
npm install --save bootstrap
npm install --save nodemon
npm install --save lite-server
npm audit fix

```

## Files

In with the lecture:

```
2_contract.js
EmployeeVest.sol
package.json
truffle-config.js
EmployeeVest_001_getDays.js
EmployeeVest_002_StartVest.js
```

Create the project with

```
$ truffle init assignment
$ cd assignment
```

Then move the fiels to:


| File Name | Path in ./assignment |
|:----------|:-----------------------------|
|     2_contract.js       |   ./migrations/2_contract.js   |
|     EmployeeVest.sol       |   ./contracts/EmployeeVest.sol   |
|     package.json       |   ./package.json   |
|     truffle-config.js       |   ./truffle-config.js   |
|     EmployeeVest_001_getDays.js       |   ./test/EmployeeVest_001_getDays.js   |
|     EmployeeVest_002_StartVest.js       |   ./test/EmployeeVest_002_StartVest.js   |



