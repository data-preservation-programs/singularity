#!/usr/bin/env bash
set -euo pipefail

# Initialize Go modules
go mod download

if ! command -v mysqld >/dev/null 2>&1; then
  sudo apt-get update
  sudo DEBIAN_FRONTEND=noninteractive apt-get install -y mariadb-server mariadb-client >/dev/null
  sudo mkdir -p /var/run/mysqld
  sudo chown -R mysql:mysql /var/run/mysqld || true
  sudo chmod 775 /var/run/mysqld || true
fi


