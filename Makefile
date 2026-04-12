BIN := templer
BUILD_DIR := build

.PHONY: help build test test-all clean

help:
	@echo "make build        - build binary"
	@echo "make release 	 - release version(use for CI)"
	@echo "make test-all     - all tests"
	@echo "make clean        - clean build"

build:
	@scripts/build.sh

release:
	@scripts/release.sh

test-all:
	@scripts/test.sh ./...

clean:
	rm -rf $(BUILD_DIR)