
PP=`pwd`
FN=Lect-04
DIR=../../../F20-1015
IMG= IncomeForecast1.pdf IncomeForecast1.png Lect-04.html.pdf little-endian.png
PY= Lect-04_00.go Lect-04_01.go h1 h2 war-and-peace.txt

all: gen val ${FN}.md ${FN}.html 

gen:
	-go run Lect-04_00.go >Lect_04_02.txt 2>&1

val:
	go build Lect-04_01.go

%.md: %.raw.md $(PY) $(IMG) Makefile
	m4 -P $< >$@

%.html: %.md
	blackfriday-tool ./$< $@
	echo cat ./${DIR}/md.css $@ >/tmp/$@
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/md.css ./${DIR}/css/hpre $@ ./${DIR}/css/hpost >/tmp/$@
	mv /tmp/$@ ./$@
	echo "file://${PP}/$@" >>open.1

