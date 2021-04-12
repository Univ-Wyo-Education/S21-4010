// SPDX-License-Identifier: MIT

const EmployeeVest = artifacts.require("EmployeeVest");

module.exports = function (deployer) {
  deployer.deploy(EmployeeVest,10000000,0);
};
