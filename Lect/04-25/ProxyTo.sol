// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <=0.9.0;

import "@openzeppelin/contracts/proxy/Proxy.sol";

/**
 * @title Proxy
 * @dev Gives the possibility to delegate any call to a foreign implementation.
 */
contract ProxyToXXX is Proxy {

	address private owner;
	address private GrouceCreditImpl;
	uint256 public version;

	modifier onlyOwner() {
		require( msg.sender == owner, "Sender not authorized.");
		// Do not forget the "_;"! It will be replaced by the actual function
		// body when the modifier is used.
		_;
	}

    constructor( address _addr, uint256 _version ) Proxy() {
		GrouceCreditImpl = _addr;
		owner = msg.sender;
		version = _version;
	}

	/**
    * @dev Make `_newOwner` the new owner of this contract.
	*/
	function changeOwner(address _newOwner) public onlyOwner() {
		owner = _newOwner;
	}
	
	/**
    * @dev Upgrade the underling contract to a new version. `_newAddr` is the address of
	* the new contract. `_newVersion` is the new version number.
	*/
	function upgradeContract(address _newAddr, uint256 _newVersion) public onlyOwner() {
		GrouceCreditImpl = _newAddr;
		version = _newVersion;
	}

	/**
    * @dev Tells the address of the implementation where every call will be delegated.
    * @return address of the implementation to which it will be delegated
    */
	function _implementation() internal view override  returns (address) {
		return ( GrouceCreditImpl );
	}

}
