---
title: Key Management
type: concept
domain: 3
tags: [key-management, HSM, TPM, key-escrow, dual-control, split-knowledge, crypto-shredding, kerckhoffs]
sources: [destination-cissp]
updated: 2026-05-13
---

# Key Management

## Definition

Key management encompasses all activities related to the lifecycle of cryptographic keys: generation, distribution, storage, rotation, recovery, and destruction. Per **Kerckhoffs's Principle**, if the key is secure, the entire cryptosystem is secure — even if the algorithm, ciphertext, IV, and all other system details are publicly known.

## Key Management Activities

| Activity | Description |
|---|---|
| **Generation / Creation** | Fully automated (humans create patterns); pseudorandom number generators used; keys chosen randomly from entire key space |
| **Distribution** | Out-of-band (in-person, phone, letter) or key wrapping (KEK) |
| **Storage** | TPM (single device) or HSM (organizational); most critical aspect |
| **Change / Rotation** | Frequency based on asset value; higher value = more frequent rotation |
| **Recovery** | Split knowledge, dual control, or key escrow |
| **Disposition / Destruction** | Crypto shredding or physical destruction |

## Key Storage Systems

### Trusted Platform Module (TPM)

- A **microchip installed on the motherboard** of individual laptops/servers.
- Stores encryption keys for the device it is installed on only.
- Enables full-disk encryption (e.g., BitLocker) without storing keys on the encrypted drive itself.
- A **crypto processor** — generates and stores keys in a tamper-resistant chip.

### Hardware Security Module (HSM)

- A **dedicated physical device** (looks like a server) connected to the organization's network.
- Stores and manages encryption keys for an **entire organization**.
- Hardened device with only one purpose: generate and store keys securely.
- Required for FIPS 140-2 Level 3+ compliance.
- Used by CAs to protect root CA private keys.

| | TPM | HSM |
|---|---|---|
| Scope | Single device | Entire organization |
| Form factor | Chip on motherboard | Network-connected device |
| Primary use | Disk encryption, device attestation | Enterprise key management, CA operations |

## Key Distribution Methods

| Method | Description |
|---|---|
| **Out-of-band** | Key shared via different channel than data (in-person, phone, letter) |
| **Key wrapping (KEK)** | Many session keys are encrypted (wrapped) with a single Key Encrypting Key; only the KEK is distributed out-of-band |
| **Diffie–Hellman** | Derives a shared secret (session key) without transmitting it; see [Asymmetric Crypto](asymmetric-crypto.md) |
| **Public key wrapping** | Encrypt symmetric key with recipient's public key; only recipient's private key can unwrap |

## Key Recovery Methods

| Method | Description |
|---|---|
| **Split knowledge** | Key is divided into pieces (e.g., cut in half); each piece held by a different person; no single person knows the complete key |
| **Dual control** | Two (or more) authorized individuals must cooperate simultaneously to access the key; analogous to nuclear launch protocols |
| **Key escrow** | Complete key stored with a trusted third party; used in cloud computing; legally mandated in some countries |

## Key Disposition / Destruction

### Crypto Shredding

Used in cloud environments where physical destruction of media is not possible:
1. Encrypt all sensitive data with a strong algorithm.
2. Destroy the encryption key.
3. Without the key, encrypted data is computationally inaccessible — effectively "destroyed."

Useful when moving cloud providers (e.g., AWS to Azure) where you cannot physically destroy their storage media.

### Physical Destruction

Shredding, melting, or degaussing storage media. Most secure option when media is under your physical control.

## Key Rotation

- Frequency of key rotation should match the **value of the asset** being protected.
- Session keys (from DH) are ephemeral — new key per session, not reused.
- **Crypto shredding in the cloud** is often an alternative to proving key rotation/destruction.

## Exam Traps

- **Key management is the most critical aspect** of any cryptographic system (Kerckhoffs's principle).
- TPM = single device chip; HSM = organizational network device. Commonly confused.
- Key escrow means keys are stored with a trusted third party — this is NOT the same as dual control.
- Crypto shredding is the cloud-native way to "destroy" data — destroy the key, not the drive.
- Pseudorandom number generators are not truly random, but they are sufficient for cryptographic purposes.
- Keys must be chosen **randomly from the entire key space** — not sequentially or predictably.

## Cross-References

- [Cryptography Fundamentals](cryptography-fundamentals.md) — Kerckhoffs's principle
- [Symmetric Crypto](symmetric-crypto.md) — key distribution problem
- [Asymmetric Crypto](asymmetric-crypto.md) — Diffie–Hellman key exchange
- [PKI](pki.md) — certificate enrollment and CSR lifecycle

## Sources

- destination-cissp §3.6.12 (pages 0440–0446)
