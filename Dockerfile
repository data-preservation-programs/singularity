FROM golang:1.23.6-bullseye as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o singularity .
FROM debian:bullseye-slim
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/singularity /app/singularity
ENTRYPOINT ["/app/singularity"]
