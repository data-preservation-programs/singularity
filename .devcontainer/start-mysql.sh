#!/usr/bin/env bash
set -euo pipefail

# Paths (homebrew mysql in devcontainer)
MYSQLD="${MYSQLD:-/home/linuxbrew/.linuxbrew/opt/mysql/bin/mysqld}"
MYSQLADMIN="${MYSQLADMIN:-/home/linuxbrew/.linuxbrew/bin/mysqladmin}"

# Runtime locations under /tmp so tests are ephemeral and permissions are simple
DATADIR="${MYSQL_DATADIR:-/tmp/singularity-mysql-data}"
SOCKET="${MYSQL_SOCKET:-/tmp/singularity-mysql.sock}"
PORT="${MYSQL_PORT:-3306}"
PIDFILE="${MYSQL_PIDFILE:-/tmp/singularity-mysql.pid}"
LOGFILE="${MYSQL_LOGFILE:-/tmp/singularity-mysql.log}"

if ! command -v "$MYSQLD" >/dev/null 2>&1; then
  echo "mysqld not found at $MYSQLD; skipping mysql start"
  exit 0
fi

mkdir -p "$(dirname "$SOCKET")" "$DATADIR" "$(dirname "$PIDFILE")" "$(dirname "$LOGFILE")"

# Initialize data directory if empty (insecure root, no password)
if [ -z "$(ls -A "$DATADIR" 2>/dev/null || true)" ]; then
  "$MYSQLD" --initialize-insecure --datadir="$DATADIR" --user="$(id -un)" >/dev/null 2>&1 || true
fi

# Already up?
if "$MYSQLADMIN" --socket="$SOCKET" --user=root ping >/dev/null 2>&1; then
  exit 0
fi

# Start daemonized bound to 127.0.0.1
nohup "$MYSQLD" \
  --datadir="$DATADIR" \
  --socket="$SOCKET" \
  --port="$PORT" \
  --bind-address=127.0.0.1 \
  --pid-file="$PIDFILE" \
  --log-error="$LOGFILE" \
  --user="$(id -un)" \
  --daemonize >/dev/null 2>&1 || true

# Wait up to 60s for readiness
for i in {1..60}; do
  if "$MYSQLADMIN" --socket="$SOCKET" --user=root ping >/dev/null 2>&1; then
    exit 0
  fi
  sleep 1
done

echo "mysqld failed to start; see $LOGFILE"
exit 0


