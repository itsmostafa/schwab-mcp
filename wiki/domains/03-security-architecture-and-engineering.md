---
title: "Domain 3 — Security Architecture and Engineering"
type: domain
domain: 3
tags: [cryptography, architecture, secure-design, ics, cloud, iot]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 3 — Security Architecture and Engineering

Exam weight: **13%**. Covers secure design principles, security models, cryptographic
solutions, and physical/facility security.

## Subtopics (from CISSP Exam Outline)

### 3.1 Research, implement and manage engineering processes using secure design principles

Security must be embedded from the very beginning of the engineering life cycle ("security by design"), not added as an afterthought. The risk management process drives control selection: identify valuable assets, assess risks, select cost-effective controls. Key principles include [Secure Design Principles](../concepts/secure-design-principles.md): Zero Trust (trust nothing; authenticate and authorize every user, device, and service), Privacy by Design (seven foundational principles with privacy as a proactive default), Shared Responsibility (cloud customers retain accountability even when delegating responsibility), and the Cyber Kill Chain (reconnaissance → weaponization → delivery → exploitation → installation → command and control → actions on objectives).

- Threat modeling
- Least privilege
- Defense in depth
- Secure defaults
- Fail securely
- Segregation of Duties (SoD)

### 3.2 Understand the fundamental concepts of security models

Security models are formal representations of what security must look like in an architecture; they fall into two categories: lattice-based (Bell–LaPadula and Biba) and rule-based (Clark–Wilson, Brewer–Nash, Graham–Denning, Harrison–Ruzzo–Ullman). See [Security Models](../concepts/security-models.md). Evaluation criteria systems — [TCSEC](../standards/tcsec.md) (Orange Book, confidentiality only, DoD 1980s), ITSEC (European successor, separate functional and assurance ratings E0–E6), and [Common Criteria](../concepts/common-criteria.md) (ISO/IEC 15408, current standard, EAL 1–7) — provide independent, objective measurement of vendor products. Certification is the technical analysis of a solution; accreditation is management's official sign-off for a predetermined period.

- Biba, Star Model, Bell-LaPadula

### 3.3 Select controls based upon systems security requirements

Control selection is driven by risk management: identify assets, assess risks, select cost-effective mitigating controls. Security control frameworks provide best-practice guidance and can be combined to meet organizational needs. Key frameworks applicable to Domain 3 include: COBIT (IT assurance and audits, from ISACA), ITIL (IT service management), NIST SP 800-53 (cybersecurity controls best practices), PCI DSS (payment card industry), ISO 27001/27002 (ISMS requirements and implementation guidance; organizations can be certified against 27001), HIPAA (healthcare PHI protection), FISMA (US federal agency security programs), FedRAMP (cloud services for US federal data), COSO (enterprise risk management), and SOX (financial fraud prevention). Organizations typically rationalize overlapping frameworks into a unified control set and test controls once rather than repeatedly.

### 3.4 Understand security capabilities of Information Systems (IS)

All security within information systems operates in terms of subjects (active entities accessing) and objects (passive entities being accessed). The [Trusted Computing Base (TCB)](../concepts/trusted-computing-base.md) is the totality of all protection mechanisms; the Reference Monitor Concept (RMC) defines the theory of mediated access, while the security kernel is its implementation (requiring completeness, isolation, and verifiability). The Ring Protection Model layers CPU privilege — Ring 0 (most trusted: firmware, OS kernel) through Ring 3 (user apps). The [TPM](../concepts/tpm.md) (ISO/IEC 11889) is a hardware chip providing cryptographic services and platform integrity: binding ties encrypted data to a specific TPM's hardware; sealing ties decryption to a specific system state. Defense in depth (layered controls) and abstraction/virtualization are additional security capabilities covered in this section.

- Memory protection
- Trusted Platform Module (TPM)
- Encryption/decryption

### 3.5 Assess and mitigate the vulnerabilities of security architectures, designs, and solution elements

Every system type has specific vulnerabilities; the mitigation approach centers on hardening (reducing attack surface by securing each component based on its value and function). Key cross-cutting vulnerabilities include: single points of failure (mitigated by redundancy where cost-justified), bypass controls (intentional but must be protected with SoD, logging, and physical security), TOCTOU/race conditions (mitigated by more frequent access checks), and emanations (mitigated by TEMPEST shielding, white noise, or control zones). [ICS/SCADA systems](../concepts/ics-scada-security.md) run critical infrastructure on legacy hardware and are best protected by air gapping; [IoT devices](../concepts/iot-security.md) are cheap, unpatched, and mass-connected — segment them on isolated networks. [Cloud security](../concepts/cloud-security-models.md) requires understanding the shared responsibility model across IaaS/PaaS/SaaS/CaaS/FaaS, and [virtualization](../concepts/virtualization-security.md) requires hypervisor hardening to prevent VM escape. Web vulnerabilities include XSS (target: user browser; prevent with server-side input validation and WAF) and CSRF (target: web server; mitigate with cookie expiry).

