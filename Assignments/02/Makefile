
PP=`pwd`
FN=syllabus
FN1=syllabus-5010
DIR=../../../F20-1015
IMG=
PY=


all: Hash.html build

build:
	( cd main ; go build )


#Hash.html: Hash.md
#	markdown-cli --input=./Hash.md --output=Hash.html
#	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Hash.html ${FR}/css/hpost >/tmp/Hash.html
#	mv /tmp/Hash.html ./Hash.html


%.md: %.raw.md $(PY) $(IMG)
	m4 -P $< >$@

%.html: %.md
	blackfriday-tool ./$< $@
	echo cat ./${DIR}/md.css $@ >/tmp/$@
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/md.css ./${DIR}/css/hpre $@ ./${DIR}/css/hpost >/tmp/$@
	mv /tmp/$@ ./$@
	echo "file://${PP}/$@" >>open.1
