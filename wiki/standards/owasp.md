---
title: "OWASP — Open Web Application Security Project"
type: standard
domain: 8
tags: [owasp, web-security, api-security, samm, top-10, non-profit, community]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# OWASP — Open Web Application Security Project

## Overview

**OWASP (Open Web Application Security Project)** is a nonprofit, community-driven organization dedicated to improving software security. OWASP produces freely available tools, documentation, and standards that help organizations develop, purchase, and maintain software they can trust. All OWASP content is available under open licenses.

*(source: destination-cissp §8.5.1)*

---

## Key OWASP Projects (CISSP-Relevant)

| Project | Description |
|---|---|
| **OWASP Top 10 (Web)** | The most well-known OWASP project. Consensus list of the top 10 most critical web application security risks. Updated every 3–4 years. 2021 edition is current. See [OWASP Top 10](../concepts/owasp-top-10.md). |
| **OWASP API Security Top 10** | Equivalent list for API-specific risks. Most recent: 2023 edition. Covers broken object-level authorization, broken authentication, excessive data exposure, SSRF, etc. See [API Security](../concepts/api-security.md). |
| **OWASP SAMM** | Software Assurance Maturity Model — a maturity model for building security into software development programs. See [SAMM](cmm-samm.md). |
| **OWASP Testing Guide** | Comprehensive methodology for testing web application security. Covers manual and automated testing approaches across all vulnerability classes. |
| **OWASP Mobile Security Testing Guide (MASTG)** | Security testing methodology for mobile applications (iOS, Android). Companion: OWASP Mobile Top 10. |
| **OWASP Application Security Verification Standard (ASVS)** | A framework of security requirements for web application design, development, and testing. Defines security verification levels. |
| **OWASP Cheat Sheet Series** | Short, actionable guidance on specific secure development topics (e.g., SQL injection prevention, authentication cheat sheet). |

---

## OWASP Top 10 (Web) — 2021 Edition at a Glance

| Rank | Category |
|---|---|
| A01 | Broken Access Control |
| A02 | Cryptographic Failures |
| A03 | Injection |
| A04 | Insecure Design |
| A05 | Security Misconfiguration |
| A06 | Vulnerable and Outdated Components |
| A07 | Identification and Authentication Failures |
| A08 | Software and Data Integrity Failures |
| A09 | Security Logging and Monitoring Failures |
| A10 | Server-Side Request Forgery (SSRF) |

---

## OWASP API Security Top 10 — 2023 Edition at a Glance

| Rank | Category |
|---|---|
| API1 | Broken Object Level Authorization |
| API2 | Broken Authentication |
| API3 | Broken Object Property Level Authorization |
| API4 | Unrestricted Resource Consumption |
| API5 | Broken Function Level Authorization |
| API6 | Unrestricted Access to Sensitive Business Flows |
| API7 | Server Side Request Forgery |
| API8 | Security Misconfiguration |
| API9 | Improper Inventory Management |
| API10 | Unsafe Consumption of APIs |

---

## Exam-Relevant Nuance

- OWASP is **non-profit and community-driven** — all resources are free.
- The **OWASP Top 10 is the most frequently cited** CISSP reference for web application security risks.
- OWASP does not certify products or organizations — it produces guidance and tools.
- **SAMM** is OWASP's contribution to software development maturity modeling (distinct from CMMI).
- The exam may test awareness that OWASP produces multiple Top 10 lists for different contexts (web, API, mobile).

---

## Cross-Links

- [OWASP Top 10](../concepts/owasp-top-10.md) — detailed breakdown of each web category
- [API Security](../concepts/api-security.md) — OWASP API Security Top 10
- [CMMI/SAMM](cmm-samm.md) — OWASP SAMM maturity model
- [Secure Coding Practices](../concepts/secure-coding-practices.md) — OWASP guidance applied

## Sources

- destination-cissp §8.5.1 (OWASP reference) (p. 0965)
- destination-cissp §8.4.1 (OWASP SAMM reference) (p. 0977)
- cissp-exam-outline (Domain 8.5: Secure coding guidelines and standards)
- OWASP website (owasp.org) — authoritative source for all OWASP projects
