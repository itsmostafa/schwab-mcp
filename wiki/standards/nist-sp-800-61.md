---
title: "NIST SP 800-61 — Computer Security Incident Handling Guide"
type: standard
domain: 7
tags: [NIST, incident-response, IR, CSIRT, 800-61]
sources: [cissp-exam-outline]
updated: 2026-05-13
---

# NIST SP 800-61 — Computer Security Incident Handling Guide

## Bibliographic Information

- **Full title:** Computer Security Incident Handling Guide
- **Publication:** NIST Special Publication 800-61, Revision 2
- **Issuing body:** National Institute of Standards and Technology (NIST), U.S. Dept. of Commerce
- **Current revision:** Rev. 2 (2012) — Rev. 3 draft circulated but Rev. 2 remains the current authoritative version for CISSP purposes.
- **Purpose:** Provides guidelines for establishing an effective incident response (IR) capability and handling security incidents.

---

## Four-Phase Incident Response Model

NIST SP 800-61 defines incident response as a **four-phase process**:

| Phase | Key Activities |
|---|---|
| **1. Preparation** | Establish IR capability: develop IR policy, plan, and procedures; build and train the IR team; acquire tools (forensic software, network analyzers, documentation templates); establish communication channels; conduct exercises. |
| **2. Detection and Analysis** | Identify and validate incidents using precursors (advance warning) and indicators (signs an incident may have occurred or is occurring); categorize and prioritize; assign severity; document initial findings and notify appropriate parties. |
| **3. Containment, Eradication, and Recovery** | Contain the incident (isolate affected systems); collect evidence; eradicate root cause (remove malware, close vulnerabilities, disable compromised accounts); recover systems (restore from clean backups, patch, verify). |
| **4. Post-Incident Activity** | Lessons learned meeting (held within two weeks of incident closure); update IR procedures; evidence retention; provide metrics to management; close out the incident record. |

> **Exam note:** The phases are cyclical — lessons learned from Phase 4 feed back into improving Phase 1 (Preparation).

---

## IR Team Models

NIST SP 800-61 defines three organizational models for IR teams:

| Model | Description | Best For |
|---|---|---|
| **Central IR team** | Single team handles all incidents for the entire organization | Small/medium organizations |
| **Distributed IR teams** | Multiple teams, each handling IR for their own department/business unit | Large, geographically dispersed organizations |
| **Coordinating IR team** | Central team provides guidance and coordination to distributed teams; does not directly handle incidents | Very large organizations, national CERTs |

---

## Incident Categories

NIST SP 800-61 categorizes incidents by type (not exhaustive):
- Denial of Service (DoS/DDoS)
- Malicious code (malware)
- Unauthorized access
- Inappropriate usage
- Scans, probes, attempted access
- Investigations (of unknown incidents)

Incidents should also be prioritized by **functional impact** (effect on business operations), **information impact** (effect on data confidentiality/integrity), and **recoverability** (ease of recovery).

---

## Evidence Retention

NIST recommends defining an evidence retention policy before incidents occur. Factors include:
- Legal requirements (statutes of limitations for prosecution)
- Regulatory requirements
- Business need for post-incident analysis

---

## Relationship to destination-cissp 8-Step Model

The destination-cissp 8-step model (Preparation → Detection → Response → Mitigation → Reporting → Recovery → Remediation → Lessons Learned) maps to NIST's 4 phases:

| NIST Phase | destination-cissp Steps |
|---|---|
| Preparation | Preparation |
| Detection and Analysis | Detection, Response (impact assessment) |
| Containment, Eradication, Recovery | Mitigation, Recovery, Remediation |
| Post-Incident Activity | Reporting, Lessons Learned |

Both models are testable on the CISSP exam. Prefer the NIST 4-phase framing when the question references a "standard" or "guide."

---

## Exam-Relevant Nuance

- **Preparation is Phase 1** in both models — the exam may try to trick you into thinking it comes last.
- The **lessons learned meeting** is a formal Post-Incident Activity, not optional.
- NIST emphasizes that **both precursors and indicators** should be monitored — many organizations only watch for indicators.
- A **CSIRT** (Computer Security Incident Response Team) is the team model; **CERT** (Computer Emergency Response Team) is often used for government/national-level teams.
- NIST 800-61 is the primary CISSP-referenced standard for incident response — know it more deeply than other IR frameworks.

---

## Cross-Links

- [Incident Management](../concepts/incident-management.md) — concept page covering both models
- [Forensics](../concepts/forensics.md) — forensic evidence collection during Phase 3
- [Chain of Custody](../concepts/chain-of-custody.md) — evidence handling in Phase 3

## Sources

- NIST SP 800-61 Rev. 2 (2012) — primary source
- cissp-exam-outline (Domain 7: incident management)
