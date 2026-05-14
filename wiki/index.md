---
title: "Wiki Index"
type: overview
domain: cross
tags: [index, navigation]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# CISSP Wiki Index

Catalog of all pages in this wiki. Updated on every ingest and whenever a page is created.

## Reference

- [Overview](overview.md) — CISSP exam map: 8 domains, weights, top-level structure
- [Log](log.md) — chronological event log (ingests, queries, lint passes)
- [Cheatsheet](cheatsheet.md) — living weak-areas doc organized by domain
- [Acronyms](acronyms.md) — alphabetized acronym glossary (all 8 domains)
- [Mnemonics](mnemonics.md) — curated CISSP mnemonics organized by domain

## Domains

- [01 — Security and Risk Management](domains/01-security-and-risk-management.md) — governance, compliance, risk frameworks (16%)
- [02 — Asset Security](domains/02-asset-security.md) — data classification, lifecycle, roles, privacy (10%)
- [03 — Security Architecture and Engineering](domains/03-security-architecture-and-engineering.md) — design principles, models, cryptography, physical security (13%)
- [04 — Communication and Network Security](domains/04-communication-and-network-security.md) — network protocols, OSI model, firewalls, VPNs (13%)
- [05 — Identity and Access Management](domains/05-identity-and-access-management.md) — authentication, access control models, federation, IAM (13%)
- [06 — Security Assessment and Testing](domains/06-security-assessment-and-testing.md) — vulnerability assessment, pen testing, audits, metrics (12%)
- [07 — Security Operations](domains/07-security-operations.md) — incident response, forensics, DR/BCP, patch management (13%)
- [08 — Software Development Security](domains/08-software-development-security.md) — SDLC, secure coding, OWASP, DevSecOps (10%)

## Sources

- [CISSP Exam Outline](sources/cissp-exam-outline.md) — official ISC2 domain/subtopic taxonomy and weights; structural backbone of the wiki
- [CISSP Ultimate Guide RB](sources/cissp-ultimate-guide-rb.md) — comprehensive study guide; not yet ingested
- [Destination CISSP](sources/destination-cissp.md) — concise exam prep guide covering all 8 domains

## Concepts

### A–B
- [AAA — Authentication, Authorization, Accounting](concepts/aaa.md) — identification, authentication, authorization, and accounting framework
- [Access Control Models](concepts/access-control-models.md) — DAC, MAC, RBAC, ABAC, rule-based; how authorization decisions are made
- [Accountability vs. Responsibility](concepts/accountability-vs-responsibility.md) — governance distinction; accountability cannot be delegated
- [API Security](concepts/api-security.md) — REST vs. SOAP, OAuth for APIs, rate limiting, API gateways
- [Assessment vs Audit vs Penetration Test](concepts/assessment-vs-audit.md) — key differences between the three evaluation approaches
- [Asymmetric Cryptography](concepts/asymmetric-crypto.md) — public/private key pairs, RSA, ECC, Diffie-Hellman
- [Audit Types and SOC Reports](concepts/audit-types.md) — internal, external, third-party audits; SAS 70 → SSAE 18 evolution
- [Authentication Factors](concepts/authentication-factors.md) — knowledge, ownership, characteristic; MFA requirements
- [Backup Strategies](concepts/backup-strategies.md) — full, differential, incremental; archive bit behavior; RTO/RPO implications
- [BCP/DRP Operations](concepts/bcp-drp-operations.md) — operational continuity; BIA; BCM; test types
- [Business Continuity Planning and Disaster Recovery](concepts/bcp-drp.md) — BCP vs. DRP; scope differences; key metrics (MTD, RTO, RPO)
- [Bell–LaPadula Model](concepts/bell-lapadula.md) — lattice-based confidentiality; no read up, no write down
- [Biba Model](concepts/biba.md) — lattice-based integrity; no read down, no write up
- [Biometrics](concepts/biometrics.md) — FRR, FAR, CER; physiological vs. behavioral; retina vs. iris

