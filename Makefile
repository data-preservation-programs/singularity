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

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./api/docs

test:
	go test -coverprofile=coverage.out -coverpkg=./... ./...

model2ts:
	go run ./buildtool/model2ts/main.go
