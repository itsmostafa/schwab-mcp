---
title: Weak-Areas Cheatsheet
type: cheatsheet
domain: cross
tags: [cheatsheet, weak-areas, review]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Weak-Areas Cheatsheet

This is a living document for topics that are easy to confuse or forget. Each entry
is a minimum-viable summary — just enough to jog your memory. For full context, follow
the linked concept page.

---

## Domain 1 — Security and Risk Management

### ISC2 Ethics Canon Order

1. Protect **society** (common good, public trust, infrastructure)
2. Act **honorably** (honestly, justly, responsibly, legally)
3. Serve **principals** (diligently and competently)
4. Advance **the profession**

- Exam trap: society ALWAYS beats principals. If your employer asks you to do something harmful to society, Canon 1 overrides Canon 3.
- Canon order = conflict resolution order.

[Full concept page](concepts/professional-ethics.md)

---

### Due Care vs. Due Diligence

| Due Care | Due Diligence |
|---|---|
| Doing the right thing | Proving you did the right thing |
| Implementing controls | Reporting, auditing, demonstrating |
| The pentest itself | The pentest report to management |

- Exam trap: "The company implemented encryption" = due care. "The company had encryption verified by a third-party auditor and reported to the board" = due diligence.

[Full concept page](concepts/due-care-due-diligence.md)

---

### ALE Formula (Quantitative Risk Math)

**SLE = AV × EF**
**ALE = SLE × ARO**

- AV = Asset Value (monetary)
- EF = Exposure Factor (% of asset value lost per incident; always 0–100%)
- SLE = cost of one incident
- ARO = times per year it's expected to happen
- ALE = annual cost of the risk

**Rule**: Never spend more on a control than the ALE. If control cost > ALE → Accept.

[Full concept page](concepts/risk-management.md)

---

### Risk Treatment Options

| Option | When | Key Fact |
|---|---|---|
| Avoid | Stop the risky activity | Loses opportunity too (opportunity cost) |
| Transfer | Buy cyber insurance | Transfers **responsibility**, NOT accountability |
| Mitigate | Implement controls | Reduces to residual risk (never to zero) |
| Accept | Cost of control > ALE | Must be made by **asset owner**, not security |

- Exam trap: "Ignore" is NOT a valid option. Risk ignorance = negligence.
- Exam trap: After transfer, the data owner is STILL accountable.

[Full concept page](concepts/risk-management.md)

---

### Security Document Hierarchy

```
Policy (corporate law; mandatory; CEO/Board)
  └── Standard (specific product/solution; mandatory)
  └── Procedure (step-by-step; mandatory)
  └── Baseline (minimum config levels; mandatory)
  └── Guideline (recommendation; optional — no audit findings)
```

- Exam trap: Guidelines are **NOT mandatory** — you won't fail an audit for not following them.
- Policies ≠ Procedures. Policies say what; procedures say how.

[Full concept page](concepts/security-policies.md)

---

### Security Control Types

| Pre-Event (Safeguards) | Post-Event (Countermeasures) |
|---|---|
| Directive, Deterrent, Preventive, Compensating | Detective, Corrective, Recovery |

**Complete Control** = Preventive + Detective + Corrective (minimum)

Categories: Administrative / Logical-Technical / Physical

- Exam trap: Detective controls operate **after** the event.
- Exam trap: CCTV can be both **deterrent** (sign warning) and **detective** (recording).

[Full concept page](concepts/security-controls-types.md)

---

### BCP vs. DRP

| BCP | DRP |
|---|---|
| Keep business running DURING a disruption | RESTORE IT systems AFTER a disruption |
| Business-wide | IT/technical focused |
| BCP is the parent | DRP is a subset of BCP |

**Key metrics**:
- **MTD** = maximum tolerable downtime (sets the ceiling)
- **RTO** = recovery time objective (must be ≤ MTD)
- **RPO** = recovery point objective (max acceptable data loss in time)
- Formula: RPO ≤ RTO ≤ MTD

**Test types** (least → most disruptive): Tabletop → Walkthrough → Simulation → Parallel → Full Interruption/Cutover

[Full concept page](concepts/bcp-drp.md)

---

### Accountability vs. Responsibility

| Accountability | Responsibility |
|---|---|
| Cannot be delegated | Can be delegated |
| Only ONE person | Multiple people |
| "Where the buck stops" | "The doer" |

- Exam trap: Cloud provider stores your data → You are STILL accountable.
- Exam trap: The CEO is ultimately accountable when no other owner is identified.

[Full concept page](concepts/accountability-vs-responsibility.md)

---

### STRIDE Threat-to-Pillar Mapping

| STRIDE Letter | Threat | Pillar Violated |
|---|---|---|
| S | Spoofing | Authentication |
| T | Tampering | Integrity |
| R | Repudiation | Nonrepudiation |
| I | Information Disclosure | Confidentiality |
| D | Denial of Service | Availability |
| E | Elevation of Privilege | Authorization |

- Exam trap: "E" = **Authorization** (not authentication). Elevation = already authenticated, gaining more access.

[Full concept page](concepts/threat-modeling.md)

---

### Tailgating vs. Piggybacking

| Tailgating | Piggybacking |
|---|---|
| Attacker has a **fake badge** (looks real) | Attacker has **no badge** |
| Both = physical social engineering | Both = following someone through a door |

[Full concept page](concepts/social-engineering.md)

---

### NIST RMF 7 Steps (SP 800-37 Rev. 2)

1. **Prepare** — context and strategy
2. **Categorize** — what do we have? how sensitive?
3. **Select** — choose controls
4. **Implement** — deploy controls
5. **Assess** — are controls working?
6. **Authorize** — senior management risk acceptance
7. **Monitor** — continuous monitoring

- Exam trap: Step 6 (Authorize) = **senior management** decides, not the security team.
- Rev. 2 added "Prepare" (Step 1) — key differentiator from earlier versions.

[Full standard page](standards/nist-rmf.md)

