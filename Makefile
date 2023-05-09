build:
	go build -o singularity ./cmd

buildall:
	go build ./...

gen:
	go generate ./replication/internal/types.go

lint:
	gofmt -s -w .
	golangci-lint run

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d ./cmd,./api,./handler -o ./api/docs
