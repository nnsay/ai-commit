build:
	go build -o ./bin

buildlinux:
	GOOS=linux GOARCH=arm64 go build -o ./bin/linux-arm64-ai-commit

buildmac:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin-arm64-ai-commit

release: buildlinux buildmac
	cd ./bin; \
	shasum -a 256 linux-arm64-ai-commit darwin-arm64-ai-commit > checksumfile; \
	shasum -a 256 -c checksumfile; \
	cat checksumfile; \
	cd -

zip:
	cd ./bin; \
	zip relelase.zip checksumfile darwin-arm64-ai-commit linux-arm64-ai-commit; \
	cd -

install: build
	mv bin/ai-commit ~/.local/bin

clear:
	rm -rf bin/*

.PHONY: release