### C–D
- [Chain of Custody](concepts/chain-of-custody.md) — evidence handling; five rules of evidence; MOM framework
- [Change Management (Software Context)](concepts/change-management-software.md) — SCM, versioning, RFC in software development
- [Change Management](concepts/change-management.md) — CAB, RFC, normal/emergency changes; ITIL integration
- [CI/CD Security](concepts/ci-cd-security.md) — pipeline security, SBOM, SCA, infrastructure as code
- [CIA Triad and the Five Pillars of Information Security](concepts/cia-triad.md) — confidentiality, integrity, availability (+ nonrepudiation, authenticity)
- [Clark–Wilson Model](concepts/clark-wilson.md) — rule-based integrity; CDI, UDI, TP, IVP; three goals of integrity
- [Cloud Security Models](concepts/cloud-security-models.md) — SaaS/PaaS/IaaS/FaaS responsibility splits; deployment models
- [Code Review for Security](concepts/code-review-security.md) — manual vs. automated review; security-focused code review practices
- [Common Criteria](concepts/common-criteria.md) — ISO/IEC 15408; EAL 1–7; PP, ST, TOE; evaluation assurance
- [Compliance Requirements](concepts/compliance-requirements.md) — GDPR, HIPAA, SOX, GLBA, FERPA, PIPEDA, CCPA, ITAR, EAR
- [Configuration Management](concepts/configuration-management.md) — CMDB, baseline, change control integration
- [Continuous Monitoring](concepts/continuous-monitoring.md) — ongoing security assessment; SIEM integration; automation
- [Cryptanalysis Attacks](concepts/cryptanalysis-attacks.md) — ciphertext-only, known-plaintext, chosen-plaintext, MITM, birthday attacks
- [Cryptography Fundamentals](concepts/cryptography-fundamentals.md) — five services, block/stream ciphers, cipher modes (ECB, CBC, CTR, GCM)
- [Data Classification](concepts/data-classification.md) — government and private sector classification; labeling vs. marking; CIA dimensions
- [Data Lifecycle](concepts/data-lifecycle.md) — create, store, use, share, archive, destroy
- [Data Loss Prevention (DLP)](concepts/dlp.md) — network, endpoint, cloud DLP; deep content inspection
- [Data Retention](concepts/data-retention.md) — legal holds, retention schedules, disposal requirements
- [Data Roles](concepts/data-roles.md) — owner, custodian, steward, user, controller, processor
- [Data Security Controls](concepts/data-security-controls.md) — encryption, DRM, IRM, media sanitization
- [Data States](concepts/data-states.md) — at rest, in transit, in use; protections for each state
- [Database Security](concepts/database-security.md) — ACID, SQL injection, parameterized queries, polyinstantiation
- [DevSecOps](concepts/devsecops.md) — security in CI/CD; IPT; shift-left; canary deployment
- [Digital Signatures](concepts/digital-signatures.md) — signing with private key; verification with public key; non-repudiation
- [Directory Services](concepts/directory-services.md) — LDAP, X.500, Active Directory; DN, OU
- [Disaster Recovery Sites](concepts/disaster-recovery-sites.md) — cold, warm, hot, mobile, redundant sites; recovery time comparison
- [Due Care and Due Diligence](concepts/due-care-due-diligence.md) — doing the right thing vs. proving you did it

### F–H
- [Facility Security](concepts/facility-security.md) — HVAC, power systems, fire suppression, ASHRAE guidelines
- [Failure Modes](concepts/failure-modes.md) — fail-safe, fail-secure, fail-soft; implications for doors and firewalls
- [Identity Federation](concepts/federation.md) — cross-org trust relationships; IdP, SP/RP; IDaaS
- [Firewalls](concepts/firewalls.md) — packet filter, stateful, circuit proxy, application proxy, NGFW, WAF; architectures
- [Digital Forensics](concepts/forensics.md) — order of volatility; evidence collection; imaging; chain of custody
- [Security Governance Principles](concepts/governance.md) — governance vs. management; board/executive accountability; frameworks
- [Hash Functions](concepts/hash-functions.md) — MD5, SHA-1, SHA-2, SHA-3; HMAC; integrity vs. confidentiality

