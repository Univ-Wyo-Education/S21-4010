
all: 2020-02-25.html

FR=../01/

2020-02-25.html: 2020-02-25.md
	markdown-cli --input=./2020-02-25.md --output=2020-02-25.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre 2020-02-25.html ${FR}/css/hpost >/tmp/2020-02-25.html
	mv /tmp/2020-02-25.html ./2020-02-25.html

