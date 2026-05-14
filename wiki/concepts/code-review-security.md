---
title: "Code Review for Security"
type: concept
domain: 6
tags: [code-review, sast, static-analysis, white-box, black-box, secure-sdlc]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Code Review for Security

## Definition

Code review from a security perspective is the examination of application source code to identify defects, vulnerabilities, and non-conformities with security requirements. It can be performed manually (human reviewer reading code) or through automated static analysis tooling.

## Code Review Perspectives

Code review is considered from two knowledge perspectives (destination-cissp §6.2.1):

| Perspective | Also Called | Access |
|---|---|---|
| **No source code access** | Black box | Reviewer examines behavior and outputs only. Cannot read code. |
| **Source code access** | White box | Reviewer has full visibility into the codebase. |

White box code review is the most thorough. It enables:
- Identification of hardcoded credentials, secrets, and API keys.
- Detection of insecure function calls (e.g., `strcpy`, `sprintf`, SQL string concatenation).
- Logic flow analysis for access control bypasses.
- Architecture-level design flaw identification.

## Static Analysis (SAST)

Static Application Security Testing (SAST) is automated white box code review. The tool scans source code while the application is not running. See [Security Testing Types](security-testing-types.md) for the full SAST/DAST/Fuzz comparison.

Common SAST findings:
- Undefined or uninitialized variables.
- Use of deprecated/unsafe functions.
- Injection vulnerabilities (SQL, command, LDAP).
- Buffer overflows.
- Insecure random number generation.

## Fuzz Testing

Fuzz testing is a dynamic (runtime) complement to code review. Rather than reading code, it throws random or malformed inputs at the running application to provoke unexpected behavior. Useful for finding input validation flaws that static analysis may miss. See [Security Testing Types](security-testing-types.md).

## When to Conduct Code Review in the SDLC

Code review is most effective — and cheapest to remediate — when performed during development, before code reaches production. The principle: the later a defect is found, the more expensive it is to fix.

- **Development phase**: automated SAST runs on every commit (CI pipeline integration).
- **Pre-deployment**: manual review of high-risk modules; review of security-relevant code paths.
- **Post-incident**: code review of affected components after a security incident.

## Test Coverage Consideration

Test coverage analysis measures what percentage of source code is exercised by tests. Security-critical code paths should be targeted for high coverage:

> Coverage (%) = (lines of code covered) / (total lines of code) × 100

## Cross-links

- [Security Testing Types](security-testing-types.md) — SAST, DAST, fuzz testing comparison
- [Penetration Testing](penetration-testing.md) — exploitation-focused testing

## Sources

- destination-cissp §6.2.1 (code review perspectives, SAST, fuzz testing, coverage analysis)
- cissp-exam-outline (domain 6 subtopic list)
