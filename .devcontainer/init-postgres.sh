#!/usr/bin/env bash
set -euo pipefail

# If postgres isn't present in this image, exit quietly
if ! command -v pg_isready >/dev/null 2>&1; then
  exit 0
fi

# Wait for server readiness (best effort)
for i in {1..60}; do
  if pg_isready -q; then
    break
  fi
  sleep 1
done

# If we still can't connect, bail (container may not include a running server)
if ! psql -U postgres -d postgres -tc "SELECT 1" >/dev/null 2>&1; then
  exit 0
fi

# Create role if missing (idempotent)
psql -U postgres -d postgres -v ON_ERROR_STOP=1 -c "DO $$ BEGIN IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'singularity') THEN CREATE ROLE singularity WITH LOGIN SUPERUSER PASSWORD 'singularity'; END IF; END $$;"

# Create database if missing (idempotent)
if ! psql -U postgres -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname = 'singularity'" | grep -q 1; then
  createdb -U postgres singularity
fi


