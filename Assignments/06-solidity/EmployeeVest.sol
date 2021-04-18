// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <=0.9.0;

/**
 * @title EmployeeVest
 * @dev Combining a simple ERC777 token with an employee vesting plan.
 * 
 * Employees are vested for ownership over a 4 year period.  This is
 * done on a per-day basis.   If the employee leaves the company before
 * the 4 years they loose all vested tokens.
 * 
 * Each employee gets a percentage of ownership (1/4th pect, 1/2pct).
 */
contract EmployeeVest {

	struct Employee { 
		string name;
		uint256 startdate;
		uint256 pctToVest;
		bool isVested;
		bool isTransfered;
		uint8 found;
	}

    address private ownerAddress;
    uint256 ntoken;
    uint256 daysToVest;
    mapping (address => Employee) internal vesting;
	address [] vest_list;

	event EmployeeRegistered ( address indexed employeeAddr, string name, uint256 startDate );
	event EmployeeFullyVested ( address indexed employeeAddr, bool isVested );
	event EmployeeSeparated ( address indexed employeeAddr, bool isVested );
	event EmployeeError ( address indexed employeeAddr, string msg );

	modifier onlyOwner() {
		require(msg.sender == ownerAddress, "Can only be called by the contract owner.");
		_;
	}


    /**
     * @dev Constructor that saves msg.sender (owner), total tokens and vesting time 
	 * with a default vesting of 4 years if _daysToVest is passed as 0.
	 * @param _ntoken number of tokens to mint.
	 * @param _daysToVest number days before employee is considered vested.
     */
    constructor(uint256 _ntoken, uint256 _daysToVest) {
        ownerAddress = msg.sender;
		// ERC777 - call
		// Should et # of tokens from ERC777 contract.
		ntoken = _ntoken;
		if ( _daysToVest == 0 ) {
			daysToVest = 365 + 365 + 365 + 366;	// 4 years
		} else {
			daysToVest = _daysToVest;
		}
    }

    /**
     * @dev GetDaysToVest
	 * returns the number of days the contract is set to vest in.
     */
	function GetDaysToVest() public view returns(uint256) {
		return(daysToVest);
	}

    /**
     * @dev StartVest is called when an employee starts and sets the beiginning date
	 * to the current block number time.
	 * @param _acct the employees ethereum account number.
	 * @param _name the employees name.
	 * @param _tvest the percentage times 1 million of the stock that the employee receives.
     */
	function StartVest(address _acct, string memory _name, uint256 _tvest) public onlyOwner returns(bool) {	
		// TODO check that this account is not already used.  IF it is then fail a require.
		// TODO populate a struct, populate startdate with block.timestamp
		// TODO save the new employee
		// TODO emit EmployeeRegistered event.
		// TODO return true
	}

    /**
     * @dev StartVestDate is called when an employee starts and sets the beiginning date
	 * to the current block number time.
	 * @param _acct the employees ethereum account number.
	 * @param _name the employees name.
	 * @param _tvest the percentage times 1 million of the stock that the employee receives.
	 * @param _startDate the date that the employee was hired on in unix timestamp format.  Seconds since 1970.
     */
	function StartVestDate(address _acct, string memory _name, uint256 _tvest, uint256 _startDate) public onlyOwner returns(bool) {	
		// TODO check that this account is not already used.  IF it is then fail a require.
		// TODO populate a struct, populate startdate with _startDate
		// TODO save the new employee
		// TODO emit EmployeeRegistered event.
		// TODO return true
	}


    /**
     * @dev SetEmployeeVested marks an employee as being vested.  This is usually for founders of the company.
	 * @param _acct the employees ethereum account number.
     */
	function SetEmployeeVested (address _acct) public onlyOwner returns(bool) {
		// TODO check that this account is used.  
		// TODO set flag isVested to true.
		// TODO save data.
		// TODO return true
	}

    /**
     * @dev EmployeeSeparation Employee leaves the company.  If employee is not vested
	 * then all stock (tokens) are lost.
	 * @param _acct the employees ethereum account number.
     */
	function EmployeeSeparation (address _acct) public onlyOwner returns(bool) {
		// TODO check that this account is used.  
		// TODO if employee IsVested - then
		// TODO emit event EmployeeSeperated with a "true" flag.
		//      else set pctToVest to 0 for employee
		//         TODO save data.
		//         TODO emit event EmployeeSeperated with a "false" flag.
	}

    /**
     * @dev CheckEmployeeVested Return true if the employee is vested.
	 * @param _acct the employees ethereum account number.
     */
	function CheckEmployeeVested(address _acct) public onlyOwner returns(bool) {
		// TODO check that this account is used.  
		// TODO xyzzy
		if ( IsVested(_acct) ) {
			// ERC777 - call
			// Should transfer tokens to employee at this point.
			// Set e.isTransfered = true; // and update data.
			emit EmployeeFullyVested ( _acct, true );
			return (true);
		} else {
			emit EmployeeFullyVested ( _acct, false );
			return (false);
		}
	}
	
    /**
     * @dev GetName - return the name of the employee.
	 * @param _acct the employees ethereum account number.
     */
	function GetName(address _acct) public view returns(string memory) {
		Employee memory f;
		string memory name;
		name = "";
		f = vesting[_acct];
		if ( f.found == 1 ) {
			name = f.name;
		}
		return(name);
	}

    /**
     * @dev GetStartDate return the start date in unix timesamp format.
	 * @param _acct the employees ethereum account number.
     */
	function GetStartDate(address _acct) public view returns(uint256) {
		Employee memory f;
		f = vesting[_acct];
		if ( f.found == 1 ) {
			return(f.startdate);
		}
		return(0);
	}

    /**
     * @dev GetNthEmployeeAcct returns the Nth value in the vest_list.  0 is
	 * returns if the value of _n is out of range.
	 * @param _n the subscript for the array
     */
	function GetNthEmployeeAcct(uint _n) public view returns(address) {
		// TODO check that _n is in range, if so return [_n]
	}

    /**
     * @dev GetEmployeeAcctLength returns the number of values in the vest_list.
     */
	function GetEmployeeAcctLength() public view returns(uint) {
		// TODO -- return length of array
	}

    /**
     * @dev DaysEmployeed Calculate number of days emploeed.
	 * @param _acct the employees ethereum account number.
     */
	function DaysEmployeed(address _acct) public view returns(uint256) {
		// TODO check that this account is used.  
		// TODO xyzzy
		Employee memory f;
		f = vesting[_acct];
		require(f.found == 1, "E-197: Employee not found.");
		if ( f.found == 1 ) {
			uint256 diff = (block.timestamp - f.startdate) / 60 / 60 / 24; 
			return (diff);
		}
		return(0);
	}

    /**
     * @dev IsVested returns true if employee is fully vested.
	 * @param _acct the employees ethereum account number.
     */
	function IsVested(address _acct) public view returns(bool) {
		// TODO check that this account is used.  
		// TODO xyzzy
		Employee memory f;
		f = vesting[_acct];
		require(f.found == 1, "E-212: Employee not found.");
		if ( f.isVested ) {
			return (true);
		}
		uint256 dv = DaysEmployeed(_acct);
		if ( dv >= daysToVest ) {
			return (true);
		}
		return (false);
	}

    /**
     * @dev ProportionVested returns the current proprtion of vesting that the employee has in number of tokens.
	 * @param _acct the employees ethereum account number.
     */
	function ProportionVested(address _acct) public view returns(uint256) {
		// TODO check that this account is used.  
		// TODO xyzzy
		Employee memory f;
		f = vesting[_acct];
		uint256 dv = DaysEmployeed(_acct);
		if ( dv >= daysToVest ) {
			uint256 pd0 = ( f.pctToVest * ntoken );
			pd0 = pd0 / 1000000;
			return ( pd0 / 100 );
		}
		uint256 pd = ( f.pctToVest * ntoken );
		// pd = pd / 1000000;
		pd = pd / daysToVest;
		// Suppose that an employee gets 0.25%, as a decimal this is 0.0025.
		// Multiply by Stock at issue of company, 10 million.
		//        Pct                * TotalStock 
		// a = ( 0.0025 * 1000000 ) * 10000000 = 	25,000,000,000
		// a /= 1,000,000                            // Take back out the multiplier for the PCT 
		// a                                         // is now # of shares vested over 4 years
		// a / 4y
		// Devides are saved till the end so accuracy is not lost.
		return ( ( ( dv * pd ) / 1000000 ) / 100 ) ;
	}

	// TODO - change name of person (Marriage etc)

	
}