---

### GDPR Key Facts

- 7 principles: Lawfulness/fairness/transparency, Purpose limitation, Data minimization, Accuracy, Storage limitation, Integrity & confidentiality, Accountability
- **72-hour** breach notification to Supervisory Authority
- Each EU member state has an independent **Supervisory Authority (SA)**
- Data subjects can lodge complaints with SA
- Applies to personal data of EU citizens regardless of where processed

[Full concept page](concepts/privacy.md)

---

### Personnel Security Quick Reference

- **Job Rotation** → detects/prevents fraud + provides cross-training
- **Mandatory Vacation** → forces someone else to cover; exposes hidden fraud
- **Separation of Duties** → requires ≥2 people for critical tasks (prevents single-person fraud)
- **Least Privilege** → minimum permissions needed; nothing more
- **Need-to-Know** → access only if required for the job

**Involuntary termination** = higher risk than voluntary → immediate access removal; possible escort.

[Full concept page](concepts/personnel-security.md)

---

## Domain 2 — Asset Security

### Owner vs. Custodian (HIGHEST YIELD TRAP)
- **Owner** = business role; determines value; classifies; approves access; **accountable** (cannot be delegated)
- **Custodian** = technical role (IT/DBA); backup, restore, operations; **responsible** only
- Exam trap: "IT manages the database" — IT is the **custodian**, not the owner. The HR Director who uses and knows the data is the owner.

[Data Roles](concepts/data-roles.md)

### Owner vs. Controller (GDPR TRAP)
- In GDPR: **Controller** = legal entity determining purpose of processing; **Processor** = entity doing the processing on behalf of controller
- Internally: **Owner** = person accountable for the asset; **Controller** is an organizational/legal concept
- destination-cissp conflates these in one row; treat separately in GDPR exam context

[Data Roles](concepts/data-roles.md)

### Classification vs. Categorization
- **Classification** = the system of classes (Top Secret, Secret, Confidential, Unclassified)
- **Categorization** = the act of assigning an asset to a class
- Exam trap: "Who categorizes an asset?" — the **owner** does (they perform the act of categorization using the classification system)

[Data Classification](concepts/data-classification.md)

### Labeling vs. Marking
- **Labeling** = system-readable (metadata, RFID, barcode) → machine enforces policy
- **Marking** = human-readable (document headers, stamps) → human enforces handling

[Data Classification](concepts/data-classification.md)

