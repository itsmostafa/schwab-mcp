# Task 01 — Project scaffold and configuration

**Read first:** `.claude/commands/tasks/00-conventions.md`.

**Depends on:** nothing — this is the entry point.

## Objective

Stand up a minimal, running Astro + TypeScript project with the configurable
external data directory wired up, and nothing else. At the end of this task
`npm run dev` should serve a placeholder homepage, and the app should be able
to read `SIA_DATA_DIR` and print/confirm where it thinks the data directory
is.

## Scope

In scope:
- `npm create astro@latest` (or equivalent) scaffold, TypeScript strict mode.
- `sia.config.ts` reading `SIA_DATA_DIR` with the dev fallback from `00-conventions.md`, roughly:

  ```ts
  export default {
    dataDir: process.env.SIA_DATA_DIR,
    search: { semantic: false },
    ai: { provider: "none" },
    network: { trustProxy: true },
  }
  ```
- `src/lib/config` module: resolves and validates the data directory (creates `inbox/`, `knowledge/`, `runtime/` under it if missing), exports a typed config object.
- `.gitignore` additions for the dev fallback data dir, `node_modules`, Astro build output, `.env`.
- `examples/knowledge/` directory with 2-3 sample OKF documents (hand-written, matching the frontmatter table in `00-conventions.md`) — enough for later tasks to load against.
- One placeholder Astro page that renders "Sia" and the resolved data directory path (proves config wiring end-to-end).
- README updated: install/dev instructions matching the spec's quickstart (`git clone`, `npm install`, `npm run dev`), and a short note on `SIA_DATA_DIR`.

Out of scope (later tasks): document loading/validation (Task 02), any real
navigation or knowledge pages (Task 05), search, capture, MCP.

## Deliverables

- Running Astro project at repo root.
- `src/lib/config/index.ts` (or similar) with a typed `SiaConfig`.
- `examples/knowledge/*.md` sample documents.
- Updated `README.md`, `.gitignore`.

## Acceptance criteria

- [ ] `npm install && npm run dev` works from a clean clone with no `SIA_DATA_DIR` set, using the gitignored dev fallback, and logs a visible warning that it's a dev-only location.
- [ ] Setting `SIA_DATA_DIR=/some/path npm run dev` uses that path instead, creating `inbox/`, `knowledge/`, `runtime/` if absent.
- [ ] No user data or private paths are hardcoded or committed.
- [ ] `examples/knowledge/` documents are valid per the frontmatter table in `00-conventions.md` (this will be re-validated programmatically in Task 02).

## Notes / assumptions

- Keep dependencies minimal — this task should not pull in a UI framework, state library, or CSS framework beyond Astro's defaults. Simple yet powerful: a config module and a working dev server, nothing more.
