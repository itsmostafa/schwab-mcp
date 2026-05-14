---
title: "CMMI and SAMM — Software Development Maturity Models"
type: standard
domain: 8
tags: [cmmi, samm, bsimm, maturity-model, software-assurance, owasp, capability-maturity]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# CMMI and SAMM — Software Development Maturity Models

## Overview

Maturity models provide a structured framework for measuring and improving the capability of an organization's development processes. In the CISSP context, two are most important:
- **CMMI** (Capability Maturity Model Integration) — a general-purpose process improvement model applicable to software development and other disciplines.
- **SAMM** (Software Assurance Maturity Model) — OWASP's maturity model specifically for software security practices.

---

## CMMI — Capability Maturity Model Integration

*(source: destination-cissp §8.1.4)*

CMMI is a set of best practices that helps organizations understand the maturity of their processes, identify strengths and weaknesses, and drive performance improvement. Originally developed at Carnegie Mellon University's Software Engineering Institute (SEI).

### The Six CMMI Maturity Levels

> **Contradiction note:** Some older CISSP materials describe only five CMMI levels (1–5). destination-cissp explicitly defines **six levels (0–5)**, with Level 0 (Incomplete) added. Where wiki content elsewhere describes five levels, this page reflects the 6-level model from destination-cissp. The user should use the 6-level model for exam purposes.

| Level | Name | Characteristics |
|---|---|---|
| **0** | **Incomplete** | Processes are unknown or ad hoc; work may not get completed at all. |
| **1** | **Initial** | Reactive and unpredictable. Work gets done, but often over budget and late. No consistent process. |
| **2** | **Managed** | Projects are planned and managed. Key metrics are tracked. Work is controlled and performed to plan. |
| **3** | **Defined** | Proactive (not reactive). Standardized processes exist across the organization; portfolios, programs, and projects follow them. |
| **4** | **Quantitatively Managed** | Data-driven. Performance objectives are measured quantitatively; processes are predictable and meet stakeholder needs. |
| **5** | **Optimizing** | Flexible, stable, and continuously improving. Organization can adapt and innovate; stability allows agility when opportunity presents. |

**Memory aid for levels 0–5:** "**I Can't Do Quantitative Optimization** — so I'm still at Level 0." (Incomplete, Cannot, Defined, Quantitative, Optimizing)

**Key exam point:** CMMI helps organizations understand *how mature their processes are* — not just whether they're "good." Level 3 (Defined) is the threshold between reactive and proactive.

---

## SAMM — Software Assurance Maturity Model (OWASP)

*(source: destination-cissp §8.1.4)*

**OWASP SAMM** is a maturity model specifically designed for software security. In OWASP's words, SAMM aims "to be the prime maturity model for software assurance that provides an effective and measurable way for all types of organizations to analyze and improve their software security posture."

**Key characteristics:**
- **Supports the complete software life cycle** — both development and acquisition
- **Technology and process agnostic** — works regardless of language, platform, or methodology
- **Evolutive and risk-driven** — organizations start where they are and improve incrementally
- Maturity levels can be measured from both a **coverage** and a **quality** perspective using a software assurance scoring model

### SAMM Maturity Levels

| Level | Name | Description |
|---|---|---|
| **1** | **Initial Implementation** | Basic security practices are in place but incomplete; ad hoc application. |
| **2** | **Structured Realization** | Practices are formalized; consistent application across teams; measured. |
| **3** | **Optimized Operation** | Processes are optimized; continuous improvement; security integrated throughout. |

### SAMM Five Business Functions

SAMM organizes software security activities across five high-level business functions:

| Function | Security Activities |
|---|---|
| **1. Governance** | Strategy, policy, training, compliance |
| **2. Design** | Threat assessment, security requirements, secure architecture |
| **3. Implementation** | Secure build, secure deployment, defect management |
| **4. Verification** | Architecture review, requirements testing, security testing |
| **5. Operations** | Incident management, environment management, operational management |

---

## BSIMM — Building Security In Maturity Model

> **Coverage gap:** BSIMM is not covered in destination-cissp. Included here for cissp-exam-outline completeness.

**BSIMM** (pronounced "bee-simm") is a study of real-world software security practices based on observations across many organizations. Unlike SAMM (prescriptive — tells you what to do), BSIMM is **descriptive** — it describes what organizations actually do, allowing benchmarking against peers.

- Used for: understanding where your software security program stands relative to industry
- Not a prescription; a benchmark
- Updated annually based on new participant data
- Organized into four domains: Governance, Intelligence, SSDL Touchpoints, Deployment

---

## CMMI vs. SAMM Comparison

| Aspect | CMMI | SAMM |
|---|---|---|
| **Focus** | General process improvement (any domain) | Specifically software security |
| **Levels** | 6 levels (0–5) | 3 levels (1–3) |
| **Owner** | ISACA (formerly SEI/CMU) | OWASP |
| **Scope** | Software development, services, product development | Software development and acquisition |
| **Approach** | Certification/appraisal possible | Self-assessment and scoring |

---

## Exam-Relevant Nuance

- **Know all six CMMI levels by name** — the exam may ask which level is described by a scenario.
- **Level 3 (Defined) = proactive** — organizations with standardized processes. This is often the target for security programs.
- **SAMM is OWASP's model**; it has 3 levels and 5 business functions. Different from CMMI's structure.
- CMMI Level 0 (Incomplete) is sometimes omitted in older sources — expect the 6-level model on the current exam.
- **BSIMM** = descriptive/benchmarking; **SAMM** = prescriptive/guidance. Knowing this distinction matters.

---

## Cross-Links

- [SDLC](../concepts/sdlc.md) — maturity models applied to improve SDLC process
- [OWASP](owasp.md) — OWASP organization and other projects
- [DevSecOps](../concepts/devsecops.md) — DevSecOps reflects a mature, integrated development approach

## Sources

- destination-cissp §8.1.4 (CMMI six levels, SAMM three levels and five functions) (pp. 0936–0939)
- cissp-exam-outline (Domain 8.1: Maturity models — CMM, SAMM)
- OWASP SAMM (owaspsamm.org) — authoritative source
- BSIMM (bsimm.com) — coverage gap noted
