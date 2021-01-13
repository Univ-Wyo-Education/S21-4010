
all: Merkle.html build

build:
	( cd merkle ; go test )

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
Merkle.html: Merkle.md
	markdown-cli --input=./Merkle.md --output=Merkle.html

