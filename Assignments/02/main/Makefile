
all:
	go build

test:
	go build
	mkdir -p data
	rm -f data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json 
	rm -f data/index.json
	rm -f data/TestWriteBlock.json
	./main --create-genesis
	diff data/index.json ref
	diff data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json ref
	./main --test-read-block
	./main --test-write-block
	diff data/TestWriteBlock.json ref
	@echo 
	@echo PASS

