
Insurance Companies
===================

Videos:

[https://youtu.be/pVCqrFMs7zA - Lect-22-4010-5010-cosc-Pt1.mp4](https://youtu.be/pVCqrFMs7zA)<br>
[https://youtu.be/Hr_8yYii9Mc - Lect-22-4010-5010-pt2-Insurance.mp4](https://youtu.be/Hr_8yYii9Mc)<br>
[https://youtu.be/IBNtM5GzgIE - Lect-22-4010-5010-pt3-KYC-KYI-Tracking.mp4](https://youtu.be/IBNtM5GzgIE)<br>
[https://youtu.be/fG0iVUIeFiI - Lect-22-4010-5010-pt4-hw.mp4](https://youtu.be/fG0iVUIeFiI)<br>

From Amazon S3 - for download (same as youtube videos)

[http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-cosc-Pt1.mp4](http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-cosc-Pt1.mp4)<br>
[http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt2-Insurance.mp4](http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt2-Insurance.mp4)<br>
[http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt3-KYC-KYI-Tracking.mp4](http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt3-KYC-KYI-Tracking.mp4)<br>
[http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt4-hw.mp4](http://uw-s20-2015.s3.amazonaws.com/Lect-22-4010-5010-pt4-hw.mp4)<br>

## News

1. The World Economic Forum Highlighted Blockchain and Digitization to Address COVID-19 Supply Chain Disruption
2. Older article on China and Blockchain.
The Washington Post has an article:
[https://www.washingtonpost.com/opinions/global-opinions/china-is-racing-ahead-of-the-united-states-on-blockchain/2019/03/07/c1e7776a-4116-11e9-9361-301ffb5bd5e6_story.html](https://www.washingtonpost.com/opinions/global-opinions/china-is-racing-ahead-of-the-united-states-on-blockchain/2019/03/07/c1e7776a-4116-11e9-9361-301ffb5bd5e6_story.html)

"The grand strategic competition between the United States and China will be won
or lost based on who controls the rules and systems that govern 21st-century
commerce, communications and security. Washington can no longer ignore that
China is way ahead on the core technology that will underpin those systems:
blockchain."


# Homework 6  - Document Metadata Contract

There are an entire class of blockchain applications that are based on "meta" data
and proof of accurate data - where the actual data is stored off-chain but the
"Proof" of accuracy is on chain.  For data that is not small - video, sound, images
or data that has special searching requirements like ground-truth (GPS coordinate)
data this is a very viable solution.

Data in the database then needs to be immutable - and then chunks of it are hashed
and stored in some known location.  For example Amazon S3 - Simple Storage Subsystem
or IPFS.    Then an auditor can go to the chain, find all the hashes of the data,
take the files of data that are store and roll a database forward from the beginning
to the current and audit the correctness of the data.

The current work between BeefChain - a Wyoming company and IOHK on supply chain for
animals works this way. 

There is an ISO 9000 / IOS 9002 document company that uses this approach.

See the homework - for the interface.

The other common classes of blockchain applications are "token" based and "payment"
based. (ERC-20, or ERC-721).  For example the Colorado inmate parole system is
based on ERC-721 contracts.   Rational land sale is based on an ERC-721 contract.
These contracts are used for tracking individual "things" or counts of "things".



## Insurance and Non-Fungible-Tokens (ERC-721 | NFT)

What is Insurance? What do Insurance Companies do?

In insurance, the insurance policy is a contract (generally a standard form contract) between the insurer and the
insured, known as the policyholder, which determines the claims which the insurer is legally required to pay. In
exchange for an initial payment, known as the premium, the insurer promises to pay for loss caused by perils covered
under the policy language.

1. Identify Risks and Markets
2. Manage payments
2. Provide Initial Capital 
4. Pay out claims - Pay out on terms of insurance

90% of what they do is all overhead.

Medicare - Overhead(Operating costs) less than 2%, Private 17.8%.


| Percent  |  Item |
|----------|----------------------------------------------------|
|   22.1%  |    Drugs											|
|   20.0%  |     Physical Services								|
|   19.8%  |     Outpatient Services							|
|   15.8%  |     Inpatient Services								|
|   17.8%  |     Insurance Overhead								|
|    3.0%  |     Insurance Profit								|

856B * 17% - save 145B										



Term Life Insurance



| Percent |  Item |
|---------|----------------------------------------------------------|
|   61%   |      Pay out                                             |
|   11%   |      Product development                                 |
|    9%   |      Operations                                          |
|   11%   |      IT                                                  |
|    8%   |      Sales/Costs (across industry only a 19% variance,   |
|         |        6.56% to 9.52%)                                   |
|         |       (Size, Sales Channels, Product Mix, Geography)     |
|    0%   |       Profit                                             |



$155 Billion a year - $48 billion opportunity - save minimum of $36B - probably closer to $45B
    



How it Works?  Why are they still in finesse?

3 Parts with Smart contracts.


1. Future Sales Prediction Markets
2. Self managing "overhead"
3. Payout based on reliable "observers" - reputation based pay out.


## Banking and KNC.

1. $100 billion at large banks for KNC
2. 17% increase in cost 2016, 18% increase in cost 2017.
3. 19% of customers say switched banks because of KNC
4. A "token" that identifies customers accurately.




www.bcg.com - First blockchain only insurance company.



# Copyright

Copyright (C) University of Wyoming, 2020.


