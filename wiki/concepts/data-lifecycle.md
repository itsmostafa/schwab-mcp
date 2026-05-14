---
title: "Data Lifecycle"
type: concept
domain: 2
tags: [data-lifecycle, data-management, retention, destruction, archiving]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Lifecycle

The data lifecycle describes the stages through which data passes from creation to final destruction. Each stage requires appropriate security controls. Classification assigned at creation drives decisions at every subsequent stage.

## Lifecycle phases

| Phase | Definition | Key security considerations |
|---|---|---|
| **Create** | Generation of new digital content, or alteration/updating of existing content | Assign classification immediately upon creation; owner is designated |
| **Store** | Committing data to a storage repository (occurs nearly simultaneously with creation) | Encryption at rest, access controls, media security; classification-driven storage requirements |
| **Use** | Data viewed or processed in some activity (not including modification) | RBAC, DLP, homomorphic encryption for sensitive data; data in use protections |
| **Share** | Information made accessible to others (internal users, customers, partners) | Need-to-know enforcement, DRM/IRM for IP, data classification must follow the data |
| **Archive** | Data leaves active use and enters long-term storage | Retention policy compliance, media format longevity, protecting archived data (encryption, access control) |
| **Destroy** | Data permanently destroyed using physical or digital means | Sanitization categories: Destroy > Purge > Clear; crypto shredding for cloud |

Source: destination-cissp Table 2-4, §2.4.1 (pages 0205–0207).

## Key principles

- **Classification must be assigned immediately upon creation** and propagates through all subsequent stages
- As information moves through the lifecycle, its classification level **may increase or decrease** in value — reassess periodically
- Each stage may require **different protective measures** — the same data asset needs different controls depending on its current state (at rest vs. in transit vs. in use)
- The owner retains accountability **through the entire lifecycle**, including final destruction

## Archiving considerations

Archiving is governed by legal, regulatory, industry, and business requirements. Key questions the archiving policy must answer:
- How long must data be retained?
- What media type ensures readability over that period?
- What security (encryption, access control) applies to archived data?
- What are availability requirements for the archived data?
- What are the disposal requirements at end of retention period?

Some retention requirements can span 150 years (health records, pension records). This creates a challenge: digital formats readable today may not be readable at that time horizon. (destination-cissp §2.5.1)

## Cross-links

- [Data Roles](data-roles.md) — owner accountability spans the full lifecycle
- [Data Retention](data-retention.md) — archiving and retention policy details
- [Data Security Controls](data-security-controls.md) — sanitization and destruction methods
- [Data States](data-states.md) — protections differ by data state (rest/transit/use)
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.4.1, §2.5.1, Table 2-4 (pages 0205–0207, 0211–0212)
- cissp-exam-outline §2.4 (data lifecycle subtopics)
