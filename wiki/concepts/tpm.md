---
title: "Trusted Platform Module (TPM)"
type: concept
domain: 3
tags: [tpm, hardware-security, cryptography, binding, sealing, platform-integrity]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Trusted Platform Module (TPM)

A Trusted Platform Module (TPM) is a **hardware chip** that implements ISO/IEC 11889 and provides a hardware-based root of trust for computing devices. TPMs are typically soldered onto the motherboard and perform cryptographic operations independently of the operating system.

## Definition

A TPM is a dedicated microcontroller that:
- Generates and stores cryptographic keys
- Performs cryptographic operations (encryption, decryption, hashing)
- Verifies platform integrity at boot time
- Maintains a secure, tamper-resistant boundary

## Key Characteristics

- **Hardware-based**: A physical chip, not software — cannot be bypassed by OS-level attacks
- **Independent**: Has its own internal circuits and firmware; does not rely on the OS for processing
- **Unique**: Every TPM has a unique, factory-burned **Endorsement Key (EK)** — a special-purpose RSA key that never leaves the chip and is used only for encryption/authentication
- **Black box**: Commands can be sent to the TPM, but data stored within it cannot be extracted

## Core Operations

### Binding
A cryptographic operation that encrypts data (such as another encryption key) using the TPM's endorsement key, tying it to that specific TPM's hardware configuration. Bound data can **only** be decrypted by that exact TPM.

- Purpose: Protect keys from being disclosed if removed from the hardware
- Behavior: If the TPM is transferred to different hardware, bound data is inaccessible

### Sealing
A cryptographic operation that encrypts data with a **conditional** decryption policy. Unlike binding (which ties to the hardware), sealing ties decryption to a specific **software/firmware state** or **user authentication condition**.

- Example: Data sealed such that it can only be decrypted if the system boots with a specific OS and firmware configuration
- Purpose: Detect tampering — if system state changes (e.g., malware alters boot chain), the sealed data cannot be decrypted

| Operation | Tied to | Purpose |
|---|---|---|
| **Binding** | Specific TPM hardware | Protect key from disclosure outside that hardware |
| **Sealing** | Specific system state/conditions | Detect tampering; enforce pre-conditions for access |

## Boot Integrity

At boot, the TPM measures (hashes) the state of firmware, bootloader, and OS components and stores these measurements in **Platform Configuration Registers (PCRs)**. If any measured component has changed since the sealed state was created, the TPM refuses to unseal the data — signaling a potential tampering event.

## Exam-Relevant Nuance

- TPM is a **hardware** security mechanism — not purely software or firmware.
- The Endorsement Key is unique per chip and never leaves the TPM.
- Binding vs. sealing is a high-frequency exam distinction: binding = hardware-tied, sealing = state-conditional.
- TPM enables technologies like **BitLocker** (Windows full-disk encryption) and **Measured Boot**.
- ISO/IEC 11889 is the standard — knowing the standard number is useful but less frequently tested than the concepts.

## Cross-links

- [Trusted Computing Base](trusted-computing-base.md) — TPM is a TCB component
- [Virtualization Security](virtualization-security.md) — virtual TPMs in cloud/VM environments
- [Cloud Security Models](cloud-security-models.md) — TPM availability in cloud context

## Sources

- destination-cissp §3.4.10 (pp. 297–298)
- cissp-exam-outline §3.4