### Sanitization Order (CRITICAL)
- **Destroy > Purge > Clear** (most effective to least)
- Overwriting (any # of passes) = **Clear** (not Purge)
- Crypto shredding = **Purge** IF key is provably destroyed; **Clear** if key might be recoverable
- Degaussing = **Purge/Destroy boundary** (destroys data AND often the media)
- Formatting = **Clear (weakest)**
- Best for SSDs: vendor tools or **physical destruction** (overwriting doesn't work reliably on flash)
- Best for cloud: **crypto shredding**

[Data Security Controls](concepts/data-security-controls.md) | [NIST SP 800-88](standards/nist-sp-800-88.md)

### Data States and Protections
- **At Rest**: encryption (AES-256), access control, backup
- **In Transit**: end-to-end encryption (VPN), link encryption (per-hop, every node is a risk point), onion (TOR — anonymity + confidentiality, slow)
- **In Use**: homomorphic encryption, RBAC, DLP/DRP
- Exam: "HTTPS protects data in ___": **transit**
- Exam: "Homomorphic encryption protects data in ___": **use**

[Data States](concepts/data-states.md)

### DLP vs. DRM
- **DRM**: protects IP (content rights, licensing, DMCA) — narrower
- **DLP**: identifies, monitors, protects *any* sensitive data across all 3 states — broader

[DLP](concepts/dlp.md) | [Data Security Controls](concepts/data-security-controls.md)

### Data Lifecycle Phases
- Create → Store → Use → Share → Archive → Destroy
- Classification assigned at **Create**; drives all subsequent phases
- Owner remains **accountable** through the entire lifecycle including destruction

[Data Lifecycle](concepts/data-lifecycle.md)

### CIA-Based Classification (non-obvious)
- Most people classify data only on **confidentiality** (sensitivity)
- Modern best practice and exam guidance: classify on **all three CIA dimensions** — confidentiality (sensitivity), integrity (accuracy), availability (criticality)

[Data Classification](concepts/data-classification.md)

### Content Gaps (topics on exam outline NOT in destination-cissp)
- **CASB**: Cloud Access Security Broker — not covered in source; exam outline §2.6
- **Anonymization / Pseudonymization / Tokenization**: not directly covered in source; exam outline §2.6
- **Data location / cross-border transfer**: exam outline §2.4; not in source
- **NIST SP 800-60**: data categorization mapping; not in source

---

## Domain 3 — Security Architecture and Engineering

### Security Model Comparison Table

| Model | Property | No Read | No Write | Type | Commercial? |
|---|---|---|---|---|---|
| **Bell–LaPadula** | Confidentiality | Up | Down | Lattice | No (military) |
| **Biba** | Integrity | Down | Up | Lattice | Limited |
| **Clark–Wilson** | Integrity | N/A (rule-based) | N/A | Rule | Yes |
| **Brewer–Nash** | Confidentiality | Conflict-of-interest blocked | Same | Rule | Yes (financial) |
| **Lipner** | Conf. + Integrity | Combines BLP + Biba | Same | Implementation | Yes |

Biba and BLP are exact mirrors. Clark–Wilson adds authorized-bad-changes and consistency goals that Biba misses.

Full pages: [Bell–LaPadula](concepts/bell-lapadula.md) | [Biba](concepts/biba.md) | [Clark–Wilson](concepts/clark-wilson.md) | [Security Models hub](concepts/security-models.md)

---

### Clark–Wilson — Goals and Rules

**3 Goals of integrity:**
1. Prevent unauthorized subjects from making changes (Biba also)
2. Prevent authorized subjects from making bad changes (CW only)
3. Maintain system consistency (CW only)

**3 Rules:**
- Well-formed transactions
- Separation of duties
- Access Triple (Subject → TP → CDI; no direct access)

**Key terms:** CDI (protected data), UDI (untrusted input), TP (authorized program), IVP (integrity checker)

---

### Common Criteria EAL Levels

| EAL | Description | Who uses it |
|---|---|---|
| 7 | Formally verified | Rare; military/government |
| 6 | Semi-formally verified | Rare |
| 5 | Semi-formally designed + tested | High security |
| **4** | Methodically designed, tested, reviewed | **Firewalls; commercial ceiling** |
| **3** | Methodically tested + checked | **Operating systems** |
| 2 | Structurally tested | Basic |
| 1 | Functionally tested | Minimal |

- EAL does NOT change with patches; only major functionality changes trigger re-evaluation
- Higher is not always better — EAL 7 products are often too complex to use properly
- CC = ISO/IEC 15408

Full page: [Common Criteria](concepts/common-criteria.md)

---

### Certification vs. Accreditation

| | Certification | Accreditation |
|---|---|---|
| **What** | Technical analysis confirming solution meets needs | Management sign-off to use the certified solution |
| **Who** | Security/technical team | Asset owner / management |
| **Duration** | Point-in-time | Time-limited period |
| **Renewal** | As needed | At expiration of accreditation period |

Accreditation is management's job. Security certifies; management accredits.

---

### TCB / Reference Monitor / Security Kernel

```
Concept:   Reference Monitor Concept (RMC)
              |
Implementation: Security Kernel
              |
Totality:  Trusted Computing Base (TCB) ← includes ALL protection mechanisms
```

**Security Kernel must have (CIV):**
- **C**ompleteness — cannot be bypassed
- **I**solation — rules are tamper-proof
- **V**erifiability — logging + monitoring confirms correct operation

**System Kernel ≠ Security Kernel** (common trap)

**Ring Protection Model:**
- Ring 0 = most trusted (firmware, OS kernel)
- Ring 3 = least trusted (user applications)

Full page: [Trusted Computing Base](concepts/trusted-computing-base.md)

---

### TPM — Binding vs. Sealing

| | Binding | Sealing |
|---|---|---|
| **Ties to** | Specific TPM hardware (endorsement key) | Specific system state/conditions |
| **Result** | Data only accessible on that TPM | Data only accessible if conditions met (e.g., correct boot state) |
| **Purpose** | Key protection from disclosure | Tamper detection |

TPM is hardware (ISO/IEC 11889). Every TPM has a unique Endorsement Key burned in at manufacture.

Full page: [TPM](concepts/tpm.md)

---

### Cloud Service Models — Responsibility Split

| Model | Provider manages | Customer manages |
|---|---|---|
| **SaaS** | Physical → data (everything) | Users, access control |
| **PaaS** | Physical → runtime | Applications, data |
| **IaaS** | Physical → hypervisor | OS, apps, networking, data |
| **FaaS** | Everything (serverless) | Code/functions only |

**The cloud customer is ALWAYS accountable for their data — accountability cannot be delegated.**

Full page: [Cloud Security Models](concepts/cloud-security-models.md)

---

### Cloud Deployment Models — Quick Reference

| Model | Who accesses | Who owns hardware |
|---|---|---|
| Public | Everyone | Provider |
| Private | One org (trusted) | Org or provider (dedicated) |
| Community | Specific group (e.g., gov, hospitals) | Org or provider |
| Hybrid | Mix | Mix |

Multi-tenancy = public cloud characteristic only. Private cloud = no multi-tenancy.

---

### ICS/SCADA — Key Facts

| Type | Remote control? | Scope |
|---|---|---|
| SCADA | Yes | Large-scale; infrastructure |
| DCS | No (local only) | Large processing facilities |
| PLC | N/A (device-level) | Manufacturing processes |

**Best ICS protection = Air Gap** (isolate from internet and corporate network entirely)

When air gap isn't possible: VLANs, logging, anomaly detection, vulnerability assessments

OT > ICS > SCADA/DCS/PLC (hierarchy)

Full page: [ICS/SCADA Security](concepts/ics-scada-security.md)

---

### Mobile Security — MDM vs. MAM

| | MDM | MAM |
|---|---|---|
| **Scope** | Full device | Applications only |
| **Best for** | COPE (corporate-owned) | BYOD (personal device) |
| **Remote wipe** | Full device wipe | App/data wipe only |

Remote wipe requires device to be online — attacker can defeat by keeping it offline.

Full page: [Mobile Security](concepts/mobile-security.md)

---

### XSS vs. CSRF

| | XSS | CSRF |
|---|---|---|
| **Target** | User's browser | Web server |
| **Mechanism** | Malicious JS injected into trusted site | Cookie persistence exploited to forge requests |
| **Types** | Stored (persistent), Reflected (most common), DOM-based | — |
| **Prevention** | Server-side input validation; WAF | Expire cookies/session tokens frequently |

**Reflected XSS is the most common type.** Stored XSS affects all visitors; reflected only affects the specific victim.

---

### Symmetric Algorithm Comparison Table

| Algorithm | Key Length | Block Size | Status | Type |
|---|---|---|---|---|
| RC2-40 | 40 bits | 64 bits | Deprecated | Block |
| **DES** | **56 bits** | **64 bits** | Deprecated | Block (16 rounds) |
| Skipjack | 80 bits | 64 bits | Deprecated | Block |
| IDEA | 128 bits | 64 bits | Secure | Block |
| Blowfish | 128 bits | 64 bits | Secure | Block |
| **3DES** | 168 (nominal) / **112 effective** | 64 bits | **Disallowed by NIST** | Block (48 rounds) |
| RC4 | Variable | N/A (stream) | Deprecated | **Stream** |
| Twofish | 256 bits | 128 bits | Secure | Block |
| **AES** | **128/192/256 bits** | **128 bits** | Current standard | Block (10/12/14 rounds) |
| ChaCha20 | 256 bits | N/A (stream) | Secure | **Stream** |

> Exam trap: AES block size is ALWAYS 128 bits regardless of key size.

---

### Hash Algorithm Output Sizes

| Algorithm | Output Size | Status |
|---|---|---|
| MD5 | **128 bits** | **Broken** — collisions trivially found |
| SHA-1 | **160 bits** | **Deprecated** — practical collisions demonstrated |
| SHA-256 | **256 bits** | Secure — current standard |
| SHA-512 | **512 bits** | Secure |
| SHA-3 | 224/256/384/512 bits | Secure — different algorithm (Keccak) |

> Memorize: **MD5=128, SHA-1=160** (the two broken ones have "unusual" sizes).

---

### Asymmetric Algorithm Quick Reference

| Algorithm | Hard Math | Primary Use |
|---|---|---|
| **RSA** | **Factoring** | Encryption, digital signatures |
| **ECC** | **Discrete log** | Encryption, signatures (smaller keys than RSA) |
| **Diffie–Hellman** | **Discrete log** | **Session key exchange ONLY** (not message encryption) |
| DSA | Discrete log | Digital signatures only |

---

### Key Direction Rules (Memorize)

| Goal | Use This Key |
|---|---|
| Encrypt message for Bob (confidentiality) | **Bob's PUBLIC key** |
| Decrypt message from Alice | Your own **PRIVATE key** |
| Create digital signature (prove you sent it) | Your own **PRIVATE key** |
| Verify Alice's digital signature | **Alice's PUBLIC key** |

---

### Block Cipher Mode Comparison

| Mode | IV? | Speed | Best Use | Gotcha |
|---|---|---|---|---|
| ECB | **No** | Fastest | Short, random, non-repeating (PIN codes) | Same plaintext → same ciphertext — patterns emerge |
| CBC | Yes | Moderate | Email, bulk messages | Sequential processing only |
| CFB | Yes | Moderate | Email, streaming | — |
| OFB | Yes | Moderate | Email | Bit errors don't propagate |
| **CTR** | Yes (counter) | **Fast** | **Long messages — most used** | Parallelizable |
| GCM | Yes | Fast | TLS 1.3, IPsec | AEAD — adds authentication |

---

### Physical Security Layers

```
Perimeter (fence, lighting, bollards, grading, landscaping)
    ↓
Building exterior (CCTV, mantraps, doors)
    ↓
Interior access (locks, biometrics, card readers)
    ↓
High-value rooms (server rooms, wiring closets)
```

**Primary goal: Safety and protection of human life** (not data)

---

### Fire Suppression Quick Reference

| System | Type | Data Center? | Notes |
|---|---|---|---|
| Wet pipe | Water | No | Always pressurized; freeze/leak risk |
| Dry pipe | Water | No | Gas until activated |
| Pre-action | Water | **Best water option** | Requires detection signal; targeted |
| Deluge | Water | **Never** | All heads open; use only for explosives |
| INERGEN/Argonite/FM-200/Aero-K | Gas | **Yes** | Halon replacements; safe for people |
| Halon | Gas | **ILLEGAL** | Ozone depleting; never use |
| CO2 | Gas | Caution | Effective but lethal in enclosed spaces |

---

### Crypto Common Exam Traps

1. **3DES effective key = 112 bits** (not 168 — meet-in-the-middle removes 56 bits)
2. **AES block size is ALWAYS 128 bits** regardless of key size
3. **ECB has NO IV** — only block mode without one; never use for repeating data
4. **Digital signatures provide NO confidentiality** — message body is readable
5. **Halon is ILLEGAL** — any answer involving new Halon installation is wrong
6. **CO2 is dangerous to people** in enclosed spaces — not ideal for occupied data centers
7. **Pre-action is the best water-based system** for data centers
8. **Deluge is never appropriate** for data centers (only explosives/fireworks)
9. **VESDA detects at incipient stage** — best and most expensive early detection
10. **CRL is the old method; OCSP is the new better method** for revocation checking
11. **Root CA should be OFFLINE** — if online, entire PKI trust hierarchy is at risk
12. **Locks are DELAY controls** — they do not prevent access
13. **Primary goal of physical security = human life** (not systems or data)
14. **CCTV is primarily a DETECTIVE control** (secondary: deterrent)

---

## Domain 4 — Communication and Network Security

### OSI Layers at a Glance

| # | Layer | PDU | Key Devices | Key Protocols | Security Device |
|---|-------|-----|-------------|---------------|-----------------|
| 7 | Application | Data | Gateway | HTTP/S, FTP, DNS, SSH, SMTP, SNMP | Application-proxy firewall |
| 6 | Presentation | Data | — | JPEG, XML, codecs | — |
| 5 | Session | Data | Circuit proxy FW | PAP, CHAP, EAP, NetBIOS, RPC | Circuit-proxy firewall |
| 4 | Transport | Segment/Datagram | — | TCP, UDP, SSL/TLS | Stateful FW (partial) |
| 3 | Network | Packet | Router | IP, ICMP, IPsec, OSPF | Packet-filtering FW |
| 2 | Data Link | Frame | Switch, Bridge | ARP, L2TP, PPTP | — |
| 1 | Physical | Bits | Hub, Repeater, NIC | — | — |

Mnemonic (top→bottom): **All People Seem To Need Data Processing**

---

### Firewall Types Comparison

| Type | OSI Layer | Inspects | Speed | Notes |
|------|-----------|---------|-------|-------|
| Packet Filter | 3 | Header (IP, port) | Fastest | Limited intelligence; uses ACLs |
| Stateful | 3–4 | Header + state table | Fast | Tracks connections; catches unexpected packets |
| Circuit Proxy | 5 | TCP handshake/session | Medium | Hides internal IPs via NAT; no payload inspection |
| Application Proxy | 7 | Full payload (DPI) | Slowest | Most intelligent; separate proxy per service |

**Rule:** Higher layer = more intelligence, more latency.

---

### IDS vs IPS

| | IDS | IPS |
|--|-----|-----|
| Action | Detect + alert | Detect + alert + **block** |
| Placement | **Passive** (mirror/span port) | **Inline** (in the traffic path) |
| Traffic impact | None | Can drop/modify packets |
| Failure mode | Traffic flows unaffected | May block legit traffic if misconfigured |

Detection: **Signature** (known threats, low FP) vs **Anomaly** (unknown threats, high FP).

Worst alert: **False Negative** (attack occurring, no alert).

---

### IPsec: AH vs ESP vs Modes

| Component | Integrity | Auth | Replay | Encryption |
|-----------|-----------|------|--------|-----------|
| AH | Yes | Yes | Yes | **No** |
| ESP | Yes | Yes | Yes | **Yes** |

| Mode | Original IP Header | What's encrypted |
|------|-------------------|-----------------|
| **Transport** | Kept | Payload only |
| **Tunnel** | Wrapped in new header | Original header + payload |

Site-to-site VPN → **Tunnel mode**. Host-to-host → Transport mode.

SA count (AH + ESP, bidirectional): **4 SAs total**.

IKE provides the symmetric key exchange (based on Diffie-Hellman).

---

### Wireless Security Protocol Evolution

| | WEP | WPA | WPA2 | WPA3 |
|--|-----|-----|------|------|
| Year | 1997 | 2003 | 2004 | 2018 |
| Encryption | RC4 (weak IV) | TKIP/RC4 | **CCMP/AES** | CCMP or GCMP |
| Integrity | None | Michael MIC | CCMP | CCMP/GCMP |
| Status | **Broken** | Deprecated | Current | Latest |

TKIP = stopgap fix for WEP; still RC4-based; **not secure**.
EAP-TLS = strongest EAP (mutual cert auth). LEAP = deprecated.

---

### Key Port Numbers

| Port | Protocol | Secure alt |
|------|----------|-----------|
| 20/21 | FTP | SFTP (22) |
| 22 | SSH / SFTP | — |
| 23 | Telnet | SSH (22) |
| 25 | SMTP | — |
| 53 | DNS | DNSSEC |
| 69 | TFTP (UDP) | SFTP (22) |
| 80 | HTTP | HTTPS (443) |
| 110 | POP3 | — |
| 143 | IMAP | — |
| 161/162 | SNMP (UDP) | SNMPv3 |
| 389 | LDAP | LDAPS (636) |
| 443 | HTTPS | — |
| 3389 | RDP | — |

---

### RADIUS vs TACACS+

| | RADIUS | TACACS+ |
|--|--------|---------|
| Transport | **UDP** | **TCP** |
| Encryption | Password only (poorly) | **Full packet** |
| AAA | Combined | **Separated** |
| Vendor | IETF standard | Cisco proprietary |
| Successor | Diameter | — |

---

### VPN Tunneling Protocols Cheat

- **GRE:** Encapsulation only. **No encryption. Not a VPN.**
- **L2TP:** Layer 2. No encryption. Pair with IPsec → VPN.
- **PPTP:** Deprecated.
- **IPsec:** Layer 3. Preferred for site-to-site. AH + ESP.
- **TLS:** Layer 4. Encrypts by default. Easier to configure.
- **SSH:** Layer 7. Secures Telnet, FTP, etc. inside tunnel.

VPN = Tunnel + Encryption. Without encryption = just a tunnel.

Full concept pages: [OSI Model](concepts/osi-model.md) | [TCP/IP Model](concepts/tcp-ip-model.md) | [Firewalls](concepts/firewalls.md) | [IDS and IPS](concepts/ids-ips.md) | [VPN](concepts/vpn.md) | [Wireless Security](concepts/wireless-security.md)

---

## Domain 5 — Identity and Access Management

### Access Control Model Comparison (DAC / MAC / RBAC / ABAC)

| Model | Who decides | Basis | Typical use | Key trait |
|---|---|---|---|---|
| **DAC** | Asset owner | Owner's discretion | General enterprise | Owner is accountable |
| **Rule-Based** | Owner (via rules) | ACL / rule table | Granular per-resource | High admin overhead |
| **RBAC** | Organization | Job function / role | Enterprises with clear roles | Best practice; mirrors org chart |
| **ABAC** | Policy engine | User + env attributes | Cloud, zero-trust | Most flexible; XACML standard |
| **MAC** | System | Labels (clearance/classification) | Government / military | Confidentiality focus; rare in private sector |

**ABAC enabler:** XACML (eXtensible Access Control Markup Language).

Full detail: [Access Control Models](concepts/access-control-models.md)

---

### Biometric Error Rates

| Term | Type | Severity | Meaning |
|---|---|---|---|
| FRR | Type 1 error | Low | Valid user rejected |
| FAR | Type 2 error | **HIGH** | Invalid user accepted |
| **CER** | Crossover point | — | Where FRR = FAR; lower = more accurate system |

- FAR ↑ when you loosen sensitivity (more false accepts, fewer false rejects).
- FRR ↑ when you tighten sensitivity (fewer false accepts, more false rejects).
- Most accurate biometric: **Retina** (invasive). Most common: **Fingerprint**.
- Iris ≠ Retina: Iris = colored ring (non-contact); Retina = back of eye (invasive).

Full detail: [Biometrics](concepts/biometrics.md)

---

### Kerberos Components

| Component | Role |
|---|---|
| **KDC** | Contains AS + TGS |
| **AS** (Authentication Service) | Verifies identity; issues TGT |
| **TGS** (Ticket Granting Service) | Issues service tickets |
| **TGT** | Token Alice gets from AS; encrypted with TGS key (Alice can't read it) |
| **Service Ticket** | Token from TGS; grants access to target service |

Weaknesses:
- **Symmetric encryption only** → key distribution challenges
- **TOCTOU vulnerability** → mitigate by increasing re-authentication frequency

Full detail: [Kerberos](concepts/kerberos.md)

---

### SAML vs. OAuth/OIDC

| | SAML | OAuth 2.0 | OpenID Connect | OpenID (orig.) |
|---|---|---|---|---|
| Authentication | Yes | **No** | Yes | Yes |
| Authorization | Yes | **Yes** | Yes (via OAuth) | No |
| Format | XML | JSON/tokens | JSON/JWT | — |
| Use case | Enterprise FIM/SSO | API delegation | Consumer identity | Decentralized auth |

**Critical exam point:** OAuth 2.0 = authorization ONLY. OIDC = authentication layer on OAuth.

Full detail: [SAML](concepts/saml.md) | [OAuth 2.0 and OIDC](concepts/oauth-oidc.md)

---

### MFA Factor Types (Exam Traps)

Three factor families:
1. **Something you know** — password, passphrase, security questions
2. **Something you have** — OTP token (hard/soft), smart card, passkey device
3. **Something you are** — biometrics (physiological + behavioral)

**Trap 1:** Two "know" factors (password + security question) = **single-factor** authentication (SFA). Factors must cross family lines.

**Trap 2:** RSA token + Microsoft token = both "have" = **still SFA**.

Full detail: [Authentication Factors](concepts/authentication-factors.md)

---

### AAA / Access Control Services

**Sequence:** Identification → Authentication → Authorization → Accounting

**Principle of Access Control = Accountability.** Requires all 4 steps.

Shared accounts = **undermine accountability** (cannot attribute actions to individuals).

**SSO vs. FIM:**
- SSO = one org, multiple systems
- FIM = multiple orgs, multiple systems (trust relationships + IdP/SP)

Full detail: [AAA](concepts/aaa.md) | [Federation](concepts/federation.md)

---

### Identity Lifecycle

**Provisioning → Review → Revocation**

| Phase | Trigger | Key action |
|---|---|---|
| Provisioning | New hire or role change | Background check, identity proof, grant access |
| Review | Periodic (at minimum annually; privileged accounts more often) | Owner reviews appropriateness |
| Revocation | Departure or role change | Remove all access; prevent orphaned accounts |

**Privilege creep:** Accumulated rights from multiple role changes. Prevented by full revoke-and-reprovision at each move.

Full detail: [Identity Lifecycle](concepts/identity-lifecycle.md) | [Privileged Access Management](concepts/privileged-access-management.md)

---

### AAL Levels (NIST SP 800-63B)

| Level | Auth Requirement | Confidence |
|---|---|---|
| AAL1 | Single-factor + secure protocol | Some assurance |
| AAL2 | MFA + secure protocol + approved crypto | High confidence |
| AAL3 | MFA + secure protocol + hard cryptographic authenticator + impersonation resistance | Very high confidence |

Full detail: [NIST SP 800-63B](standards/nist-sp-800-63b.md)

---

## Domain 6 — Security Assessment and Testing

### Assessment vs Audit vs Penetration Test

| | Vulnerability Assessment | Penetration Test | Security Audit |
|---|---|---|---|
| Exploitation? | **No** | **Yes** (key differentiator) | No |
| Primarily automated? | Yes | No (manual-driven) | No |
| Speed | Fast (minutes–days) | Slow (several days) | Weeks–months |
| Output | Findings report | Exploit proof + recommendations | Formal attestation |
| Audience | Internal security | Security + management | Management, customers, regulators |
| Stops at | Vuln identification | Exploitation confirmation | Control design/effectiveness |

Full page: [Assessment vs Audit vs Penetration Test](concepts/assessment-vs-audit.md)

---

### SAST / DAST / Fuzz / IAST / RASP

| | SAST | DAST | Fuzz | IAST | RASP |
|---|---|---|---|---|---|
| App running? | No | Yes | Yes | Yes | Yes (production) |
| Box type | White | Black | Black | Gray/both | N/A (internal) |
| Source code needed? | Yes | No | No | No | No |
| Finds | Code flaws | Runtime flaws | Edge cases (chaos) | Both | Blocks attacks live |

Full page: [Security Testing Types](concepts/security-testing-types.md)

---

### SOC Report Types

| | SOC 1 | SOC 2 | SOC 3 |
|---|---|---|---|
| Focus | Financial controls | 5 Trust Services Criteria | Same as SOC 2 (sanitized) |
| Audience | Financial auditors | Security professionals | General public |
| Distribution | Restricted | Restricted | **Public** |
| Has Type 1/2? | Yes | Yes | No |
| Compliance use? | Yes | Yes | **No** |

**Trust Services Criteria (SOC 2):** Security*, Availability*, Confidentiality*, Processing Integrity (opt), Privacy (opt). (*Always required.)

Full page: [SOC Reports](standards/soc-reports.md) | [Audit Types](concepts/audit-types.md)

---

### SOC Type 1 vs Type 2

| | Type 1 | Type 2 |
|---|---|---|
| Scope | Point in time | Period of time (~1 year) |
| Examines | Control design | Design + operating effectiveness |
| Operational assurance? | No | **Yes** |
| Gold standard? | No | **Yes — SOC 2 Type 2** |

---

### CVSS Scoring Ranges (v3.1)

| Score | Rating |
|---|---|
| 9.0–10.0 | Critical |
| 7.0–8.9 | High |
| 4.0–6.9 | Medium |
| 0.1–3.9 | Low |
| 0.0 | None |

**CVE** = identity (what the vuln is). **CVSS** = severity (how bad it is). Used together.

Full page: [CVSS](standards/cvss.md) | [Vulnerability Assessment](concepts/vulnerability-assessment.md)

---

### KPI vs KRI

| | KPI | KRI |
|---|---|---|
| Direction | **Backward-looking** | **Forward-looking** |
| Measures | Achievement of past performance targets | Current/emerging risk exposure |
| Examples | Patch rate, training completion, MTTR | Unpatched vuln trend, phishing click rate trend |

Full page: [Security Metrics](concepts/security-metrics.md)

---

### Pen Test Phases (5-phase model, destination-cissp)

1. **Reconnaissance** — passive; OSINT; target cannot detect.
2. **Enumeration** — active; target can detect; IPs, ports, services, hostnames.
3. **Vulnerability Analysis** — THE FORK: vuln assessment stops here, pen test continues.
4. **Exploitation** — pen test only; confirms true positives.
5. **Reporting** — both; prioritized findings, remediation steps, minimize false positives.

Full page: [Penetration Testing](concepts/penetration-testing.md)

---

### Log File Size Management

| Method | How it works | Deletes data? | Better for security investigations? |
|---|---|---|---|
| **Circular overwrite** | Overwrites oldest entries when limit reached | Yes | **No** — breach evidence may be lost |
| **Clipping levels** | Only logs events that exceed a threshold | No | **Yes** — relevant events preserved |

Full page: [Log Management and SIEM](concepts/log-management-siem.md)

---

## Domain 7 — Security Operations

### RTO / RPO / WRT / MTD Relationships

| Metric | Full Name | Measures | Formula Role |
|---|---|---|---|
| **RPO** | Recovery Point Objective | Max data loss (backward-looking) | Drives backup frequency |
| **RTO** | Recovery Time Objective | Max time to reach a defined service level | Drives recovery site type |
| **WRT** | Work Recovery Time | Time to verify system integrity after restore | Part of MTD |
| **MTD** | Maximum Tolerable Downtime | Max total disruption before org fails | **MTD = RTO + WRT** |

**Key rule:** RTO must always be less than MTD. (MTD > RTO)

Lower RPO/RTO = more expensive. Higher RPO/RTO = less expensive.

Full page: [RTO/RPO/MTD](concepts/rto-rpo-mtd.md)

---

### DR Site Types — Comparison Table

| Site | Has Hardware? | Has Data? | Recovery Time | Cost |
|---|---|---|---|---|
| **Cold** | No | No | Weeks | $ |
| **Warm** | Basic equipment only | No | Days | $$ |
| **Hot** | Yes | No (restore needed) | Hours | $$$ |
| **Mobile** | Yes (on wheels) | No (restore needed) | Days–Hours | $$$ |
| **Redundant** | Yes | Yes (live sync) | Instant/Seconds | $$$$ |

Full page: [Disaster Recovery Sites](concepts/disaster-recovery-sites.md)

---

### NIST SP 800-61 Incident Response Phases

1. **Preparation** — build IR capability, train team, create plans
2. **Detection & Analysis** — identify and validate the incident, prioritize
3. **Containment, Eradication & Recovery** — contain, remove root cause, restore
4. **Post-Incident Activity** — lessons learned, evidence retention, reporting

Preparation is Phase 1, not last. Lessons learned feeds back into Preparation.

Full page: [NIST SP 800-61](standards/nist-sp-800-61.md)

---

### Order of Volatility (Most → Least Volatile)

1. CPU registers, cache
2. RAM (routing tables, ARP cache, process list, network connections)
3. Swap space / paging file
4. Hard disk / SSD
5. Remote logging / monitoring data
6. Archived media (tapes, optical)

**Collect most volatile first** — it disappears when the system is powered off.

Full page: [Forensics](concepts/forensics.md)

---

### Backup Type Comparison

| Type | Archive Bit Reset? | Data Backed Up | Backup Speed | Restore Speed | Storage |
|---|---|---|---|---|---|
| **Full** | Yes (→0) | Everything | Slowest | Fast (1 tape) | High |
| **Differential** | No (stays 1) | Since last **full** | Moderate→slow | Moderate (2 tapes) | Moderate |
| **Incremental** | Yes (→0) | Since last **backup** | Fast | Slow (many tapes) | Lowest |
| **Mirror** | N/A | Exact copy | Fastest | Fastest | Highest |

**Key distinction:** Differential does NOT reset archive bit → grows larger each night. Incremental resets → stays small each night but harder to restore.

Full page: [Backup Strategies](concepts/backup-strategies.md)

---

### Failure Modes

| Mode | Security Priority | Example |
|---|---|---|
| **Fail-safe** | People first | Fire door unlocks automatically |
| **Fail-secure (fail-closed)** | Security first | Firewall blocks all traffic on failure |
| **Fail-soft (fail-open)** | Availability first | Firewall allows all traffic on failure |

Fail-safe ≠ fail-secure. Exam will test this distinction.

Full page: [Failure Modes](concepts/failure-modes.md)

---

### DRP Test Types (Least → Most Disruptive)

| Test Type | Paper-Based? | Affects Production? |
|---|---|---|
| **Read-through / Checklist** | Yes | No |
| **Walkthrough** | Yes (tabletop) | No |
| **Simulation** | Yes (facilitated scenario) | No |
| **Parallel** | No | No (backup systems only) |
| **Full-interruption / Full-scale** | No | YES (production affected) |

Full-interruption test requires **management approval** and successful completion of all prior tests.

---

### Malware Types — Key Exam Distinctions

| Type | Key Trait |
|---|---|
| **Virus** | Requires user action to trigger |
| **Worm** | Self-propagates, no user interaction |
| **Logic bomb** | Executes on a condition (time, event) |
| **Rootkit** | Hides presence; collection of attacker tools |
| **Zero-day** | No signatures exist yet |
| **Ransomware** | Encrypts data, demands payment |
| **Polymorphic** | Changes form to evade signature detection |

Full page: [Malware Analysis](concepts/malware-analysis.md)

---

### Five Rules of Evidence

**A**uthentic — **A**ccurate — **C**omplete — **C**onvincing (Reliable) — **A**dmissible

Chain of custody supports all five but does not guarantee admissibility.

Full page: [Chain of Custody](concepts/chain-of-custody.md)

---

### RAID Quick Reference

| Level | Type | Redundancy? | Min Drives | Best For |
|---|---|---|---|---|
| RAID 0 | Striping | No | 2 | Speed |
| RAID 1 | Mirroring | Yes | 2 | Availability |
| RAID 5 | Parity | Yes | 3 | Balance |
| RAID 10 | Stripe + Mirror | Yes | 4 | Speed + Availability |

---

## Domain 8 — Software Development Security

### SDLC Phases + Security Activities

| Phase | Security Activity | Exam Hook |
|---|---|---|
| Planning / Initiation | Security requirements; asset classification | Security in scope from day 1 |
| Requirements | **Risk analysis** | Risk analysis = requirements phase |
| Design | **Threat modeling** | Threat modeling = design phase |
| Development | Secure coding; unit SAST | Coding + early SAST |
| Testing | SAST + DAST + fuzz; pen test; **certification** | Comprehensive: normal + malicious use |
| Deployment | **Accreditation** (management sign-off) | Cert = technical; Accred = management |
| Operations | Monitoring; patching; **change management with security review** | Security must be on CAB |
| Disposal | Secure media sanitization; archival | Covered in Domain 2 also |

Full page: [SDLC](concepts/sdlc.md)

---

### SAST vs. DAST vs. IAST vs. Fuzzing

| Property | SAST | DAST | IAST | Fuzz Testing |
|---|---|---|---|---|
| App running? | No | Yes | Yes | Yes |
| Box type | White box | Black box | Gray box | Black box |
| Source code? | Yes | No | Yes (agent) | No |
| When | Development | QA/staging | QA/staging | QA/staging |
| Finds | Code bugs | Runtime flaws | Both | Edge cases |

Full page: [Security Testing Types](concepts/security-testing-types.md)

---

### OWASP Top 10 (2021) Quick Reference

| Rank | Name | One-Line |
|---|---|---|
| A01 | Broken Access Control | Users exceed their permissions |
| A02 | Cryptographic Failures | Weak or missing crypto; plaintext data |
| A03 | Injection | User input interpreted as commands (SQL, LDAP, OS) |
| A04 | Insecure Design | Missing security in the blueprint/design |
| A05 | Security Misconfiguration | Insecure defaults; unnecessary features |
| A06 | Vulnerable & Outdated Components | Using libraries with known CVEs |
| A07 | Identification & Auth Failures | Weak login; no MFA; bad session management |
| A08 | Software & Data Integrity Failures | Tampered updates; insecure CI/CD; bad deserialization |
| A09 | Logging & Monitoring Failures | Attacks go undetected; no forensic trail |
| A10 | SSRF | Server fetches attacker-controlled URL → reaches internal resources |

Full page: [OWASP Top 10](concepts/owasp-top-10.md)

---

### CMMI Six Maturity Levels

| Level | Name | One-Line |
|---|---|---|
| 0 | Incomplete | Ad hoc; work may not get done |
| 1 | Initial | Reactive; over budget; unpredictable |
| 2 | Managed | Planned; metrics tracked; controlled |
| 3 | Defined | **Proactive**; org-wide standards ← threshold |
| 4 | Quantitatively Managed | Data-driven; measurable; predictable |
| 5 | Optimizing | Continuous improvement; innovation |

> Level 3 = proactive. Below 3 = reactive. Above 3 = data-driven and optimizing.

Full page: [CMMI/SAMM](standards/cmm-samm.md)

---

### SQL Injection Prevention

| Defense | Description |
|---|---|
| **Parameterized queries (primary)** | Query structure fixed; user input always treated as data, never SQL |
| Stored procedures | Equivalent protection if properly implemented |
| Input validation | Validate type, length, content — secondary defense |
| Least privilege DB accounts | App DB user has only SELECT/INSERT — not DROP |
| WAF | Detects common patterns in transit — not a substitute for parameterized queries |

Full page: [Database Security](concepts/database-security.md)

---

### Buffer Overflow Mitigations (in order of importance for exam)

1. **ASLR** — randomizes memory layout → attacker can't predict locations
2. **DEP / NX bit** — marks stack/heap as non-executable → injected code won't run
3. **Stack canaries** — detects overflow before return
4. **Bounds / parameter checking** — prevents overflow from occurring
5. **Safe languages / libraries** — architectural solution

Full page: [Memory Safety](concepts/memory-safety.md)

---

### Code Acquisition Decision Matrix

| Source | White-box possible? | Key Risks | Mitigation |
|---|---|---|---|
| In-house | Yes | Developer errors; insecure practices | SSDLC; code review; testing |
| COTS | No | Vendor gone; missing features; CVEs | Escrow; black-box testing; SLA |
| Open source | Yes | Heartbleed-type supply chain | Code review; SCA scanning; treat like in-house |
| Third-party | Varies | Same as COTS + contractor quality | Contractual security requirements; SDLC applied |
| Cloud/SaaS | No | Shared responsibility; data sovereignty | SOC reports; shared responsibility model |

Full page: [Software Acquisition Security](concepts/software-acquisition-security.md)

---

### API Security Key Points

| Topic | Key Fact |
|---|---|
| REST vs. SOAP | REST = flexible/modern/multi-format. SOAP = rigid/XML-only/strong error handling |
| Primary auth method | OAuth 2.0 (access tokens) for REST APIs |
| Rate limiting | Prevents DoS and data harvesting against APIs |
| API gateway | Centralizes auth, validation, throttling, logging for all APIs |
| OWASP API Top #1 | Broken Object Level Authorization (accessing other users' resources by ID manipulation) |

Full page: [API Security](concepts/api-security.md)

---

### DevSecOps Key Points

- IPT = DevOps (exam language)
- DevSecOps = security integral from day 1; security is planned for, not bolted on
- Traditional pen tests are too slow for DevOps velocity → automate SAST/DAST in CI/CD
- Canary deployment = release to subset of users first
- Smoke testing = quick check of basic functions after a change

Full page: [DevSecOps](concepts/devsecops.md)
