
all: Intro-To-Go.html

FR=../../Lectures/01

Intro-To-Go.html: Intro-To-Go.md
	markdown-cli --input=./Intro-To-Go.md --output=Intro-To-Go.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Intro-To-Go.html ${FR}/css/hpost >/tmp/Intro-To-Go.html
	mv /tmp/Intro-To-Go.html ./Intro-To-Go.html

test: test01 test02
.PHONEY: test test01 test02

test01:
	( cd hello-world ; go build ; ./hello-world ) >,a
	diff ,a ref/hw.out

test02:
	( cd echo1 ; go build ; ./echo1 a b c X ) >,b
	diff ,b ref/echo1.out

clean:
	rm ,a ,b echo1/echo1 hello-world/hello-world ksum/ksum


