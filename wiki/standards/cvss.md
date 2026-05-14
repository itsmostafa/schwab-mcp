---
title: "CVSS — Common Vulnerability Scoring System"
type: standard
domain: 6
tags: [cvss, cve, vulnerability-scoring, risk-prioritization, nvd]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# CVSS — Common Vulnerability Scoring System

## Overview

The **Common Vulnerability Scoring System (CVSS)** is an open framework for communicating the characteristics and severity of software vulnerabilities. It produces a numerical score (0.0–10.0) — the higher the score, the more severe and critical the vulnerability.

**Current version**: CVSS v3.1 (the version aligned with CISSP exam content).

**Maintained by**: FIRST (Forum of Incident Response and Security Teams), with the NVD (National Vulnerability Database, maintained by NIST) being the primary distribution point.

**Used with**: CVE identifiers. Scanner reports pair a CVE ID (identifying which vulnerability) with a CVSS score (quantifying severity). See [Vulnerability Assessment](../concepts/vulnerability-assessment.md).

## Score Ranges (CVSS v3.1)

| Score Range | Rating |
|---|---|
| 9.0 – 10.0 | **Critical** |
| 7.0 – 8.9 | **High** |
| 4.0 – 6.9 | **Medium** |
| 0.1 – 3.9 | **Low** |
| 0.0 | None |

## CVSS Metric Groups (v3.1)

The CVSS score is composed of three metric groups. Note: destination-cissp only covers the top-line scoring concept; the metric breakdown below is standard CISSP exam knowledge per cissp-exam-outline.

### Base Score Metrics

Represent inherent characteristics of the vulnerability, independent of time or deployment environment. The Base Score is what is most commonly reported.

**Exploitability sub-metrics:**

| Metric | Description | Values |
|---|---|---|
| **Attack Vector (AV)** | How the vulnerability can be exploited | Network (N), Adjacent (A), Local (L), Physical (P) |
| **Attack Complexity (AC)** | Conditions beyond attacker control required for exploitation | Low (L), High (H) |
| **Privileges Required (PR)** | Level of privileges the attacker must already have | None (N), Low (L), High (H) |
| **User Interaction (UI)** | Whether user action is required | None (N), Required (R) |
| **Scope (S)** | Does exploitation affect only the vulnerable component or can it impact other components? | Unchanged (U), Changed (C) |

**Impact sub-metrics (CIA):**

| Metric | Description | Values |
|---|---|---|
| **Confidentiality Impact (C)** | Impact on confidentiality of information | None (N), Low (L), High (H) |
| **Integrity Impact (I)** | Impact on integrity of information | None (N), Low (L), High (H) |
| **Availability Impact (A)** | Impact on availability of affected system | None (N), Low (L), High (H) |

### Temporal Score Metrics

Modify the Base Score based on the current state of exploit techniques and mitigations:
- **Exploit Code Maturity** — availability and reliability of exploit code.
- **Remediation Level** — availability of a fix (official fix, workaround, unavailable).
- **Report Confidence** — degree of confidence in the vulnerability's existence and details.

### Environmental Score Metrics

Allow organizations to customize the score based on their specific environment:
- **Modified Base Metrics** — adjust Base metrics for their infrastructure.
- **Confidentiality/Integrity/Availability Requirements** — weight CIA impacts based on asset criticality in their environment.

## CVE vs CVSS

| | CVE | CVSS |
|---|---|---|
| **Full name** | Common Vulnerabilities and Exposures | Common Vulnerability Scoring System |
| **Purpose** | Unique identifier + description for each known vulnerability | Severity score (0–10) for each vulnerability |
| **Maintained by** | MITRE (with sponsorship from CISA/US-CERT) | FIRST |
| **Distributed via** | NVD (NIST), CVE.org | NVD, vendor advisories |
| **Together they provide** | Identity + severity = full vulnerability intelligence |

## Practical Use

1. Vulnerability scanner identifies a weakness in a scanned system.
2. The finding is mapped to a **CVE ID** (e.g., CVE-2021-44228 = Log4Shell).
3. The **CVSS Base Score** for that CVE (e.g., 10.0 — Critical) is pulled from NVD.
4. Organization applies **Environmental metrics** to adjust for their context (if they have compensating controls, the environmental score may be lower).
5. Patches and remediations are **prioritized** by CVSS score (Critical → High → Medium → Low).

## Cross-links

- [Vulnerability Assessment](../concepts/vulnerability-assessment.md)

## Sources

- destination-cissp §6.2.4 (CVE and CVSS definitions, 0–10 scoring concept)
- cissp-exam-outline (CVSS v3.1 metric groups, score ranges, NVD)
