pragma solidity >=0.4.21 <0.6.0;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

// Proxy contract.  Based on Zeppelin OS.

// version 9001003004 === 9 001 003 004 - is semantic version
// v1.3.4 - with 3 digit encode and a leading 9.  This has the
// advantage of being numerically comparable for versions.

contract SignedData is Ownable {

	// Storage position of the address of the current implementation
	bytes32 private constant implementationPosition = keccak256("ethereum.plumbing.SignedData.proxy.implementation");
	uint256 private proxyToContractVersion;

	event ContractUpgradeEvent(uint256 proxyToContractVersion, address proxyToAddress);

	constructor(uint256 _version, address _implementation) public {
		require(_implementation != address(0), "Implementation address can't be zero.");
		proxyToContractVersion = _version;
		setImplementation(_implementation);
	}

	/**
	 * @dev getCurrentContractVersion returns the current version of the contarct.
	 */
	function getCurrentContractVersion() public view returns ( uint256 ) {
		return ( proxyToContractVersion );
	}

	/**
	 * @dev getToAddress returns the address of the current proxied TO contract.
	 */
	function getToAddress() public onlyOwner view returns ( address ) {
		address _impl = implementation();
		return ( _impl );
	}

	/**
	 * @dev Gets the address of the current implementation.
	 * @return address of the current implementation.
	*/
	function implementation() public view returns (address _implementation) {
		bytes32 position = implementationPosition;
		/* solium-disable-next-line */
		assembly {
			_implementation := sload(position)
		}
	}

	/**
	 * @dev Sets the address of the current implementation.
	 * @param _implementation address representing the new implementation to be set.
	 */
	function setImplementation(address _implementation) internal {
		bytes32 position = implementationPosition;
		/* solium-disable-next-line */
		assembly {
			sstore(position, _implementation)
		}
	}

	/**
	 * @dev Delegate call to the current implementation contract.
	 */
	function() external payable {
		address _impl = implementation();
		/* solium-disable-next-line */
		assembly {
			let ptr := mload(0x40)
			calldatacopy(ptr, 0, calldatasize)
			let result := delegatecall(gas, _impl, ptr, calldatasize, 0, 0)
			let size := returndatasize
			returndatacopy(ptr, 0, size)

			switch result
			case 0 { revert(ptr, size) }
			default { return(ptr, size) }
		}
	}

	/**
	 * @dev Upgrade current implementation.
	 * @param _implementation Address of the new implementation contract.
	 */
	function upgradeTo(uint256 _newVersion, address _implementation) public onlyOwner {
		address currentImplementation = implementation();
		require(_implementation != address(0), "Implementation address can't be zero.");
		require(_implementation != currentImplementation, "Implementation address must be different from the current one.");
		require(_newVersion > proxyToContractVersion, "Software version must increase to be an upgrade.");
		proxyToContractVersion = _newVersion;
		setImplementation(_implementation);
		emit ContractUpgradeEvent(_newVersion,_implementation);
	}
}
