VERSION=$(shell git describe --tags 2>/dev/null|| echo "$$(git rev-parse HEAD)@$$(git rev-parse --abbrev-ref HEAD)")
EXE_PATH=./bin/newpwd

.PHONY: build test version
build:
	bash build.sh $(VERSION)
test:
	go test ./...
version:
	@echo $(VERSION)