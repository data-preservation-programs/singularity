ARG GO_VERSION=1.25
ARG DEBIAN_CODENAME=trixie

FROM mcr.microsoft.com/devcontainers/go:${GO_VERSION}-${DEBIAN_CODENAME}

ARG PG_MAJOR=16

USER root

# ensure /tmp is writable for apt GPG verification under rootless podman
RUN chmod 1777 /tmp

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
       wget ca-certificates gnupg curl \
    && wget -qO- https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg \
    && echo "deb http://apt.postgresql.org/pub/repos/apt $(. /etc/os-release && echo $VERSION_CODENAME)-pgdg main" > /etc/apt/sources.list.d/pgdg.list \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
       postgresql-${PG_MAJOR} postgresql-client-${PG_MAJOR} \
    && rm -rf /var/lib/apt/lists/*

COPY --from=ghcr.io/foundry-rs/foundry:latest /usr/local/bin/anvil /usr/local/bin/anvil
# forge is needed by `go generate ./...` via util/testutil/sol/Makefile
# (fvm_precompiles.go go:generate rebuilds the mock bytecode .txt files)
COPY --from=ghcr.io/foundry-rs/foundry:latest /usr/local/bin/forge /usr/local/bin/forge

RUN mkdir -p /home/vscode/.local/share/pg/pgdata \
    && mkdir -p /home/vscode/.local/share/tmp \
    && chown -R vscode:vscode /home/vscode/.local/share

ENV PG_BIN_DIR=/usr/lib/postgresql/${PG_MAJOR}/bin

USER vscode
