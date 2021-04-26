
// Improper usage of integers
function withdraw(uint _amount) {
    require(balances[msg.sender] - _amount > 0);
    msg.sender.transfer(_amount);
    balances[msg.sender] -= _amount;
}


// One corrct way
function withdraw(uint _amount) {
    require(balances[msg.sender] >= _amount);
    msg.sender.transfer(_amount);
    balances[msg.sender] -= _amount;
}

// Anotehr
function withdraw(uint _amount) {
    require(_amount >= 0 && balances[msg.sender] >= 0 &&
		balances[msg.sender] >= _amount);
    msg.sender.transfer(_amount);
    balances[msg.sender] -= _amount;
}

