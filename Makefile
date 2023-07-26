build:
	go build -o singularity .

buildall:
	go build ./...

generate:
	go generate ./...

lint:
	gofmt -s -w .
	golangci-lint run --fix
	staticcheck ./...

test:
	go test -race -coverprofile=coverage.out -coverpkg=./... ./...
