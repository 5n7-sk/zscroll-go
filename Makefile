.PHONY: build
build:
	go build -o bin/zscroll cmd/zscroll/main.go

.PHONY: format
format:
	goimports -w .

.PHONY: lint
lint:
	golangci-lint run --tests ./...

.PHONY: test
test:
	go test -v ./...