### I–K
- [ICS/SCADA Security](concepts/ics-scada-security.md) — SCADA, DCS, PLC; air gap; OT hierarchy
- [Identity Lifecycle](concepts/identity-lifecycle.md) — provisioning, review, revocation; privilege creep; orphaned accounts
- [IDS and IPS](concepts/ids-ips.md) — signature vs. anomaly detection; HIDS vs. NIDS; false positive/negative
- [Incident Management](concepts/incident-management.md) — NIST IR phases; CSIRT; triage; reporting
- [Information Lifecycle Management](concepts/information-lifecycle-management.md) — ILM policies; EOL vs. EOS; retention and disposal
- [Information Obfuscation](concepts/information-obfuscation.md) — steganography, watermarking, lexical/control/data obfuscation
- [Intellectual Property](concepts/intellectual-property.md) — copyright, patent, trademark, trade secret; DMCA; ITAR/EAR
- [IoT Security](concepts/iot-security.md) — constrained devices; firmware update challenges; network segmentation
- [IP Addressing](concepts/ip-addressing.md) — IPv4 classes, CIDR, NAT/PAT, private ranges, IPv6 overview
- [Information System Lifecycle (IS Lifecycle)](concepts/is-lifecycle.md) — initiation through disposal; security at each phase
- [Kerberos](concepts/kerberos.md) — KDC, AS, TGS, TGT, service ticket; TOCTOU weakness; SESAME

### K–M
- [Key Management](concepts/key-management.md) — key lifecycle; HSM; TPM; KEK; escrow
- [Log Management and SIEM](concepts/log-management-siem.md) — log aggregation; clipping levels; circular overwrite; NTP
- [Logging and Monitoring](concepts/logging-monitoring.md) — SIEM, SOAR, UBA/UEBA; ISAC threat intelligence

### M–O
- [Malware Analysis](concepts/malware-analysis.md) — virus vs. worm vs. rootkit vs. ransomware; IOC; C2; FIM
- [Memory Safety and Buffer Overflow](concepts/memory-safety.md) — ASLR, DEP, stack canaries, bounds checking; safe languages
- [Mobile Security](concepts/mobile-security.md) — MDM vs. MAM; BYOD, COPE, CYOD; OWASP Mobile Top 10
- [Network Access Control](concepts/network-access-control.md) — NAC; 802.1X; RADIUS vs. TACACS+; PAP, CHAP, EAP
- [Network Attacks](concepts/network-attacks.md) — DoS/DDoS, MITM, ARP poisoning, SYN flood, smurf
- [Network Devices](concepts/network-devices.md) — hub, switch, router, bridge; VLAN; STP; NAT
- [Network Protocols](concepts/network-protocols.md) — TCP/IP stack; DNS, DHCP, SMTP, FTP, SNMP; port numbers
- [OAuth 2.0 and OpenID Connect (OIDC)](concepts/oauth-oidc.md) — authorization vs. authentication; access tokens; ID tokens
- [OSI Model](concepts/osi-model.md) — 7 layers; PDU types; devices and protocols per layer
- [OWASP Top 10 (2021)](concepts/owasp-top-10.md) — A01–A10; SSRF, injection, cryptographic failures, broken access control

### P–R
- [Patch Management](concepts/patch-management.md) — patch lifecycle; WSUS; priority based on CVSS; testing before deployment
- [Penetration Testing](concepts/penetration-testing.md) — 5 phases; black/white/gray box; scope and authorization
- [Personnel Security Policies and Procedures](concepts/personnel-security.md) — hiring, termination, job rotation, mandatory vacation, separation of duties
- [Physical Security](concepts/physical-security.md) — perimeter layers; CPTED; locks as delay controls; CCTV as detective control
- [Public Key Infrastructure (PKI)](concepts/pki.md) — CA hierarchy; CRL vs. OCSP; CSR; root CA offline rule
- [Privacy and PII](concepts/privacy-pii.md) — PII definition; linkage; data minimization
- [Privacy and Data Protection](concepts/privacy.md) — GDPR 7 principles; PbD; DPIA; supervisory authorities
- [Privileged Access Management (PAM)](concepts/privileged-access-management.md) — PAM tools; JIT access; dual accounts; privileged account review
- [Professional Ethics and ISC2 Code of Ethics](concepts/professional-ethics.md) — 4 canons in priority order; conflict resolution
- [Risk Management](concepts/risk-management.md) — qualitative vs. quantitative; ALE formula; risk treatment options; residual risk

