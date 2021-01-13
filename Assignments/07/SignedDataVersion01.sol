pragma solidity >=0.4.21 <0.6.0;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract SignedDataVersion01 is Ownable {

    address payable owner_address;
	uint256 private minPayment;

	mapping(uint256 => mapping(uint256 => bytes32)) dDataR;
	mapping(uint256 => mapping(uint256 => bytes32)) dDataS;
	mapping(uint256 => mapping(uint256 => bytes2)) dDataV;
	mapping(uint256 => mapping(uint256 => address)) dOowner;
	event DataChange(uint256 App, uint256 Name, bytes32 ValueR, bytes32 ValueS, bytes2 ValueV, address By);

	event ReceivedFunds(address sender, uint256 value, uint256 application, uint256 payFor);
	event Withdrawn(address to, uint256 amount);

	constructor() public {
        owner_address = msg.sender;
		minPayment = 1000;
	}

	modifier needMinPayment {
		require(msg.value >= minPayment, "Insufficient payment.  Must send more than minPayment.");
		_;
	}

	function init() public {
		minPayment = 1000;
	}

	function setMinPayment( uint256 _minPayment ) public onlyOwner {
		minPayment = _minPayment;
	}

	function getMinPayment() public onlyOwner view returns ( uint256 ) {
		return ( minPayment );
	}

	// ----------------------------------------------------------------------------------------------------------------------

	/**
	 * @dev TODO if the data is empty, or if the msg.sender is the original createor of the data:
	 *      then : save the msg.sender into dOwner, save the data into dData
     *             create a DataChange event.
     *      else : revert an error.
	 */
	function setData ( uint256 _app, uint256 _name, bytes32 _data_r, bytes32 _data_s, bytes2 _data_v  ) public needMinPayment payable {
		// TODO-start - code for students to implement
		// 1. See if dOwner[_app][_name] is alredy set.  If it is not set then it will be address 0,
		//    if it is msg.sender then they should be able to update the data.
		// 2. If they can update then:
		//    	1. Save the msg.sender into dOwner[_app][_name]
		//    	2. Save the 3 chunks of data (the signature) into dDataR, S, V
		//		3. Emit a DataChange event with the paramters.
		//    Else:
		//		revert with an appropriate error.
		// TODO-end
	}

	/**
	 * @dev TODO return the data by looking up _app and _name in dData.
	 */
	function getData ( uint256 _app, uint256 _name ) public view returns ( bytes32, bytes32, bytes2 ) {
		// TODO-start - code for students to implement
		// Return the 3 values for the signature, R, S, V
		// TODO-end
	}

	// ----------------------------------------------------------------------------------------------------------------------

	/**
	 * @dev payable fallback
	 */
	function () external payable {
		emit ReceivedFunds(msg.sender, msg.value, 0, 1);
	}

	/**
	 * @dev genReceiveFunds - generate a receive funds event.
	 */
	function genReceivedFunds ( uint256 application, uint256 payFor ) public payable {
		emit ReceivedFunds(msg.sender, msg.value, application, payFor);
	}

	/**
	 * @dev Withdraw contract value amount.
	 */
	function withdraw( uint256 amount ) public onlyOwner returns(bool) {
		address(owner_address).transfer(amount);
		// owner_address.send(amount);
		emit Withdrawn(owner_address, amount);
		return true;
	}

	/**
	 * @dev How much do I got?
	 */
	function getBalanceContract() public view onlyOwner returns(uint256){
		return address(this).balance;
	}

	/**
	 * @dev For futute to end the contract, take the value.
	 */
	function kill() public onlyOwner {
		emit Withdrawn(owner_address, address(this).balance);
		selfdestruct(owner_address);
	}
}
