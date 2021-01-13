require('babel-register');
require('babel-polyfill');

module.exports = {
	networks: {
		development: {
			host: "127.0.0.1",
			port: 9545,
			network_id: "*" // Match any network id
		}
		, test3: {
			host: "192.168.0.199",
			port: 8545,
			network_id: "*",
		    gas: 4712388,
			from: "0x6ffba2d0f4c8fd7961f516af43c55fe2d56f6044"
		}
		, local: {
			host: "127.0.0.1",
			port: 7545,
			network_id: "*",
		    gas: 4712388,
			from: "0x6ffba2d0f4c8fd7961f516af43c55fe2d56f6044"
		}
	}
};
