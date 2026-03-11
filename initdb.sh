#!/usr/bin/env bash
# initdb.sh — Create the 110claw database and all tables.
#
# Usage:
#   ./initdb.sh                          # use defaults / backend/.env
#   ./initdb.sh -h 127.0.0.1 -P 3306 \
#               -u root -p mypass \
#               -d 110claw               # explicit flags
#
# Connection priority (highest → lowest):
#   1. CLI flags  (-h -P -u -p -d)
#   2. DB_DSN in backend/.env
#   3. Built-in defaults
#
# Requires: mysql client in PATH.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCHEMA="$SCRIPT_DIR/db/schema.sql"

# ── defaults ────────────────────────────────────────────────
DB_HOST="127.0.0.1"
DB_PORT="3306"
DB_USER="root"
DB_PASS="password"
DB_NAME="110claw"

# ── parse DSN from backend/.env if present ──────────────────
ENV_FILE="$SCRIPT_DIR/backend/.env"
if [[ -f "$ENV_FILE" ]]; then
  DSN=$(grep -E '^DB_DSN=' "$ENV_FILE" | head -1 | cut -d= -f2- | tr -d '"' | tr -d "'")
  if [[ -n "$DSN" ]]; then
    # DSN format: user:pass@tcp(host:port)/dbname?options
    _USER="${DSN%%:*}"
    _AFTER_USER="${DSN#*:}"
    _PASS="${_AFTER_USER%%@*}"
    _AFTER_PASS="${_AFTER_USER#*@}"          # tcp(host:port)/dbname?options
    _HOST_PORT="${_AFTER_PASS#tcp(}"         # host:port)/dbname?options
    _HOST_PORT="${_HOST_PORT%%)*}"           # host:port
    _HOST="${_HOST_PORT%%:*}"
    _PORT="${_HOST_PORT##*:}"
    _DB="${_AFTER_PASS#*/}"
    _DB="${_DB%%\?*}"

    [[ -n "$_USER" ]] && DB_USER="$_USER"
    [[ -n "$_PASS" ]] && DB_PASS="$_PASS"
    [[ -n "$_HOST" ]] && DB_HOST="$_HOST"
    [[ -n "$_PORT" ]] && DB_PORT="$_PORT"
    [[ -n "$_DB"   ]] && DB_NAME="$_DB"
  fi
fi

# ── CLI flag overrides ───────────────────────────────────────
while getopts ":h:P:u:p:d:" opt; do
  case $opt in
    h) DB_HOST="$OPTARG" ;;
    P) DB_PORT="$OPTARG" ;;
    u) DB_USER="$OPTARG" ;;
    p) DB_PASS="$OPTARG" ;;
    d) DB_NAME="$OPTARG" ;;
    :) echo "ERROR: option -$OPTARG requires an argument" >&2; exit 1 ;;
    ?) echo "ERROR: unknown option -$OPTARG" >&2; exit 1 ;;
  esac
done

# ── preflight ───────────────────────────────────────────────
if ! command -v mysql &>/dev/null; then
  echo "ERROR: mysql client not found in PATH." >&2
  exit 1
fi

if [[ ! -f "$SCHEMA" ]]; then
  echo "ERROR: schema file not found: $SCHEMA" >&2
  exit 1
fi

MYSQL_OPTS=(-h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER")
if [[ -n "$DB_PASS" ]]; then
  MYSQL_OPTS+=("-p${DB_PASS}")
fi

echo "──────────────────────────────────────────"
echo "  Host   : $DB_HOST:$DB_PORT"
echo "  User   : $DB_USER"
echo "  Database: $DB_NAME"
echo "──────────────────────────────────────────"

# ── create database if not exists ───────────────────────────
echo "[1/2] Creating database '$DB_NAME' if not exists..."
mysql "${MYSQL_OPTS[@]}" \
  --execute="CREATE DATABASE IF NOT EXISTS \`${DB_NAME}\`
             CHARACTER SET utf8mb4
             COLLATE utf8mb4_unicode_ci;"

# ── apply schema ─────────────────────────────────────────────
echo "[2/2] Applying schema from db/schema.sql..."
mysql "${MYSQL_OPTS[@]}" "$DB_NAME" < "$SCHEMA"

echo ""
echo "Done. Tables created in '$DB_NAME':"
mysql "${MYSQL_OPTS[@]}" "$DB_NAME" \
  --execute="SELECT TABLE_NAME, TABLE_COMMENT
             FROM information_schema.TABLES
             WHERE TABLE_SCHEMA = '${DB_NAME}'
             ORDER BY TABLE_NAME;" \
  --table
