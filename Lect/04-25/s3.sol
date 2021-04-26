

function withdraw(uint _amount) {
    require(balances[msg.sender] >= _amount);
    msg.sender.call.value(_amount)();			// bad
    balances[msg.sender] -= _amount;
}

