---
title: "NIST SP 800-64 — Security Considerations in the SDLC"
type: standard
domain: 8
tags: [nist, sdlc, security-considerations, software-development, sp-800-64]
sources: [cissp-exam-outline]
updated: 2026-05-13
---

# NIST SP 800-64 — Security Considerations in the System Development Life Cycle

## Overview

**NIST Special Publication 800-64, Revision 2** — *Security Considerations in the System Development Life Cycle* — provides guidance on integrating information technology security into the SDLC. It establishes a framework for selecting, implementing, and verifying security controls at each SDLC phase.

> **Coverage gap:** destination-cissp does not reference NIST SP 800-64 explicitly in Domain 8. This page is a stub based on cissp-exam-outline scope and NIST's published content. The source confirms NIST and OWASP as key organizations providing secure software development frameworks (§8.5.1).

**Note:** NIST SP 800-64 was **withdrawn in 2022** and superseded by content in **NIST SP 800-160** (Systems Security Engineering) and guidance integrated into the **NIST Cybersecurity Framework (CSF)**. For exam purposes, SP 800-64 may still appear as a reference; understand its core message.

---

## Core Message

- Security must be integrated at every phase of the SDLC — not added at the end.
- The cost of fixing security vulnerabilities increases dramatically at later SDLC phases. A defect found in design costs far less to fix than one found in production.
- Each SDLC phase has specific security activities, responsibilities, and deliverables.

---

## SDLC Phases (NIST SP 800-64 Framework)

| Phase | Key Security Activities |
|---|---|
| **Initiation** | Categorize the system (FIPS 199); conduct preliminary risk assessment; identify security requirements. |
| **Acquisition / Development** | Conduct risk assessment; perform security functional requirements analysis; design security controls; incorporate security into design. |
| **Implementation** | Integrate security controls; conduct security testing; perform certification (C&A) activities; receive authorization to operate (ATO). |
| **Operations / Maintenance** | Operate within ATO boundaries; continuous monitoring; change management; configuration management; incident handling. |
| **Disposal** | Archive records; sanitize media; close system; execute disposal plan. |

---

## Relation to RMF

NIST SP 800-64 aligns with the **Risk Management Framework (RMF)** defined in NIST SP 800-37. The RMF's steps (Categorize → Select → Implement → Assess → Authorize → Monitor) map to SDLC phases and provide the current guidance framework for federal systems.

---

## Exam-Relevant Nuance

- Know that NIST produces guidance for SDLC security (even if SP 800-64 is withdrawn; its concepts persist in NIST SP 800-37/160).
- The principle that security costs less when built in early (shift left) is the core message of all SDLC security guidance.
- Certification and Accreditation (C&A) — now called Assessment and Authorization (A&A) under the RMF — occur during the Implementation phase.

---

## Cross-Links

- [SDLC](../concepts/sdlc.md) — SDLC phases and security activities
- [CMMI/SAMM](cmm-samm.md) — maturity models for SDLC improvement

## Sources

- destination-cissp §8.5.1 (NIST as source of secure software development frameworks) (p. 0965)
- cissp-exam-outline (Domain 8: SDLC security)
- NIST SP 800-64 Rev. 2 (csrc.nist.gov) — withdrawn 2022; content now in NIST SP 800-160 and CSF
