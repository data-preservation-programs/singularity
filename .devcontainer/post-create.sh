#!/usr/bin/env bash
set -euo pipefail

# Initialize Go modules
go mod download

chmod +x .devcontainer/*.sh || true

echo "Setting up databases..."

mkdir -p /home/vscode/.local/share/pg/pgdata /home/vscode/.local/share/pg /home/vscode/.local/share/tmp

# Initialize Postgres
if [ ! -f "/home/vscode/.local/share/pg/pgdata/PG_VERSION" ]; then
  echo "Initializing Postgres..."
  /usr/lib/postgresql/16/bin/initdb -D /home/vscode/.local/share/pg/pgdata --auth trust --auth-local trust --auth-host trust --encoding=UTF8 --locale=C.UTF-8
  echo "listen_addresses='localhost'" >> /home/vscode/.local/share/pg/pgdata/postgresql.conf
  {
    echo 'host    all all 127.0.0.1/32     trust'
    echo 'host    all all ::1/128          trust'
    echo 'local   all all                  trust'
  } >> /home/vscode/.local/share/pg/pgdata/pg_hba.conf
fi

echo "Starting database server..."
.devcontainer/start-postgres.sh

echo "Creating database users..."
psql -h localhost -p 55432 -d postgres -c "CREATE USER singularity WITH SUPERUSER CREATEDB CREATEROLE LOGIN;"

echo "Database setup completed successfully"
