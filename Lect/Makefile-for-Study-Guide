
all: Lect-15.html 

FR=../01/

Lect-15.html: Lect-15.md
	markdown-cli --input=./Lect-15.md --output=Lect-15.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-15.html ${FR}/css/hpost >/tmp/Lect-15.html
	mv /tmp/Lect-15.html ./Lect-15.html

S19-Study-Guide.html: S19-Study-Guide.md
	markdown-cli --input=./S19-Study-Guide.md --output=S19-Study-Guide.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre S19-Study-Guide.html ${FR}/css/hpost >/tmp/S19-Study-Guide.html
	mv /tmp/S19-Study-Guide.html ./S19-Study-Guide.html

