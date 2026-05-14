---
title: PKCS Standards (Public-Key Cryptography Standards)
type: standard
domain: 3
tags: [PKCS, RSA, PKI, certificate, standards, cryptography]
sources: [destination-cissp]
updated: 2026-05-13
---

# PKCS Standards (Public-Key Cryptography Standards)

## Overview

PKCS (Public-Key Cryptography Standards) is a family of standards originally designed by RSA Laboratories and now maintained by OASIS. They define formats and protocols for various aspects of public-key cryptography, enabling interoperability across systems.

## Key PKCS Standards for CISSP

| Standard | Name | Purpose |
|---|---|---|
| **PKCS#1** | RSA Cryptography Standard | Defines RSA public/private key format and encryption/signature schemes (RSA-OAEP, RSA-PSS) |
| **PKCS#7 (CMS)** | Cryptographic Message Syntax | Defines format for signed/enveloped data; used in S/MIME for email encryption and signing |
| **PKCS#10** | Certification Request Syntax | Defines format for **Certificate Signing Requests (CSRs)** — sent to CA to request a digital certificate |
| **PKCS#11** | Cryptographic Token Interface (Cryptoki) | API for hardware cryptographic tokens (HSMs, smart cards); defines how applications interact with crypto hardware |
| **PKCS#12 (PFX)** | Personal Information Exchange | Defines a portable format for storing a **certificate + private key bundle** (commonly exported as .pfx or .p12 files) |

## Exam-Relevant Facts

- **PKCS#10** = CSR format (what you send to a CA when requesting a certificate).
- **PKCS#12 / PFX** = the file that contains BOTH your certificate AND your private key together (password-protected).
- **PKCS#11** = the API that lets applications talk to HSMs and smart cards.
- **PKCS#7** underlies S/MIME email security.
- **PKCS#1** defines RSA operations — RSA-OAEP for encryption, RSA-PSS for signatures.

## Cross-References

- [PKI](../concepts/pki.md) — certificate enrollment uses PKCS#10 CSR
- [Asymmetric Crypto](../concepts/asymmetric-crypto.md) — RSA (PKCS#1)
- [Key Management](../concepts/key-management.md) — PKCS#11 for HSM access, PKCS#12 for certificate/key transport

## Sources

- destination-cissp §3.6.11 (PKI components context)
- RSA Laboratories PKCS specification series
