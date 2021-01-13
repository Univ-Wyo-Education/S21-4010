
all: Lect-02.html

FR=../01/

Lect-02.html: Lect-02.md
	markdown-cli --input=./Lect-02.md --output=Lect-02.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-02.html ${FR}/css/hpost >/tmp/Lect-02.html
	mv /tmp/Lect-02.html ./Lect-02.html

