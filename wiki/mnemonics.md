---
title: Mnemonics
type: glossary
domain: cross
tags: [mnemonics, memory-aids]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Mnemonics

Curated memory aids for CISSP topics. Only high-value ones that actually help.
Format: the trick → what it stands for → domain/concept it applies to.

## Domain 1 — Security and Risk Management

### "Protect, Act, Provide, Advance"
**Stands for**: ISC2 Code of Ethics — 4 Canons in order
1. **Protect** society, the common good, public trust, and infrastructure
2. **Act** honorably, honestly, justly, responsibly, and legally
3. **Provide** diligent and competent service to principals
4. **Advance** and protect the profession

**Subtopic**: D1.1 — ISC2 Code of Professional Ethics

**Notes**: The verb sequence (Protect → Act → Provide → Advance) is a natural escalation from the macro (society) to the micro (profession) and matches the canon priority order when conflicts arise. Society ALWAYS beats principals.

---

### "Please Call Steve If Anybody Asks Me"
**Stands for**: NIST RMF 7 steps in order (SP 800-37 Rev. 2)
- **P**lease → **Prepare**
- **C**all → **Categorize**
- **S**teve → **Select**
- **I**f → **Implement**
- **A**nybody → **Assess**
- **A**sks → **Authorize**
- **M**e → **Monitor**

**Subtopic**: D1.9 — Risk Management Frameworks (NIST SP 800-37)

**Notes**: Rev. 2 added "Prepare" (Step 1) — key differentiator from earlier versions. Step 6 (Authorize) = senior management decision.

---

### "STRIDE Maps to CIA + AuthAuth"
**Stands for**: STRIDE threat categories and the security pillar each violates
- **S**poofing → Authentication
- **T**ampering → Integrity
- **R**epudiation → Nonrepudiation
- **I**nformation Disclosure → Confidentiality
- **D**enial of Service → Availability
- **E**levation of Privilege → Authorization

**Subtopic**: D1.10 — Threat Modeling (STRIDE)

**Notes**: "E" = **Authorization** (not authentication). Elevation = already authenticated, gaining more access.

---

### "Avoid, Transfer, Mitigate, Accept — in that order of assertiveness"
**Stands for**: The 4 valid risk treatment options (spectrum from most to least assertive)
- **Avoid** — eliminate the risk by stopping the activity
- **Transfer** — shift financial burden via insurance (accountability stays with you)
- **Mitigate** — implement controls to reduce risk to an acceptable residual level
- **Accept** — consciously decide to live with the risk (must be made by the asset owner)

**Subtopic**: D1.9 — Risk Response/Treatment

**Notes**: If control cost > ALE, the answer is Accept, not Mitigate. "Ignore" is NOT a valid option.

---

### "SLE times ARO equals ALE"
**Stands for**: The quantitative risk formula chain
- SLE = AV × EF (Single Loss Expectancy = Asset Value × Exposure Factor)
- ALE = SLE × ARO (Annualized Loss Expectancy = SLE × Annualized Rate of Occurrence)

**Subtopic**: D1.9 — ALE Calculation

**Notes**: Calculate SLE first (one occurrence), then multiply by how often it happens (ARO) to get the annual cost (ALE).

---

### "P-Do-C-Act spins like a wheel"
**Stands for**: Deming/PDCA Cycle for continuous improvement
- **P**lan → determine which controls to implement
- **D**o → implement the controls
- **C**heck → monitor/assess effectiveness
- **A**ct → take corrective action → back to Plan

**Subtopic**: D1.9 — Continuous Improvement (Risk Maturity)

**Notes**: The wheel metaphor reinforces that PDCA is cyclical and never-ending. Risk management is continuous, not one-time.

---

## Domain 2 — Asset Security

### "Destroy beats Purge beats Clear — DPC like a security degree"
**Stands for**: The sanitization effectiveness order: **D**estroy > **P**urge > **C**lear

**Subtopic**: D2.4 — Data Destruction / Data Security Controls

**Notes**: Overwriting = Clear. Crypto-shredding = Purge (if key is gone). Incineration = Destroy. Best for SSDs: physical destruction. Best for cloud: crypto-shredding.

---

### "Owners Accountable, Custodians Responsible — OACR"
**Stands for**: **O**wners are **A**ccountable; **C**ustodians are **R**esponsible (for technical implementation)

