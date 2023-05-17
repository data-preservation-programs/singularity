build:
	go build -o singularity ./cmd/singularity

buildall:
	go build ./...

gen:
	go generate ./replication/internal/proposal110/types.go
	go generate ./replication/internal/proposal120/types.go

lint:
	gofmt -s -w .
	golangci-lint run

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d ./cmd,./api,./handler -o ./api/docs

