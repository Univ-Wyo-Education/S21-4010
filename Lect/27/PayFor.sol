pragma solidity >=0.4.21 <0.6.0;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract PayFor is Ownable {

    address payable owner_address;
	event ReceivedFunds(address sender, uint256 value, uint256 application, uint256 loc);
	event Withdrawn(address to, uint256 amount);

    uint256 internal nPayments;
    uint256 internal paymentID;

    address[] private listOfPayedBy;
    uint256[] private listOfPayments;
    uint256[] private payFor;

    mapping (address => uint256) internal totalByAccount;

	constructor() public {
        owner_address = msg.sender;
    	nPayments = 0;
	}

	function ReceiveFunds(uint256 forProduct) public payable returns(bool) {
		nPayments++;
		uint256 pos;
		pos = listOfPayments.length;
    	listOfPayedBy.push(msg.sender);
    	listOfPayments.push(msg.value);
    	payFor.push(forProduct);
		uint256 tot;
		tot = totalByAccount[msg.sender];
		totalByAccount[msg.sender] = tot + msg.value;
    	emit ReceivedFunds(msg.sender, msg.value, forProduct, pos);
		return true;
	}

	function getNPayments() public onlyOwner payable returns(uint256) {
		return ( nPayments );
	}

	function getPaymentInfo(uint256 n) public onlyOwner payable returns(address, uint256, uint256) {
		return ( listOfPayedBy[n], listOfPayments[n], payFor[n] );
	}
	
	function withdraw( uint256 amount ) public onlyOwner returns(bool) {
		require(amount <= address(this).balance, "Insufficient funds for witdrawl");
		address(owner_address).transfer(amount);
		emit Withdrawn(owner_address, amount);
		return true;
	}

	function getBalanceContract() public view onlyOwner returns(uint256){
		return address(this).balance;
	}
}
