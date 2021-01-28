
PP=`pwd`
FN=Lect-06
DIR=../../../F20-1015
IMG= 1920px-Hash_Tree.svg.png Hash_Tree.svg.png
PY= 

all: gen val ${FN}.md ${FN}.html 

gen:

val:

%.md: %.raw.md $(PY) $(IMG) Makefile
	m4 -P $< >$@

%.html: %.md
	blackfriday-tool ./$< $@
	echo cat ./${DIR}/md.css $@ >/tmp/$@
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/md.css ./${DIR}/css/hpre $@ ./${DIR}/css/hpost >/tmp/$@
	mv /tmp/$@ ./$@
	echo "file://${PP}/$@" >>open.1

