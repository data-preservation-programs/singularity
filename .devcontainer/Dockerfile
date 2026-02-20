ARG GO_VERSION=1.24
ARG DEBIAN_CODENAME=trixie

FROM mcr.microsoft.com/devcontainers/go:${GO_VERSION}-${DEBIAN_CODENAME}

ARG PG_MAJOR=16
ARG MARIADB_MAJOR=12.0.2

USER root

# skip man pages and docs for all future installs
RUN printf 'path-exclude=/usr/share/man/*\npath-exclude=/usr/share/doc/*\n' \
    > /etc/dpkg/dpkg.cfg.d/excludes

# add PostgreSQL and MariaDB official repositories
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
       wget ca-certificates gnupg curl \
    && wget -qO- https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg \
    && echo "deb http://apt.postgresql.org/pub/repos/apt $(. /etc/os-release && echo $VERSION_CODENAME)-pgdg main" > /etc/apt/sources.list.d/pgdg.list \
    && curl -LsSO https://r.mariadb.com/downloads/mariadb_repo_setup \
    && chmod +x mariadb_repo_setup \
    && ./mariadb_repo_setup --mariadb-server-version="mariadb-${MARIADB_MAJOR}" --skip-maxscale --skip-tools \
    && rm -f mariadb_repo_setup \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
       postgresql-${PG_MAJOR} postgresql-client-${PG_MAJOR} \
       mariadb-server mariadb-client \
    && rm -rf /var/lib/apt/lists/*

# purge docs and man pages inherited from base image
RUN rm -rf /usr/share/man/* /usr/share/doc/* /usr/share/info/* /usr/share/lintian/* \
    && apt-get purge -y --auto-remove man-db manpages 2>/dev/null || true

# Prepare user-owned data dirs for rootless startup
RUN mkdir -p /home/vscode/.local/share/pg/pgdata \
    && mkdir -p /home/vscode/.local/share/mysql \
    && chown -R vscode:vscode /home/vscode/.local/share

# Set environment variables based on build args
ENV PG_BIN_DIR=/usr/lib/postgresql/${PG_MAJOR}/bin

USER vscode