- Client-based systems
- Server-based systems
- Database systems
- Cryptographic systems
- Industrial Control Systems (ICS)
- Cloud-based systems (SaaS, IaaS, PaaS)
- Distributed systems
- IoT
- Microservices / API
- Containerization
- Serverless
- Embedded systems
- High-Performance Computing systems
- Edge computing systems
- Virtualized systems

### 3.6 Select and determine cryptographic solutions

Cryptography ("secret writing") provides up to five services: **confidentiality, integrity, authenticity, non-repudiation, and access control**. The most critical aspect of any cryptographic system is **key management** — an attacker who knows the ciphertext, algorithm, and IV but not the key still cannot break the system (Kerckhoffs's principle). Modern cryptography is electronic, relying on algorithms such as DES, AES, and RSA. Foundational terms include plaintext/ciphertext, encryption/decryption key, key clustering (two different keys producing the same ciphertext — bad), work factor (effort to break), initialization vector/nonce (random value preventing patterns), and the **avalanche effect** (≥50% of ciphertext bits change if a single bit of plaintext or key changes; driven by **confusion** — key–ciphertext relationship — and **diffusion** — plaintext–ciphertext relationship).

**Substitution and transposition** are the two primitive methods of encryption; modern algorithms apply many rounds of both. Patterns in ciphertext are a primary weakness (frequency analysis exploits them); polyalphabetic and one-time-pad ciphers eliminate patterns. Stream ciphers (RC4, ChaCha20) encrypt one bit at a time via XOR with a keystream — fast, favored at hardware/network layer. Block ciphers (DES, AES) work on fixed-size chunks. Block cipher modes: **ECB** (no IV, fastest, insecure for repeating data), **CBC/CFB/OFB** (all use IV, good for messages), **CTR** (counter as IV, fastest and most used for long messages), **GCM** (adds authentication).

**Symmetric cryptography** advantages: extremely fast, strong. Disadvantages: key distribution and scalability problems (formula: n × (n−1) / 2 keys). Notable symmetric algorithms ranked weak→strong: RC2-40 (40-bit), DES (56-bit, 16 rounds, 64-bit block, deprecated), Skipjack (80-bit), IDEA (128-bit), Blowfish (128-bit, 64-bit block), 3DES (168-bit nominal, **112-bit effective** due to meet-in-the-middle attack, now disallowed by NIST), Twofish (256-bit), Rijndael/AES (128/192/256-bit key, always 128-bit block, 10/12/14 rounds). ChaCha20-Poly1305 (256-bit stream cipher, used by Google/Cloudflare as AEAD alternative to AES-GCM).

**Asymmetric cryptography** solves the key distribution problem using mathematically linked key pairs (public + private). Two hard math problems underlie all asymmetric algorithms: **factoring** (RSA) and **discrete logarithms** (Diffie–Hellman, ECC, ElGamal). Rule: encrypt with recipient's public key → decrypt with recipient's private key; sign with sender's private key → verify with sender's public key. Asymmetric is significantly slower than symmetric; RSA keys are moving toward 2048+ bits. **ECC** achieves equivalent security to RSA with shorter keys (discrete log), making it faster and preferred on constrained devices. **Diffie–Hellman** is used almost exclusively for symmetric session key exchange (not message encryption). **Hybrid cryptography** (used in TLS/SSL) combines symmetric encryption for bulk data and asymmetric for key exchange.

**Hashing** produces a fixed-length digest regardless of input length; it is one-way and deterministic. Properties: collision resistance, preimage resistance, avalanche effect. Key algorithms: MD5 (128-bit, broken), SHA-1 (160-bit, deprecated), SHA-2 (224/256/384/512-bit), SHA-3 (224/256/384/512-bit). A **collision** (two inputs produce the same digest) undermines integrity; **birthday attacks** exploit collision probability mathematically. **HMAC** adds a symmetric key to a hash for keyed message authentication. **Digital signatures**: sender hashes the message then encrypts the hash with their private key → provides integrity + authenticity + non-repudiation (NOT confidentiality). Process to verify: decrypt with sender's public key → re-hash received message → compare. Uses: document signing, code signing.

**Digital certificates** bind an entity to their public key; issued per the **X.509** standard. A **Certificate Authority (CA)** signs the certificate with its private key. CA hierarchy: Root CA (offline, self-signed, "root of trust") → Intermediate CAs → Issuing CAs → Entity certificates (chain of trust). Certificate lifecycle: Enrollment (CSR) → Issuance (identity proofing, signed by root/intermediate) → Validation → Revocation → Renewal. Revocation methods: **CRL** (full list download, slow) vs **OCSP** (single-certificate query, fast). **Certificate pinning** caches a trusted certificate to prevent MITM substitution.

**Key management** activities: generation (automated, pseudorandom), distribution (out-of-band or key wrapping/KEK), storage (TPM for single device; HSM for organizational keys), rotation (frequency based on asset value), recovery (split knowledge, dual control, key escrow), and disposition/destruction (crypto shredding — encrypt then destroy key; physical destruction).

See concept pages: [Cryptography Fundamentals](../concepts/cryptography-fundamentals.md), [Symmetric Crypto](../concepts/symmetric-crypto.md), [Asymmetric Crypto](../concepts/asymmetric-crypto.md), [Hash Functions](../concepts/hash-functions.md), [Digital Signatures](../concepts/digital-signatures.md), [PKI](../concepts/pki.md), [Key Management](../concepts/key-management.md).

### 3.7 Understand methods of cryptanalytic attacks

Cryptanalysis is the science of cracking codes, breaking cryptographic protocols, and finding or deducing encryption keys. The **primary goal** of a cryptanalytic attack is always to determine the key. Two categories exist: **cryptanalytic attacks** (focus on the key via mathematical analysis) and **cryptographic attacks** (broader goals including interception, forgery, and system compromise).

**Cryptanalytic attacks** include: brute-force (try all keys; ineffective for ≥80-bit keys), ciphertext-only (hardest — attacker has only ciphertext), known-plaintext (attacker has matching plaintext+ciphertext pairs), chosen-plaintext (attacker feeds chosen plaintext, studies ciphertext; easiest), chosen-ciphertext (attacker feeds chosen ciphertext, studies resulting plaintext). Linear cryptanalysis uses known-plaintext; differential cryptanalysis uses chosen-plaintext. Factoring attacks target RSA by attempting to factor the large composite number used to generate the private key. Meet-in-the-middle attacks reduce 2-DES effective key length from 112 to 56 bits (making it no stronger than DES).

**Cryptographic attacks** include: man-in-the-middle (attacker impersonates both parties), replay (captured credentials/tokens replayed later), pass-the-hash (stolen hash used directly for authentication without cracking the password), implementation attacks (exploit weak algorithm implementation — e.g., WEP's short IV allows RC4 to be cracked), side-channel attacks (timing, power consumption, radiation), dictionary attacks (try likely passwords from wordlists), rainbow tables (precomputed hash→password lookups; mitigated by salting), birthday attacks (exploit hash collision probability), social engineering, Kerberos attacks (golden ticket — forged TGT via KRBTGT hash; silver ticket — forged TGS for a specific service), ransomware, and fault injection attacks (deliberate hardware/software faults to exploit access controls).

See concept page: [Cryptanalysis Attacks](../concepts/cryptanalysis-attacks.md).

### 3.8 Apply security principles to site and facility design

Physical security extends information security's CIA triad to the physical world, protecting from the perimeter inward. The **primary goal of physical security is the safety and protection of human life** — all control decisions must prioritize this. Physical security controls map to: **Deter/Prevent** (fences, signs), **Delay** (locks), **Detect** (CCTV, motion sensors, alarms), **Assess**, and **Respond**. Defense-in-depth (layered) is the model: outer perimeter → building perimeter → interior zones → high-value areas (server rooms, wiring closets, media storage, evidence storage, restricted work areas).

See concept page: [Physical Security](../concepts/physical-security.md).

### 3.9 Design site and facility security controls

Physical security controls are selected via a **security/site survey**. **CPTED** (Crime Prevention Through Environmental Design) provides guidelines for how building and landscaping design can naturally deter criminal activity. Key controls: bollards (vehicle attack prevention), CCTV (detective), PIR sensors (motion detection), mantraps (double-door airlocks prevent tailgating; two authentication factors), locks (delay control — biometric most accurate), walls extending from true floor to true ceiling, raised floors for cooling airflow.

**Infrastructure**: UPS (short-term battery bridging + power conditioning), generators (long-term, ~1-minute spin-up). Power disruption types: fault/blackout, sag-dip/brownout, spike/surge. **HVAC**: ASHRAE TC 9.9 — temperature 64.4–80.6°F, humidity 40–60% RH; positive pressurization. **Fire suppression**: wet pipe (always pressurized), dry pipe (gas-filled until activation), **pre-action** (detection required first — best water-based option for data centers), deluge (all heads open). Gas agents: INERGEN, Argonite, FM-200 (all preferred for data centers); Halon is **illegal** (ozone damage). Extinguisher classes: A (combustibles), B (liquids), C (electrical), D (metals), K (kitchen).

See concept pages: [Facility Security](../concepts/facility-security.md), [Physical Security](../concepts/physical-security.md).

### 3.10 Manage the information system lifecycle

The information system lifecycle encompasses the complete lifespan of an information system from conceptualization through decommissioning. It is functionally equivalent to the SDLC covered in Domain 8. Security must be integrated at every phase — not bolted on after deployment. The lifecycle provides the governance framework ensuring systems are designed, built, tested, operated, and retired securely.

See concept page: [IS Lifecycle](../concepts/is-lifecycle.md). For detailed SDLC security integration, see [Domain 8 — Software Development Security](../domains/08-software-development-security.md).

## Key concepts

**Architecture & models (§3.1–3.5):**
- [Secure Design Principles](../concepts/secure-design-principles.md) — least privilege, defense in depth, zero trust, fail-secure, privacy by design, shared responsibility
- [Security Models](../concepts/security-models.md) — lattice-based vs. rule-based model hub
- [Bell–LaPadula](../concepts/bell-lapadula.md) — confidentiality model; no read up, no write down
- [Biba](../concepts/biba.md) — integrity model; no read down, no write up
- [Clark–Wilson](../concepts/clark-wilson.md) — commercial integrity model; well-formed transactions, SoD, access triple
- [Common Criteria](../concepts/common-criteria.md) — ISO/IEC 15408; EAL 1–7; PP/TOE/ST; certification vs. accreditation
- [Trusted Computing Base](../concepts/trusted-computing-base.md) — TCB, RMC, security kernel, ring protection model, process isolation
- [TPM](../concepts/tpm.md) — hardware security chip; binding and sealing
- [Virtualization Security](../concepts/virtualization-security.md) — Type 1/2 hypervisors, VM escape, containers
- [Cloud Security Models](../concepts/cloud-security-models.md) — IaaS/PaaS/SaaS shared responsibility; deployment models; cloud roles
- [ICS/SCADA Security](../concepts/ics-scada-security.md) — OT/ICS/SCADA/DCS/PLC; air gapping; legacy patching challenges
- [IoT Security](../concepts/iot-security.md) — mass-produced insecure devices; DDoS via botnet; network segmentation
- [Mobile Security](../concepts/mobile-security.md) — MDM vs. MAM; BYOD/COPE/CYOD; OWASP Mobile Top 10

**Cryptography & physical (§3.6–3.10):**
- [Cryptography Fundamentals](../concepts/cryptography-fundamentals.md) — five services, Kerckhoffs's principle, stream vs. block, modes of operation
- [Symmetric Crypto](../concepts/symmetric-crypto.md) — DES/3DES/AES, algorithm comparison table, block cipher modes
- [Asymmetric Crypto](../concepts/asymmetric-crypto.md) — RSA/ECC/DH, key direction rules, hybrid cryptography
- [Hash Functions](../concepts/hash-functions.md) — MD5/SHA-1/SHA-2/SHA-3, digest sizes, HMAC, birthday attack
- [Digital Signatures](../concepts/digital-signatures.md) — integrity + authenticity + non-repudiation; sign with private, verify with public
- [PKI](../concepts/pki.md) — CA hierarchy, X.509, CRL vs. OCSP, certificate pinning
- [Key Management](../concepts/key-management.md) — lifecycle, HSM, split knowledge, dual control, key escrow, crypto shredding
- [Cryptanalysis Attacks](../concepts/cryptanalysis-attacks.md) — full attack taxonomy: brute-force, side-channel, rainbow tables, golden/silver tickets
- [Physical Security](../concepts/physical-security.md) — CPTED, layered defense, CCTV, mantraps, locks
- [Facility Security](../concepts/facility-security.md) — power (UPS/generator), HVAC, fire detection/suppression systems
- [IS Lifecycle](../concepts/is-lifecycle.md) — system lifecycle phases; see also Domain 8 SDLC

**Standards:**
- [Common Criteria Standard](../standards/common-criteria-standard.md) — ISO/IEC 15408 reference
- [TCSEC](../standards/tcsec.md) — Orange Book; historical predecessor to CC
- [FIPS 140](../standards/fips-140.md) — cryptographic module security levels 1–4; HSM certification
- [AES Standard](../standards/aes-standard.md) — FIPS 197; Rijndael; 10/12/14 rounds; approved modes
- [PKCS Standards](../standards/pkcs-standards.md) — PKCS#1/7/10/11/12

## Sources
- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §3.1–3.10
