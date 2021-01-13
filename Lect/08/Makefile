
all: Lect-08.html

FR=../01/

Lect-08.html: Lect-08.md
	markdown-cli --input=./Lect-08.md --output=Lect-08.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-08.html ${FR}/css/hpost >/tmp/Lect-08.html
	mv /tmp/Lect-08.html ./Lect-08.html


