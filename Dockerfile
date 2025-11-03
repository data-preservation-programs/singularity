FROM golang:1.24.9-bookworm AS builder
WORKDIR /app

# Copy dependency files first (changes infrequently)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code after (changes frequently)
COPY . ./

# Build binary
RUN go build -o singularity .

FROM gcr.io/distroless/cc-debian12

COPY --from=builder /app/singularity /app/singularity
ENTRYPOINT ["/app/singularity"]
