#!/usr/bin/env bash
set -euo pipefail

# If mysql client isn't present, exit quietly
if ! command -v mysql >/dev/null 2>&1; then
  exit 0
fi

# Determine root auth flags
MYSQL_ROOT_FLAGS=("-uroot")
if [ -n "${MYSQL_ROOT_PASSWORD:-}" ]; then
  MYSQL_ROOT_FLAGS+=("-p${MYSQL_ROOT_PASSWORD}")
fi

# Wait for server readiness (best effort)
for i in {1..60}; do
  if mysqladmin ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

# Bail if still unreachable
if ! mysqladmin ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
  exit 0
fi

# Create database and user idempotently (MySQL 8+ supports IF NOT EXISTS for users)
DB=${MYSQL_DATABASE:-singularity}
USER=${MYSQL_USER:-singularity}
PASS=${MYSQL_PASSWORD:-singularity}

mysql "${MYSQL_ROOT_FLAGS[@]}" <<SQL
CREATE DATABASE IF NOT EXISTS \`${DB}\`;
CREATE USER IF NOT EXISTS '${USER}'@'%' IDENTIFIED BY '${PASS}';
GRANT ALL PRIVILEGES ON \`${DB}\`.* TO '${USER}'@'%';
FLUSH PRIVILEGES;
SQL


