
FN=syllabus
FN1=syllabus-5010
DIR=.

all: ${FN}.html ${FN1}.html 

${FN}.html: ${FN}.md
	blackfriday-tool ./${FN}.md ${FN}.html
	echo cat ./${DIR}/css/md.css ${FN}.html >/tmp/${FN}.html
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/css/md.css ./${DIR}/css/hpre ${FN}.html ./${DIR}/css/hpost >/tmp/${FN}.html
	mv /tmp/${FN}.html ./${FN}.html


${FN1}.html: ${FN1}.md
	blackfriday-tool ./${FN1}.md ${FN1}.html
	echo cat ./${DIR}/css/md.css ${FN1}.html >/tmp/${FN1}.html
	cat ./${DIR}/css/pre ./${DIR}/css/markdown.css ./${DIR}/css/post ./${DIR}/css/md.css ./${DIR}/css/hpre ${FN1}.html ./${DIR}/css/hpost >/tmp/${FN1}.html
	mv /tmp/${FN1}.html ./${FN1}.html


