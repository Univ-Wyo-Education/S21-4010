
1. Intro - Eth
2. Install Truffle
3. Build Someting

4. Language Features
5. Language - bad features -

6. ERC-20 superseeded by ERC-777
7. An ERC-721 Smart Contract using OpenZepplin

8. An employee vesting contract
- simpler not tied back to the ERC777
- we will get to that later
9. Build 1st part of it
9. start ganache-cli
9. 1st migration
9. first test
9.  Assignment

9. Integrating vesting contract with ERC777
- get 777 to work
- mint 10000000 tokens
- get address
- initialize w/ address
- call from contract to contract
- add call for after - mint

9. Multi token ERC1155
- set of tokens both fungable and nonfungable
- a derived ERC1155
- game tokens

9. ERC721 based Grouse Credit



git@github.com:ProjectOpenSea/opensea-erc1155.git
	ERC-721, 1155 collectable marketplace



  function() payable external {
        if(!lib.delegatecall(msg.data)) {
          revert();
        }
    }


https://michalzalecki.com/ethereum-test-driven-introduction-to-solidity-part-2/


/home/pschlump/go/src/github.com/Univ-Wyo-Education/S21-4010/Eth/node_modules/@openzeppelin/contracts/token/ERC1155
