
all: Lect-18.html

FR=../01/

Lect-18.html: Lect-18.md
	markdown-cli --input=./Lect-18.md --output=Lect-18.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-18.html ${FR}/css/hpost >/tmp/Lect-18.html
	mv /tmp/Lect-18.html ./Lect-18.html


