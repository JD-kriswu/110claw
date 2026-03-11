# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

```
110claw/
├── frontend/         # Vue 3 + Vite app
├── backend/          # Go/Gin API server
│   ├── main.go
│   ├── go.mod        # module: github.com/110claw/backend
│   ├── config/       # env-based config (config.Load())
│   ├── router/       # route registration + CORS middleware wiring
│   ├── handler/      # HTTP handlers (one file per domain)
│   ├── middleware/   # reusable middleware (CORS, auth, …)
│   ├── model/        # GORM models + AutoMigrate entry point
│   └── store/        # DB (GORM/MySQL) and Redis client singletons
├── bin/              # compiled binaries (git-ignored)
├── run/              # PID + log files at runtime (git-ignored)
├── build.sh          # build frontend and/or backend
└── manage.sh         # start / stop / restart / status
```

## Commands

### Frontend (run from `frontend/`)

```bash
npm run dev      # dev server at http://localhost:8090
npm run build    # production build → frontend/dist/
npm run preview  # preview production build
```

### Backend (run from `backend/`)

```bash
go run .         # dev: run in place (requires MySQL + Redis)
go build -buildvcs=false -o ../bin/backend .   # compile
```

### Build scripts (run from repo root)

```bash
./build.sh frontend   # npm install + npm run build
./build.sh backend    # go build → bin/backend
./build.sh all        # both in sequence (default)
```

### Service management (run from repo root)

```bash
./manage.sh start   [frontend|backend|all]
./manage.sh stop    [frontend|backend|all]
./manage.sh restart [frontend|backend|all]
./manage.sh status  [frontend|backend|all]   # default action
```

PIDs are stored in `run/frontend.pid` / `run/backend.pid`.
Logs are written to `run/frontend.log` / `run/backend.log`.

## Architecture

### Frontend

- **Vue 3** with `<script setup>` SFC syntax
- **Vite** with `@` aliased to `src/`
- **Element Plus** UI library — all icons are globally registered in `main.js`
- **Pinia** for state management — stores go in `src/stores/`
- **Vue Router** — routes go in `src/router/`

The `src/views/`, `src/router/`, and `src/stores/` directories exist but are currently empty scaffolds.

### Backend

- **Gin** HTTP framework (`github.com/gin-gonic/gin`)
- **CORS** via `github.com/gin-contrib/cors` — allowed origins: `localhost:8090`, `localhost:5173`
- **GORM** + MySQL driver for the database (`gorm.io/gorm`, `gorm.io/driver/mysql`)
- **go-redis v9** for Redis (`github.com/redis/go-redis/v9`)
- **Viper** for config (`github.com/spf13/viper`)

Route layout:
- `GET /api/health` — health check (no auth)
- `/api/v1/...` — versioned API group

#### Auth (`/api/v1/auth`)
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/auth/register` | No | Register (username + password) |
| POST | `/api/v1/auth/login` | No | Login → returns `token` |
| POST | `/api/v1/auth/logout` | Yes | Invalidate current session |
| GET | `/api/v1/auth/me` | Yes | Current user info |

#### Content (public)
| Method | Path | Query params |
|--------|------|-------------|
| GET | `/api/v1/news` | `page`, `page_size`, `tag` |
| GET | `/api/v1/news/:id` | — |
| GET | `/api/v1/learn` | — |
| GET | `/api/v1/learn/:day` | day 1–7 |
| GET | `/api/v1/skills` | `page`, `page_size`, `tag`, `sort` (rating/downloads) |
| GET | `/api/v1/skills/:id` | — |

### Session / Auth design

- **No JWT** — server-side sessions stored in Redis
- Key: `session:{token}` (token = 32-byte random hex)
- Value: JSON `{"user_id", "username", "role", "expires_at"}`
- TTL: 24 h (7 days with `remember_me: true`); sliding renewal on every request
- Client sends `Authorization: Bearer {token}` header
- Middleware: `middleware.Auth()` — validates session and injects `user_id`, `username`, `role` into `gin.Context`

### Data models (GORM / MySQL)

`users`, `news`, `learn_steps`, `skills`, `machines`, `prizes`, `orders` — all defined in `backend/model/model.go` and auto-migrated on startup.

## Configuration

The backend reads environment variables (with defaults):

| Variable         | Default                                                        | Description              |
|------------------|----------------------------------------------------------------|--------------------------|
| `SERVER_PORT`    | `8080`                                                         | HTTP listen port         |
| `DB_DRIVER`      | `mysql`                                                        | Database driver          |
| `DB_DSN`         | `root:password@tcp(127.0.0.1:3306)/110claw?…`                 | GORM DSN                 |
| `REDIS_ADDR`     | `127.0.0.1:6379`                                               | Redis address            |
| `REDIS_PASSWORD` | *(empty)*                                                      | Redis auth password      |
| `REDIS_DB`       | `0`                                                            | Redis DB index           |

Variables can also be placed in a `backend/.env` file (Viper loads it automatically).

## Dev Workflow

1. Start MySQL and Redis locally.
2. (Optional) Create `backend/.env` with your DSN / Redis settings.
3. Run backend: `cd backend && go run .`
4. Run frontend: `cd frontend && npm run dev`
5. Frontend dev server proxies or calls `http://localhost:8080` directly.

There is no test runner configured.
