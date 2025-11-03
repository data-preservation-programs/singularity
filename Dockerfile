FROM golang:1.24.9-bookworm AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o singularity .
FROM gcr.io/distroless/cc-debian12

COPY --from=builder /app/singularity /app/singularity
ENTRYPOINT ["/app/singularity"]
