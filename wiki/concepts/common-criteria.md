---
title: "Common Criteria"
type: concept
domain: 3
tags: [common-criteria, evaluation-criteria, eal, protection-profile, certification, accreditation, iso-15408]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Common Criteria

Common Criteria (CC) is the globally recognized standard for evaluating the security capabilities of IT products. Formally designated **ISO/IEC 15408**, it superseded both TCSEC (Orange Book) and ITSEC. Common Criteria enables vendors to have their products independently evaluated and rated, and enables consumers to make informed purchasing decisions based on objective, internationally recognized ratings.

## Purpose

Security product claims by vendors are unverifiable without an independent evaluation system. Common Criteria solves this by:
- Providing a globally accepted, objective evaluation framework
- Enabling independent, licensed organizations to evaluate and rate products
- Producing documentation (Security Targets) that any potential consumer can examine
- Assigning a standardized EAL rating that communicates security assurance level

## Certification vs. Accreditation

Before diving into CC mechanics, understand this distinction:

| Term | Definition | Who performs it |
|---|---|---|
| **Certification** | Comprehensive technical analysis of a solution to confirm it meets desired needs | Security / technical function |
| **Accreditation** | Official management decision to use the certified solution for a predetermined period | Asset owner / management (not security) |

Accreditation is time-limited — at expiration, the certification/accreditation process repeats. Accreditation is not performed by security staff.

## Common Criteria Components

The CC evaluation process involves four components working together:

| Component | Abbreviation | Description |
|---|---|---|
| **Protection Profile** | PP | Specification of functional and assurance requirements for a *category* of security product (e.g., "all firewalls should support 2FA, VPN, 128-bit encryption, and secure logging") |
| **Target of Evaluation** | TOE | The specific vendor product being evaluated |
| **Security Targets** | ST | Written statement by the vendor describing how the product's security capabilities meet the PP requirements |
| **Evaluation** | — | Independent testing of functional and assurance security requirements against the ST and PP |

### Evaluation Dimensions

Each ST is evaluated against two dimensions:
- **Security Functional Requirements** — what features exist and how well they work relative to expected security behavior
- **Security Assurance Requirements** — confidence that the claimed security functionality is correctly implemented and the evaluation process aligns with CC

## Evaluation Assurance Levels (EAL 1–7)

After evaluation, an overall **EAL** is assigned:

| EAL | Name | Description |
|---|---|---|
| **EAL 7** | Formally verified, designed, and tested | Highest rigor; mathematical proof of correctness |
| **EAL 6** | Semi-formally verified, designed, and tested | |
| **EAL 5** | Semi-formally designed and tested | |
| **EAL 4** | Methodically designed, tested, and reviewed | Most common ceiling for commercial products |
| **EAL 3** | Methodically tested and checked | Typical for operating systems |
| **EAL 2** | Structurally tested | |
| **EAL 1** | Functionally tested | Lowest meaningful assurance |

**Memory aid:** "Formally Verified, Semi-Formally V/D, Semi-Formally D, Methodically DTR, Methodically TC, Structurally T, Functionally T" — from EAL 7 down to 1.

### Practical Note on EAL Levels

Higher is not always better in practice:
- Most organizations will not purchase products rated **above EAL 4** — higher-rated products are more complex, harder to maintain, and more expensive
- Common benchmarks: operating systems ≈ EAL 3, firewalls ≈ EAL 4
- An EAL 7 product may actually increase risk if its complexity causes administrators to misconfigure or under-use its features

### EAL Persistence

Once assigned, an EAL level **does not change** when patches or minor software updates are applied. Only a **major change in functionality** triggers re-evaluation (at the vendor's discretion and cost). CC evaluation is voluntary.

## Comparison: TCSEC → ITSEC → Common Criteria

| Feature | TCSEC (Orange Book) | ITSEC | Common Criteria |
|---|---|---|---|
| Origin | US DoD, 1980s | European, post-Orange Book | International (multi-country joint) |
| Scope | Confidentiality only; single-box | Confidentiality + CIA; networked | All security properties; global |
| Ratings | D, C1, C2, B1, B2, B3, A1 | F levels + E levels (E0–E6) | EAL 1–7 |
| Status | Superseded | Replaced by CC in 2005 | Current standard (ISO/IEC 15408) |

See [TCSEC](../standards/d3a/tcsec.md) for the full Orange Book rating table.

## Exam-Relevant Nuance

- Common Criteria = ISO/IEC 15408 — know both names.
- PP defines requirements for a **product category**; ST defines requirements for a **specific product**.
- TOE is the product being evaluated — not the organization.
- Accreditation is management's job, not the security team's job.
- EAL does not change with patches — only major functionality changes warrant re-evaluation.
- Most products commercially target EAL 4 — high enough to signal quality, not so high as to be impractical.

## Cross-links

- [Security Models](security-models.md) — models that CC evaluates implementations of
- [Trusted Computing Base](trusted-computing-base.md) — the TCB is what CC products must protect
- [TCSEC Standard](../standards/d3a/tcsec.md) — historical predecessor
- [Common Criteria Standard](../standards/d3a/common-criteria-standard.md) — standard page with ISO reference

## Sources

- destination-cissp §3.2.5–3.2.7 (pp. 257–273)
- cissp-exam-outline §3.2, §3.3
