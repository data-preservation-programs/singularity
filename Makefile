build:
	go build -o singularity .

buildall:
	go build ./...

gen:
	go generate ./replication/internal/proposal110/types.go
	go generate ./replication/internal/proposal120/types.go
	go run handler/datasource/generate/add.go

lint:
	gofmt -s -w .
	golangci-lint run

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./api/docs

test:
	gotestsum --format testname -- -coverprofile=coverage.out -coverpkg=./... ./...

gendoc:
	go run ./docgen/clireference/main.go

genwebdoc:
	go run ./docgen/webapireference/main.go

translate:
	go run ./docgen/translate/main.go
