# Repository Guidelines

## Project Structure & Module Organization
This repository is currently documentation-first. The primary project content is in `README.md`, with repository-level settings in `.gitignore`.

When adding material, keep files focused and easy to scan:
- Keep interview prep content in Markdown (`*.md`).
- Use one topic per file when content grows (for example, `go-http.md`, `gorm-notes.md`).
- Prefer shallow structure unless sections become large.

## Build, Test, and Development Commands
There is no formal build system configured in this repository yet. Use these commands for day-to-day work:

```bash
git status              # Check local changes
git log --oneline -n 10 # Review recent commit style
rg "Gin|GORM|OTel" .    # Search topic coverage quickly
```

If available in your environment, run Markdown linting before opening a PR:

```bash
markdownlint "**/*.md"
```

## Coding Style & Naming Conventions
- Use clear, direct technical writing with short paragraphs and bullet lists.
- Prefer ATX headings (`#`, `##`, `###`) and consistent heading hierarchy.
- Use fenced code blocks with language hints (for example, ` ```bash `, ` ```go `).
- File names should be lowercase with hyphens (for example, `otel-tracing.md`).

## Testing Guidelines
No automated test suite is currently defined. Validate contributions by:
- Checking Markdown renders correctly in your editor or GitHub preview.
- Verifying commands and examples are copy-paste runnable.
- Re-reading for factual consistency across related notes.

## Commit & Pull Request Guidelines
Recent history uses short, imperative, lowercase commit messages (for example, `code cleanup`, `add python projects`). Follow that pattern:
- Keep commit subject concise and action-oriented.
- Group related documentation updates in one commit.

For pull requests:
- Include a brief summary of what changed and why.
- List affected files/sections.
- Link any relevant issue or discussion when applicable.
