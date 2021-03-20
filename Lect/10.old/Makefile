
all: Lect-10.html

FR=../01/

Lect-10.html: Lect-10.md
	markdown-cli --input=./Lect-10.md --output=Lect-10.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-10.html ${FR}/css/hpost >/tmp/Lect-10.html
	mv /tmp/Lect-10.html ./Lect-10.html


