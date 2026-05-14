---
title: Hash Functions
type: concept
domain: 3
tags: [cryptography, hashing, MD5, SHA, HMAC, collision, birthday-attack, integrity]
sources: [destination-cissp]
updated: 2026-05-13
---

# Hash Functions

## Definition

A **cryptographic hash function** takes an input of any length and produces a **fixed-length output** called a **message digest** (or hash value). Hashing is a **one-way** operation: you cannot reverse a digest to recover the original input. Hashing is the primary mechanism for providing **data integrity**.

## Key Properties of a Good Hash Function

| Property | Description |
|---|---|
| **Fixed-length digest** | Any input length → same output length |
| **One-way** | Cannot derive input from output |
| **Deterministic** | Same input always produces same output |
| **Calculated on entire message** | Must process the whole input, not a portion |
| **Uniformly distributed** | Hash values spread evenly across output space |
| **Collision resistant** | Hard to find two different inputs that produce the same output |
| **Avalanche effect** | Changing even one bit of input should change ~50% of output bits |

## Hash Algorithm Digest Sizes (Critical to Memorize)

| Algorithm | Digest Size | Status |
|---|---|---|
| **MD5** | **128 bits** | Broken — collisions found; do not use for security |
| **SHA-1** | **160 bits** | Deprecated — collisions found; do not use |
| **SHA-2 (SHA-224)** | 224 bits | Secure |
| **SHA-2 (SHA-256)** | **256 bits** | Widely used; current standard |
| **SHA-2 (SHA-384)** | 384 bits | Secure |
| **SHA-2 (SHA-512)** | **512 bits** | Strongest SHA-2 variant |
| **SHA-3 (all variants)** | 224/256/384/512 bits | Different underlying algorithm (Keccak); equally secure |

> Rule: longer digest = fewer possible collisions = stronger integrity guarantee.

## Collisions

A **collision** occurs when two different inputs produce the **same hash digest**. This is catastrophic for integrity: a MITM could substitute a message and attach a new hash without detection — if the hashing algorithm is weak.

**Collision resistance levels**:
- MD5: broken (collisions demonstrated in seconds)
- SHA-1: broken (practical collision attacks demonstrated by Google's SHAttered, 2017)
- SHA-2/SHA-3: secure (no practical collisions)

## Birthday Attack

The **birthday attack** exploits the **birthday paradox** statistics: in a group of just 23 people, there is a ~50% chance two share the same birthday. Applied to hashing: the number of hashes needed to find a collision is proportional to 2^(n/2), not 2^n (where n = digest size). This is why:
- MD5 (128-bit) is more vulnerable than SHA-256 (256-bit)
- Longer digests exponentially increase collision resistance

## HMAC (Hash-based Message Authentication Code)

HMAC combines a hash function with a **symmetric key** to produce a **keyed digest**:

```
HMAC = H(key || message)
```

- Provides **integrity AND authentication** (proves the message came from someone with the key)
- Does NOT provide non-repudiation (symmetric key shared between parties; either could have generated it)
- Used in TLS, VPNs, API authentication (HMAC-SHA256 is common)

## Hashing vs. Encryption vs. Digital Signatures

| Property | Hashing | Encryption | Digital Signature |
|---|---|---|---|
| Reversible? | No | Yes | No (hash is one-way) |
| Key required? | No (HMAC needs one) | Yes | Yes (private key) |
| Provides integrity? | Yes | No (alone) | Yes |
| Provides authenticity? | No | No (alone) | Yes |
| Provides non-repudiation? | No | No | Yes |

## Exam Traps

- MD5 = **128 bits**, SHA-1 = **160 bits** — these are frequently tested exact numbers.
- SHA-2 and SHA-3 are **different algorithms** but produce the **same digest lengths**. SHA-3 is based on Keccak sponge construction; SHA-2 is based on Merkle-Damgård.
- Hashing alone does **NOT prevent MITM** — a MITM can replace both the message and the hash. Digital signatures solve this.
- HMAC provides authentication but **not non-repudiation** (both parties share the key).
- "Birthday attack" → collisions → hashing → integrity. These four concepts are tightly linked in exam questions.

## Cross-References

- [Cryptography Fundamentals](cryptography-fundamentals.md)
- [Digital Signatures](digital-signatures.md)
- [Cryptanalysis Attacks](cryptanalysis-attacks.md) — birthday attacks, rainbow tables

## Sources

- destination-cissp §3.6.8 (pages 0414–0420)
