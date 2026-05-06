---
title: "Source: CISSP Exam Outline"
type: source
domain: cross
tags: [source, official, exam-outline]
sources: [cissp-exam-outline]
updated: 2026-05-05
---

# Source: CISSP Exam Outline

**File**: `raw/CISSP-Exam-Outline.pdf` (extracted: `raw/CISSP-Exam-Outline.md`)
**Publisher**: ISC2
**Purpose**: Official authoritative source for CISSP exam domain structure, subtopic
taxonomy, and exam weights. Used as the structural spine of this entire wiki.

## Summary

The Exam Outline defines the eight CISSP domains, their subtopics, and the percentage
weight each domain carries in the CAT exam. It is the source-of-truth for naming conventions
across all wiki pages. Every domain page in `wiki/domains/` was seeded from this document.

Exam format: CAT, 125–175 questions, 4 hours. Requires 5 years experience in 2+ domains
(1 year waived with qualifying degree or approved ISC2 credential).

## Wiki pages informed by this source

- [overview.md](../overview.md)
- [domains/01-security-and-risk-management.md](../domains/01-security-and-risk-management.md)
- [domains/02-asset-security.md](../domains/02-asset-security.md)
- [domains/03-security-architecture-and-engineering.md](../domains/03-security-architecture-and-engineering.md)
- [domains/04-communication-and-network-security.md](../domains/04-communication-and-network-security.md)
- [domains/05-identity-and-access-management.md](../domains/05-identity-and-access-management.md)
- [domains/06-security-assessment-and-testing.md](../domains/06-security-assessment-and-testing.md)
- [domains/07-security-operations.md](../domains/07-security-operations.md)
- [domains/08-software-development-security.md](../domains/08-software-development-security.md)

## Notes on the raw extraction

The `raw/CISSP-Exam-Outline.md` file is very large (inline base64 images from the PDF
inflate the token count). Use `grep` to find line numbers, then read with `offset`/`limit`:

- Domain 1 starts at line ~70
- Domain 2 starts at line ~149
- Domain 3 starts at line ~178
- Domain 4 starts at line ~260
- Domain 5 starts at line ~301
- Domain 6 starts at line ~349
- Domain 7 starts at line ~399
- Domain 8 starts at line ~493
