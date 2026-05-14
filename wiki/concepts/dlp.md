---
title: "Data Loss Prevention (DLP)"
type: concept
domain: 2
tags: [dlp, data-loss-prevention, endpoint, network, cloud, data-exfiltration]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Loss Prevention (DLP)

**NIST definition**: "A system's ability to identify, monitor, and protect data in use (e.g., endpoint actions), data in motion (e.g., network actions) and data at rest (e.g., data storage) through deep packet content inspection, contextual security analysis of transaction (attributes of the originator, data object, medium, timing, recipient/destination, etc.), within a centralized management framework."

DLP is broader and more encompassing than DRM: DRM protects IP rights; DLP protects *any* sensitive data across all three states.

Source: destination-cissp §2.6.6 (pages 0220–0221).

## Key facts

- DLP operates across all three data states: **data in use**, **data in motion**, **data at rest**
- Uses **deep packet content inspection** and **contextual analysis** (who is sending, what, to whom, when, via which medium)
- Managed from a **centralized management framework**
- Can operate in **monitoring mode** (log and alert only) or **blocking mode** (block/redirect the transfer)

## DLP deployment contexts

| Context | Data State | How DLP works |
|---|---|---|
| **Endpoint DLP** | Data in use | Monitors actions on endpoints (copy to USB, print, email attachment). Enforces acceptable use policies. |
| **Network DLP** | Data in motion | Inspects data traversing the network; blocks or alerts on suspicious data movement |
| **Storage DLP (Cloud DLP)** | Data at rest | Scans data stored on employee endpoints and network/cloud storage locations for sensitive content |

## What DLP protects

Data protection matters for multiple categories:
- Organizational proprietary data (trade secrets, R&D)
- Customer PII
- Vendor/partner data
- Employee data

Increasingly stringent regulations (GDPR, PCI DSS, HIPAA, CCPA) and litigation risk make DLP a compliance necessity, not just a best practice.

## DLP vs. DRM distinction (exam)

| | DRM | DLP |
|---|---|---|
| Scope | Intellectual property (IP), media content | Any sensitive data |
| Mechanism | Licensing, encryption, digital tags | Content inspection, contextual analysis |
| Focus | Usage rights and distribution control | Data loss/exfiltration prevention |
| Coverage | Primarily data sharing/distribution | Data in use, in transit, at rest |

## Exam-relevant nuance

- DLP tools work *with* acceptable use policies — policy defines what is allowed; DLP enforces it technically
- DLP and DRP (Digital Rights Protection) can limit specific actions users take on data even when access is authorized
- CASB (Cloud Access Security Broker) extends DLP-like capabilities to cloud services — **not covered in destination-cissp** but listed in cissp-exam-outline §2.6. (Content gap — see report.)

## Cross-links

- [Data States](data-states.md) — DLP operates across all three states
- [Data Security Controls](data-security-controls.md) — DLP as part of the control baseline
- [Privacy and PII](privacy-pii.md) — PII is a primary DLP target
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.6.6 (pages 0220–0221)
- cissp-exam-outline §2.6 (data protection methods: DLP, CASB)
