
all: Mining.html build

build:
	( cd main ; go build )

# You can get markdown-cli and compile it yourself.
#
# $ cd ~/go/src/github.com/
# $ mkdir pschlump
# $ cd pschlump
# $ git pull https://github.com/pschlump/markdown-cli.git
# $ go get
# $ cd markdown-cli
# $ go build
#
FR=../../Lectures/01

Mining.html: Mining.md
	markdown-cli --input=./Mining.md --output=Mining.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Mining.html ${FR}/css/hpost >/tmp/Mining.html
	mv /tmp/Mining.html ./Mining.html

