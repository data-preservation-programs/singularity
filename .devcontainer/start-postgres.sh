#!/usr/bin/env bash
set -euo pipefail

# Postgres server configuration
PGDATA_DIR="${PGDATA:-/home/vscode/.local/share/pg/pgdata}"
LOG_FILE="${PGDATA_DIR}/postgres.log"
PGSOCK_DIR="${PGSOCK_DIR:-/home/vscode/.local/share/pg}"
PGPORT="${PGPORT:-55432}"
PG_BIN_DIR="${PG_BIN_DIR:-/usr/lib/postgresql/16/bin}"

# Check if already running
if "${PG_BIN_DIR}/pg_ctl" -D "$PGDATA_DIR" status >/dev/null 2>&1; then
  echo "Postgres already running"
  exit 0
fi

# Start Postgres server
echo "Starting Postgres server"
"${PG_BIN_DIR}/pg_ctl" -D "$PGDATA_DIR" -l "$LOG_FILE" -w -o "-p ${PGPORT} -k ${PGSOCK_DIR}" start


