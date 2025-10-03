#!/usr/bin/env bash
set -euo pipefail

# MariaDB client is required for init

# Resolve socket path; default to user-owned socket
SOCKET="${MYSQL_SOCKET:-${HOME}/.local/share/mysql/mysql.sock}"

## Removed one-time init guard; operations below are idempotent

# Determine root auth flags
MYSQL_ROOT_FLAGS=("-uroot")
if [ -n "${MYSQL_ROOT_PASSWORD:-}" ]; then
  MYSQL_ROOT_FLAGS+=("-p${MYSQL_ROOT_PASSWORD}")
fi

# Wait for server readiness (best effort)
echo "Waiting for MySQL server at socket: $SOCKET"
for i in {1..60}; do
  if mariadb-admin --socket="$SOCKET" ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
    echo "MySQL server is ready for init"
    break
  fi
  sleep 1
done

# Bail if still unreachable
if ! mariadb-admin --socket="$SOCKET" ping "${MYSQL_ROOT_FLAGS[@]}" >/dev/null 2>&1; then
  echo "MySQL server not reachable, init failed"
  exit 1
fi

# Create database and user idempotently (MySQL 8+ supports IF NOT EXISTS for users)
DB=${MYSQL_DATABASE:-singularity}
USER=${MYSQL_USER:-singularity}
PASS=${MYSQL_PASSWORD:-singularity}

echo "Creating database and user: ${USER}@localhost and ${USER}@%"
mariadb --socket="$SOCKET" "${MYSQL_ROOT_FLAGS[@]}" <<SQL
CREATE DATABASE IF NOT EXISTS \`${DB}\`;
CREATE USER IF NOT EXISTS '${USER}'@'localhost' IDENTIFIED BY '${PASS}';
CREATE USER IF NOT EXISTS '${USER}'@'%' IDENTIFIED BY '${PASS}';
-- tests create per-run databases, so grant global privileges for dev only
GRANT ALL PRIVILEGES ON *.* TO '${USER}'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO '${USER}'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON \`${DB}\`.* TO '${USER}'@'localhost';
GRANT ALL PRIVILEGES ON \`${DB}\`.* TO '${USER}'@'%';
FLUSH PRIVILEGES;
SQL

echo "MySQL init completed successfully"

## No marker file; script can run safely on every start


