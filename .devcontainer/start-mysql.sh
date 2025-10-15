#!/usr/bin/env bash
set -euo pipefail

# MariaDB server configuration
MYSQL_BASE="${HOME}/.local/share/mysql"
DATA_DIR="${MYSQL_BASE}/data"
SOCKET="${MYSQL_SOCKET:-${MYSQL_BASE}/mysql.sock}"
PID_FILE="${MYSQL_BASE}/mysql.pid"
PORT="${MYSQL_PORT:-3306}"
LOG_FILE="${MYSQL_BASE}/mysql.err"

# Check if already running
if [ -S "${SOCKET}" ] && mariadb-admin --socket="${SOCKET}" ping >/dev/null 2>&1; then
  echo "MySQL already running"
  exit 0
fi

# Start MariaDB server
echo "Starting MySQL server"
touch "${LOG_FILE}"
nohup mariadbd \
  --datadir="${DATA_DIR}" \
  --socket="${SOCKET}" \
  --pid-file="${PID_FILE}" \
  --bind-address=127.0.0.1 \
  --port="${PORT}" \
  --skip-name-resolve \
  --log-error="${LOG_FILE}" \
  >/dev/null 2>&1 &

# Wait for MySQL to be ready
for i in {1..60}; do
  if [ -S "${SOCKET}" ] && grep -q "ready for connections" "${LOG_FILE}" >/dev/null 2>&1; then
    echo "MySQL server is ready"
    exit 0
  fi
  sleep 1
done

echo "MySQL server failed to start"
if [ -f "${LOG_FILE}" ]; then
  echo "--- Begin MariaDB error log ---"
  tail -n 200 "${LOG_FILE}" || true
  echo "--- End MariaDB error log ---"
fi
exit 1