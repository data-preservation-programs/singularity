# runtime base is overridable so container-publish can also build a `-debug`
# image off the distroless :debug tag (busybox shell at /busybox/sh) for exec-in
# debugging of the otherwise shell-less distroless runtime.
ARG RUNTIME_BASE=gcr.io/distroless/cc-debian12

FROM golang:1.26.5-bookworm AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o singularity .

FROM ${RUNTIME_BASE}

COPY --from=builder /app/singularity /app/singularity
ENTRYPOINT ["/app/singularity"]