**Subtopic**: D2.3 — Data Roles

**Notes**: Accountability can never be delegated. "Who is ultimately accountable?" = always the owner. "Who does the technical work?" = custodian.

---

### "Labels Speak to Systems; Marks Speak to Humans — LSSM"
**Stands for**: **L**abeling = **S**ystem-readable; **M**arking = **H**uman-readable

**Subtopic**: D2.1 — Labeling and Marking

**Notes**: Labeling uses metadata, barcodes, RFID — machine reads it. Marking uses document headers, stamps, physical notices — humans read and act on it.

---

### "Can Some Users Share Archives Deliberately?"
**Stands for**: **C**reate → **S**tore → **U**se → **S**hare → **A**rchive → **D**estroy

**Subtopic**: D2.4 — Information Life Cycle

**Notes**: Locks in the six data lifecycle phases in order. Pay attention to the Use/Share order — Use comes before Share.

---

## Domain 3 — Security Architecture and Engineering

### Bell–LaPadula vs. Biba — "Mirror Images"
**Stands for**: Direction rules for the two most confused security models

| | No Read | No Write |
|---|---|---|
| **BLP** | Up | Down |
| **Biba** | Down | Up |

- BLP: protect secrets from leaking down → "no read up, no write down"
- Biba: protect quality from being polluted from below → "no read down, no write up"

**Subtopic**: D3.2 — Bell–LaPadula and Biba Models

**Notes**: "BLP protects secrets from leaking down. Biba protects quality from being polluted from below." They are exact mirrors.

---

### Clark–Wilson Three Goals — "UAC"
**Stands for**: Three integrity goals (1) **U**nauthorized changes blocked, (2) **A**uthorized-bad changes blocked, (3) **C**onsistency maintained

**Subtopic**: D3.2 — Clark–Wilson

**Notes**: Biba only covers #1. Clark–Wilson covers all three.

---

### Clark–Wilson Three Rules — "WAS"
**Stands for**: (W) **W**ell-formed transactions, (A) **A**ccess Triple (Subject → Program → Object), (S) **S**eparation of Duties

**Subtopic**: D3.2 — Clark–Wilson

---

### Common Criteria EAL — Anchor Points
**Stands for**: Key EAL levels to memorize
- **EAL 1** = lowest (functionally tested only)
- **EAL 3** = operating systems (methodically tested + checked)
- **EAL 4** = commercial ceiling; firewalls (methodically designed, tested, reviewed)
- **EAL 7** = highest (formally verified; rare/military)

**Subtopic**: D3.3 — Common Criteria

**Notes**: "Formally verified" language = high assurance (EAL 5–7).

---

### TCB Security Kernel — "CIV"
**Stands for**: **C**ompleteness (cannot bypass), **I**solation (tamper-proof), **V**erifiability (logging confirms operation)

**Subtopic**: D3.1 — Trusted Computing Base

---

### TCSEC Divisions — "Dogs Can Be Able"
**Stands for**: **D**ivision D (minimal) → **C** (C1 weak, C2 strict login) → **B** (labeled, covert channel checks, startup) → **A** (verified design)

**Subtopic**: D3.3 — TCSEC

---

### "Public for Privacy, Private for Proof"
**Stands for**: Asymmetric key direction rules
- Encrypt with recipient's **public** key → only recipient's **private** key can decrypt → **confidentiality**
- Sign with sender's **private** key → anyone with sender's **public** key can verify → **authenticity + non-repudiation**

**Subtopic**: D3.6 — Asymmetric Cryptography

**Notes**: Concept: [Asymmetric Crypto](concepts/asymmetric-crypto.md), [Digital Signatures](concepts/digital-signatures.md)

---

### Hash Output Sizes — "MD5=128, SHA-1=160"
**Stands for**: The two broken hash algorithms have "odd" sizes relative to powers of 2 — easy to forget

| Algorithm | Size |
|---|---|
| MD5 | 128 bits |
| SHA-1 | 160 bits |
| SHA-256 | 256 bits |
| SHA-512 | 512 bits |

**Subtopic**: D3.6 — Hash Functions

**Notes**: SHA-2 and SHA-3 share size options (224/256/384/512). [Hash Functions](concepts/hash-functions.md)

---

