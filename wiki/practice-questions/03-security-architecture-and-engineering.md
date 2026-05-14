---
title: "Practice Questions — Domain 03: Security Architecture And Engineering"
type: practice
domain: 03
tags: [practice, cryptography, physical-security, security-models, secure-design]
sources: [destination-cissp]
updated: 2026-05-13
---

# Practice Questions — Domain 03: Security Architecture And Engineering

## Questions

### Q1: Which symmetric encryption algorithm has an effective key length of 112 bits despite using three 56-bit keys?

**Answer:** 3DES (Triple DES)

**Why:** 3DES nominally uses three 56-bit keys = 168 bits. However, the meet-in-the-middle attack reduces its effective key strength by 56 bits, leaving an effective key length of 112 bits. NIST has since disallowed 3DES; AES-256 is the current standard.

*Subtopic: 3.6 — [Symmetric Crypto](../concepts/d3b/symmetric-crypto.md)*

---

### Q2: An attacker intercepts a software update file, modifies it, and replaces the developer's hash value with a new hash they computed from the modified file. Which security service would have prevented the recipient from being deceived?

**Answer:** A digital signature (not just hashing alone)

**Why:** A standalone hash value provides no protection against a MITM who can simply recompute the hash for the modified file. A digital signature — hash encrypted with the sender's private key — prevents this because the MITM cannot forge the signature without the sender's private key. The recipient verifies the signature using the sender's public key, proving authenticity and protecting integrity.

*Subtopic: 3.6 — [Digital Signatures](../concepts/d3b/digital-signatures.md)*

---

### Q3: What are the three services provided by a digital signature?

**Answer:** Integrity, Authenticity (proof of origin), and Non-repudiation.

**Why:** A digital signature = hash of the message encrypted with the sender's private key. The receiver decrypts with the sender's public key (authenticity), re-hashes the message and compares (integrity), and since both are verified, neither party can deny their role (non-repudiation). Note: digital signatures do NOT provide confidentiality — anyone can read the message body.

*Subtopic: 3.6 — [Digital Signatures](../concepts/d3b/digital-signatures.md)*

---

### Q4: Alice wants to send a confidential message to Bob using asymmetric cryptography. Which key should she use to encrypt the message?

**Answer:** Bob's public key

**Why:** For confidentiality using asymmetric cryptography, the sender encrypts with the *recipient's public key*. Only the recipient's private key (which only Bob possesses) can decrypt it. If Alice used her own private key instead, she would be creating a digital signature (proving she sent it), not providing confidentiality.

*Subtopic: 3.6 — [Asymmetric Crypto](../concepts/d3b/asymmetric-crypto.md)*

---

### Q5: Which asymmetric algorithm uses factoring of large prime numbers as its hard math problem, and which uses discrete logarithms?

**Answer:** RSA uses factoring; Diffie–Hellman and ECC use discrete logarithms.

**Why:** RSA's security depends on the difficulty of factoring a large composite number back into its two prime factors. Diffie–Hellman and ECC rely on the discrete logarithm problem (computing an exponent given a base and result). ECC's use of discrete logarithms allows it to achieve the same security as RSA with shorter keys, making it faster and more efficient.

*Subtopic: 3.6 — [Asymmetric Crypto](../concepts/d3b/asymmetric-crypto.md)*

---

### Q6: Which block cipher mode of operation does NOT use an initialization vector (IV)?

**Answer:** ECB (Electronic Codebook)

**Why:** ECB is the only standard block cipher mode without an IV. Because the same plaintext block always produces the same ciphertext block, ECB reveals patterns in encrypted data (the famous "ECB penguin" image demonstrates this). ECB should only be used for short, random, non-repeating data such as PIN codes. For any structured or repeating data, a mode that uses an IV (CBC, CFB, OFB, or CTR) must be used.

*Subtopic: 3.6 — [Cryptography Fundamentals](../concepts/d3b/cryptography-fundamentals.md)*

---

### Q7: Hash algorithm MD5 produces what size digest? SHA-1? SHA-256?

**Answer:** MD5 = 128 bits; SHA-1 = 160 bits; SHA-256 = 256 bits.

**Why:** These are the fixed-length outputs regardless of input size. MD5 and SHA-1 are both broken due to demonstrated collision attacks and should not be used for security purposes. SHA-256 (part of the SHA-2 family) remains secure and is the current standard. SHA-3 produces digests in the same sizes (224/256/384/512) but uses a fundamentally different algorithm (Keccak).

*Subtopic: 3.6 — [Hash Functions](../concepts/d3b/hash-functions.md)*

---

### Q8: An organization stores password hashes without salting. An attacker steals the password database. What attack is most effective, and what control would have mitigated it?

**Answer:** Rainbow table attack; mitigation is salting (adding a unique random value to each password before hashing).

**Why:** A rainbow table is a precomputed database of hash values for common passwords. Without salting, an attacker can directly look up stolen hashes in the table. Salting defeats this by making every user's hash unique — even two users with the same password will have different hashes. Salts are per-user and unique; peppers are the same for all users (less secure).

