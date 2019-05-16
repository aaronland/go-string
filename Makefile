vendor-deps:
	go mod vendor

fmt:
	go fmt cmd/*.go
	go fmt dsn/*.go
	go fmt random/*.go
	go fmt *.go

tools: 	
	@GOPATH=$(GOPATH) go build -o bin/randomstr cmd/randomstr/main.go
