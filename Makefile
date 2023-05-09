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

test:
	rm -f ~/.singularity/singularity.db
	make build
	./singularity init
	./singularity dataset create -m 80MiB -M 100MiB test
	./singularity dataset add-source --s3-region us-west-2 test s3://public-dataset-test/subfolder/15M.random
	#./singularity dataset add-source test /mnt/e/test
	#./singularity run dataset-worker
