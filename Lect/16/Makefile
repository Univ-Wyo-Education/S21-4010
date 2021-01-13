
all: Lect-16.html S20-Study-Guide.html backup.html

FR=../01/

Lect-16.html: Lect-16.md
	markdown-cli --input=./Lect-16.md --output=Lect-16.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-16.html ${FR}/css/hpost >/tmp/Lect-16.html
	mv /tmp/Lect-16.html ./Lect-16.html


S20-Study-Guide.html: S20-Study-Guide.md
	markdown-cli --input=./S20-Study-Guide.md --output=S20-Study-Guide.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre S20-Study-Guide.html ${FR}/css/hpost >/tmp/S20-Study-Guide.html
	mv /tmp/S20-Study-Guide.html ./S20-Study-Guide.html

backup.html: backup.md
	markdown-cli --input=./backup.md --output=backup.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre backup.html ${FR}/css/hpost >/tmp/backup.html
	mv /tmp/backup.html ./backup.html

