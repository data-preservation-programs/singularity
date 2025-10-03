#!/usr/bin/env bash
set -euo pipefail

# Start MariaDB as an unprivileged user using a user-owned data directory

MYSQL_BASE="${HOME}/.local/share/mysql"
DATA_DIR="${MYSQL_BASE}/data"
SOCKET="${MYSQL_SOCKET:-${MYSQL_BASE}/mysql.sock}"
PID_FILE="${MYSQL_BASE}/mysql.pid"
PORT="${MYSQL_PORT:-3306}"
LOG_FILE="${MYSQL_BASE}/mysql.err"

mkdir -p "${DATA_DIR}" "${MYSQL_BASE}"

# Initialize data dir if missing
if [ ! -d "${DATA_DIR}/mysql" ]; then
  echo "Initializing MariaDB data directory"
  mariadb-install-db --datadir="${DATA_DIR}" --auth-root-authentication-method=normal --skip-test-db >/dev/null
fi

# If already running, exit
if [ -S "${SOCKET}" ] && mysqladmin --socket="${SOCKET}" ping >/dev/null 2>&1; then
  echo "MySQL already running"
  exit 0
fi

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

# Wait for MySQL to be ready (log-based + socket existence)
for i in {1..60}; do
  if [ -S "${SOCKET}" ] && grep -q "ready for connections" "${LOG_FILE}" >/dev/null 2>&1; then
    echo "MySQL server is ready"
    echo "Socket exists at: ${SOCKET}"
    echo "Continuing to init script..."
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