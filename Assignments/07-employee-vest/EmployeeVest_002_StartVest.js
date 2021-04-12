// SPDX-License-Identifier: MIT

const { BN, ether, expectEvent, expectRevert, time } = require('@openzeppelin/test-helpers');
const truffleAssert = require('truffle-assertions');

const EmployeeVest = artifacts.require("EmployeeVest");

contract('EmployeeVest', (accounts) => {

	const deployer = accounts[0];
	var instance;

	before("T1: Setup Contract", async function() {
		instance = await EmployeeVest.deployed();
	}); 

	it("T2: should return a Tx - Create 'Bob the Builder'", async () => {
		let tx = await instance.StartVest(accounts[4], "Bob the Builder", 1.25 * 1000000);
		// console.log ( tx );
		expectEvent ( tx, "EmployeeRegistered", {} );
		assert.equal(tx.receipt.status, true);
	});

	it("T9: should return a Tx - Faile to registerd 'Bob the Builder'", async () => {
		let tx = await instance.StartVest(accounts[4], "Bob the Builder", 1.25 * 1000000);
		expectEvent ( tx, "EmployeeError", {} );
		assert.equal(tx.receipt.status, true);
	});

});

