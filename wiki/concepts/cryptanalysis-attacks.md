---
title: Cryptanalysis Attacks
type: concept
domain: 3
tags: [cryptanalysis, brute-force, MITM, replay, side-channel, rainbow-table, birthday-attack, pass-the-hash, ransomware]
sources: [destination-cissp]
updated: 2026-05-13
---

# Cryptanalysis Attacks

## Definition

Cryptanalysis is the science of cracking codes, decoding secrets, and breaking cryptographic protocols. It encompasses two categories: **cryptanalytic attacks** (primary goal: determine the key via mathematical analysis) and **cryptographic attacks** (broader goals including interception, forgery, identity theft, and system disruption).

## Category 1: Cryptanalytic Attacks (Goal: Determine the Key)

| Attack | Attacker Has | Difficulty | Notes |
|---|---|---|---|
| **Brute Force** | Nothing but time | Low (if key is short) | Try every possible key; ineffective for keys ≥80 bits |
| **Ciphertext Only** | Ciphertext only | Hardest | Must deduce key purely from ciphertext patterns |
| **Known Plaintext** | Both plaintext and ciphertext | Moderate | Uses both to deduce key; once found, all other ciphertext can be decoded |
| **Chosen Plaintext** | Can feed plaintext, observe ciphertext | Easiest | Attacker controls input; deduces key from input/output pairs |
| **Chosen Ciphertext** | Can feed ciphertext, observe plaintext | Easiest | Reverse of chosen plaintext; attacker controls decryption input |
| **Linear Cryptanalysis** | Known-plaintext samples | Moderate-Hard | Statistical approach; uses many known plaintext/ciphertext pairs |
| **Differential Cryptanalysis** | Chosen-plaintext pairs | Moderate-Hard | Analyzes differences in ciphertext from related plaintexts |
| **Factoring Attack** | Public key | Hard | Targets RSA; attempts to factor the product of two large primes to determine private key |
| **Meet-in-the-Middle** | Plaintext + ciphertext | Moderate | Attacks double encryption (e.g., 2-DES); works from both ends of key space; reduces 2-DES effective key from 112→56 bits |

### Brute-Force Key Space and Time

| Key Length | Key Space | Attack Time (1999 $250K system) |
|---|---|---|
| 56 bits | 7.2 × 10¹⁶ | 20 hours |
| 80 bits | 1.2 × 10²⁴ | 54,800 years |
| 128 bits | 3.4 × 10³⁸ | 1.5 × 10¹⁹ years |
| 256 bits | 1.15 × 10⁷⁷ | 5.2 × 10⁵⁷ years |

## Category 2: Cryptographic Attacks (Broader Goals)

### Man-in-the-Middle (MITM)
- Attacker inserts themselves between two communicating parties.
- Intercepts, reads, and potentially modifies messages — each party thinks they are communicating directly with the other.
- Mitigated by PKI / mutual authentication / certificate verification.

### Replay Attack
- Attacker captures authentication tokens or session identifiers and retransmits them later to gain unauthorized access.
- Example: capturing a hashed password from the network, then replaying the hash to authenticate as that user.
- Mitigated by timestamps, nonces, or sequence numbers in authentication protocols.

### Pass-the-Hash Attack
- Attacker captures a password hash (does not need to crack the actual password).
- Presents the hash directly to the target resource, which accepts it as valid authentication.
- Extends to Kerberos: the stolen hash can be used to create valid Kerberos tickets.

### Temporary Files Attack
- During encryption/decryption, keys are temporarily placed in RAM or volatile memory.
- If attacker gains system access at that moment, they can read the key from memory.

### Implementation Attack
- Targets a weakness in how an algorithm is *implemented*, not the algorithm itself.
- Classic example: **WEP** — uses RC4 (sound algorithm) but with short, repeated IVs, making it trivially breakable. The algorithm isn't broken; the implementation is.

### Side-Channel Attack
- Does not attack the algorithm or key directly — instead monitors physical emissions of the system during operation.
- Only used by sophisticated adversaries (nation-states, APTs, security researchers).

| Type | Measurement |
|---|---|
| **Timing** | How long operations take |
| **Power analysis** | How much power is consumed |
| **Radiation emissions** | Electromagnetic emissions from the device |

Often combined with fault injection attacks to weaken countermeasures first.

### Fault Injection Attack
- Deliberately injects hardware or software faults to change normal behavior.
- Used to bypass access controls or expose key material.
- Often used in conjunction with side-channel attacks.

### Dictionary Attack
- Attacker uses a precompiled list of likely passwords (dictionaries, leaked credential dumps).
- More efficient than brute force — targets human-chosen passwords.

### Rainbow Table Attack
- Precomputed table of hash values for common passwords using common hashing algorithms.
- Attacker steals a password hash file and looks up matches in the rainbow table.
- **Mitigation**: salting — a unique random value is appended to each password before hashing, making precomputed tables useless.
- **Pepper**: same random value added to all users (less secure than salt; salt is unique per user).

### Birthday Attack
- Exploits the birthday paradox to find hash collisions.
- Goal: find two different inputs with the same hash digest.
- Related to: collision resistance, hashing, integrity.
- See [Hash Functions](hash-functions.md) for details.

### Kerberos-Specific Attacks

| Attack | Description |
|---|---|
| **Pass-the-Ticket** | Using stolen Kerberos ticket instead of cracking credentials |
| **Golden Ticket** | Attacker obtains KRBTGT account hash → can forge any TGT → access to any AD resource |
| **Silver Ticket** | Attacker obtains service account hash → forge TGS tickets for that specific service only; no KDC interaction needed |

### Social Engineering (Cryptographic Context)
- **Purchase key attack (bribery)**: paying someone to hand over a cryptographic key.
- **Rubber hose attack**: using duress or torture to extract a key.

### Ransomware
- Attacker encrypts victim's files and demands payment (usually cryptocurrency) for the decryption key.
- Cryptocurrency enables anonymity for attackers.
- Even paying the ransom does not guarantee file recovery.

## Exam Traps

- The **primary goal** of cryptanalytic attacks is to determine the key. Cryptographic attacks may have other goals.
- **Chosen plaintext** and **chosen ciphertext** are the easiest attacks because the attacker controls inputs.
- **Ciphertext only** is the hardest attack (least information).
- WEP is broken due to an **implementation attack** (short IV in RC4), not an inherent algorithm flaw.
- **Meet-in-the-middle** reduces 2-DES effective key from 112 to 56 bits — this is why 2-DES provides no improvement over DES.
- **Rainbow tables** are defeated by salting; salting is better than peppering.
- **Golden ticket** = full AD compromise (KRBTGT hash); **silver ticket** = single service (service account hash).

## Cross-References

- [Cryptography Fundamentals](cryptography-fundamentals.md) — IV, key space, work factor
- [Symmetric Crypto](symmetric-crypto.md) — DES/3DES, meet-in-the-middle
- [Hash Functions](hash-functions.md) — birthday attack, rainbow tables, salting
- [Asymmetric Crypto](asymmetric-crypto.md) — factoring attacks on RSA

## Sources

- destination-cissp §3.7.1–3.7.3 (pages 0451–0460)
