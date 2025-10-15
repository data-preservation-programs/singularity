#!/usr/bin/env bash
set -euo pipefail

# Initialize Go modules
go mod download

chmod +x .devcontainer/*.sh || true

# Initialize database systems
echo "Setting up databases..."

# Create data directories
mkdir -p /home/vscode/.local/share/pg/pgdata /home/vscode/.local/share/pg /home/vscode/.local/share/mysql/data /home/vscode/.local/share/mysql

# Initialize Postgres
if [ ! -f "/home/vscode/.local/share/pg/pgdata/PG_VERSION" ]; then
  echo "Initializing Postgres..."
  /usr/lib/postgresql/16/bin/initdb -D /home/vscode/.local/share/pg/pgdata --auth trust --auth-local trust --auth-host trust
  echo "listen_addresses='localhost'" >> /home/vscode/.local/share/pg/pgdata/postgresql.conf
  {
    echo 'host    all all 127.0.0.1/32     trust'
    echo 'host    all all ::1/128          trust'
    echo 'local   all all                  trust'
  } >> /home/vscode/.local/share/pg/pgdata/pg_hba.conf
fi

# Initialize MariaDB
if [ ! -d "/home/vscode/.local/share/mysql/data/mysql" ]; then
  echo "Initializing MariaDB..."
  mariadb-install-db --datadir=/home/vscode/.local/share/mysql/data --auth-root-authentication-method=normal --skip-test-db >/dev/null
fi

# Start both servers
echo "Starting database servers..."
.devcontainer/start-postgres.sh
.devcontainer/start-mysql.sh

# Create users (databases will be created during testing as needed)
echo "Creating database users..."

# Postgres setup
psql -h localhost -p 55432 -d postgres -c "CREATE USER singularity WITH SUPERUSER CREATEDB CREATEROLE LOGIN;"

# MySQL setup  
mariadb --socket=/home/vscode/.local/share/mysql/mysql.sock -uroot <<SQL
CREATE USER 'singularity'@'localhost' IDENTIFIED BY 'singularity';
CREATE USER 'singularity'@'%' IDENTIFIED BY 'singularity';
GRANT ALL PRIVILEGES ON *.* TO 'singularity'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'singularity'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
SQL

echo "Database setup completed successfully"


