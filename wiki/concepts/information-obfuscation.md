---
title: "Information Obfuscation"
type: concept
domain: 2
tags: [obfuscation, data-masking, privacy, pii, non-production]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Information Obfuscation

Information obfuscation (also called data masking) makes data obscure, unclear, or unintelligible — hiding it from unauthorized parties while often still allowing systems or authorized users to function.

**Why it is used:**
- Impair malicious actors' ability to understand data
- Protect proprietary information
- Meet compliance requirements (e.g., GDPR, PCI DSS)
- Enable functional testing without exposing real PII

Source: destination-cissp §2.6.4 (pages 0218–0219).

## Obfuscation methods (Table 2-8)

| Method | Description | Typical use case |
|---|---|---|
| **Concealing data** | Completely removes access and visibility; the attribute field does not appear on screens or reports at all | Maximum protection; user has no knowledge the attribute exists |
| **Pruning data** | Removes the *value* from an attribute in non-production environments; field remains visible but unpopulated | Dev/test environments; attribute structure is preserved without real data |
| **Fabricating data** | Replaces real data with realistic fake data | Functional testing with full data structure; preventing exposure of sensitive records |
| **Trimming data** | Partially masks an attribute; e.g., showing only last 4 digits of SSN or credit card | Identification without full disclosure; common in PCI DSS environments |
| **Encrypting data** | Creates ciphertext; reversible with the correct key; can be applied at attribute, table, or database level | Storage and transit of credit card numbers, passwords, sensitive fields |

## Key distinctions for the exam

- **Concealing vs. Pruning**: Concealing = field is invisible; Pruning = field is visible but empty
- **Pruning vs. Trimming**: Pruning removes the whole value; Trimming removes *part* of the value
- **Fabricating vs. Anonymizing**: Fabrication replaces data with fake data (field structure preserved); anonymization removes identifying information such that re-identification is impossible
- **Encrypting** is the only method that is *reversible* given the key — all other masking methods are permanent or one-way for that attribute value

## Cross-links

- [Privacy and PII](privacy-pii.md) — anonymization, pseudonymization, tokenization
- [Data Security Controls](data-security-controls.md) — encryption as a security control
- [DLP](dlp.md) — DLP can detect unmasked PII in data streams
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.6.4, Table 2-8 (pages 0218–0219)
