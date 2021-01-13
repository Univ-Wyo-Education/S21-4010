More on Finance
==


News
--


1. Supply Chain: Shipping / Mercedes-Benz / Boeing
2. Proof of Authentecity: New Balance shews and IOHK.  Serious Problem: 
[https://thewirecutter.com/blog/i-bought-these-things-from-amazon-prime-can-you-tell-which-ones-are-real/](https://thewirecutter.com/blog/i-bought-these-things-from-amazon-prime-can-you-tell-which-ones-are-real/)
3. Australian Universities have move to a Blockchain Base solution for Classes.



Notes
--

1. What is an Escrow or Escrow Account
1. Calculate the PE for a company
	1. Price to Earnings Radio - a P/E ratio is the cost of share of the company divided by Earnings of that particular share of a company.
	2. EPS = earnings ÷ total shares outstanding.
	2. Price Per Share / EPS
	2. 100 / 5 = PE of 20.   Microsoft 101.57 (Dec 31 price) / 4.31 for 2018 = 23.56
2. What is a basis point (BPS), Example 50 BSP = 0.5%
2. Insider Trading (GOP Rep Chris Collins of NY)
2. Asset Allocation
2. Prospectus
2. Pro-Forma
2. KYI - Know Your Investor (See SEC 506(d)) - Know your Customer KYC
2. Certified Investor
2. Going public - Make a public offering.  Cost etc.
2. Money Laundering
2. Proof-of-Work, Proof-of-Stake

# Reading

A simple ERC-20 Contract:
[https://www.toptal.com/ethereum/create-erc20-token-tutorial](https://www.toptal.com/ethereum/create-erc20-token-tutorial)

Zepplin ERC20:
[https://forum.openzeppelin.com/t/how-to-implement-erc20-supply-mechanisms/226](https://forum.openzeppelin.com/t/how-to-implement-erc20-supply-mechanisms/226)

```
contract ERC20FixedSupply is ERC20 {
    constructor() public {
        _mint(msg.sender, 1000);
    }
}
```
