#!/usr/bin/env bash
set -euo pipefail

# Start MariaDB/MySQL via system service

# If already running, exit
if pgrep -x mysqld >/dev/null 2>&1; then
  echo "MySQL already running"
  exit 0
fi

echo "Starting MySQL server..."
sudo service mysql start >/dev/null 2>&1 || sudo service mariadb start >/dev/null 2>&1 || true

# Debian socket path
SOCKET="/var/run/mysqld/mysqld.sock"

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