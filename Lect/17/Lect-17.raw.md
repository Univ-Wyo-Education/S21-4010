
m4_include(../../setup.m4)

# Lecture 17 - Non Fungible Tokens

[Truffle MetaCoin Walk-through - https://youtu.be/hj2Oxo5YpFo](https://youtu.be/hj2Oxo5YpFo)<br>
[NFTs ERC 721 contract - https://youtu.be/21GzrWBBVIk](https://youtu.be/21GzrWBBVIk)<br>
[NFT Sales Process - 4010-L17-ERC-721-part-2-buy-sell-them.mp4](4010-L17-ERC-721-part-2-buy-sell-them.mp4)<br>

From Amazon S3 - for download (same as youtube videos)

[Truffle MetaCoin Walk-through](http://uw-s20-2015.s3.amazonaws.com/4010-L17-truffle-walkthrough.mp4)<br>
[NFTs ERC 721 contract](http://uw-s20-2015.s3.amazonaws.com/4010-L17-ERC-721-part-1-the-721-contract.mp4)<br>
[NFT Sales Process](http://uw-s20-2015.s3.amazonaws.com/https://youtu.be/xDxsHXl6XpU)<br>

Non-Fungible Tokens or (NFTs) are in the news with some art tided
to a NFT selling for $69 million.

Aft has been plagued with forgeries and theft for ages.  The nice
thing about a NFT tied to some art work is that the NFT can show
the ownership of the art to be a specific account on the chain.

Also the NFT is a a unique "value" for each smart contact.
Usually this is then tied to a set of off-chain data
that describes what the NFT means.

In the land of Solidity/Ethereum we can build a simple non-fungable token 
contract using the ERC-721 base and extending it.

```
m4_include(SimpleERC721.sol)
```

You can buy, sell, trade ERC721 tokens that support the standard interface
on exchanges like Kraken and Coinbase.  When you transfer tokens you have to
be certain that the destination contract that you are sending the token to
can understand a "NFT" token.

The important difference between fungible and non-fungible is that there is
a unique ID for non-fungible.  This ID can be used to lookup outside unique
information about the token.

Notice the tokenURI that is saved with the token.  Let's say the token is 
132 and it's unique URI is
`https://www.register-tokens/d?type=CO2CREDIT&id=132`.
When we access this URL we can get back meta-information about this
token - usually in JSON format.

```
{
	"NFTID": 132,
	"Type": "Co2 Credit",
	"Issuer: "Clear Air, Swuamish BC Canada",
	"Date": "2021-04-18T11:22:33",
	"Signed": "23821903801283902183091283021839021839021830921890128903801280398210382103809123"
}
```

Let's take a look at the underlying 721 contract and what it provides.



