# Course AI Backend Go

Minimal Go backend scaffold for rebuilding the Course AI API.

## Requirements

- Go 1.26+
- PostgreSQL running from the root `docker-compose.yml`
- A local `.env` file copied from `.env.example`

## Setup

```powershell
cd backend-go
Copy-Item .env.example .env
go run ./cmd/api
```

## Endpoints

- `GET /health` returns service status.
- `GET /prompts` returns loaded prompt names.

## Project Layout

```txt
cmd/api              application entrypoint
internal/config      environment loading and app config
internal/envfile     minimal .env loader
internal/httpapi     HTTP routes and handlers
internal/prompts     prompt loading from markdown files
prompts              system prompts used by generation pipeline
```

This scaffold intentionally uses the Go standard library first. Add framework or database dependencies only when the core boundaries are clear.
