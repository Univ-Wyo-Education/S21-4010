// SPDX-License-Identifier: MIT
const EmployeeVest = artifacts.require("EmployeeVest");
 
contract('EmployeeVest', (accounts) => {
 
  it("should return an int", async ()=> {
    const instance = await EmployeeVest.deployed();
    const value = await instance.GetDaysToVest();
 
    assert.equal(value, 4*365+1);
  });
});
