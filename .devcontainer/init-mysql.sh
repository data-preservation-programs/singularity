#!/usr/bin/env bash
set -euo pipefail

# If mysql client isn't present, exit quietly
if ! command -v mysql >/dev/null 2>&1; then
  exit 0
fi

# Use same socket as mysqld starter
SOCKET="${MYSQL_SOCKET:-/tmp/singularity-mysql.sock}"

# One-time init guard
MARKER="/workspace/.devcontainer/.mysql-initialized"
if [ -f "$MARKER" ]; then
  exit 0
fi

# Determine root auth flags
MYSQL_ROOT_FLAGS=("-uroot")
if [ -n "${MYSQL_ROOT_PASSWORD:-}" ]; then
  MYSQL_ROOT_FLAGS+=("-p${MYSQL_ROOT_PASSWORD}")
fi

# Wait for server readiness (best effort)
for i in {1..60}; do
  if mysqladmin --socket="$SOCKET" ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

# Bail if still unreachable
if ! mysqladmin --socket="$SOCKET" ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
  exit 0
fi

# Create database and user idempotently (MySQL 8+ supports IF NOT EXISTS for users)
DB=${MYSQL_DATABASE:-singularity}
USER=${MYSQL_USER:-singularity}
PASS=${MYSQL_PASSWORD:-singularity}

mysql --socket="$SOCKET" "${MYSQL_ROOT_FLAGS[@]}" <<SQL
CREATE DATABASE IF NOT EXISTS \`${DB}\`;
CREATE USER IF NOT EXISTS '${USER}'@'%' IDENTIFIED BY '${PASS}';
-- tests create per-run databases, so grant global privileges for dev only
GRANT ALL PRIVILEGES ON *.* TO '${USER}'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON \`${DB}\`.* TO '${USER}'@'%';
FLUSH PRIVILEGES;
SQL

# mark initialized
touch "$MARKER"


