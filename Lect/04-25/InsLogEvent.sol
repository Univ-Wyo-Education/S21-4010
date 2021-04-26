// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <=0.9.0;

contract InsLogEvent {

	address payable owner;

	event AnEvent ( address indexed account, string msg );

	constructor() {
		owner = payable(msg.sender);
	}

	modifier onlyOwner() {
		require( msg.sender == owner, "Sender not authorized.");
		// Do not forget the "_;"! It will be replaced by the actual function
		// body when the modifier is used.
		_;
	}

	// Make `_newOwner` the new owner of this contract.
	function changeOwner(address payable _newOwner) public onlyOwner() {
		owner = _newOwner;
	}
	
	function IndexedEvent ( address _acct, string memory _msg ) public returns ( bool ) {
		// TODO - emit event
		emit AnEvent ( _acct, _msg );
		return (true);
	}

	// This function is called for all messages sent to
	// this contract, except plain Ether transfers
	// (there is no other function except the receive function).
	// Any call with non-empty calldata to this contract will execute
	// the fallback function (even if Ether is sent along with the call).
	fallback() external payable {}

	// This function is called for plain Ether transfers, i.e.
	// for every call with empty calldata.
	receive() external payable {}

	function withdraw( uint256 _amount ) public onlyOwner() {
		owner.transfer(_amount);
	}

	// destroy the contract and reclaim the leftover funds.
    function kill() public onlyOwner() {
		//	Calling selfdestruct(address) sends all of the contract's current balance to address.
        selfdestruct(payable(msg.sender));
    }
}

