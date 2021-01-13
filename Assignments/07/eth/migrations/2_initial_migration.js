const SignedData = artifacts.require("./SignedData.sol");
const SignedDataVersion01 = artifacts.require("./SignedDataVersion01.sol");

module.exports = function(deployer) {
  deployer.deploy(SignedDataVersion01)
    .then(function() {
      return deployer.deploy(SignedData, 10000, SignedDataVersion01.address);
    })
	;
};

