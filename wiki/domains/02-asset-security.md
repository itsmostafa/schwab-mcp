---
title: "Domain 2 — Asset Security"
type: domain
domain: 2
tags: [asset-management, data-classification, data-lifecycle, dlp]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 2 — Asset Security

Exam weight: **10%**. Covers classification, ownership, handling, lifecycle management, and
protection of information and physical assets.

## Subtopics (from CISSP Exam Outline)

### 2.1 Identify and classify information and assets
- Data classification
- Asset classification

Asset classification is a system of ordered classes used to differentiate asset values and determine protection levels. **Classification** is the system; **categorization** is the act of assigning an asset to a class. Classification must be driven by asset owners, applies across all three CIA dimensions (not just confidentiality), and is an ongoing process — asset values change over time. Once classified, assets must be **labeled** (system-readable metadata/RFID/barcodes) and **marked** (human-readable handling instructions), per NIST SP 800-53A definitions. See [Data Classification](../concepts/data-classification.md).

### 2.2 Establish information and asset handling requirements

Handling requirements are based on **classification level, not media type** — the same requirements apply regardless of whether the asset is on tape, hard drive, or paper. Only designated individuals should have access to sensitive media; the owner defines who is authorized. Storage of top-secret data mandates encrypted storage (e.g., AES-256) in a physically secured location. Retention and destruction schedules are governed by regulatory requirements such as PCI DSS (audit logs ≥1 year; 90 days immediately available). See [Data Security Controls](../concepts/data-security-controls.md) and [Data Retention](../concepts/data-retention.md).

### 2.3 Provision information and assets securely
- Information and asset ownership
- Asset inventory (tangible, intangible)
- Asset management

The **data owner** is the business person (e.g., an HR Director for an HR database) who is **accountable** for an asset's classification, access approval, and protection throughout its lifecycle — including final destruction. **Accountability cannot be delegated**, although responsibility for day-to-day technical management can be passed to a **data custodian** (IT). A continuously updated asset inventory of all tangible and intangible assets is the prerequisite for any classification program. Shadow IT (untracked cloud services) is a common gap. See [Data Roles](../concepts/data-roles.md) and [Information Lifecycle Management](../concepts/information-lifecycle-management.md).

### 2.4 Manage data lifecycle
- Data roles: owners, controllers, custodians, processors, users/subjects
- Data collection
- Data location
- Data maintenance
- Data retention
- Data remanence
- Data destruction

The six lifecycle phases are: **Create → Store → Use → Share → Archive → Destroy**. Classification must be assigned at the Create phase and drives protections at every subsequent phase. **Data remanence** — residual data remaining after deletion — is addressed through three sanitization categories: **Destroy** (physical destruction; most effective), **Purge** (logical/physical; data unrecoverable), and **Clear** (logical; data may be recoverable). Overwriting, regardless of pass count, is classified as Clear. Crypto shredding is the recommended method for cloud environments. See [Data Lifecycle](../concepts/data-lifecycle.md) and [Data Security Controls](../concepts/data-security-controls.md).

### 2.5 Ensure appropriate asset retention (e.g., End of Life (EOL), End of Support)

Retention requirements are driven by law, regulation, industry standards, and business needs — not organizational preference. Health and financial records may require retention periods measured in decades or longer (up to 150 years in some cases), creating media longevity challenges. **EOL** (End of Life) means no further product development; **EOS** (End of Support) means no more security patches — EOS represents the higher security risk and requires compensating controls or replacement. Archiving policies must specify media type, security requirements, availability, retention period, and disposal procedures. See [Data Retention](../concepts/data-retention.md).

### 2.6 Determine data security controls and compliance requirements
- Data states: in use, in transit, at rest
- Scoping and tailoring
- Standards selection
- Data protection methods: Digital Rights Management (DRM), data loss prevention (DLP),
  cloud access security broker (CASB)

Controls are selected based on the asset's classification level and its current data state. **Data at rest** is protected by encryption, access controls, and backup/restore. **Data in transit** is protected by end-to-end encryption (routing visible; e.g., VPN), link encryption (full packet encrypted per hop but plaintext at each node), or onion routing (full anonymity, lower performance; e.g., Tor). **Data in use** is protected by homomorphic encryption, RBAC, and DLP/DRP. **DRM** controls intellectual property usage rights; **DLP** is broader — identifying, monitoring, and protecting sensitive data across all three states via deep content inspection. **CASB** extends DLP-like controls to cloud services. Information obfuscation (concealing, pruning, fabricating, trimming, encrypting) masks PII for compliance and testing. See [Data States](../concepts/data-states.md), [DLP](../concepts/dlp.md), and [Information Obfuscation](../concepts/information-obfuscation.md).

> **Content gap**: CASB is listed in the exam outline and this domain page but is not covered in destination-cissp. Anonymization/pseudonymization/tokenization is also undertreated in the source. See [Privacy and PII](../concepts/privacy-pii.md) for current coverage.

## Key concepts

- [Data Classification](../concepts/data-classification.md) — classification levels, categorization, labeling/marking, owner-driven process
- [Data Roles](../concepts/data-roles.md) — owner, custodian, steward, controller, processor, data subject — EXAM CRITICAL
- [Data Lifecycle](../concepts/data-lifecycle.md) — Create/Store/Use/Share/Archive/Destroy; protections per phase
- [Data States](../concepts/data-states.md) — at rest, in transit, in use; encryption and access control options per state
- [Data Security Controls](../concepts/data-security-controls.md) — sanitization categories (Destroy/Purge/Clear), DRM, IRM, media handling
- [Data Retention](../concepts/data-retention.md) — retention policies, archiving, legal holds, EOL/EOS
- [Privacy and PII](../concepts/privacy-pii.md) — PII, GDPR roles, anonymization/pseudonymization (content gap flagged)
- [DLP](../concepts/dlp.md) — Data Loss Prevention across all three data states
- [Information Obfuscation](../concepts/information-obfuscation.md) — masking techniques (concealing, pruning, fabricating, trimming, encrypting)
- [Information Lifecycle Management](../concepts/information-lifecycle-management.md) — asset inventory, EOL/EOS, scoping and tailoring

## Sources
- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §2.1–2.6 (pages 0184–0223)