### Symmetric Algorithm Ranking — "Rats Don't Skip Right To Amazing"
**Stands for**: **R**C2-40 → **D**ES → **S**kipjack → RC2-128/IDEA/Blowfish → **T**wofish/RC6 → **A**ES (Rijndael)

**Subtopic**: D3.6 — Symmetric Cryptography

**Notes**: Weak → Medium → Strong → Very Strong. [Symmetric Crypto](concepts/symmetric-crypto.md)

---

### 3DES Effective Key — "168 minus 56 = 112"
**Stands for**: 3DES uses three 56-bit keys = 168 bits nominal. The meet-in-the-middle attack removes one key's worth (56 bits). Effective = **112 bits**.

**Subtopic**: D3.6 — Symmetric Cryptography / Cryptanalysis

---

### AES Rounds — "10-12-14 = 128-192-256"
**Stands for**: Add 2 rounds for each step up in key size

| Key Size | Rounds |
|---|---|
| 128-bit | 10 rounds |
| 192-bit | 12 rounds |
| 256-bit | 14 rounds |

**Subtopic**: D3.6 — AES

---

### Hard Math Problems — "RSA Factors, ECC and DH Log"
**Stands for**: **RSA** = Factoring large primes; **ECC** and **Diffie–Hellman** = Discrete Logarithm

**Subtopic**: D3.6 — Asymmetric Cryptography

---

### Fire Triangle — "FOH" (Fuel, Oxygen, Heat)
**Stands for**: **F**uel + **O**xygen + **H**eat = Fire. Remove any one → fire goes out.

**Subtopic**: D3.9 — Facility Security

---

### Water-Based Fire Suppression — "Deluge Drowns Data; Pre-action Plays it Safe"
**Stands for**: Order worst to best for data center use: **Deluge** → **Wet Pipe** → **Dry Pipe** → **Pre-action** (best)

**Subtopic**: D3.9 — Facility Security

---

### Gas Fire Suppression — Halon Replacements "IAFA"
**Stands for**: **I**NERGEN → **A**rgonite → **F**M-200 → **A**ero-K

**Subtopic**: D3.9 — Facility Security

**Notes**: All are Halon replacements. All are safe for data centers. **Halon = ILLEGAL** (ozone damage).

---

### Fire Extinguisher Classes — "Awesome Blazes Catch Dangerous Kitchens"
**Stands for**: **A** = common combustibles, **B** = liquids, **C** = electrical, **D** = combustible metals, **K** = commercial kitchens

**Subtopic**: D3.9 — Facility Security

---

### Physical Security Primary Goal — "Life First, Then Data"
**Stands for**: Safety and protection of human life is always the #1 goal of physical security — not data protection.

**Subtopic**: D3.9 — Physical Security

---

### Power Disruptions — Duration Table
**Stands for**: Short vs. long duration, high vs. low vs. no power

|  | Short | Long |
|---|---|---|
| No power | **Fault** | **Blackout** |
| Low voltage | **Sag/Dip** | **Brownout** |
| High voltage | **Spike** | **Surge** |

**Subtopic**: D3.9 — Facility Security

---

### PKI CA Hierarchy — "ROOT IS OFFLINE"
**Stands for**: Root CA must not be directly accessible to prevent catastrophic compromise.

**Path**: Root CA (self-signed, offline) → Intermediate CA → Issuing CA → Entity Certificate

**Subtopic**: D3.7 — PKI

---

## Domain 4 — Communication and Network Security

### OSI Layer Order (Top→Bottom) — "All People Seem To Need Data Processing"
**Stands for**: Application (7) → Presentation (6) → Session (5) → Transport (4) → Network (3) → Data Link (2) → Physical (1)

**Subtopic**: D4.1 — OSI Model

**Notes**: Concept: [OSI Model](concepts/osi-model.md)

---

### OSI Layer Order (Bottom→Top) — "Please Do Not Throw Sausage Pizza Away"
**Stands for**: Physical (1) → Data Link (2) → Network (3) → Transport (4) → Session (5) → Presentation (6) → Application (7)

**Subtopic**: D4.1 — OSI Model

---

### TCP Three-Way Handshake — "See you, See you back, Acknowledged"
**Stands for**: **SYN, SYN-ACK, ACK**

**Subtopic**: D4.2 — TCP/IP Model

**Notes**: Graceful close: **FIN-ACK, FIN-ACK** — two pairs (four packets total).

---

