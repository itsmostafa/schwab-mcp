---
title: "Security Testing Types: SAST, DAST, and Fuzz Testing"
type: concept
domain: 6
tags: [sast, dast, fuzz-testing, static-analysis, dynamic-analysis, white-box, black-box, code-review]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Security Testing Types: SAST, DAST, and Fuzz Testing

## Overview

Application security testing falls into two broad categories: **static** (application not running) and **dynamic** (application running). Both can be manual or automated, and testing effectiveness is maximized by using them in combination.

**Note on IAST and RASP:** The source (destination-cissp) does not cover Interactive Application Security Testing (IAST) or Runtime Application Self-Protection (RASP). These are covered in other CISSP study materials (cissp-exam-outline scope). Brief definitions for completeness:
- **IAST** — agents inside the running application observe behavior from within; combines SAST and DAST benefits.
- **RASP** — security built into the application runtime that detects and blocks attacks in real time.

## Static Application Security Testing (SAST)

- **Runtime state**: Application is NOT running.
- **What is examined**: The underlying source code directly.
- **Testing perspective**: White box — source code is visible and accessible.
- **When used**: During development (before deployment).
- **Advantage**: Finds bugs and vulnerabilities before they reach runtime. Can scan 100% of the codebase.
- **Example use**: Automated tool scans source code for hardcoded credentials, SQL injection patterns, buffer overflow risks.

## Dynamic Application Security Testing (DAST)

- **Runtime state**: Application IS running.
- **What is examined**: The application's behavior and responses — not the source code.
- **Testing perspective**: Black box — code is not visible; focus is entirely on inputs and outputs.
- **When used**: After deployment to a test or staging environment; mimics attacker behavior.
- **Advantage**: Finds runtime vulnerabilities that only manifest during execution (e.g., authentication bypass, session management flaws).
- **Example use**: Automated vulnerability scanner probing web application endpoints.

## Fuzz Testing

- **Type**: A form of dynamic testing.
- **Core principle**: Chaos. Random, unexpected, or malformed inputs are thrown at the application to see how it responds ("breaks").
- **Why effective**: Developers tend to be logical; fuzz testing exposes edge cases that logical code paths never reach.

Two types of fuzz testing (destination-cissp §6.2.1):

| Type | Also Called | How it Works |
|---|---|---|
| **Mutation** | Dumb fuzzer | Randomly mutates valid input (bit flipping, appending random bytes). No understanding of input structure. |
| **Generation** | Intelligent/smart fuzzer | Generates new inputs from scratch based on knowledge of the file format or protocol. Understands input structure. |

## Comparison Table

| Property | SAST | DAST | Fuzz Testing |
|---|---|---|---|
| App running? | No | Yes | Yes |
| Box type | White box | Black box | Black box (dynamic) |
| Source code needed? | Yes | No | No |
| Finds issues at | Code level | Runtime | Runtime (edge cases) |
| Automation | High | High | High |
| Best phase | Development | QA/staging | QA/staging |
| Key limitation | Misses runtime flaws | Misses code-level defects | Non-deterministic; coverage varies |

## Manual vs Automated Testing

The source distinguishes two testing method dimensions (destination-cissp §6.2.1):

- **Manual testing** — a person actively tests (examining code, entering inputs, using the UI).
- **Automated testing** — scripts and batch files executed by software. Vulnerability scanners are a common automated testing tool.

Best practice: use both — manual catches nuanced logic flaws automated tools miss; automated provides breadth and speed.

Combinations:
- *Automated static white box* — tool scans source code for common errors, undefined variables.
- *Automated dynamic black box* — vulnerability scanner probing a live app from outside (no source access).

## Code Review Access Levels

| Level | Also Called | Description |
|---|---|---|
| No source code access | Black box | Tester examines behavior only. |
| Source code access | White box | Tester can read and analyze code directly. |

## Software Testing Progression (SDLC)

Security testing must be embedded at every phase (destination-cissp §6.2.0):

| Phase | Testing Activity |
|---|---|
| Planning | Capture and validate security requirements |
| Design | Confirm required controls are in the design |
| Develop | Unit, interface, integration, system, acceptance testing; SAST; vulnerability assessments |
| Deploy | Usability, performance testing; log review; vulnerability assessment |
| Operate | Configuration management reviews; vulnerability management; log analysis |
| Retire | Confirm data migration is secure; verify data destruction |

Software testing builds progressively:

1. **Unit testing** — individual components/functions tested in isolation.
2. **Interface testing** — verifies that individual components connect and communicate correctly.
3. **Integration testing** — tests groups of units working together.
4. **System testing** — tests the fully integrated system.

## Coverage Analysis

**Test coverage** = (lines of code covered by tests) / (total lines of code). Expressed as a percentage. Higher coverage = greater confidence that bugs have been found. 100% coverage is rarely achieved but is the goal for high-criticality systems.

## Cross-links

- [Code Review and Security](code-review-security.md)
- [Penetration Testing](penetration-testing.md)
- [Assessment vs Audit](assessment-vs-audit.md)

## Sources

- destination-cissp §6.2.0 (software testing lifecycle phases)
- destination-cissp §6.2.1 (SAST, DAST, fuzz testing, manual/automated, code review, test types, boundary/equivalence, coverage)
- cissp-exam-outline (domain 6 subtopic list)
