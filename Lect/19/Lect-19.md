# Lecture 19 - security - and signatures

# News

1. The stimulus bill included a "digital currency".
[https://www.forbes.com/sites/michaeldelcastillo/2020/03/25/trillion-dollar-stimulus-jumpstarts-project-to-issue-central-bank-currency-on-ethereum/#69f5d60c47bc](https://www.forbes.com/sites/michaeldelcastillo/2020/03/25/trillion-dollar-stimulus-jumpstarts-project-to-issue-central-bank-currency-on-ethereum/#69f5d60c47bc)
2. Atari is using Ethereum to trade game assets
[https://cointelegraph.com/news/gaming-giant-atari-to-feature-in-the-sandboxs-forthcoming-blockchain-platform](https://cointelegraph.com/news/gaming-giant-atari-to-feature-in-the-sandboxs-forthcoming-blockchain-platform)

# Videos

[ Lect-19-4010-pt1.mp4: https://youtu.be/bPAwT39YY9Q ](https://youtu.be/bPAwT39YY9Q) <br>
[ Lect-19-4010-pt2.mp4: https://youtu.be/ZJ9pAP8So8o ](https://youtu.be/ZJ9pAP8So8o) <br>
[ Lect-19-4010-pt3.mp4: https://youtu.be/txbFyoe5uXA ](https://youtu.be/txbFyoe5uXA) <br>
[ Lect-19-4010-pt4.mp4: https://youtu.be/XK5FlYDXWc8 ](https://youtu.be/XK5FlYDXWc8) <br>
[ Lect-19-4010-pt5.mp4: https://youtu.be/D7R-FxFyqGs ](https://youtu.be/D7R-FxFyqGs) <br>

Amazon S3 - only if you need to download videos (these are the same as the YouTube
versions - just setup so that you can download them - this is primarily intended
for people that have a slow connection or need to view the videos offline):

[Lecture 19 - quick zoom overview]([http://uw-s20-2015.s3.amazonaws.com/zoom01.mp4) <br>
[Lecture 19 - part 1](http://uw-s20-2015.s3.amazonaws.com/Lect-19-4010-pt1.mp4) <br>
[Lecture 19 - part 2](http://uw-s20-2015.s3.amazonaws.com/Lect-19-4010-pt2.mp4) <br>
[Lecture 19 - part 3](http://uw-s20-2015.s3.amazonaws.com/Lect-19-4010-pt3.mp4) <br>
[Lecture 19 - part 4](http://uw-s20-2015.s3.amazonaws.com/Lect-19-4010-pt4.mp4) <br>
[Lecture 19 - part 5](http://uw-s20-2015.s3.amazonaws.com/Lect-19-4010-pt5.mp4) <br>

# Personal Security - Why

Systems that survive pay out to all the responsible parties in proportion
to the level of responsibility.

## Friction vs Access

Friction to get to the information you want is security to prevent unauthorized
people from getting to the information that you don't want them to get to.

People will choose convenience over security every time.

Data ownership.  Equafax and what data you have as public.

Individual are not motivated to protect systems.

### Passwords

Top 10 List
```
		123456
		123456789
		qwerty
		12345678
		111111
		1234567890
		1234567
		password
		123123
		987654321
```
Represent 76% of passwords in use.

Top 10000 passwords are 92% of passwords in use.

Hash passwords: MD5, SHA1

Rainbow Tables

Salt and then Hash

use a hash that is appropriate: bcrypt, scrypt, or PBKDF2.

1. Generated Passwords ( 1password or other similar tools ).
```
#!/bin/bash
openssl rand -base64 20
```
2. Never change passwords
3. Write it down.
4. Use a system that verifies that your password is not pwned.


### Other security systems good/bad.

1. Auth0 - JWT
2. Password-less login systems
3. Biometrics
4. Biometrics that works
5. NFC
6. Wireless WiFi and security
7. IOT
8. Back doors into systems
9. Chipping dogs - why not people?
10. Software is never 100% fixed.  Security everywhere.  No tootsy-pop solutions.  Hard & Crunchy on the Outside, Soft & Chewy on the Inside

### Security == Friction

"Making Password Cracking Harder: Slow Hash Functions"
Proof of Work: PBKDF2.

Risk of pandemic is really an airline - high speed transpiration system.

Speed Limits: highways - approaches to airports.
Drugs - require prescriptions.

Digital information moves limitlessly. The same design philosophy that
accelerated the flow of correspondence, news, and commerce also accelerate
the flow of phishing, ransom-ware, and disinformation.

### Free isn't free.

Sherman Antitrust Act: 1989 Supreme Court Decision - that evaluation of 
an illegitimate merger be based on price of goods produced.  
We are now in a world where 0 price is common.

So if you give away your product (e.g. Facebook) then you purchase some
other entity (Instagram) and give it away for free then the SAA will not
apply to you.

### Who Pays?

During the height of the recession, in 2009, we wasted 6.3 billion hours on the road.   The problem is that the employer is not paying for
commute time.  They get it for "free".  Now suddenly in 2020 we are seeing everybody work from home for all the worst reasons.   Maybe
this will change the daily commute.   Commuting is terrible for people - not just the environment - People that commute more than 1 hour
each way have a 40% increase in divorces.

Another example of free but not good for people is classes that uses PDFs for handouts - instead of printing the handouts and handing
them out.  This save the author time and money - but costs the students.  Free to the author - not free in the real world.

A lot of the BlockChain tracking is to make systems where cheating  is taking place more "fair" where everybody contributes and everybody
is payed proportionately to the contribution.  The difficulty in designing systems like this is to make it "fair".


### Real Security Solutions - public private keys and signatures

This is our homework - to sign a document and validate it on the server.
I provided the server and part of the client (you get to make the library
calls to do the "sign" part) and I provided a command line sign tool.
(Hint - it has the sign code in it - look in the homework it will
direct you to the correct file and functions)

For authentication you can use a zero-knowledge proof system called
SRP6a - search for that "Secure Remote Password SRP6a" and there are
implementations for most languages. 

Interesting side note: the Wikipedia page on this has a Python program
that sort of works to do SRP6a.  It will work in sending stuff to
a duplicate copy of the program and it has the correct algorithm in it.
In one 4 line section of code it also has 5 type conversion errors
that will cause about 1 in 16 messages to get the wrong results.
YMMV on example code! (the code also appears in the standard for
SRP6a!)

SRP6a is used in SSH for remote logins! so most of you have used it already.
It has been around for a long time.   The good side - the server has a
validation number instead of your password - so - it can never leek your
password.  Also this is a "strong" authentication system based on
cryptography - not just pass the password over the wire authentication.







