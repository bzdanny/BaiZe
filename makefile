.PHONY: all build run gotool clean help

BINARY="baize"


init:
	go get github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/nicksnyder/go-i18n/v2/goi18n

wire:
	cd app/cmd/ && wire

swag:
	cd app/cmd/ && swag init

build:
	make wire;
	make swag;
	cd app/cmd/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}


help:
	@echo "make - 格式化 Go 代码,并编译成二进制文件"