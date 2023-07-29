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
	go test -coverprofile=coverage.out -coverpkg=./... ./...

diagram: build
	./singularity admin init
	schemacrawler.sh --server=sqlite --database=./singularity.db --command=schema --output-format=svg --output-file=docs/database-diagram.svg --info-level=maximum
