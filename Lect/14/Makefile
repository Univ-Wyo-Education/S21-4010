
all: Lect-14.html note.2.html

FR=../01/

Lect-14.html: Lect-14.md
	markdown-cli --input=./Lect-14.md --output=Lect-14.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-14.html ${FR}/css/hpost >/tmp/Lect-14.html
	mv /tmp/Lect-14.html ./Lect-14.html

note.2.html: note.2.md
	markdown-cli --input=./note.2.md --output=note.2.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre note.2.html ${FR}/css/hpost >/tmp/note.2.html
	mv /tmp/note.2.html ./note.2.html

