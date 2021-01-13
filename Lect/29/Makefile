
all: Lect-29.html

FR=../01/

Lect-29.html: Lect-29.md
	markdown-cli --input=./Lect-29.md --output=Lect-29.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-29.html ${FR}/css/hpost >/tmp/Lect-29.html
	mv /tmp/Lect-29.html ./Lect-29.html


