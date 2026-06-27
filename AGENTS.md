# Mnet Auto — AGENTS.md

## Quick start
- **Build:** `go build -o mnet-auto.exe .`
- **Run:** `.\mnet-auto.exe`

## Architecture
- **Single-file Go CLI** (no frameworks, no external dependencies, standard library only)
- **Entrypoint:** `main.go` — boots the while-true menu loop
- **2 modes** (set via menu or change `mode` in `config.json`):
  - `auto` (default) — auto temp email + API registration, retry 3 times + domain rotation
  - `manual` — user enters email + API registration
- **Config persistence:** `config.json` is auto-created in project root on first run (stores `password`, `mode`)
- **Accounts saved to:** `accounts.txt` in project root (appended with timestamped session headers)

## Key packages
- All in `package main` — no sub-packages. Source files are organized by domain (e.g., `mnet.go`, `tempmail.go`, `orchestrator.go`)

## Gotchas
- API auto mode retries 3 times with domain rotation + blacklist
- Human-like behavior is intentional (random delay between requests) — slow by design
- `config.json` and `accounts.txt` are gitignored; will be created at runtime
- Default password: `nmixx0222-`
