build:
	go build ./...
	go build -o singularity ./cmd

gen:
	go generate ./replication/internal/types.go

lint:
	gofmt -s -w .
	golangci-lint run
