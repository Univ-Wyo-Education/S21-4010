
m4_include(../../setup.m4)

# Assignment 06 - Install and get truffle to work

pts 50 - part 1

This is a "hello world" for truffle/solidity.   Go through the install process and get truffle to work.

At the end you should have a working systems with:


```
npm install --save -g ganache-cli
npm install --save -g truffle
npm install --save truffle-assertions
npm install --save @openzeppelin/contracts
npm install --save @openzeppelin/test-helpers
npm install --save @nomiclabs/hardhat-web3 web3
npm install --save truffle-contract
npm install --save bootstrap
npm install --save nodemon
npm install --save lite-server
npm audit fix
```

Turn in the results of running the tests on the "metacoin" example from the truffle site.
Upload a .png or .jpg file as a screen capture of "truffle test".


# Just generate an event in Solidity

pts 50 - part 2

Write a small contract that will create an event - implement the tests
to show that the event works.

contaract LogEvent {

	event AnEvent ( address indexed addr, string msg );

	// TODO - constructor
	
	function IndexedEvent ( address _acct, string _msg ) public view returns ( bool ) {
		// emit event
		emit AnEvent ( _acct, _msg );
		return (true);
	}
}


Write tests in the ./test direcotry that check for the event and verity that you 
generated an event.

Turn in your LogEvent contract (.sol file) and the test that you used from the ./test
direcotry (.js file)



