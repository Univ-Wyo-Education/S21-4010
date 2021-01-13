
all: Lect-06.html

FR=../01/

Lect-06.html: Lect-06.md
	markdown-cli --input=./Lect-06.md --output=Lect-06.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-06.html ${FR}/css/hpost >/tmp/Lect-06.html
	mv /tmp/Lect-06.html ./Lect-06.html

run:
	./run.sh

