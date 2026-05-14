---
title: AES — Advanced Encryption Standard (FIPS 197)
type: standard
domain: 3
tags: [AES, Rijndael, FIPS197, symmetric, NIST, encryption-standard]
sources: [destination-cissp]
updated: 2026-05-13
---

# AES — Advanced Encryption Standard (FIPS 197)

## Overview

AES (Advanced Encryption Standard) was published by NIST as **FIPS PUB 197** in November 2001. It is the current US government and global industry standard for symmetric encryption, replacing DES and 3DES.

AES is based on the **Rijndael** algorithm, designed by Belgian cryptographers Joan Daemen and Vincent Rijmen. Rijndael won a public NIST competition (1997–2001) against other finalists including Twofish, Serpent, RC6, and MARS.

## Key Specifications

| Parameter | AES-128 | AES-192 | AES-256 |
|---|---|---|---|
| **Key length** | 128 bits | 192 bits | 256 bits |
| **Block size** | 128 bits | 128 bits | 128 bits |
| **Rounds** | **10** | **12** | **14** |

> Note: Rijndael supports variable block sizes (128/192/256), but NIST standardized **only 128-bit blocks** for AES.

## Current NIST Position

- **3DES** is now **disallowed** by NIST (as of 2023).
- **AES-256** is the current NIST recommended symmetric standard.
- AES is used in FIPS-approved modes: CBC, CTR, GCM (AEAD), CFB, OFB, XTS (disk encryption).

## Approved Modes (NIST SP 800-38 series)

| Mode | Used For |
|---|---|
| ECB | Short, random, non-repeating data only (generally avoid) |
| CBC | Bulk data, requires IV |
| CTR | Parallel processing, most commonly used |
| GCM | Authenticated encryption (TLS 1.3, IPsec) |
| XTS | Disk encryption (BitLocker, FileVault) |
| CCM | IoT and wireless (IEEE 802.11) |

## Exam-Relevant Facts

- AES key sizes: **128, 192, 256** bits.
- AES block size: always **128 bits**.
- Rounds: 10 (128-bit key), 12 (192-bit key), 14 (256-bit key).
- Rijndael is the **algorithm**; AES is the **NIST standard** (FIPS 197) based on Rijndael.
- 3DES was the **transitional bridge** between DES and AES — it is NOT the same as AES. 3DES is now deprecated.
- NIST did not adopt Rijndael's larger block sizes — only 128-bit blocks are standardized.

## Cross-References

- [Symmetric Crypto](../concepts/symmetric-crypto.md) — full algorithm comparison table
- [FIPS 140](fips-140.md) — cryptographic module validation
- [PKCS Standards](pkcs-standards.md)

## Sources

- destination-cissp §3.6.5 (pages 0391–0394)
- NIST FIPS PUB 197 (2001)
