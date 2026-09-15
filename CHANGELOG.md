# Changelog

## [0.2.1](https://github.com/itsmostafa/schwab-mcp/compare/v0.2.0...v0.2.1) (2026-09-15)


### ⚠ BREAKING CHANGES

* **cli:** `schwab` no longer starts the MCP server and `schwab serve` is removed. Re-run `schwab mcp setup`, or add `mcp` to the args of manually configured clients.

### Features

* **cli:** run the MCP server with `schwab mcp` ([ca3a500](https://github.com/itsmostafa/schwab-mcp/commit/ca3a500583c917d8588040082942d0ab731828b4))
* **setup:** register schwab with Claude Desktop ([17e6ee9](https://github.com/itsmostafa/schwab-mcp/commit/17e6ee9828a61d5bee49d94e6c931afec485cb40))
* **setup:** register schwab with Claude Desktop ([baef53c](https://github.com/itsmostafa/schwab-mcp/commit/baef53c84b3a2ef11bcc2ceaadc83a42e3a94639))
* **setup:** show detected clients and a completion message ([add5f39](https://github.com/itsmostafa/schwab-mcp/commit/add5f396b58dd7fbcf629011330980fc261ce63d))


### Bug Fixes

* **setup:** handle null Claude Desktop config ([b1687e9](https://github.com/itsmostafa/schwab-mcp/commit/b1687e94d97524f48b500166ac22b499a0efe9b7))


### Miscellaneous Chores

* release 0.2.1 ([a3f4d2f](https://github.com/itsmostafa/schwab-mcp/commit/a3f4d2f7b086f6f5532119bf4b2f329b4dc7dc46))

## [0.2.0](https://github.com/itsmostafa/schwab-mcp/compare/v0.1.0...v0.2.0) (2026-09-15)


### ⚠ BREAKING CHANGES

* **trading:** with SCHWAB_ALLOW_TRADING=true, MARKET, STOP and other uncapped order types are rejected unless SCHWAB_ALLOW_MARKET_ORDERS=true is set. LIMIT orders that cross the live market by more than 50 bps are rejected; tune with SCHWAB_MAX_PRICE_DEVIATION_BPS.

### Features

* **cli:** add schwab mcp setup to register with Claude Code and Codex ([78e5717](https://github.com/itsmostafa/schwab-mcp/commit/78e57170f3bbfbc410e91bad82e912e93d22edce))
* **cli:** add schwab mcp setup to register with Claude Code and Codex ([0725504](https://github.com/itsmostafa/schwab-mcp/commit/0725504d29015f3975a7e81ff20c101176920993))
* **config:** cobra CLI and ~/.config/schwab credentials file ([03ceb4e](https://github.com/itsmostafa/schwab-mcp/commit/03ceb4e61094b5f29d9d8257d32a213707ed3aa7))
* **config:** read app credentials from ~/.config/schwab/config ([79c36dd](https://github.com/itsmostafa/schwab-mcp/commit/79c36dd53f70317523fbb6430deff136eff39d96))
* **trading:** re-quote and reject orders built on stale prices ([59bea93](https://github.com/itsmostafa/schwab-mcp/commit/59bea93bba06cb9521caca1bbe748e065384d3fb))


### Bug Fixes

* **cli:** restore Claude entry if replacement fails ([cb513ca](https://github.com/itsmostafa/schwab-mcp/commit/cb513ca59d584e4b05a42cf6fc59070854ba308f))
* **trading:** check order types and quotes in TRIGGER children ([85bcc0d](https://github.com/itsmostafa/schwab-mcp/commit/85bcc0dc445573c15f312b421a3794de63bba43b))

## 0.1.0 (2026-09-15)


### Features

* **accounts:** add account, order and transaction tools ([0c0a3f2](https://github.com/itsmostafa/schwab-mcp/commit/0c0a3f26746e5d59c77c0aaf1f4364ad66b8cfda))
* add Schwab Trader API MCP server ([0bd0fdf](https://github.com/itsmostafa/schwab-mcp/commit/0bd0fdfc099e5c54c1b5139ccccc083f5f367645))
* add versioning, GitHub releases, install script and schwab update ([79cda3e](https://github.com/itsmostafa/schwab-mcp/commit/79cda3ea02655a9faf2ed8e9e95f7411ce27d89a))
* **cli:** add version command and schwab update self-update ([34b2dc9](https://github.com/itsmostafa/schwab-mcp/commit/34b2dc961112d309c2ef4f7f56cc754c7cfb6937))
* **install:** add install script for release binaries ([9a37a66](https://github.com/itsmostafa/schwab-mcp/commit/9a37a667ab4e20cf312384b1926124829d2c6c82))
* **market:** add market data tools ([84c4028](https://github.com/itsmostafa/schwab-mcp/commit/84c4028f41968dde694aa7d07f321a356ee5f9c9))
* **server:** add schwab MCP stdio server with OAuth login ([4c99b05](https://github.com/itsmostafa/schwab-mcp/commit/4c99b05bbfbf0c74eed17cf78441ed1a2598b7bc))


### Bug Fixes

* **client:** treat order previews as retryable, not mutations ([f09412c](https://github.com/itsmostafa/schwab-mcp/commit/f09412c8500c7d615289c7363a4c55c8696ec1dc))

## Changelog
