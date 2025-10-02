#!/usr/bin/env bash
set -euo pipefail

PGDATA_DIR="${PGDATA:-/home/vscode/.local/share/pg/pgdata}"
LOG_FILE="${PGDATA_DIR}/postgres.log"
PGSOCK_DIR="${PGSOCK_DIR:-/home/vscode/.local/share/pg}"
PGPORT="${PGPORT:-55432}"
PG_BIN_DIR="${PG_BIN_DIR:-/usr/lib/postgresql/16/bin}"

mkdir -p "$PGDATA_DIR" "$PGSOCK_DIR"

if [ ! -f "$PGDATA_DIR/PG_VERSION" ]; then
  "${PG_BIN_DIR}/initdb" -D "$PGDATA_DIR" --auth trust --auth-local trust --auth-host trust
  echo "listen_addresses='localhost'" >> "$PGDATA_DIR/postgresql.conf"
  {
    echo 'host    all all 127.0.0.1/32     trust'
    echo 'host    all all ::1/128          trust'
    echo 'local   all all                  trust'
  } >> "$PGDATA_DIR/pg_hba.conf"
fi

"${PG_BIN_DIR}/pg_ctl" -D "$PGDATA_DIR" -l "$LOG_FILE" -w -o "-p ${PGPORT} -k ${PGSOCK_DIR}" start


