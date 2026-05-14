---
title: Digital Signatures
type: concept
domain: 3
tags: [cryptography, digital-signature, non-repudiation, integrity, authenticity, code-signing, DSA, ECDSA]
sources: [destination-cissp]
updated: 2026-05-13
---

# Digital Signatures

## Definition

A **digital signature** is an encrypted hash value created by the sender using their **private key**. It provides three security services: **integrity**, **authenticity (proof of origin)**, and **non-repudiation**. It does **NOT** provide confidentiality.

## Three Services of Digital Signatures

| Service | Mechanism |
|---|---|
| **Integrity** | Receiver re-hashes the message and compares with the decrypted digest |
| **Authenticity (proof of origin)** | Receiver decrypts the signature with sender's public key — only sender's private key could have created it |
| **Non-repudiation** | When both integrity and authenticity are verified, sender cannot deny sending the message |

> Critical exam point: Digital signatures do **NOT** provide confidentiality — anyone can read the message body.

## Creating a Digital Signature

1. Sender computes a hash of the message (produces fixed-length digest).
2. Sender encrypts the hash with the sender's **private key** → this encrypted hash IS the digital signature.
3. Sender attaches the digital signature to the message and sends both.

## Verifying a Digital Signature

1. Receiver decrypts the digital signature using the sender's **public key** → recovers the original hash digest. (If decryption succeeds → **authenticity** confirmed.)
2. Receiver independently hashes the received message.
3. Receiver compares both digests. If they match → **integrity** confirmed.
4. Both verified → **non-repudiation** established.

## Why Hashing Alone Is Insufficient

Without a digital signature, a MITM can:
1. Intercept the message
2. Modify the content
3. Compute a new hash of the modified message
4. Send the modified message + new hash to the recipient

The recipient would falsely believe integrity is confirmed. Digital signatures prevent this because the MITM cannot encrypt a new hash with the sender's private key (they don't have it).

## Five Services — Full Example (Alice → Bob)

For confidentiality + integrity + authenticity + non-repudiation + access control:
1. Alice and Bob exchange digital certificates (verified public keys).
2. Alice generates a symmetric key; encrypts it with Bob's public key → sends encrypted key to Bob (solves key distribution; Bob's private key decrypts it).
3. Alice encrypts the large message with the symmetric key → sends ciphertext (confidentiality, access control).
4. Alice hashes the message, encrypts hash with her private key → sends digital signature (integrity + authenticity + non-repudiation of origin).
5. Bob hashes the message, encrypts hash with his private key → sends back to Alice (non-repudiation of delivery).

## Common Uses of Digital Signatures

- **Document signing**: legal equivalence to handwritten signature; much harder to forge (requires private key).
- **Code signing**: verifying that software (OS updates, apps) came from the legitimate vendor (Apple, Microsoft) and was not modified in transit.
- **S/MIME**: email authentication and integrity via PKI-based digital signatures.

## Digital Signature Algorithms

| Algorithm | Notes |
|---|---|
| **DSA** (Digital Signature Algorithm) | FIPS-standardized; uses discrete logarithm; for signing only (not encryption) |
| **RSA-PSS** | RSA with Probabilistic Signature Scheme; used for signing |
| **ECDSA** (Elliptic Curve DSA) | Shorter keys than RSA; used in TLS certificates, Bitcoin |

## Exam Traps

- Digital signature = hash encrypted with **private key** (NOT the message itself encrypted).
- Digital signatures provide authenticity and non-repudiation, but **NOT confidentiality**.
- To get confidentiality, the message must also be encrypted (separately) with the **recipient's public key**.
- Non-repudiation requires both integrity AND authenticity to be verified.

## Cross-References

- [Hash Functions](hash-functions.md) — hashing is the foundation
- [Asymmetric Crypto](asymmetric-crypto.md) — private/public key mechanics
- [PKI](pki.md) — certificate infrastructure that makes public key distribution trustworthy
- [Key Management](key-management.md)

## Sources

- destination-cissp §3.6.9, §3.6.13 (pages 0420–0451)