### IPsec: AH vs ESP — "AH = Authentic Header (no Hide); ESP = Encrypt + Secret Payload"
**Stands for**:
- AH: Integrity + authentication + replay protection. **No encryption.**
- ESP: Everything AH does **plus encryption (confidentiality).**

**Subtopic**: D4.3 — VPN / IPsec

**Notes**: If you need confidentiality → use ESP (or both AH + ESP).

---

### IPsec Modes: Transport vs Tunnel — "T = Thin; T-T = Thick Tunnel"
**Stands for**:
- Transport mode: original IP header visible; payload encrypted (thin — keeps original header)
- Tunnel mode: entire original packet (header + payload) encrypted inside new outer packet (thick — wraps everything)

**Subtopic**: D4.3 — VPN / IPsec

**Notes**: Tunnel mode used for site-to-site VPNs.

---

### Wireless Security Evolution — "WEP Wept, WPA Patched, WPA2 Rules, WPA3 Rising"
**Stands for**: WEP (broken/RC4) → WPA (stopgap/TKIP/deprecated) → WPA2 (current/CCMP/AES) → WPA3 (latest/GCMP)

**Subtopic**: D4.5 — Wireless Security

---

### Firewall Intelligence vs Speed — "The higher the floor, the smarter but slower the security guard"
**Stands for**:
- Layer 3 (packet filter): fastest, dumbest (reads header only)
- Layer 5 (circuit proxy): middle
- Layer 7 (application proxy): smartest, slowest (inspects payload)

**Subtopic**: D4.4 — Firewalls

---

### IDS vs IPS Placement — "IDS Detects, Sits aside; IPS Prevents, sits In-Path"
**Stands for**: IDS = passive/mirror port; IPS = inline

**Subtopic**: D4.4 — IDS and IPS

---

### False Negative is Worst — "FN = Fatal Negative"
**Stands for**: Attack happening, no alert. The security team is blind.

**Subtopic**: D4.4 — IDS and IPS

**Notes**: False Positive = noisy but not dangerous. False Negative = dangerous.

---

## Domain 5 — Identity and Access Management

### AAA Sequence — "I Am Authorized And Accountable"
**Stands for**: **I**dentification → **A**uthentication → **A**uthorization → **A**ccounting

**Subtopic**: D5.1 — AAA

**Notes**: Concept: [AAA](concepts/aaa.md)

---

### Biometric Error Rates — "Type 2 is too bad"
**Stands for**:
- **FRR** = False Rejection Rate = Type **1** error = less dangerous (frustrated users)
- **FAR** = False Acceptance Rate = Type **2** error = more dangerous (bad actor gets in)
- **CER** = Crossover Error Rate; lower CER = better system

**Subtopic**: D5.4 — Biometrics

**Notes**: Tighten system → FAR goes down, FRR goes up. "Tight out, frustrated in." [Biometrics](concepts/biometrics.md)

---

### Kerberos Components — "KDC = AS + TGS; AS gives TGT, TGS gives the key"
**Stands for**:
- **KDC** contains **AS** + **TGS**
- AS issues the **TGT**; TGS issues the **Service Ticket**

**Flow**: "Alice Asks, AS Answers with TGT. TGS Takes TGT, Gives Service Ticket."

**Subtopic**: D5.3 — Kerberos

---

### SAML vs. OAuth — "SAML Says who, OAuth Opens what"
**Stands for**:
- **SAML** — covers both authentication (who you are) and authorization (what you can do)
- **OAuth** — authorization only (opens access to what resources)
- **OIDC** — adds authentication on top of OAuth ("OAuth + identity = OIDC")

**Subtopic**: D5.5 — Federation / SAML / OAuth

---

### MFA Factor Types — "Know, Have, Are"
**Stands for**: **Know** = Knowledge, **Have** = Ownership, **Are** = Characteristic (biometrics)

**Subtopic**: D5.2 — Authentication Factors

**Notes**: Two "know" factors = still **single-factor**. Factors must be from **different families**.

---

### Access Control Models — "Owner, System, Role, Attribute"
**Stands for**:
- **DAC** = *Owner* decides
- **MAC** = *System* decides (labels/clearances)
- **RBAC** = *Role* decides (job function)
- **ABAC** = *Attribute* decides (most granular)

**Subtopic**: D5.4 — Access Control Models

**Notes**: Ordering by flexibility: DAC < RBAC < ABAC (ABAC most flexible/granular)

