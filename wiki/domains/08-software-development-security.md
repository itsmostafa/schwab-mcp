---
title: "Domain 8 — Software Development Security"
type: domain
domain: 8
tags: [sdlc, devsecops, secure-coding, sast, dast, supply-chain, api-security]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 8 — Software Development Security

Exam weight: **10%**. Covers integrating security into the SDLC, securing development
ecosystems, assessing software security, and secure coding practices.

## Subtopics (from CISSP Exam Outline)

### 8.1 Understand and integrate security in the Software Development Life Cycle (SDLC)
- Development methodologies (Agile, Waterfall, DevOps, DevSecOps, SAFe)
- Maturity models (CMM, SAMM)
- Operation and maintenance
- Change management
- Integrated Product Team

Security must be embedded at every SDLC phase — from initiation through disposal — because retrofitting security is costlier and less effective than building it in from the start. Threat modeling belongs in the design phase; SAST/DAST/fuzz testing belong in the testing phase; certification (technical) and accreditation (management sign-off) occur at deployment. The Integrated Product Team (IPT) is the exam name for DevOps; DevSecOps (or SecDevOps) is DevOps with security integral from day one. Maturity models — CMMI (six levels, 0–5) and OWASP SAMM (three levels, five business functions) — measure and guide improvement of the development process.

### 8.2 Identify and apply security controls in software development ecosystems
- Programming languages
- Libraries
- Tool sets
- Integrated Development Environment (IDE)
- Runtime
- Continuous Integration and Continuous Delivery (CI/CD)
- Code repositories
- Software configuration management (CM)
- Managed services (enterprise applications)
- Cloud services (SaaS, IaaS, PaaS)

The development ecosystem includes programming languages (generations 1–5 from machine code to natural language), static and dynamic libraries, SDKs, IDEs, and the CI/CD pipeline that automates integration, testing, and deployment. Software Configuration Management (SCM) establishes baselines and revision control for all code changes. Code repositories (GitHub, GitLab, etc.) must be secured with access controls, branch protection, and audit logging. Database environments require attention to RDBMS concepts (attributes/tuples, primary/foreign keys, ACID properties), concurrency/locking controls, and code obfuscation techniques (lexical, data, control-flow). Development, test, QA, and production environments must be segregated.

### 8.3 Assess the effectiveness of software security
- Auditing and logging of changes
- Risk analysis and mitigation

Assessing software security effectiveness requires ongoing monitoring of logs, periodic risk analysis and mitigation, and formal assessment activities including white-box and black-box testing, threat modeling, and penetration testing. Certification (technical evaluation) and accreditation (management authorization) provide formal assurance at deployment. Security must participate in procurement, internal and external audit, and the change management process throughout the operational life of the software.

### 8.4 Assess security impact of acquired software
- Commercial-off-the-shelf (COTS)
- Open source
- Third-party

Acquired software requires the same security rigor as internally developed software. The four acquisition phases are: planning/requirements, contracting, acceptance, and monitoring/follow-on. COTS software does not allow white-box testing — black-box acceptance testing and software escrow (for source code access if the vendor fails) are key controls. Open source software must be treated as if developed in house (the Heartbleed/OpenSSL incident is the canonical risk example). Third-party developers must be held to contractual security requirements. Cloud services fall under the shared responsibility model, with the customer always remaining ultimately accountable.

### 8.5 Define and apply secure coding guidelines and standards
- Security weaknesses and vulnerabilities at the source-code level
- Security of APIs
- Secure coding practices
- Software-defined security
- Application security testing:
  - Static Application Security Testing (SAST)
  - Dynamic Application Security Testing (DAST)
  - Software composition analysis
  - Interactive Application Security Testing (IAST)

Secure coding practices include input validation, proper error handling (never expose internals to users), session management, cryptographic best practices, and memory management. Source-level vulnerabilities to know: buffer overflow, TOCTOU (race conditions), covert channels (timing and storage types), memory/object reuse, executable mobile code, backdoors/trapdoors, malformed input, and citizen developer risks. APIs (REST vs. SOAP) must be secured with authentication/authorization (OAuth), TLS, rate limiting, data validation, and API gateways. Polyinstantiation prevents unauthorized inference in multi-level security database environments. Low coupling + high cohesion is the target for well-designed, secure code.

## Key Concepts

- [SDLC](../concepts/sdlc.md) — phases, security activities, methodologies, certification/accreditation
- [DevSecOps](../concepts/devsecops.md) — IPT, DevOps vs DevSecOps, canary testing
- [Secure Coding Practices](../concepts/secure-coding-practices.md) — input validation, coupling/cohesion, polyinstantiation, source-level vulnerabilities
- [Memory Safety](../concepts/memory-safety.md) — buffer overflow, ASLR, DEP, stack canaries, integer overflow, use-after-free
- [Database Security](../concepts/database-security.md) — DBMS, ACID, SQL injection, parameterized queries, polyinstantiation, inference/aggregation attacks
- [API Security](../concepts/api-security.md) — REST vs. SOAP, OAuth, TLS, rate limiting, OWASP API Top 10
- [Software Acquisition Security](../concepts/software-acquisition-security.md) — COTS, open source, third-party, escrow, SBOM, shared responsibility
- [CI/CD Security](../concepts/ci-cd-security.md) — pipeline security, SCM, code signing, supply chain attacks, SBOM
- [Change Management (Software)](../concepts/change-management-software.md) — RFC process, SCM, version control security, CAB
- [OWASP Top 10](../concepts/owasp-top-10.md) — 2021 list, definitions, exam mapping
- [Security Testing Types (D6)](../concepts/security-testing-types.md) — SAST, DAST, IAST, fuzz testing (cross-domain: also covers 8.5)

## Standards

- [OWASP](../standards/owasp.md) — OWASP organization; Top 10 web and API; SAMM; Testing Guide
- [CMMI / SAMM](../standards/cmm-samm.md) — six CMMI levels (0–5); SAMM three levels and five business functions; BSIMM
- [NIST SP 800-64](../standards/nist-sp-800-64.md) — SDLC security guidance (withdrawn 2022; superseded by NIST SP 800-160)

## Sources
- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §8.1–8.5 (pp. 0925–0972)
