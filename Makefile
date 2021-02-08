CURRENT_PATH := $(shell pwd)

run:
	go get github.com/jschneider98/jgovalidator
	go get github.com/jschneider98/jgoweb
	go build -o app
	./app

unit-test:
	go test -v -tags=unit -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

int-test:
	go test -v -tags=integration -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

test:
	go test -v -tags="integration unit" -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

build-linux:
	GOOS=linux GOARCH=386 go build -o app

