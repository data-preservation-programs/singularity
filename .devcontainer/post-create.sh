#!/usr/bin/env bash
set -euo pipefail

# Initialize Go modules
go mod download

chmod +x .devcontainer/*.sh || true


