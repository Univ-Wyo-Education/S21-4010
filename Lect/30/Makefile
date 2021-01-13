

all: Lect-30.html

FR=../01/

Lect-30.html: Lect-30.md
	markdown-cli --input=./Lect-30.md --output=Lect-30.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-30.html ${FR}/css/hpost >/tmp/Lect-30.html
	mv /tmp/Lect-30.html ./Lect-30.html


