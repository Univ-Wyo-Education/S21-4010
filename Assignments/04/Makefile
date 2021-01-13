
all: install-truffle.html 

FR=../../Lectures/01

install-truffle.html: install-truffle.md
	markdown-cli --input=./install-truffle.md --output=install-truffle.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre install-truffle.html ${FR}/css/hpost >/tmp/install-truffle.html
	mv /tmp/install-truffle.html ./install-truffle.html



