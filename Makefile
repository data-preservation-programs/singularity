build:
	go build -o singularity ./cmd

buildall:
	go build ./...

gen:
	go generate ./replication/internal/proposal120/types.go

lint:
	gofmt -s -w .
	golangci-lint run

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d ./cmd,./api,./handler -o ./api/docs

test:
	rm -f ~/.singularity/singularity.db
	make build
	./singularity init
	./singularity dataset create -m 8MiB -M 10MiB test
	#./singularity dataset add-source --s3-region us-west-2 test s3://public-dataset-test
	#./singularity dataset add-source test /Users/xinanxu/test
	./singularity dataset add-source test http://127.0.0.1
	./singularity run dataset-worker
