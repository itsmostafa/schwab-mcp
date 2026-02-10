# Repository Guidelines

## Project Structure & Module Organization
This repository is a Go backend interview-prep workspace, organized by numbered topics and a planned capstone project.

- `01-go-fundamentals/` ... `08-kubernetes/`: Topic folders, each with `notes.md`.
- `project/`: Design and planning for an "Integration Sync Service" practice microservice.
- Root docs: `README.md` (study flow), `CLAUDE.md` (agent-focused repo context).

Keep topic directories zero-padded and sequential (`01-`, `02-`, etc.) so the learning path stays ordered.

## Build, Test, and Development Commands
There is no build pipeline yet for the notes themselves. Use these commands for contributor workflow:

- `rg --files`: Quick inventory of tracked content.
- `rg "TODO|Interview Questions|Resources" 0*-*/notes.md`: Check template sections across topics.
- `git log --oneline -n 10`: Review recent commit message style.

If/when Go code is added under `project/` or topic `examples/` folders:

- `go test ./...`: Run unit tests.
- `go fmt ./...`: Apply standard Go formatting.

## Coding Style & Naming Conventions
For notes and docs:

- Use clear Markdown headings (`#`, `##`) and short bullet lists.
- Keep each `notes.md` aligned to the same section pattern: `Key Concepts`, `Interview Questions`, `Resources`.

For future Go code:

- Follow standard Go formatting (`gofmt`) and idiomatic naming (`camelCase` locals, `PascalCase` exported identifiers).
- Prefer small packages under `internal/` for non-public logic.

## Testing Guidelines
Current state: no mandatory automated test suite for Markdown notes.

Contributor checks:

- Verify links and Markdown rendering before opening a PR.
- For new Go examples or project code, include table-driven tests where practical and run `go test ./...`.

## Commit & Pull Request Guidelines
Recent history favors short, imperative commit messages (for example: `add golang notes and practice code`, `Remove old study materials...`).

- Keep commits focused to one topic or concern.
- Use imperative subjects and concise scope.
- In PRs, include: purpose, changed paths (for example `03-gin/notes.md`), and any follow-up TODOs.
- Link related issues/tasks when applicable.
