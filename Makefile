build:
	go build -o singularity ./cmd

gen:
	go generate ./replication/internal/types.go
