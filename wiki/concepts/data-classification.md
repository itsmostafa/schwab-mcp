---
title: "Data Classification"
type: concept
domain: 2
tags: [classification, categorization, labeling, marking, asset-management]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Classification

Classification is a system of ordered classes that represents the level of protection an asset requires, based on its value to the organization. Categorization is the *act* of assigning an asset to a class within that system. The two terms are often confused on the exam.

## Key facts

- **Classification = the system**; **Categorization = the assignment act** (destination-cissp §2.1.3)
- Value is assessed across all three CIA dimensions: confidentiality (sensitivity), integrity (accuracy), and availability (criticality) — not just confidentiality (destination-cissp §2.1.2)
- **Government classification levels** (high to low): Top Secret → Secret → Confidential → Sensitive But Unclassified (SBU) → Unclassified
- **Commercial classification levels** (typical, high to low): Proprietary/Restricted → Confidential/Company Restricted → Internal → Public. *Labels vary per organization; the value assigned to each label also varies.*
- Classification is **driven by the asset owner**, not security or IT
- Classification is an **ongoing process** — values change due to age, legal requirements, competitive factors; periodic review is required
- Classification also drives archiving, retention, and destruction requirements
- A **data classification board or committee** can provide objectivity when individual owners over- or under-classify assets

## Classification process steps

1. Maintain a continually updated **asset inventory** (know what you have)
2. Identify and assign an **asset owner** (accountability must be explicit)
3. Owner **classifies** the asset based on value (CIA dimensions)
4. Apply **handling requirements** and controls appropriate to the classification level
5. **Review and reassess** classification periodically — assets are re-classified as value changes

## Labeling vs. Marking (common exam trap)

| | Labeling | Marking |
|---|---|---|
| Readable by | **System** (machine-readable) | **Human** (human-readable) |
| What it encodes | Security attributes (sensitivity, access authorizations, expiry) | Handling instructions for humans |
| Enforcement type | System-based (automated policy) | Process-based (human procedure) |
| Examples | Metadata, barcodes, QR codes, RFID, GPS tags | "Do not remove from premises," document headers |

Source: NIST SP 800-53A definitions, cited in destination-cissp §2.1.4.

Cost-effectiveness of labeling technology: GPS tags > RFID > QR codes > barcodes (highest to lowest cost). Choose based on asset value and tracking requirements.

## Exam-relevant nuance

- **Same label, different value**: Two organizations may use the label "Confidential" but it can represent very different asset values. Security must educate all users on what each classification level means internally.
- **Accountability cannot be delegated**: an owner can delegate *responsibility* for day-to-day custodial tasks, but the owner is always accountable for protection. This is a frequent exam trap.
- Classification is based on **value to the organization**, not simply sensitivity. A highly available but non-sensitive system (e.g., a critical production queue) may be classified high for availability, even if the data itself is unclassified.
- NIST SP 800-60 maps information types to security categories (CIA impact levels) — not covered in destination-cissp but referenced in exam outline §2.1. (Content gap — see report.)

## Cross-links

- [Data Roles](data-roles.md) — who classifies, who is accountable
- [Data Security Controls](data-security-controls.md) — how classification drives control selection
- [Data Lifecycle](data-lifecycle.md) — classification drives retention and destruction decisions
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.1.1, §2.1.2, §2.1.3, §2.1.4 (pages 0185–0198)
- cissp-exam-outline §2.1 (domain subtopic list)
