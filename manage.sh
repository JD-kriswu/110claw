#!/usr/bin/env bash
# Usage: ./manage.sh [start|stop|restart|status] [frontend|backend|all]

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
RUN_DIR="$REPO_ROOT/run"
LOG_DIR="$REPO_ROOT/run"

mkdir -p "$RUN_DIR"

FRONTEND_PID="$RUN_DIR/frontend.pid"
BACKEND_PID="$RUN_DIR/backend.pid"
FRONTEND_LOG="$LOG_DIR/frontend.log"
BACKEND_LOG="$LOG_DIR/backend.log"

start_frontend() {
  if [[ -f "$FRONTEND_PID" ]] && kill -0 "$(cat "$FRONTEND_PID")" 2>/dev/null; then
    echo "Frontend is already running (PID $(cat "$FRONTEND_PID"))"
    return
  fi
  echo "==> Starting frontend..."
  cd "$REPO_ROOT/frontend"
  nohup npm run dev >"$FRONTEND_LOG" 2>&1 &
  echo $! >"$FRONTEND_PID"
  echo "    Frontend started (PID $!), log: run/frontend.log"
}

start_backend() {
  if [[ -f "$BACKEND_PID" ]] && kill -0 "$(cat "$BACKEND_PID")" 2>/dev/null; then
    echo "Backend is already running (PID $(cat "$BACKEND_PID"))"
    return
  fi
  if [[ ! -f "$REPO_ROOT/bin/backend" ]]; then
    echo "ERROR: bin/backend not found. Run ./build.sh backend first." >&2
    exit 1
  fi
  echo "==> Starting backend..."
  cd "$REPO_ROOT/backend"
  nohup ../bin/backend >"$BACKEND_LOG" 2>&1 &
  echo $! >"$BACKEND_PID"
  echo "    Backend started (PID $!), log: run/backend.log"
}

stop_service() {
  local name="$1"
  local pidfile="$2"
  if [[ -f "$pidfile" ]]; then
    local pid
    pid=$(cat "$pidfile")
    if kill -0 "$pid" 2>/dev/null; then
      echo "==> Stopping $name (PID $pid)..."
      kill "$pid"
    else
      echo "$name is not running (stale PID $pid)"
    fi
    rm -f "$pidfile"
  else
    echo "$name is not running (no PID file)"
  fi
}

status_service() {
  local name="$1"
  local pidfile="$2"
  if [[ -f "$pidfile" ]]; then
    local pid
    pid=$(cat "$pidfile")
    if kill -0 "$pid" 2>/dev/null; then
      echo "$name: running (PID $pid)"
    else
      echo "$name: stopped (stale PID $pid)"
    fi
  else
    echo "$name: stopped"
  fi
}

ACTION="${1:-status}"
TARGET="${2:-all}"

case "$ACTION" in
  start)
    case "$TARGET" in
      frontend) start_frontend ;;
      backend)  start_backend  ;;
      all)      start_frontend; start_backend ;;
      *) echo "Usage: $0 start [frontend|backend|all]" >&2; exit 1 ;;
    esac
    ;;
  stop)
    case "$TARGET" in
      frontend) stop_service frontend "$FRONTEND_PID" ;;
      backend)  stop_service backend  "$BACKEND_PID"  ;;
      all)      stop_service frontend "$FRONTEND_PID"; stop_service backend "$BACKEND_PID" ;;
      *) echo "Usage: $0 stop [frontend|backend|all]" >&2; exit 1 ;;
    esac
    ;;
  restart)
    case "$TARGET" in
      frontend)
        stop_service frontend "$FRONTEND_PID"
        sleep 1
        start_frontend
        ;;
      backend)
        stop_service backend "$BACKEND_PID"
        sleep 1
        start_backend
        ;;
      all)
        stop_service frontend "$FRONTEND_PID"
        stop_service backend  "$BACKEND_PID"
        sleep 1
        start_frontend
        start_backend
        ;;
      *) echo "Usage: $0 restart [frontend|backend|all]" >&2; exit 1 ;;
    esac
    ;;
  status)
    case "$TARGET" in
      frontend) status_service frontend "$FRONTEND_PID" ;;
      backend)  status_service backend  "$BACKEND_PID"  ;;
      all)
        status_service frontend "$FRONTEND_PID"
        status_service backend  "$BACKEND_PID"
        ;;
      *) echo "Usage: $0 status [frontend|backend|all]" >&2; exit 1 ;;
    esac
    ;;
  *)
    echo "Usage: $0 [start|stop|restart|status] [frontend|backend|all]" >&2
    exit 1
    ;;
esac
