---
title: "Data Security Controls"
type: concept
domain: 2
tags: [data-destruction, sanitization, data-remanence, object-reuse, ssd, media-handling, drm, irm]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Data Security Controls

Data security controls encompass the methods for marking, handling, storing, and ultimately destroying data assets. Classification level drives the baseline of controls required. This page focuses on sanitization/destruction and the DRM/IRM protection methods covered in the source.

## Data Remanence

**Data remanence** refers to residual representations of data that persist even after attempts to securely delete it. A determined attacker with the right tools may reconstruct deleted data. This is why simply "deleting" files or formatting drives is insufficient for sensitive data.

**Defensible destruction** is the goal: being able to prove there is no possible way for anyone to recover destroyed data. (destination-cissp §2.4.2)

## Media Sanitization Categories

Three primary categories exist, ordered from most to least effective (NIST SP 800-88 Rev. 1):

| Category | Description | Key methods |
|---|---|---|
| **Destroy** | Physical destruction of media — most effective | Incineration, shredding, disintegration, drilling |
| **Purge** | Logical/physical techniques; data cannot be reconstructed | Degaussing (magnetic media), crypto shredding (when key is fully destroyed and unrecoverable) |
| **Clear** | Logical techniques; data **may** be reconstructed — least effective | Overwriting/wiping (zeroes/ones passes), formatting |

**Exam ordering rule: Destroy > Purge > Clear**

### Specific sanitization methods (most to least effective)

| Method | Category | Notes |
|---|---|---|
| **Incineration** | Destroy | Best option; renders media to molten material |
| **Shredding / Disintegration / Drilling** | Destroy | Effective but not foolproof; data may survive on platters |
| **Degaussing** | Purge/Destroy boundary | Destroys data on magnetic media; may also destroy the media itself |
| **Crypto Shredding / Crypto Erase** | Purge/Clear boundary | Encrypt with AES-256, then destroy all key copies. IF key is irretrievable → effectively purged. IF key could be found/brute-forced → only cleared. Best for cloud environments. |
| **Overwrite / Wipe / Erasure** | Clear | Writing zeroes/ones to all sectors. Any number of passes is still classified as *clearing*, not purging. |
| **Format** | Clear (weakest) | "Quick Format" (Windows) only resets the file address table; data remains and is easily recoverable with common tools |

Source: destination-cissp Table 2-6, NIST SP 800-88 Rev. 1 (pages 0208–0210).

## Object Reuse

Object reuse (originally from the Orange Book / TCSEC) refers to the reassignment of storage media to a new subject without allowing residual data to be reconstructed. The standard implementation method is overwriting. Most experts today classify any overwriting as **clearing** (not purging). (destination-cissp §2.4.2)

## Solid State Drive (SSD) Considerations

SSDs use flash memory — data cannot be overwritten the same way as magnetic hard drives. Options (in order of preference):
1. Use vendor-provided sanitization or crypto erasure tools
2. Physical destruction (always the best fallback)

Overwriting passes that work on magnetic HDDs do **not** reliably work on SSDs due to wear-leveling and the way flash memory is addressed.

## Crypto Shredding in the Cloud

Physical destruction of cloud media is typically infeasible. Crypto shredding is the recommended approach:
1. Encrypt data using a strong algorithm (AES-256)
2. Securely destroy all copies of the encryption key

After key destruction, data is effectively unrecoverable. If key recovery is ever possible (backup key found, weak algorithm broken), data was only "cleared." (destination-cissp §2.4.2, §2.6.1)

## Digital Rights Management (DRM) and Information Rights Management (IRM)

**DRM** (NIST SP 500-241): "A system of IT components and services, along with corresponding law, policies and business models, which strive to distribute and control intellectual property and its rights."

DRM protects IP assets and rights of owners. Typical protections:
- Licensing agreements restricting access
- Encryption of content
- Digital tags tied to specific license holders
- Technologies restricting copying/viewing

**DMCA (1998)** provides the US legal basis for DRM enforcement. (destination-cissp §2.6.5)

**IRM (Information Rights Management)** is a subset of DRM applied to internal sensitive documents within an organization (rather than mass-media). Same techniques, internal scope.

## Media Handling Requirements

Handling requirements are based on **classification level**, not media type. Key principles:
- Only designated individuals may access sensitive media
- Owners define who is authorized to access media
- Storage: top-secret data requires encrypted storage (AES-256) in a physically secured location
- PCI DSS example: audit logs must be retained ≥1 year, with 90 days immediately available; payment card data destroyed when no longer needed

## Cross-links

- [Data Classification](data-classification.md) — classification drives control baselines
- [Data States](data-states.md) — protections differ by state (rest/transit/use)
- [Data Retention](data-retention.md) — retention requirements drive destruction timing
- [DLP](dlp.md) — DLP as a complementary data protection control
- [Information Obfuscation](information-obfuscation.md) — obfuscation methods
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.2.1, §2.4.2, §2.6.5, Tables 2-5, 2-6 (pages 0199–0200, 0207–0211, 0219–0220)
- NIST SP 800-88 Rev. 1 (media sanitization guidelines — cited in source)
- cissp-exam-outline §2.4, §2.6
