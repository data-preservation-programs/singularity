#!/usr/bin/env bash
set -euo pipefail

PGDATA_DIR="${PGDATA:-${PWD}/.devcontainer/pgdata}"
LOG_FILE="${PGDATA_DIR}/postgres.log"

sudo mkdir -p "$PGDATA_DIR"
sudo chown -R postgres:postgres "$PGDATA_DIR"

if [ ! -f "$PGDATA_DIR/PG_VERSION" ]; then
  sudo -u postgres /usr/lib/postgresql/16/bin/initdb -D "$PGDATA_DIR" --auth trust --auth-local trust --auth-host trust
  sudo -u postgres bash -lc "echo \"listen_addresses='*'\" >> '$PGDATA_DIR/postgresql.conf'"
  sudo -u postgres bash -lc "{
    echo 'host    all all 0.0.0.0/0        trust'
    echo 'host    all all ::/0             trust'
    echo 'host    all all ::1/128          trust'
    echo 'local   all all                  trust'
  } >> '$PGDATA_DIR/pg_hba.conf'"
fi

sudo -u postgres /usr/lib/postgresql/16/bin/pg_ctl -D "$PGDATA_DIR" -l "$LOG_FILE" -w start


