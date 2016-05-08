.PHONY: all build linux release

all: build

build:
	go build server.go translate.go parse.go

linux: *.go
	GOOS=linux GOARCH=amd64 go build server.go translate.go parse.go

release: linux
	docker build -t rjocoleman/redirect.name .
	docker push rjocoleman/redirect.name
