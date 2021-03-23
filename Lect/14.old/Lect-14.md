<style>
.pagebreak { page-break-before: always; }
</style>

Lecture 14 - Start on Solidity
=======================================================================

## News

1. 14 countries are moving to a blockchain based currency.  [https://qz.com/1810727/central-banks-are-researching-digital-currencies-to-replace-cash/](https://qz.com/1810727/central-banks-are-researching-digital-currencies-to-replace-cash/)
2. Mashreq accomplishes this by removing the existing paper-based KYC procedure and replacing it with a blockchain-centric digital one. [https://cointelegraph.com/news/dubai-welcomes-global-businesses-with-first-blockchain-kyc-platform](https://cointelegraph.com/news/dubai-welcomes-global-businesses-with-first-blockchain-kyc-platform)


Solidity
===

## Version of the compiler 

```
pragma solidity >=0.4.25 <0.7.0;
```

## Library we will be using

```
import "openzeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
```

## Comments

```
/* Some Comment */ // Single line comment
```

## Classes and Inheritance

```
contract StandardToken is ERC20, BasicToken {
```

## Constructor

```
	constructor() public {
```

## Declare a string

```
	string public constant name = "SimpleToken"; // solium-disable-line uppercase
```

## Test for true statments

```
    require(_to != address(0));
```

## if
```
    if (_subtractedValue > oldValue) {
      allowed[msg.sender][_spender] = 0;
    } else {
      allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
    }
```

## Function

```
  function transferFrom( address _from, address _to, uint256 _value
  ) public returns (bool)
  {
```

## Events

```
  event Approval( address indexed owner, address indexed spender,
    uint256 value );
```

```
    emit Approval(msg.sender, _spender, _value);
```

<div class="pagebreak"></div>
## Example with a Loop

```
  1  pragma solidity >=0.4.25 <0.7.0;
  2  
  3  import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
  4  
  5  contract Loop is Ownable {
  6    struct StudentStruct {
  7      uint256 grade;        // grade is a percentage * 1000
  8      uint256 pin;        // Unique PIN code for student to get grade
  9    }
 10  
 11    mapping(address => StudentStruct) addressToStudent;
 12    address [] studentList;
 13  
 14    event LogStudentGrade(address student, uint256 studentGrade);
 15  
 16    function appendStudentGrade(address student, uint256 studentGrade,
       uint256 pin) public onlyOwner {
 17      studentList.push(student);
 18      addressToStudent[student].grade = studentGrade;
 19      addressToStudent[student].pin = pin;
 20    }
 21  
 22    function getStudentCount() public view returns(uint) {
 23      return studentList.length;
 24    }
 25  
 26    function getNthStudentGrade(uint i)
       public view onlyOwner returns(uint256) {
 27      require(i >= 0);
 28      require(i < studentList.length);
 29      return(addressToStudent[studentList[i]].grade);
 30    }
 31  
 32    function getNthStudentAddress(uint i)
       public view onlyOwner returns(address) {
 33      require(i >= 0);
 34      require(i < studentList.length);
 35      return(studentList[i]);
 36    }
```

<div class="pagebreak"></div>
```
 37  
 38    function getAGrade(address student, uint256 pin)
       public view returns(uint256) {
 39      require(msg.sender == address);
 40      require(pin == addressToStudent[student].pin);
 41      return(addressToStudent[student].grade);
 42    }
 43  
 44    function emitGrades () public onlyOwner {
 45      for (uint i=0; i<studentList.length; i++) {
 46        emit LogStudentGrade(studentList[i],
           addressToStudent[studentList[i]].grade);
 47      }
 48    }
 49  }
```