### R–S
- [RTO, RPO, WRT, and MTD](concepts/rto-rpo-mtd.md) — definitions; MTD = RTO + WRT formula; relationship to DR site selection
- [SAML — Security Assertion Markup Language](concepts/saml.md) — XML-based federated auth; assertions; roles (IdP, SP)
- [Supply Chain Risk Management (SCRM)](concepts/scrm.md) — vendor risk; SBOM; fourth-party risk; contractual controls
- [SDLC — Software Development Life Cycle](concepts/sdlc.md) — phases; security activity per phase; certification vs. accreditation
- [SDN and NFV](concepts/sdn-nfv.md) — control plane vs. data plane; virtualized network functions; VPC; MPLS
- [Secure Coding Practices](concepts/secure-coding-practices.md) — input validation, least privilege, error handling, coupling/cohesion, TOCTOU, polyinstantiation
- [Secure Design Principles](concepts/secure-design-principles.md) — least privilege, defense in depth, fail-safe, zero trust, PbD, separation of duties
- [Secure Network Design](concepts/secure-network-design.md) — DMZ, screened subnet, defense in depth, microsegmentation
- [Secure Protocols](concepts/secure-protocols.md) — SSH, HTTPS, SFTP, LDAPS, SNMPv3; replacing insecure predecessors
- [Security Awareness, Education, and Training](concepts/security-awareness.md) — awareness vs. training vs. education; phishing simulations; program metrics
- [Security Control Types and Categories](concepts/security-controls-types.md) — preventive, detective, corrective, compensating; administrative/physical/logical
- [Security Metrics: KPIs and KRIs](concepts/security-metrics.md) — SMART criteria; KPI backward vs. KRI forward; MTTD, MTTR
- [Security Models](concepts/security-models.md) — BLP, Biba, Clark–Wilson, Brewer–Nash, Lipner; comparison
- [Security Policies, Standards, Procedures, Baselines, and Guidelines](concepts/security-policies.md) — hierarchy; mandatory vs. optional; audit implications
- [Security Testing Types: SAST, DAST, and Fuzz Testing](concepts/security-testing-types.md) — static vs. dynamic vs. interactive; IAST; RASP
- [Single Sign-On (SSO)](concepts/sso.md) — SSO vs. FIM; ticket/token systems; advantages and risks
- [Social Engineering](concepts/social-engineering.md) — phishing, vishing, tailgating/piggybacking, pretexting, dumpster diving
- [Software Acquisition Security](concepts/software-acquisition-security.md) — COTS, open source, third-party; escrow; SOC reports; SLA
- [Symmetric Cryptography](concepts/symmetric-crypto.md) — DES, 3DES, AES, RC4, ChaCha20; block vs. stream; key distribution problem

### T–Z
- [TCP/IP Model](concepts/tcp-ip-model.md) — 4-layer model; TCP three-way handshake; UDP
- [Threat Modeling Concepts and Methodologies](concepts/threat-modeling.md) — STRIDE, PASTA, attack trees; when to perform; outputs
- [Trusted Platform Module (TPM)](concepts/tpm.md) — ISO/IEC 11889; endorsement key; binding vs. sealing
- [Trusted Computing Base (TCB)](concepts/trusted-computing-base.md) — RMC, security kernel, TCB; CIV properties; ring protection
- [Virtualization Security](concepts/virtualization-security.md) — hypervisor types (Type 1/2); VM escape; snapshot risks; VMM
- [VoIP Security](concepts/voip-security.md) — SIP, SRTP; toll fraud; eavesdropping; PSTN vs. VoIP risks
- [VPN](concepts/vpn.md) — IPsec (AH/ESP); transport vs. tunnel mode; L2TP, PPTP, TLS VPNs; 4 SA rule
- [Vulnerability Assessment](concepts/vulnerability-assessment.md) — credentialed vs. uncredentialed; false positive/negative; CVE + CVSS
- [Vulnerability Management Operations](concepts/vulnerability-management-ops.md) — prioritization; patching cadence; risk acceptance; remediation tracking
- [Wireless Security](concepts/wireless-security.md) — WEP/WPA/WPA2/WPA3 evolution; CCMP vs. GCMP; EAP variants; 802.11i

## Standards

