
PP=`pwd`
DIR=../../../F20-1015
IMG= go-download.png
PY= file1 file2.txt file3

all: Intro-To-Go.md Intro-To-Go.html

build:
	(cd hello-world; go run hw.go)
	(cd ans; go build)
	(cd echo1; go build)
	(cd echo2; go build)
	(cd ksum; go build)
	(cd mkPkg; go build)
	(cd print1; go build)
	(cd read-json1; go build)


%.md: %.raw.md $(PY) $(IMG)
	m4 -P $< >$@

%.html: %.md
	blackfriday-tool ./$< $@
	echo cat ./${DIR}/md.css $@ >/tmp/$@
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/md.css ./${DIR}/css/hpre $@ ./${DIR}/css/hpost >/tmp/$@
	mv /tmp/$@ ./$@
	echo "file://${PP}/$@" >>open.1

