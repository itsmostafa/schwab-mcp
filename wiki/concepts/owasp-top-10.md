---
title: "OWASP Top 10 (2021)"
type: concept
domain: 8
tags: [owasp, top-10, web-application-security, injection, xss, access-control, broken-auth, ssrf, supply-chain]
sources: [cissp-exam-outline]
updated: 2026-05-13
---

# OWASP Top 10 (2021)

## Overview

The **OWASP Top 10** is a consensus-based list of the most critical web application security risks, published by the Open Web Application Security Project (OWASP). Updated approximately every three to four years; the current authoritative version is the **2021 edition**. It is the most widely recognized baseline for web application security awareness and serves as a foundational reference for secure coding guidelines.

> **Coverage note:** destination-cissp §8.5.1 references OWASP Top 10 as a key framework and notes that "OWASP publishes a Top 10 Web Application Security Risks document that includes a list of risks, their descriptions, and mitigation strategies." However, the source does not enumerate each category. The content below reflects the 2021 OWASP Top 10 as required by cissp-exam-outline scope.

---

## The 2021 List

| Rank | Category | Core Risk | Key Mitigation |
|---|---|---|---|
| **A01** | **Broken Access Control** | Users can act outside of their intended permissions — access other users' data, escalate privileges, perform admin functions. Most common and most serious. | Enforce least privilege; deny by default; server-side access control checks; log access failures. |
| **A02** | **Cryptographic Failures** (formerly "Sensitive Data Exposure") | Failure to protect data in transit or at rest with adequate cryptography — or using weak/deprecated algorithms. Exposes passwords, credit card numbers, PHI. | Encrypt sensitive data at rest and in transit (TLS); use strong modern algorithms; never roll your own crypto. |
| **A03** | **Injection** | Attacker-supplied data is interpreted as commands or queries — SQL injection, LDAP injection, OS command injection, NoSQL injection. | Parameterized queries (primary defense); input validation; principle of least privilege for DB accounts; ORM. |
| **A04** | **Insecure Design** | Missing or inadequate security controls due to design-level failures — not implementation bugs. Reflects absent threat modeling and security requirements. | Threat modeling during design; secure design patterns; established security design principles. |
| **A05** | **Security Misconfiguration** | Insecure default configurations, incomplete configurations, open cloud storage, unnecessary features enabled, default credentials unchanged. | Hardening; disable unnecessary features; change defaults; automated configuration scanning. |
| **A06** | **Vulnerable and Outdated Components** | Using libraries, frameworks, or components with known vulnerabilities. Includes both frontend and backend dependencies. Ties directly to supply chain risk. | Inventory all dependencies (SBOM); subscribe to CVE notifications; regularly update and patch; dependency scanning (SCA). |
| **A07** | **Identification and Authentication Failures** (formerly "Broken Authentication") | Weaknesses in identity verification — weak passwords permitted, credential stuffing not prevented, session tokens not invalidated, missing MFA. | MFA; strong password policies; secure session management; rate limiting on login. |
| **A08** | **Software and Data Integrity Failures** | Code and infrastructure not protected against integrity violations — insecure deserialization, CI/CD pipeline compromises, auto-update without signature verification. | Code signing; verify integrity of dependencies (hashes); secure CI/CD; deserialization input validation. |
| **A09** | **Security Logging and Monitoring Failures** | Insufficient logging, monitoring, and alerting — attackers operate undetected; forensic evidence unavailable after breach. | Log all authentication, access control, and input validation failures; centralize logs; alert on anomalies; test monitoring. |
| **A10** | **Server-Side Request Forgery (SSRF)** | Server fetches a remote resource based on user-supplied URL, allowing attacker to reach internal services not exposed to the internet (cloud metadata APIs, internal databases). | Validate and sanitize all URLs; use allowlists for allowed destinations; block private IP ranges in outbound requests; network segmentation. |

---

## Mapping to CISSP Domain 8 Topics

| OWASP Category | Related D8 Concept |
|---|---|
| A01 Broken Access Control | [Secure Coding Practices](secure-coding-practices.md) — principle of least privilege in code |
| A02 Cryptographic Failures | [Secure Coding Practices](secure-coding-practices.md) — cryptographic practices |
| A03 Injection | [Database Security](database-security.md) — SQL injection; [Secure Coding Practices](secure-coding-practices.md) — input validation |
| A04 Insecure Design | [SDLC](sdlc.md) — threat modeling in design phase |
| A05 Security Misconfiguration | [Secure Coding Practices](secure-coding-practices.md) — secure defaults; system configuration |
| A06 Vulnerable & Outdated Components | [Software Acquisition Security](software-acquisition-security.md) — open source risk; SBOM; [CI/CD Security](ci-cd-security.md) — dependency scanning |
| A07 Authentication Failures | Cross-domain: [Domain 5 IAM](../domains/05-identity-and-access-management.md) |
| A08 Software & Data Integrity Failures | [CI/CD Security](ci-cd-security.md) — code signing, supply chain |
| A09 Logging & Monitoring Failures | Cross-domain: [Domain 6 — Security Assessment](../domains/06-security-assessment-and-testing.md); [Domain 7 — Security Operations](../domains/07-security-operations.md) |
| A10 SSRF | [API Security](api-security.md) — SSRF in API context; input validation |

---

## Exam-Relevant Nuance

- **A01 Broken Access Control** rose to #1 in 2021 (previously #5). It is the most frequently occurring risk.
- **A04 Insecure Design** is new in 2021. It emphasizes that security failures begin in the design phase — not just in implementation. Ties directly to SDLC threat modeling.
- **A08 Software & Data Integrity Failures** (new in 2021) captures the software supply chain risk category — SolarWinds-type attacks.
- **A09 Logging & Monitoring Failures** is not a vulnerability class that attackers exploit directly — it is a *detection* gap that allows other vulnerabilities to be exploited without consequence.
- The OWASP Top 10 is for **web applications** specifically. OWASP also publishes a separate **API Security Top 10** — see [API Security](api-security.md).
- All 10 categories have appeared on CISSP exams. Know each one by its 2021 name and what it means.

---

## Quick Memory Aid

**BICIISVLSS** (A01 → A10):
- **B**roken Access Control
- **C**ryptographic Failures  
- **I**njection  
- **I**nsecure Design  
- **S**ecurity Misconfiguration  
- **V**ulnerable & Outdated Components  
- **I**dentification & Authentication Failures  
- **S**oftware & Data Integrity Failures  
- **L**ogging & Monitoring Failures  
- **S**SRF

*(This mnemonic is imperfect — use only if it genuinely helps. The category names themselves are descriptive enough to remember with practice.)*

---

## Cross-Links

- [OWASP](../standards/owasp.md) — OWASP organization and full suite of projects
- [Secure Coding Practices](secure-coding-practices.md) — secure coding response to OWASP risks
- [Database Security](database-security.md) — A03 Injection detail
- [API Security](api-security.md) — OWASP API Security Top 10
- [CI/CD Security](ci-cd-security.md) — A06, A08 supply chain context
- [SDLC](sdlc.md) — A04 Insecure Design → threat modeling in SDLC

## Sources

- destination-cissp §8.5.1 (OWASP reference; does not enumerate categories) (p. 0965)
- cissp-exam-outline (Domain 8.5: Secure coding guidelines and standards)
- OWASP Top 10 2021 (owasp.org/Top10) — primary source for category enumeration
