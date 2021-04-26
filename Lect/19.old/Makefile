
all: Lect-19.html

FR=../01/

Lect-19.html: Lect-19.md
	markdown-cli --input=./Lect-19.md --output=Lect-19.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-19.html ${FR}/css/hpost >/tmp/Lect-19.html
	mv /tmp/Lect-19.html ./Lect-19.html


