---
title: FIPS 140-2 / FIPS 140-3 — Security Requirements for Cryptographic Modules
type: standard
domain: 3
tags: [FIPS, cryptographic-module, HSM, TPM, NIST, standards]
sources: [destination-cissp]
updated: 2026-05-13
---

# FIPS 140-2 / FIPS 140-3

## Overview

**FIPS 140** (Federal Information Processing Standard 140) is a NIST standard that specifies security requirements for **cryptographic modules** — hardware, software, or firmware components that implement cryptographic functions. It defines four security levels, each with increasing requirements.

- **FIPS 140-2**: Published 2001; widely referenced; being phased out.
- **FIPS 140-3**: Published 2019; aligns with ISO/IEC 19790; currently active standard.

FIPS 140 is mandatory for US federal government systems handling sensitive (non-classified) information. It is also widely adopted in regulated industries (finance, healthcare).

## Security Levels

| Level | Requirements |
|---|---|
| **Level 1** | Basic cryptographic module functionality; no physical security requirements beyond production-grade components; software-only modules qualify |
| **Level 2** | Adds tamper-evidence (seals or coatings that show signs of physical tampering); role-based authentication required |
| **Level 3** | Adds tamper-resistance (active countermeasures to zeroize keys on detected tamper attempt); identity-based authentication; physical/logical separation of interfaces |
| **Level 4** | Highest; complete physical protection envelope; detects and responds to environmental attacks (voltage, temperature); zeroizes all plaintext CSPs (Critical Security Parameters) on any detected tamper |

## HSM and FIPS 140

**Hardware Security Modules (HSMs)** are typically certified at **FIPS 140-2 Level 3** for organizational use. Enterprise HSMs (e.g., those used for CA root key protection) often achieve Level 3 or Level 4 certification.

**Trusted Platform Modules (TPMs)** are usually certified at Level 1 or Level 2.

## Exam-Relevant Facts

- FIPS 140 is the US government's cryptographic module validation program.
- Level 1 allows software-only implementations.
- Level 3 is the typical enterprise HSM baseline.
- Level 4 provides the highest physical protection — detects and responds to environmental attacks.
- Zeroization (automatic erasure of keys on tamper detection) begins at Level 3.

## Cross-References

- [Key Management](../concepts/key-management.md) — TPM and HSM
- [AES Standard](aes-standard.md) — FIPS 197

## Sources

- destination-cissp §3.6.12 (key storage context)
- NIST FIPS PUB 140-2 (2001); NIST FIPS PUB 140-3 (2019)
