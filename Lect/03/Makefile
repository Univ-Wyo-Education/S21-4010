
all: Lect-03.html

FR=../01/

Lect-03.html: Lect-03.md
	markdown-cli --input=./Lect-03.md --output=Lect-03.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-03.html ${FR}/css/hpost >/tmp/Lect-03.html
	mv /tmp/Lect-03.html ./Lect-03.html

run:
	./run.sh

