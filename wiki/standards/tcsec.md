---
title: "TCSEC — Trusted Computer System Evaluation Criteria (Orange Book)"
type: standard
domain: 3
tags: [tcsec, orange-book, evaluation-criteria, historical, confidentiality, rainbow-series]
sources: [destination-cissp]
updated: 2026-05-13
---

# TCSEC — Trusted Computer System Evaluation Criteria (Orange Book)

## Overview

TCSEC, colloquially known as the **Orange Book** (due to its orange cover), was the first evaluation criteria system for measuring the security of IT products. Published by the US Department of Defense in the 1980s as part of the **Rainbow Series** (each book in the series had a different colored cover and addressed a specific security topic). TCSEC was superseded by the Common Criteria.

## Standard Details

| Attribute | Value |
|---|---|
| **Formal name** | Trusted Computer System Evaluation Criteria (TCSEC) |
| **Common name** | Orange Book |
| **Published by** | US Department of Defense |
| **Era** | 1980s |
| **Status** | Superseded (by ITSEC, then Common Criteria) |
| **Limitation** | Measures confidentiality only; single-box (non-networked) architectures only |

## TCSEC Classification Levels

| Level | Name | Description |
|---|---|---|
| **A1** | Verified Design | Mathematical verification of design; highest assurance |
| **B3** | Security Domains | Security labels; verification of no covert channels; must remain secure at startup |
| **B2** | Structured Protection | Security labels and verification of no covert channels |
| **B1** | Labeled Security | Security labels applied |
| **C2** | Controlled Access | Strict login procedures; individual user accountability |
| **C1** | Discretionary Security | Weak protection mechanisms |
| **D1** | Minimal Protection | Failed evaluation or was not tested |

- Each level inherits all characteristics of the previous level (cumulative)
- Most commercial products were rated at **C2 or B1**
- Despite being considered outdated, TCSEC is still optimal if confidentiality is the only requirement

## Rainbow Series Context

The Rainbow Series included multiple colored books, each covering a specific security topic:
- **Orange Book** — TCSEC; security product evaluation criteria
- **Red Book** — Trusted Network Interpretation; network security
- **Light Blue Book** — Password guidelines

## TCSEC vs. ITSEC

ITSEC (Information Technology Security Evaluation Criteria) was developed by Europeans to improve on TCSEC:

| Feature | TCSEC | ITSEC |
|---|---|---|
| **Scope** | Confidentiality only | CIA triad (confidentiality, integrity, availability) |
| **Environment** | Single-box only | Networked environments |
| **Rating approach** | Unified D/C/B/A rating | Separate functional (F) and assurance (E) ratings |
| **Assurance levels** | D–A1 | E0–E6 |
| **Status** | Superseded | Replaced by CC in 2005 |

## Exam Notes

- Orange Book = TCSEC = first evaluation criteria system = confidentiality only = DoD = 1980s
- Know the division letters (D, C1, C2, B1, B2, B3, A1) in order
- TCSEC does not map to networked environments — that limitation drove ITSEC's creation
- Replaced by Common Criteria, not directly by ITSEC; CC absorbed both

## Related Pages

- [Common Criteria Standard](common-criteria-standard.md) — current evaluation standard
- [Common Criteria concept](../concepts/common-criteria.md) — detailed evaluation process

## Sources

- destination-cissp §3.2.6 (pp. 261–266)
