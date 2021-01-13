
all: Homework-06.html 

FR=../../Lectures/01

Homework-06.html: Homework-06.md
	markdown-cli --input=./Homework-06.md --output=Homework-06.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Homework-06.html ${FR}/css/hpost >/tmp/Homework-06.html
	mv /tmp/Homework-06.html ./Homework-06.html