---

## Domain 6 — Security Assessment and Testing

### SOC Report Types — "1, 2, 3: Finance, Security, Summary"
**Stands for**:
- **SOC 1** = Financial controls
- **SOC 2** = 5 Trust Services Criteria (security scope)
- **SOC 3** = Same as SOC 2, stripped down for the public

**Subtopic**: D6.4 — SOC Reports

---

### SOC Type 1 vs Type 2 — "Point vs Period"
**Stands for**:
- **Type 1** = One moment in time (design, snapshot)
- **Type 2** = Two things over time (design + operating effectiveness)

**Subtopic**: D6.4 — SOC Reports

**Notes**: Gold standard = **SOC 2, Type 2**.

---

### KPI vs KRI — "KPI = Past Performance; KRI = Risk ahead"
**Stands for**:
- **KPI** — Backward-looking. Did we hit our targets?
- **KRI** — Forward-looking. What risks are emerging?

**Subtopic**: D6.3 — Security Metrics

---

### Pen Test Phases — "REVER"
**Stands for**: **R**econ → **E**numeration → **V**ulnerability analysis → **E**xploit → **R**eport

**Subtopic**: D6.2 — Penetration Testing

**Notes**: The fork: vuln assessment stops at V (vulnerability analysis); pen test continues to exploit.

---

### SAST vs DAST — "Static Stops the app; Dynamic Does it live"
**Stands for**:
- **SAST** = Static = Source code, app NOT running, White box
- **DAST** = Dynamic = App IS running, Black box, behavior-focused

**Subtopic**: D6.2 — Security Testing Types

---

### SOC 2 Trust Services Criteria — "SAC the PProblem"
**Stands for**: **S**ecurity, **A**vailability, **C**onfidentiality (required — SAC) + **P**rocessing Integrity, **P**rivacy (optional)

**Subtopic**: D6.4 — SOC Reports

---

## Domain 7 — Security Operations

### Order of Volatility — "From the CPU out to the shelf"
**Stands for**: Registers/cache → RAM → Swap/page file → Disk → Remote logs → Long-term archives → External media

More volatile = collect first. The further from active computation, the less volatile.

**Subtopic**: D7.5 — Digital Forensics

---

### MOM = Motive, Opportunity, Means
**Stands for**: Three questions that drive an investigation: Did the suspect have the **MOM** to do it?

**Subtopic**: D7.5 — Chain of Custody

---

### DR Site Types — "Cold Cash Warms Up Hot Results"
**Stands for**:
- **Cold** site → cheap → weeks to recover
- **Warm** site → middle → days to recover
- **Hot** site → premium → hours to recover
- **Redundant** → maximum cost → instant

**Subtopic**: D7.3 — Disaster Recovery Sites

---

### RTO vs. RPO — "Time vs. Data"
**Stands for**:
- **RTO = T**ime to recover (how long can you wait?)
- **RPO = P**oint in time (how much data can you lose?)

**Subtopic**: D7.3 — RTO/RPO/MTD

**Notes**: RPO looks backward (data already lost); RTO looks forward (time until systems return).

---

### MTD Formula — "RTO + WRT = MTD"
**Stands for**: Recovery Time Objective + Work Recovery Time = Maximum Tolerable Downtime

**Subtopic**: D7.3 — RTO/RPO/MTD

---

### NIST IR Phases — "Pretty Detecting Cats Purr"
**Stands for**: **P**reparation → **D**etection & Analysis → **C**ontainment, Eradication & Recovery → **P**ost-Incident Activity

**Subtopic**: D7.4 — Incident Management

---

### BCM Priority Order — "People, Property, Profits"
**Stands for**: BCM's three goals in strict priority: 1. **People** (safety), 2. **Property** (minimize damage), 3. **Profits** (business survival)

**Subtopic**: D7.6 — BCP/DRP

---

### Backup Types — "Full Fills, Differential Does All, Incremental Is Daily"
**Stands for**:
- **Full** — backs up everything; resets all archive bits
- **Differential** — backs up since last Full; archive bit NOT reset
- **Incremental** — backs up since last backup of any kind; archive bit reset

**Subtopic**: D7.2 — Backup Strategies

**Notes**: Restore complexity: Full (1 tape) < Differential (2 tapes) < Incremental (many tapes)

---

