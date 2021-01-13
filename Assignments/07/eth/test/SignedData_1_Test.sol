pragma solidity >=0.4.21 <0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/SignedDataVersion01.sol";

contract SignedData_1_Test {

	SignedDataVersion01 ctr = new SignedDataVersion01();

	function testPlaceholder() public {
		Assert.isTrue(true, "placeholder.");
	}

}
