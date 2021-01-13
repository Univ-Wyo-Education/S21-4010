pragma solidity >=0.4.25 <0.6.0;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721Token.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

/**
 * @title Demo721
 * @dev Very simple ERC-721 Token example.  This is what Crypto-Kitties is built on.
 */
contract Demo721 is Ownable, ERC721Token {

	string public constant name = "Demo ERC-721"; // solium-disable-line uppercase
	string public constant symbol = "D721"; // solium-disable-line uppercase

	constructor() public
        ERC721Token(name, symbol)
	{
	}

	function mint(address _to, uint256 _tokenId, string _tokenURI) public onlyOwner {
		super._mint(_to,_tokenId);
		super._setTokenURI(_tokenId,_tokenURI);
	}

	function burn(address _owner, uint256 _tokenId) public onlyOwner {
		super._burn(_owner,_tokenId);
		super._setTokenURI(_tokenId,"");
	}
}