*Subtopic: 3.7 — [Cryptanalysis Attacks](../concepts/d3b/cryptanalysis-attacks.md)*

---

### Q9: What is the most effective early fire detection technology for a data center, and what stage of a fire does it detect?

**Answer:** VESDA (Very Early Smoke Detection Apparatus); it detects fire at the incipient (smoldering) stage.

**Why:** VESDA continuously samples air and can detect smoke particles before a visible fire develops. This incipient-stage detection provides the maximum response time to suppress the fire before significant damage occurs. Heat detectors and flame detectors only activate when a fire is already well advanced. VESDA is the most expensive option but provides the earliest and best warning.

*Subtopic: 3.9 — [Facility Security](../concepts/d3b/facility-security.md)*

---

### Q10: A data center manager must choose a fire suppression system. The facility houses servers worth millions of dollars and must minimize false-activation water damage. Which water-based system is most appropriate?

**Answer:** Pre-action system

**Why:** Pre-action systems require a fire detection signal before water is released into the pipes, preventing accidental discharge from a single triggered sprinkler head. This minimizes the risk of water damage from false activations. Only the zone where fire is detected receives water, and only sprinkler heads activated by heat release water in that zone. Deluge (all heads open) is appropriate for explosives/fireworks only. Gas-based systems (INERGEN, FM-200, Argonite, Aero-K) are even better for data centers if cost is not a limiting factor.

*Subtopic: 3.9 — [Facility Security](../concepts/d3b/facility-security.md)*

---

### Q11: What is the primary goal of physical security?

**Answer:** Safety and protection of human life.

**Why:** Physical security's first priority is always people, not equipment or data. This is reflected in building codes that require outward-opening emergency exit doors (less secure, but safer for egress), the minimum number of access points consistent with safe evacuation, and the principle that no physical security control should put people in danger. All other physical security objectives (protecting property, maintaining operations) are secondary.

*Subtopic: 3.8 — [Physical Security](../concepts/d3b/physical-security.md)*

---

### Q12: What is the difference between CRL and OCSP in the context of PKI certificate revocation?

**Answer:** CRL requires downloading the full list of revoked certificates; OCSP queries the CA for the status of a specific certificate and receives a simple yes/no response.

**Why:** CRL (Certificate Revocation List) is the older method — it requires the client to download a potentially very large list and search it for the certificate in question. OCSP (Online Certificate Status Protocol) is more efficient: the client sends the specific certificate's serial number and the CA/VA responds with its current status. OCSP is faster, uses less bandwidth, and is the modern preferred method.

*Subtopic: 3.6 — [PKI](../concepts/d3b/pki.md)*

---

### Q13: Why is WEP considered broken, and what does this illustrate about cryptography?

**Answer:** WEP is broken because RC4 (a sound stream cipher) was implemented with initialization vectors (IVs) that are too short and repeat too frequently, not because of a flaw in RC4 itself. This illustrates an implementation attack.

**Why:** RC4 is actually a strong algorithm. WEP's failure was in how it was used: the IVs were only 24 bits long, causing them to repeat frequently in high-traffic networks. An attacker who can collect enough ciphertext produced with a reused IV can determine the keystream and decrypt traffic. This is the classic example of an implementation attack — the algorithm is sound, but the implementation is fatally flawed.

*Subtopic: 3.7 — [Cryptanalysis Attacks](../concepts/d3b/cryptanalysis-attacks.md)*

---

### Q14: In a physical security context, what is a mantrap and what attack does it prevent?

**Answer:** A mantrap is a double-door airlock or single-person turnstile that prevents tailgating (also called piggybacking).

**Why:** Tailgating occurs when an unauthorized person follows an authorized person through a secured entrance without authenticating. A mantrap prevents this by requiring each entrant to independently authenticate (typically using two factors — e.g., access card + biometric) in an isolated vestibule before gaining access. If tailgating is detected, the person is trapped in the vestibule. Mantraps only allow one person (or a small controlled number) through at a time.

*Subtopic: 3.9 — [Physical Security](../concepts/d3b/physical-security.md)*

---

### Q15: What is the Kerckhoffs's principle and why does it matter for cryptographic system design?

**Answer:** Kerckhoffs's principle states that a cryptosystem should be secure even if everything about the system except the key is public knowledge.

**Why:** This principle means that cryptographic security must reside entirely in the secrecy of the key — not in the secrecy of the algorithm. Security through obscurity (hiding the algorithm) is not a valid security model because algorithms are eventually reverse-engineered or disclosed. Modern cryptographic algorithms (AES, RSA) are publicly known and scrutinized by the global security community; their strength comes from mathematical properties, not secrecy of design. The corollary: if an attacker knows the ciphertext, the algorithm, the IV, and everything except the key — the system should still be unbreakable.

*Subtopic: 3.6 — [Key Management](../concepts/d3b/key-management.md), [Cryptography Fundamentals](../concepts/d3b/cryptography-fundamentals.md)*
