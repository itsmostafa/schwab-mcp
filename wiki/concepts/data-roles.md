---
title: "Data Roles"
type: concept
domain: 2
tags: [data-owner, data-custodian, data-steward, data-controller, data-processor, data-subject, gdpr, roles]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Roles

Understanding the distinct responsibilities of each data role is **exam-critical**. The owner vs. custodian distinction is one of the most frequently tested traps in Domain 2.

## Role definitions

| Role | Accountability / Responsibility | Notes |
|---|---|---|
| **Data Owner / Controller** | Accountable for protection of data; holds legal rights; defines classification and access policies | In GDPR context, "controller" = entity that determines purposes and means of processing. Destination-cissp Table 2-3 combines Owner and Controller into one row — see nuance below. |
| **Data Custodian** | *Technical* responsibility for data: security, availability, capacity, backup/restore, operations | Typically IT staff. Does NOT determine value or classification — that is the owner's role. |
| **Data Steward** | *Business* responsibility for data: metadata definition, data quality, governance, compliance | Bridges business and technical sides; ensures data quality and proper governance. |
| **Data Processor** | Responsible for processing data on behalf of the owner/controller | Classic example: a cloud service provider. Has no authority to determine purpose of processing. |
| **Data Subject** | The individual to whom personal data pertains | GDPR term. Exercises rights (access, erasure, portability) against the controller. |
| **User** | Day-to-day users of data assets; must follow the policies set by the owner | Lowest accountability of all roles; responsibility is operational only. |

Source: destination-cissp Table 2-3 (page 0202).

## EXAM CRITICAL: Owner vs. Custodian

This is one of the highest-yield exam traps:

| | Data Owner | Data Custodian |
|---|---|---|
| Who | Business role (e.g., HR Director owns HR database) | Technical role (e.g., IT/DBA) |
| Determines classification | **Yes** | No |
| Accountable for protection | **Yes** | No (responsible only) |
| Day-to-day technical management | No | **Yes** |
| Can delegate accountability? | **No** — accountability cannot be delegated | N/A |

Key principle: **An owner can delegate responsibility but never accountability.** If the owner hands the asset to IT to manage, the owner is still accountable for the outcome.

## EXAM NUANCE: Owner vs. Controller

Destination-cissp conflates Data Owner and Data Controller in Table 2-3 (both in one row). However, in the GDPR framework, these are distinguishable:

- **Data Owner** is an organizational role — the business person accountable for an asset.
- **Data Controller** is a legal/regulatory concept — the entity that determines the purposes and means of data processing. An organization itself (as a legal entity) is typically the controller; the Data Owner is an internal actor within that organization.

This distinction matters for privacy law questions. When a question is in the context of GDPR, treat Controller and Processor as legal entities; when in the context of organizational security, use Owner/Custodian.

## Owner accountabilities (exam list)

An owner must:
- Classify and categorize assets
- Manage access to assets (approve who can access)
- Ensure appropriate controls are in place based on classification
- Retain accountability through the full asset life cycle, including retention and destruction

Different types of owners exist: data owners, process owners, system owners, product owners, service owners, hardware owners, application owners, IP owners — but all hold the same core accountability.

## Cross-links

- [Data Classification](data-classification.md) — classification is driven by the owner
- [Data Lifecycle](data-lifecycle.md) — owner retains accountability through full lifecycle
- [Privacy and PII](privacy-pii.md) — data subject/controller/processor in GDPR context
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.3.1, Table 2-3 (pages 0200–0202)
- cissp-exam-outline §2.3 (data roles subtopic)
