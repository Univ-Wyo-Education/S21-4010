
all: Lect-11.html

FR=../01/

Lect-11.html: Lect-11.md
	markdown-cli --input=./Lect-11.md --output=Lect-11.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-11.html ${FR}/css/hpost >/tmp/Lect-11.html
	mv /tmp/Lect-11.html ./Lect-11.html


