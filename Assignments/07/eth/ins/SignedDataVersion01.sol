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
		// xyzzy-start - code for students to implement
		address tmp = dOowner[_app][_name];
		if ( tmp == msg.sender || tmp == address(0) ) {
			dOowner[_app][_name] = msg.sender;
			dDataR[_app][_name] = _data_r;
			dDataS[_app][_name] = _data_s;
			dDataV[_app][_name] = _data_v;
			emit DataChange(_app, _name, _data_r, _data_s, _data_v, msg.sender);
		} else {
			revert("Not owner of data.");
		}
		// xyzzy-end
	}

	/**
	 * @dev TODO return the data by looking up _app and _name in dData.
	 */
	function getData ( uint256 _app, uint256 _name ) public view returns ( bytes32, bytes32, bytes2 ) {
		// xyzzy-start - code for students to implement
		return ( dDataR[_app][_name], dDataS[_app][_name], dDataV[_app][_name] );
		// xyzzy-end
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
