
PP=`pwd`
FN=Lect-09
DIR=../../../F20-1015
IMG= 1920px-EllipticCurveCatalog.svg.png EC-gb-marked.svg EC_good_bad.svg EllipticCurveCatalog.svg ec001.png ec002.png ec003.png ec004.png ec005.png ec006.png ec007.png ec008.png ec009.png ec010.png ec011.png ec012.png ec013.png ec014.png ec015.png ec016.png ec017.png
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

