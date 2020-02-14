.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./application/application

build:
	GOOS=linux GOARCH=amd64 go build -o application/application ./application 
