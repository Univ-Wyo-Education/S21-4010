
all:
	go build

.PHONY: test test001 test_sf1

test: test001 test_sf0 test_sf1

T1=8ff028fcee4fce256d8dd98bfb5184fa0e49ef9c582bbd72f3389a13dc5db899.json

test001:
	go build
	mkdir -p data
	rm -f ./data/*.json
	rm -f data/${T1}
	rm -f data/index.json
	rm -f data/TestWriteBlock.json
	./main --cfg test-genesis.json --create-genesis
	diff data/index.json ref
	diff data/${T1} ref
	rm -f ./data/*.json
	rm -f data/${T1}
	rm -f data/index.json
	rm -f data/TestWriteBlock.json
	./main --cfg test-genesis.json --create-genesis
	./main --test-read-block
	./main --test-write-block
	diff data/TestWriteBlock.json ref
	@echo 
	@echo PASS

# Setup for 2nd test - transfer between accounts
test_sf0:
	go build
	mkdir -p data
	rm -f ./data/*.json
	mkdir -p ./out
	./main --create-genesis

# 2nd Test
# You can run this test multiple times after 'test_sf0' has setup the ./data directory.
test_sf1:
	go build
	mkdir -p ./out
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 >./out/bal885.bef.out
	./main --show-balance 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 >./out/bal6E0.bef.out
	./main --test-send-funds 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 33 x x x "Tx Memo"
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 >./out/bal885.aft.out
	./main --show-balance 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 >./out/bal6E0.aft.out
	./validateBalance.sh ./out/bal885.bef.out -33 ./out/bal885.aft.out
	./validateBalance.sh ./out/bal6E0.bef.out  33 ./out/bal6E0.aft.out
	@echo 
	@echo PASS

x:
	./validateBalance.sh ./out/bal885.bef.out -33 ./out/bal885.aft.out

y:
	./validateBalance.sh ./out/bal6E0.bef.out  33 ./out/bal6E0.aft.out

test_create_genesis:
	go build
	rm -f ./data/*.json
	./main --create-genesis

test_list_acct:
	go build
	rm -f ./data/*.json
	./main --create-genesis
	./main --list-accounts

test_show_balance:
	go build
	rm -f ./data/*.json
	./main --create-genesis
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 

test_send_funds:
	go build
	rm -f ./data/*.json
	mkdir -p ./out
	./main --create-genesis
	./main --test-send-funds 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 33 x x x "Memo 02"
	./main --show-balance 0xdb180da9a8982c7bb75ca40039f959cb959c62e8 >./out/tsf.out
	./validateBal2.sh ./out/tsf.out 500200
	@echo 
	@echo PASS

w:
	go build
	./main --show-balance 0xdb180da9a8982c7bb75ca40039f959cb959c62e8
