
PP=`pwd`
FN=Lect-12
DIR=../../../F20-1015
IMG= 
PY= ../../Eth/contracts/ExampleToken.sol

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

