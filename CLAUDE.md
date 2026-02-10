# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a **study notes repository** for Go backend interview preparation, focused on an integrations team role. It contains no buildable code — only markdown notes and (planned) Go example snippets organized by topic.

## Structure

- `01-go-fundamentals/` through `08-kubernetes/` — Numbered topic directories, each with a `notes.md` and an `examples/` folder for Go code snippets
- `project/` — A planned practice microservice ("Integration Sync Service") that ties all topics together; currently a design doc only
- Topics are meant to be studied **in order** as each builds on the previous

## Topic Note Format

Each `notes.md` follows the same template:
- **Key Concepts** — topic summaries
- **Interview Questions** — Q&A for self-testing
- **Resources** — links to docs, articles, talks

## Practice Project Architecture

The planned project in `project/` is a Go microservice that syncs data from a third-party API into Postgres with event publishing and OpenTelemetry instrumentation. Layers: Gin handlers → Service layer → GORM repo / HTTP client / Event bus.

## Working in This Repo

- When adding study notes, follow the existing `notes.md` template structure
- Go example code goes in the `examples/` subdirectory of the relevant topic
- Keep topic numbering sequential with zero-padded prefixes (`01-`, `02-`, etc.)
- Use the DeepWiki MCP tool to look up relevant GitHub repositories (e.g., `golang/go`, `gin-gonic/gin`, `go-gorm/gorm`, `kubernetes/kubernetes`) when researching topics covered in this repo, to ensure notes are accurate and up to date
