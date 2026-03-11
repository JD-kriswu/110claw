#!/usr/bin/env bash
# Usage: ./build.sh [frontend|backend|all]
# Outputs: frontend/dist/  and  bin/backend

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BIN_DIR="$REPO_ROOT/bin"

build_frontend() {
  echo "==> Building frontend..."
  cd "$REPO_ROOT/frontend"
  npm install
  npm run build
  echo "==> Frontend built: frontend/dist/"
}

build_backend() {
  echo "==> Building backend..."
  mkdir -p "$BIN_DIR"
  cd "$REPO_ROOT/backend"
  go build -buildvcs=false -o "$BIN_DIR/backend" .
  echo "==> Backend built: bin/backend"
}

TARGET="${1:-all}"

case "$TARGET" in
  frontend)
    build_frontend
    ;;
  backend)
    build_backend
    ;;
  all)
    build_frontend
    build_backend
    ;;
  *)
    echo "Usage: $0 [frontend|backend|all]" >&2
    exit 1
    ;;
esac
