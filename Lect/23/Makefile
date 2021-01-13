
all: Lect-23.html

FR=../01/

Lect-23.html: Lect-23.md
	markdown-cli --input=./Lect-23.md --output=Lect-23.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-23.html ${FR}/css/hpost >/tmp/Lect-23.html
	mv /tmp/Lect-23.html ./Lect-23.html


