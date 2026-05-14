---
title: "Threat Modeling Concepts and Methodologies"
type: concept
domain: 1
tags: [threat-modeling, stride, pasta, dread, attack-trees, risk-analysis]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Threat Modeling Concepts and Methodologies

Threat modeling provides a **systematic, deliberate means of identifying, enumerating, and
prioritizing threats** to an asset. It feeds directly into risk analysis and makes risk management
more accurate and effective.

## Purpose

Threat modeling answers: "What are all the threats to this asset, and how severe are they?"
It applies to: mobile phones, servers, applications, networks, architectures, functions, processes.

## STRIDE (Microsoft — Threat-Focused)

STRIDE is less strategic/thorough than PASTA. Originally developed for applications and OS,
but applicable in broader contexts.

| Letter | Threat | CIA Pillar Violated | Definition |
|---|---|---|---|
| **S** | Spoofing | Authentication | Attacker impersonates a user/system to gain unauthorized access |
| **T** | Tampering | Integrity | Attacker modifies data at rest or in transit |
| **R** | Repudiation | Nonrepudiation | Attacker performs an action that is not attributable to them |
| **I** | Information Disclosure | Confidentiality | Attacker reads sensitive information they shouldn't access |
| **D** | Denial of Service | Availability | Attacker prevents legitimate users from accessing a service |
| **E** | Elevation of Privilege | Authorization | Attacker gains elevated (e.g., admin/root) access rights |

## PASTA (Risk-Centric, 7 Stages)

**Process for Attack Simulation and Threat Analysis**. Unlike STRIDE, PASTA is
**attacker-focused and risk-centric**. It is strategic and includes input from governance,
operations, architecture, and development — from both business and technical viewpoints (Table 1-28).

| Stage | Focus |
|---|---|
| 1 | Define Objectives — application risk profile, business impact |
| 2 | Define Technical Scope — technology stack decomposition |
| 3 | Application Decomposition — data flows among components |
| 4 | Threat Analysis — threat intelligence relevant to the service/deployment |
| 5 | Vulnerability & Weakness Analysis — design and code-level weaknesses |
| 6 | Attack Modeling — emulate attacks to determine viability |
| 7 | Risk and Impact Analysis — remediate; risk acceptance decisions |

## DREAD (Risk Scoring — used with STRIDE)

DREAD rates the **severity** of threats identified (e.g., by STRIDE). Each of 5 factors scored 1–10;
final score = sum ÷ 5 (out of 10).

| Letter | Factor | Question |
|---|---|---|
| **D** | Damage | Total damage the threat can cause? |
| **R** | Reproducibility | How easily can the threat be replicated? |
| **E** | Exploitability | How difficult is it to exploit? |
| **A** | Affected Users | How many people are affected? |
| **D** | Discoverability | How easily can the threat be discovered? |

Typical usage: STRIDE identifies what threats exist → DREAD ranks them by severity.

## Exam-Relevant Nuance

- STRIDE = **threat-focused** (identifies type of threat); PASTA = **risk-centric/attacker-focused** (full risk analysis from strategic and technical perspectives).
- DREAD is used to **rank/score** threats, not identify them.
- On the exam, if asked which methodology is "more thorough/strategic," the answer is **PASTA**.
- "Elevation of Privilege" in STRIDE maps to **authorization** (not authentication) — a common exam trap.
- Attack trees are another methodology (not in this source but may appear on exam): tree diagrams where root = goal of attacker, branches = methods.

## Cross-links

- [Risk Management](risk-management.md)
- [CIA Triad](cia-triad.md)
- [SCRM](scrm.md)

## Sources

- destination-cissp §1.10 (Tables 1-27, 1-28, 1-29)
