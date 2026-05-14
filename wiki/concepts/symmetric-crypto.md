---
title: Symmetric Cryptography
type: concept
domain: 3
tags: [cryptography, symmetric, DES, AES, 3DES, RC4, blowfish, twofish, block-cipher, stream-cipher]
sources: [destination-cissp]
updated: 2026-05-13
---

# Symmetric Cryptography

## Definition

Symmetric cryptography uses the **same key** for both encryption and decryption. Both parties must possess an identical copy of the key before secure communication can begin.

## Advantages and Disadvantages

| Advantages | Disadvantages |
|---|---|
| Extremely fast | Key distribution problem (key must be shared securely) |
| Strong with adequate key length | Scalability: n × (n−1) / 2 keys needed for n parties |
| Efficient for bulk data | No built-in authenticity, integrity, or non-repudiation |

**Scalability formula**: for 1,000 users → 1,000 × 999 / 2 = **499,500 unique keys** required.

**Key distribution**: must use out-of-band channels (in-person, phone, letter) or asymmetric key wrapping (see [Key Management](key-management.md)).

## Symmetric Algorithms: Ranked Weakest to Strongest

| Strength | Algorithm | Key Length (bits) | Block Size (bits) | Notes |
|---|---|---|---|---|
| Weak | RC2-40 | 40 | 64 | Deprecated |
| Weak | DES | 56 | 64 | 16 rounds; deprecated; brute-forced in hours |
| Weak | RC5-64/16/7 | 56 | 128 | — |
| Medium | RC5-64/16/10 | 80 | 128 | — |
| Medium | Skipjack | 80 | 64 | NSA-designed, used in Clipper chip |
| Strong | RC2-128 | 128 | 8 | — |
| Strong | IDEA | 128 | 64 | First 128-bit symmetric algo; used in PGP |
| Strong | Blowfish | 128 | 64 | — |
| Strong | 3DES | 168 nominal / **112 effective** | 64 | Disallowed by NIST; meet-in-the-middle reduces strength |
| Very Strong | Twofish | 256 | 128 | AES finalist |
| Very Strong | RC6 | 256 | 128 | AES finalist |
| Very Strong | **Rijndael (AES)** | 128, 192, or 256 | 128 | Current NIST standard; FIPS 197 |

### DES and 3DES Details

- **DES**: 56-bit key, 16 rounds of substitution + transposition, 64-bit block size. Strong confusion/diffusion but key too short for modern brute-force.
- **2-DES**: Two 56-bit keys (112-bit nominal). Vulnerable to **meet-in-the-middle attack** → effective key length reduced to 56 bits. **Not used.**
- **3DES**: Three iterations of DES with 2 or 3 keys. 168-bit nominal, **112-bit effective** (meet-in-the-middle removes 56 bits). Now **disallowed by NIST**. Current standard is AES-256.

**Meet-in-the-middle attack**: Attacker works from both ends of the key space — encrypts plaintext with all possible key1 values, decrypts ciphertext with all possible key2 values, finds the match. Halves effective key strength of any double-encryption scheme.

### AES / Rijndael Details

- Rijndael won the 2001 NIST competition for AES (FIPS 197).
- Key sizes: 128, 192, or 256 bits. Block size is always **128 bits**.
- Rounds: 10 (128-bit key), 12 (192-bit key), 14 (256-bit key).
- Rijndael natively supports 128/192/256-bit block sizes, but NIST standardized only 128-bit blocks.
- **Current NIST standard**: AES-256.

### RC4 (Stream Cipher)

- Designed by Ron Rivest. The most commonly used stream cipher.
- **Deprecated**: key scheduling algorithm has weaknesses; WEP's broken implementation used RC4 with short IVs.
- Not recommended for new systems.

### ChaCha20-Poly1305

- 256-bit stream cipher by Daniel J. Bernstein; 20 rounds.
- Combines ChaCha20 (encryption) with Poly1305 (authentication) as an **AEAD** (Authenticated Encryption with Associated Data) cipher.
- Used by Cloudflare and Google in TLS; often faster than AES on systems without hardware AES acceleration.
- Variants: ChaCha8 (8 rounds), ChaCha12, ChaCha20.

## Block Cipher Modes

See [Cryptography Fundamentals — Block Cipher Modes](cryptography-fundamentals.md#block-cipher-modes-of-operation) for the full table. Key exam point: **CTR is the most commonly used mode**; **ECB has no IV and is insecure for structured data**.

## Cross-References

- [Cryptography Fundamentals](cryptography-fundamentals.md) — stream vs. block ciphers, modes of operation
- [Asymmetric Crypto](asymmetric-crypto.md) — key exchange solution
- [Key Management](key-management.md) — out-of-band distribution, key wrapping
- [Cryptanalysis Attacks](cryptanalysis-attacks.md) — meet-in-the-middle, brute-force

## Sources

- destination-cissp §3.6.5 (pages 0386–0396)