- [IEEE 802.11 — Wireless LAN Standard](standards/802-11.md) — defines wireless communication protocols; 802.11i = WPA2
- [IEEE 802.1X — Port-Based Network Access Control](standards/802-1x.md) — EAP-based authentication to wired and wireless networks
- [AES — Advanced Encryption Standard (FIPS 197)](standards/aes-standard.md) — NIST symmetric standard; 128/192/256-bit keys; 128-bit block
- [CMMI and SAMM — Software Development Maturity Models](standards/cmm-samm.md) — CMMI 0–5 levels; SAMM 3-level 5-function model
- [COBIT — Control Objectives for Information and Related Technologies](standards/cobit.md) — ISACA IT governance framework; audit and gap assessment
- [Common Criteria (ISO/IEC 15408)](standards/common-criteria-standard.md) — international IT security evaluation standard; EAL 1–7
- [CVSS — Common Vulnerability Scoring System](standards/cvss.md) — 0–10 severity scale; base, temporal, environmental metrics
- [FIPS 140-2 / FIPS 140-3 — Security Requirements for Cryptographic Modules](standards/fips-140.md) — four security levels; covers HSM, TPM
- [ISO/IEC 27001 — Information Security Management System](standards/iso-27001.md) — ISMS standard; Plan-Do-Check-Act; certification
- [ISO/IEC 27002 — Code of Practice for Information Security Controls](standards/iso-27002.md) — control catalog supporting ISO 27001
- [ITIL v4 — IT Infrastructure Library](standards/itil-v4.md) — IT service management; change management; CAB, ECAB
- [NIST Cybersecurity Framework (CSF)](standards/nist-csf.md) — Identify, Protect, Detect, Respond, Recover
- [NIST Risk Management Framework (RMF) — SP 800-37](standards/nist-rmf.md) — 7 steps: Prepare through Monitor
- [NIST SP 800-53A — Assessing Security and Privacy Controls](standards/nist-sp-800-53a.md) — assessment procedures for security controls
- [NIST SP 800-61 — Computer Security Incident Handling Guide](standards/nist-sp-800-61.md) — 4 IR phases; preparation through post-incident
- [NIST SP 800-63B — Digital Identity Guidelines: Authentication and Lifecycle Management](standards/nist-sp-800-63b.md) — AAL 1–3; authenticator types
- [NIST SP 800-64 — Security Considerations in the SDLC](standards/nist-sp-800-64.md) — integrating security into software development lifecycle
- [NIST SP 800-88 Rev. 1 — Guidelines for Media Sanitization](standards/nist-sp-800-88.md) — Clear, Purge, Destroy; media-specific guidance
- [OWASP — Open Web Application Security Project](standards/owasp.md) — Top 10, Mobile Top 10, SAMM, MASTG/MASVS
- [PKCS Standards (Public-Key Cryptography Standards)](standards/pkcs-standards.md) — RSA Labs PKI standards; PKCS#10 (CSR), PKCS#12 (PFX)
- [SOC Reports (SOC 1 / SOC 2 / SOC 3)](standards/soc-reports.md) — AICPA attestation; Type 1 vs. Type 2; Trust Services Criteria
- [TCSEC — Trusted Computer System Evaluation Criteria (Orange Book)](standards/tcsec.md) — DoD 1980s standard; D/C/B/A divisions; confidentiality-only
- [TLS — Transport Layer Security](standards/tls-standard.md) — TLS 1.2/1.3; replaces SSL; HTTPS backbone

## Practice Questions

- [Domain 1 — Security and Risk Management](practice-questions/01-security-and-risk-management.md)
- [Domain 2 — Asset Security](practice-questions/02-asset-security.md)
- [Domain 3 — Security Architecture and Engineering](practice-questions/03-security-architecture-and-engineering.md)
- [Domain 4 — Communication and Network Security](practice-questions/04-communication-and-network-security.md)
- [Domain 5 — Identity and Access Management](practice-questions/05-identity-and-access-management.md)
- [Domain 6 — Security Assessment and Testing](practice-questions/06-security-assessment-and-testing.md)
- [Domain 7 — Security Operations](practice-questions/07-security-operations.md)
- [Domain 8 — Software Development Security](practice-questions/08-software-development-security.md)
