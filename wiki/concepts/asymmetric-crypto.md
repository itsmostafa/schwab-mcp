---
title: Asymmetric Cryptography
type: concept
domain: 3
tags: [cryptography, asymmetric, RSA, ECC, Diffie-Hellman, public-key, discrete-log, factoring, hybrid]
sources: [destination-cissp]
updated: 2026-05-13
---

# Asymmetric Cryptography

## Definition

Asymmetric cryptography (public-key cryptography) uses a mathematically linked **key pair**: a **public key** (shared freely) and a **private key** (kept secret by the owner). What one key encrypts, only the other can decrypt. Developed conceptually by Diffie and Hellman in the 1970s to solve the key distribution problem inherent in symmetric cryptography.

## Key Usage Rules (Memorize These)

| Goal | Encrypt With | Decrypt With |
|---|---|---|
| **Confidentiality** (send secret to Alice) | Alice's **public** key | Alice's **private** key |
| **Authenticity / Proof of origin** (prove you sent it) | Sender's **private** key | Sender's **public** key |
| **Digital signature** (hash + private key) | Sender's **private** key | Sender's **public** key |

> Encrypt with the **recipient's public key** for confidentiality.  
> Sign with your **own private key** for authenticity.

## Advantages and Disadvantages

| Advantages | Disadvantages |
|---|---|
| Solves key distribution problem | Significantly slower than symmetric |
| Solves scalability (only public keys need distribution) | Requires large key sizes |
| Enables digital signatures, authenticity, non-repudiation | — |
| Provides all five crypto services (in combination) | — |

## Hard Math Problems

All asymmetric algorithms rely on mathematical problems that are easy to compute in one direction but computationally infeasible to reverse:

| Problem | Algorithm(s) | Description |
|---|---|---|
| **Factoring** | RSA | Multiply two large primes = easy; factor the product = hard |
| **Discrete logarithm** | ECC, Diffie–Hellman, ElGamal | Raise a prime to a prime power = easy; determine the exponent = hard |
| ~~Knapsack~~ | Chor-Rivest, Merkle-Hellman (deprecated) | Solved by known attacks; no longer used |

## Common Asymmetric Algorithms

### RSA (Rivest, Shamir, Adleman)
- Uses **factoring** of large prime numbers.
- Most commonly used asymmetric algorithm.
- Developed late 1970s; still no significant breaks discovered.
- Key sizes increasing toward 2048+ bits as computing advances (algorithm slows as keys lengthen).
- Used for encryption and digital signatures.

### ECC (Elliptic Curve Cryptography)
- Uses **discrete logarithm** mathematics.
- Achieves the **same security as RSA with shorter keys** — faster and more efficient.
- Particularly suited to constrained environments (mobile phones, IoT, embedded systems).
- Developed early 2000s.
- ECDSA is the elliptic curve variant of digital signatures.

### Diffie–Hellman (DH) Key Exchange
- Uses **discrete logarithm** mathematics.
- **Not used for message encryption** — used almost exclusively for **symmetric session key exchange**.
- Underlying protocol for VPN session key negotiation.
- Session keys are symmetric keys generated per-session via DH; both sides derive the same shared secret without transmitting it.
- DHE (Diffie–Hellman Ephemeral) provides forward secrecy.

### Diffie–Hellman Operation (Simplified)
1. Both parties agree on a public base number.
2. Each generates a secret random number.
3. Each multiplies their secret number × the public number and exchanges results.
4. Each multiplies the received value × their own secret number → both arrive at the **same shared secret** (session key).

### ElGamal
- Uses discrete logarithm. Less commonly tested; used in PGP for key exchange.

### Quantum Key Distribution (QKD)
- Emerging technology: any eavesdropping on a quantum channel measurably disturbs the system, alerting parties to interception.
- Still experimental; not widely deployed.
- NIST is developing post-quantum public-key algorithms resistant to quantum computing attacks.

## Hybrid Cryptography

Used in TLS/SSL and most real-world secure communications:
1. **Asymmetric** is used to securely exchange a symmetric session key (Diffie–Hellman or RSA key wrapping).
2. **Symmetric** is used for bulk data encryption (fast).
3. **Hashing** + **digital signatures** are added for integrity and non-repudiation.

> Asymmetric solves key distribution; symmetric provides speed. Together they cover all five crypto services.

## Exam Traps

- ECC advantage over RSA: **smaller keys, same strength** — it uses discrete log, not factoring.
- DH is a **key exchange** protocol — it does NOT encrypt messages.
- RSA uses **factoring**; DH and ECC use **discrete logarithm**.
- The Knapsack problem algorithms are **all deprecated** — attacks exist that break them.
- Asymmetric is **significantly slower** than symmetric; this is why hybrid is used.

## Cross-References

- [Cryptography Fundamentals](cryptography-fundamentals.md)
- [Symmetric Crypto](symmetric-crypto.md)
- [Digital Signatures](digital-signatures.md)
- [PKI](pki.md)
- [Key Management](key-management.md)

## Sources

- destination-cissp §3.6.6–3.6.7 (pages 0396–0413)
