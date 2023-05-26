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

remote:
	go build -o testremote ./datasource/cmd

test:
	make build
	./singularity admin reset
	./singularity dataset create -m 1.5GB -M 2GB test
	./singularity datasource add local test ~/test/
	GOLOG_LOG_LEVEL=debug ./singularity run dataset-worker --exit-on-complete=true
