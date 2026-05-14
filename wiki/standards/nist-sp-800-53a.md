---
title: "NIST SP 800-53A — Assessing Security and Privacy Controls"
type: standard
domain: 2
tags: [nist, labeling, marking, security-controls, assessment]
sources: [destination-cissp]
updated: 2026-05-13
---

# NIST SP 800-53A — Assessing Security and Privacy Controls

**Full title**: Assessing Security and Privacy Controls in Information Systems and Organizations  
**Publisher**: NIST  
**Scope**: Provides assessment procedures for verifying security and privacy controls. Also contains key definitions for security labeling and security marking that are tested in Domain 2.

> Note: This page covers only the labeling/marking definitions relevant to Domain 2. Full control assessment methodology is out of scope here and is addressed more in Domain 1 (governance frameworks).

## Key Definitions (Domain 2 scope)

### Security Labeling
> "The association of security attributes with subjects and objects represented by internal data structures within organizational information systems, to enable information system-based enforcement of information security policies."

Security labels include: access authorizations, data lifecycle protections (encryption, expiration), nationality, affiliation, and data classification per legal/compliance requirements.

**Labeling is system-readable.** It is enforced by automated systems.

### Security Marking
> "The association of security attributes with objects in a human-readable form, to enable organizational process-based enforcement of information security policies."

Security marking directs human handling of assets. Example: a document labeled "Top Secret" might be marked "Do not remove from premises."

**Marking is human-readable.** It is enforced through organizational processes and human behavior.

## Labeling vs. Marking comparison

| Attribute | Labeling | Marking |
|---|---|---|
| Readable by | System (machine) | Human |
| Enforcement | Automated / system-based | Process-based / human |
| Represents | Internal data structure attributes | External handling instructions |
| Examples | Metadata, RFID, barcodes, QR codes, GPS tags | Document headers, physical stamps, handling notices |

## Labeling technologies (cost comparison)

- **GPS tags**: Highest cost; justified only for high-value mobile assets requiring remote tracking
- **RFID tags**: Moderate cost; effective for warehouse inventory (no individual scanning required)
- **Barcodes**: Minimal cost; printed on packaging; common in retail
- **QR codes**: Minimal cost; hold more data than barcodes; scannable with smartphones

Cost-effectiveness should be balanced against asset value and tracking requirements.

## Cross-links

- [Data Classification](../concepts/data-classification.md) — labeling and marking are post-classification steps
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.1.4, Table 2-2 (pages 0196–0198) — direct quotation and comparison table
