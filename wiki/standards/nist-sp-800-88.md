---
title: "NIST SP 800-88 Rev. 1 — Guidelines for Media Sanitization"
type: standard
domain: 2
tags: [nist, sanitization, media-destruction, data-remanence, clearing, purging]
sources: [destination-cissp]
updated: 2026-05-13
---

# NIST SP 800-88 Rev. 1 — Guidelines for Media Sanitization

**Full title**: Guidelines for Media Sanitization  
**Publication**: NIST Special Publication 800-88, Revision 1  
**Publisher**: National Institute of Standards and Technology (NIST)  
**Scope**: Provides guidance for organizations responsible for the sanitization of storage media, including hard drives, SSDs, tapes, optical media, mobile devices, and network equipment.

## Sanitization Categories (Exam-Critical)

NIST SP 800-88 Rev. 1 defines three sanitization categories, ordered from most to least effective:

| Category | Effectiveness | Data recoverable? |
|---|---|---|
| **Destroy** | Highest | No — media is physically destroyed |
| **Purge** | Middle | No — cannot be reconstructed with current technology |
| **Clear** | Lowest | Potentially yes — may be reconstructed |

## Methods by Category

### Destroy (most effective)
- **Incinerate**: Burns media to molten material; highest assurance
- **Shred / Disintegrate**: Physically fragments media; note that drilling or shredding HDDs may still leave data recoverable from platters with specialized tools
- **Drill**: Physically damages media; less effective than incineration or shredding

### Purge
- **Degauss**: Apply a strong magnetic field to erase magnetic media (HDDs, tapes). May render the media unusable. Does not work on SSDs (which have no magnetic storage).
- **Crypto Shredding / Crypto Erase**: Encrypt data with AES-256; securely destroy all key copies. Data is effectively purged *as long as* the key is provably destroyed and the algorithm remains unbroken. If key recovery is possible → downgraded to Clear.

### Clear (least effective)
- **Overwrite / Wipe / Erasure**: Write zeros, ones, or patterns to all sectors. **Any number of overwrite passes is classified as Clear, not Purge.** Research shows some data may still be recoverable.
- **Format**: Weakest method. A "Quick Format" only resets the file address table; data remains on disk until overwritten.

## SSDs and Flash Memory

Standard overwriting techniques used for magnetic HDDs do **not** work reliably on SSDs because of wear-leveling and the different addressing model for flash memory. For SSDs:
1. Use manufacturer-provided sanitization tools (if available)
2. Crypto erase (if supported)
3. Physical destruction (always reliable fallback)

## Object Reuse

Object reuse is the reassignment of storage media to a new subject such that no residual data can be accessed by the new subject. The standard implementation has been overwriting. NIST and modern guidance classify overwriting as **Clear** — NOT Purge.

## Exam applications

- "Best method to destroy data on a hard drive": **Incineration** (or physical destruction)
- "Best method for cloud data destruction": **Crypto shredding** (physical destruction infeasible)
- "Overwriting a drive 7 times": Still **Clear**, not Purge
- "Degaussing an SSD": Does **not work** — SSDs are not magnetic
- Ordering: Destroy > Purge > Clear (always)

## Cross-links

- [Data Security Controls](../concepts/data-security-controls.md) — sanitization categories in context
- [Data Lifecycle](../concepts/data-lifecycle.md) — destruction is the final lifecycle phase
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.4.2, Tables 2-5, 2-6 (pages 0207–0210) — direct citations and table
