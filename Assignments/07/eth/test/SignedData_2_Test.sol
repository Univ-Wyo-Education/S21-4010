pragma solidity >=0.4.21 <0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/SignedDataVersion01.sol";
import "../contracts/SignedData.sol";

contract SignedData_2_Test {

	SignedDataVersion01 sd = new SignedDataVersion01();
	SignedData ctr = new SignedData( 9000000004, address(sd) );

	// test through the proxy.

	// function setData ( uint256 _app, uint256 _name, bytes32 _data ) onlyOwner public {
	// function getData ( uint256 _app, uint256 _name ) public view returns ( bytes32 ) {
//	function testSetGetData01() public {
//		bytes32 d1;
//		bytes32 d2;
//		bytes32 d3;
//		d1 = hex"04000000000500000000060000000007000000000608";
//		d3 = hex"08000000000500000000060000000007000000000608";
//		ctr.setData ( 1, 44, d1 );
//		d2 = ctr.getData ( 1, 44 );
//		Assert.equal(d1, d2, "set/get data.");
//		ctr.setData ( 1, 44, d3 );
//		d2 = ctr.getData ( 1, 44 );
//		bool flag;
//		flag = true;
//		if ( d2 == d1 ) {	// if data matches then this is a problem.
//			flag = false;
//		}
//		Assert.isTrue(flag, "set/get data (2).");
//		flag = false;
//		if ( d2 == d3 ) {	// data should match now.
//			flag = true;
//		}
//		Assert.isTrue(flag, "set/get data (3).");
//	}

	function testGetVersion() public {
		uint vv;
		vv = ctr. getCurrentContractVersion();
		Assert.equal(vv, 9000000004, "get verison of proxy contract.");
	}

}
