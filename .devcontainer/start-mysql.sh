#!/usr/bin/env bash
set -euo pipefail

# Start MySQL using homebrew's mysql.server command
# This is provided by the ghcr.io/devcontainers-extra/features/mysql-homebrew feature

# Check if already running
if mysql.server status >/dev/null 2>&1; then
  echo "MySQL already running"
  exit 0
fi

echo "Starting MySQL server..."
mysql.server start

# MySQL uses /tmp/mysql.sock by default when started with mysql.server
SOCKET="/tmp/mysql.sock"

# Wait for MySQL to be ready
for i in {1..60}; do
  if mysqladmin --socket="$SOCKET" --user=root ping >/dev/null 2>&1; then
    echo "MySQL server is ready"
    exit 0
  fi
  sleep 1
done

echo "MySQL server failed to start"
exit 1