### Fail Mode Priorities — "Safety > Security > Availability"
**Stands for**:
- **Fail-safe** → people's safety first (doors open)
- **Fail-secure** → system security first (doors lock)
- **Fail-soft** → availability first (everything passes)

**Subtopic**: D7.1 — Failure Modes

---

### Five Rules of Evidence — "AACCA"
**Stands for**: **A**uthentic, **A**ccurate, **C**omplete, **C**onvincing (Reliable), **A**dmissible

**Subtopic**: D7.5 — Chain of Custody

---

## Domain 8 — Software Development Security

### ACID Database Properties — "A Consistent Individual Delivers"
**Stands for**: **A**tomicity (all or nothing), **C**onsistency (with rules), **I**solation (invisible until complete), **D**urability (done = permanent)

**Subtopic**: D8.3 — Database Security

---

### CMMI Six Maturity Levels — "Incompetent people Initially Manage processes, then they Define, Quantify, and Optimize"
**Stands for**: Level 0 Incomplete → 1 Initial → 2 Managed → 3 Defined → 4 Quantitatively Managed → 5 Optimizing

**Subtopic**: D8.7 — CMMI/SAMM

**Notes**: Level 3 = proactive. Below 3 = reactive.

---

### SDLC Phase Order — "Infants Require Amiable Developers To Deliver Often"
**Stands for**: **I**nitiation/Planning → **R**equirements → **A**rchitecture & Design → **D**evelopment → **T**esting → **D**eployment → **O**perations & Maintenance

**Subtopic**: D8.1 — SDLC

**Notes**: Threat modeling happens at Architecture & Design. Accreditation happens at Deployment.

---

### OWASP Top 10 — Description table (rank-order not reliably memorable)
**Stands for**: Focus on what each category means:
| Remember This | Category |
|---|---|
| Permissions bypass | A01 Broken Access Control |
| Weak or missing crypto | A02 Cryptographic Failures |
| SQL/command injection | A03 Injection |
| Security in the blueprint | A04 Insecure Design |
| Wrong default configs | A05 Security Misconfiguration |
| Unpatched libraries | A06 Vulnerable & Outdated Components |
| Weak login | A07 Identification & Auth Failures |
| Tampered builds/deploys | A08 Software & Data Integrity Failures |
| Can't see the attack | A09 Logging & Monitoring Failures |
| Internal network requests | A10 SSRF |

**Subtopic**: D8.5 — OWASP Top 10

---

### Coupling and Cohesion — "Low Coupling, High Cohesion — Like a Good Team"
**Stands for**: Low dependency between modules (low coupling) + high focus within a module (high cohesion) = good design

**Subtopic**: D8.4 — Secure Coding Practices

---

### REST vs. SOAP — "REST is Relaxed, SOAP is Strict"
**Stands for**: **REST** = Relaxed, flexible, modern, multiple output formats; **SOAP** = Strict, XML-only, rigid standard but stronger error handling

**Subtopic**: D8.6 — API Security

---

### Buffer Overflow Mitigations — "All Developers Should Code Safely"
**Stands for**: **A**SLR → **D**EP (NX bit) → **S**tack canaries → **C**hecking (bounds/parameter) → **S**afe languages and library functions

**Subtopic**: D8.2 — Memory Safety

---

### SAMM Business Functions — "GDIVO"
**Stands for**: **G**overnance → **D**esign → **I**mplementation → **V**erification → **O**perations

**Subtopic**: D8.7 — CMMI/SAMM

---

## Cross-Domain

### Privacy by Design: 7 Principles — Structural Pattern
**Stands for**: Remember the structure in three groups:
- **First 2**: Proactive stance (proactive not reactive; privacy as default)
- **Middle 3**: Design commitments (embedded; full functionality; end-to-end security)
- **Last 2**: Accountability (visibility + transparency; respect for user)

**Subtopic**: D3.1 / D1 cross — Privacy and Security Architecture

**Notes**: #2 "privacy as default" ≈ firewall's implicit deny. #3 "embedded" = baked in, not bolted on.

---

### Zero Trust Principles — "Know your stuff, use policies, authenticate everywhere, focus monitoring, trust nobody"
**Stands for**: Eight ZT principles (architecture, identities, health, policies, authenticate everywhere, monitor devices/services, no trusted network, choose ZT-designed services)

**Subtopic**: D3.1 — Secure Design Principles (Zero Trust)
