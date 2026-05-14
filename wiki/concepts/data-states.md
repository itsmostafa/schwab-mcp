---
title: "Data States"
type: concept
domain: 2
tags: [data-at-rest, data-in-transit, data-in-use, encryption, dlp]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data States

Data exists in one of three states at any given time. Security controls differ significantly by state. A key exam skill is matching the correct control to the correct state.

## The three states

### Data at Rest
Inactive data stored on media: hard drives, SSDs, tapes, databases, spreadsheets, USB drives.

**Protections:**
- **Encryption** (e.g., AES-256 for top-secret data; full-disk encryption, database encryption)
- **Access controls** (who can access the storage location)
- **Backup and restoration** (integrity and availability)
- Best practice for cloud migration: **encrypt data locally before migrating** to ensure confidentiality during and after the transfer

### Data in Transit (Data in Motion)
Data flowing across a network (internal LAN, internet, WAN).

**Protections:**
- **End-to-end encryption** — data encrypted at source, remains encrypted through all intermediate nodes; decrypted only at destination. Routing information (IP addresses) remains visible. Best example: VPNs. Does not provide anonymity.
- **Link encryption** — entire packet (header + data) encrypted between each node; decrypted and re-encrypted at every node. Provides routing privacy hop-by-hop but every node is a potential plaintext exposure point. Performed by service providers (NIST SP 800-12).
- **Onion network** — multiple layers of encryption; each node peels one layer revealing only next-hop address. Provides both data confidentiality and source/destination anonymity. Downside: performance overhead. Example: Tor (The Onion Router).
- Access controls and network segmentation

**End-to-end vs. Link encryption exam comparison:**

| | End-to-End | Link Encryption |
|---|---|---|
| Header encrypted? | No (routing info visible) | Yes (between hops) |
| Decryption at intermediate nodes? | No | Yes (at each hop) |
| Anonymity? | No | No (only per hop) |
| Risk point | Traffic analysis (headers visible) | Every intermediate node sees plaintext |

### Data in Use
Data being actively processed in memory or by an application.

**Protections:**
- **Homomorphic encryption** — allows computations on encrypted data without decrypting it; groundbreaking because it eliminates the need to expose plaintext during processing. Key does not need to be exposed.
- **Role-Based Access Control (RBAC)** — restricts which roles/users can access and process specific data
- **DRP (Digital Rights Protection) / DLP (Data Loss Prevention)** — limits specific actions users can take on data; monitors for unauthorized use

## Exam-relevant nuance

- HTTPS protects **data in transit** (not data in use)
- Encryption of data at rest protects **confidentiality** if media is stolen
- Homomorphic encryption is the only method that protects **data in use** while it is actively being processed (still largely in research/limited production deployment)
- Onion networks sacrifice **performance for anonymity**
- The "best way to protect confidentiality of data being migrated to the cloud" is to **encrypt locally before migrating** (data at rest protection applied pre-transit)

> Note: Detailed cryptographic mechanisms (key management, cipher modes, PKI) are covered in Domain 3. This page focuses on selection rationale for the Domain 2 exam.

## Cross-links

- [Data Security Controls](data-security-controls.md) — full control baselines per classification
- [DLP](dlp.md) — DLP operates across all three states
- [Data Lifecycle](data-lifecycle.md) — states map to lifecycle phases
- [Domain 2 — Asset Security](../domains/02-asset-security.md)
- Domain 3 (Cryptography) — cryptographic mechanisms in depth

## Sources

- destination-cissp §2.6, §2.6.1, §2.6.2, §2.6.3, Table 2-7 (pages 0212–0218)
- cissp-exam-outline §2.6 (data states subtopic)
