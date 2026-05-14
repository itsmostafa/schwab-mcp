---
title: "Data Retention"
type: concept
domain: 2
tags: [retention, archiving, legal-hold, record-management, asset-lifecycle]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Retention

Data retention governs how long an organization must keep data before it can be legally and safely destroyed. Retention requirements typically flow from laws, regulations, industry standards, and business needs — not from the organization's preference alone.

## Key facts

- Retention requirements are often mandated externally (laws, regulations, industry standards) and can span years to decades
- **Health records** and **financial records** typically have lengthy retention requirements; some organizations face retention periods of 150+ years (e.g., pension or genealogy records)
- Archiving and retention policies are typically **subsets of the overall data classification policy**
- Data classification drives what retention requirements apply to each asset

## Retention policy elements

A retention policy must address:
- **Retention period**: How long does data need to be kept?
- **Who needs access**: Access requirements during the retention period (and whether access requirements change over time)
- **Media type**: What format ensures data remains readable over the full retention period?
- **Security requirements**: Encryption, access controls, and integrity protection for archived data
- **Availability requirements**: Can archived data be offline, or must some portion be immediately accessible?
- **Disposal requirements**: How is data destroyed at end of retention period?

## Archiving challenges

Long retention periods create a challenge: digital formats readable today may not be readable decades from now. Organizations must plan for:
- Media format longevity (e.g., tape formats, proprietary file formats)
- Media degradation and migration (copying data to new media before old media fails)
- Encryption key preservation (if archives are encrypted, keys must survive alongside the data — or be separately escrowed)

## Legal holds

When litigation is anticipated or underway, organizations must apply a **legal hold** — a suspension of the normal retention/destruction schedule for affected records. Data that would otherwise be destroyed per retention policy must be preserved until the hold is lifted.

## End of retention / End of Life (EOL)

At end of retention period, data should be **defensibly destroyed** — destroyed in a way that can be proven. The appropriate sanitization method depends on media type and classification:
- Physical media: see [Data Security Controls](data-security-controls.md) for destruction/purge/clear categories
- Cloud data: crypto shredding is preferred

## PCI DSS retention example (from source)

- Audit logs: retained for **minimum 1 year**; **90 days** must be immediately available for analysis
- Payment card data: must be **destroyed as soon as it is no longer required** for business or legal purposes

Source: destination-cissp §2.2.1 (page 0200).

## Cross-links

- [Data Classification](data-classification.md) — classification drives retention requirements
- [Data Lifecycle](data-lifecycle.md) — retention is the Archive phase of the lifecycle
- [Data Security Controls](data-security-controls.md) — destruction methods at end of retention
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.2.1, §2.5.1, §2.5.2 (pages 0199–0200, 0211–0212)
- cissp-exam-outline §2.4, §2.5
