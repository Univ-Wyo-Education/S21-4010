
all: Lect-09.html

FR=../01/

Lect-09.html: Lect-09.md
	markdown-cli --input=./Lect-09.md --output=Lect-09.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-09.html ${FR}/css/hpost >/tmp/Lect-09.html
	mv /tmp/Lect-09.html ./Lect-09.html


