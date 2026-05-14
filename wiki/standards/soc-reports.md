---
title: "SOC Reports (SOC 1 / SOC 2 / SOC 3)"
type: standard
domain: 6
tags: [soc, ssae18, aicpa, isae3402, type1, type2, trust-services-criteria, third-party-audit]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# SOC Reports (SOC 1 / SOC 2 / SOC 3)

## Background and Standard

**System and Organization Controls (SOC)** reports are third-party audit reports produced under **SSAE 18** (Statement on Standards for Attestation Engagements No. 18), the current US standard issued by the **AICPA** (American Institute of Certified Public Accountants).

**Standard evolution** (destination-cissp §6.5.2):

> SAS 70 → SSAE 16 → **SSAE 18** (current)

**International equivalent**: **ISAE 3402** — similar to SSAE 16/18, with minor variations. Used outside the United States.

SOC audits exist to build trust between service organizations and their customers by providing independently verified assurance about the service provider's controls.

## SOC Report Type Comparison

| | SOC 1 | SOC 2 | SOC 3 |
|---|---|---|---|
| **Focus** | Financial reporting controls | Security, availability, confidentiality, processing integrity, privacy | Same as SOC 2 (sanitized) |
| **Standard** | SSAE 18 / AT-C Section 320 | SSAE 18 / Trust Services Criteria | SSAE 18 |
| **Audience** | Customers and their financial auditors | Security professionals, customers evaluating the provider | General public, prospective customers |
| **Confidentiality** | Restricted distribution | Restricted distribution (contains sensitive system details) | **Public** — can be published on website |
| **Detail level** | Financial controls | Detailed control descriptions and system information | Summary/marketing document |
| **Compliance use** | Yes | Yes | **No** — cannot substitute for SOC 2 in compliance |
| **Type 1/Type 2?** | Yes | Yes | No — no Type 1/2 distinction |
| **Most important for security professionals** | No | **Yes — SOC 2** | No |

## AICPA Trust Services Criteria (for SOC 2)

SOC 2 reports evaluate controls against five **Trust Services Criteria (TSC)**:

| Criterion | Required in SOC 2? | Description |
|---|---|---|
| **Security** | Always | Protection of the system against unauthorized access (logical and physical). The foundational criterion. |
| **Availability** | Always | System is available for operation and use as committed or agreed. |
| **Confidentiality** | Always | Information designated as confidential is protected as committed. |
| **Processing Integrity** | Optional | System processing is complete, valid, accurate, timely, and authorized. |
| **Privacy** | Optional | Personal information is collected, used, retained, disclosed, and disposed of according to the privacy notice. |

Note: Processing integrity and privacy are optional — a given SOC 2 may or may not include them.

## Type 1 vs Type 2 Reports

Both SOC 1 and SOC 2 come in two types:

| | **Type 1** | **Type 2** |
|---|---|---|
| **Scope** | Controls at **a point in time** | Controls over **a period of time** (typically ~1 year) |
| **What auditor examines** | Policy and procedure documentation — are controls properly *designed*? | Everything in Type 1 + sampling of actual control *operations* during the period |
| **Auditor conclusion** | Controls are appropriately designed as of [date] | Controls are both designed and operating effectively throughout [period] |
| **Provides operational assurance?** | No | **Yes** |
| **When used** | First year of auditing; identifies design gaps | Ongoing annually; demonstrates sustained control effectiveness |
| **More rigorous** | No | **Yes — Type 2 is more comprehensive** |

**Best practice progression** (destination-cissp §6.5.2):
- Year 1: SOC 2, Type 1 → auditor identifies gaps, organization remediates.
- Year 2+: SOC 2, Type 2 → demonstrates operational control continuity.

**The gold standard for a customer's security team**: SOC 2, Type 2 — attests to both the design and sustained operating effectiveness of security controls.

## When Each Report is Required

| Scenario | Recommended Report |
|---|---|
| Payroll processor, financial transaction service | SOC 1 Type 2 |
| SaaS provider being evaluated by enterprise customer | SOC 2 Type 2 |
| Cloud provider publishing general security assurance | SOC 3 |
| New provider in first year of audit program | SOC 2 Type 1 (first year), then Type 2 |
| International service provider | ISAE 3402 (equivalent to SOC under international standard) |

## Cross-links

- [Audit Types](../concepts/audit-types.md)
- [Assessment vs Audit](../concepts/assessment-vs-audit.md)

## Sources

- destination-cissp §6.5.2 (SOC 1/2/3, Type 1/Type 2, SSAE 18, AICPA, ISAE 3402, Trust Services Criteria)
- cissp-exam-outline (domain 6 audit standards)
