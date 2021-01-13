
all: Lect-13.html

FR=../01/

Lect-13.html: Lect-13.md
	markdown-cli --input=./Lect-13.md --output=Lect-13.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-13.html ${FR}/css/hpost >/tmp/Lect-13.html
	mv /tmp/Lect-13.html ./Lect-13.html


