pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title ExampleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned
 * to the creator.  Note they can later distribute these tokens as they
 * wish using `transfer` and other `ERC20` functions.
 */
contract ExampleToken is ERC20 {
    /**
     * @dev Constructor that gives msg.sender all of existing tokens.
     */
    constructor(
        uint256 initialSupply
    ) public ERC20("ExampleToken", "ETX") {
        _mint(msg.sender, initialSupply);
    }
}
