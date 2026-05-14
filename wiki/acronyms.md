---
title: Acronym Glossary
type: glossary
domain: cross
tags: [acronyms, glossary]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Acronym Glossary

Alphabetized. Format: **ABC** — Full Form. _(Short gloss. [Concept page](concepts/...).)_

## A

**AAA** — Authentication, Authorization, Accounting. (The three verification/logging components of Access Control Services; often extended to include Identification. [AAA](concepts/aaa.md).)

**AAL** — Authenticator Assurance Level. (NIST SP 800-63B metric for authentication robustness: AAL1 (some assurance, SFA) → AAL2 (high confidence, MFA) → AAL3 (very high, hardware cryptographic authenticator). [NIST SP 800-63B](standards/nist-sp-800-63b.md).)

**ABAC** — Attribute-Based Access Control. (Authorization model that grants access based on user/environment attributes like device type, IP, time of day. Most flexible model. [Access Control Models](concepts/access-control-models.md).)

**ACL** — Access Control List. (A set of rules on a firewall or router that permit or deny traffic based on source/destination IP, port, and protocol. [Firewalls](concepts/firewalls.md).)

**ACID** — Atomicity, Consistency, Isolation, Durability. (Properties required of RDBMS transactions to maintain data integrity. [Database Security](concepts/database-security.md).)

**AD** — Active Directory. (Microsoft's directory service; uses LDAP + Kerberos; required for Kerberos SSO in Windows environments. [Directory Services](concepts/directory-services.md).)

**AEAD** — Authenticated Encryption with Associated Data. (Cipher mode that provides both confidentiality and authentication; e.g., AES-GCM, ChaCha20-Poly1305.)

**AES** — Advanced Encryption Standard. (FIPS 197; 128/192/256-bit keys; 128-bit block; current NIST symmetric standard; used in CCMP for WPA2/WPA3 and in IPsec ESP. [Symmetric Crypto](concepts/symmetric-crypto.md), [AES Standard](standards/aes-standard.md).)

**AH** — Authentication Header. (IPsec sub-protocol providing integrity, data-origin authentication, and replay protection. Does **not** provide confidentiality. [VPN](concepts/vpn.md).)

**AICPA** — American Institute of Certified Public Accountants. (US governing body that issues SSAE 18 and oversees SOC audit standards. [SOC Reports](standards/soc-reports.md).)

**AIW** — Acceptable Interruption Window. (Synonym for MTD/MAD — maximum time a critical process can be disrupted. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**ALE** — Annualized Loss Expectancy. (ALE = SLE × ARO; total expected annual cost of a risk materializing. [Risk Management](concepts/risk-management.md).)

**API** — Application Programming Interface. (Standards-based interface allowing applications to communicate with each other. [API Security](concepts/api-security.md).)

**APPs** — Australian Privacy Principles. (Privacy guidelines under Australia's Privacy Act.)

**ARO** — Annualized Rate of Occurrence. (Expected number of times a risk event occurs per year. [Risk Management](concepts/risk-management.md).)

**ARP** — Address Resolution Protocol. (Maps IP addresses to MAC addresses at the Layer 2/3 boundary. Vulnerable to poisoning attacks. [Network Protocols](concepts/network-protocols.md).)

**ASLR** — Address Space Layout Randomization. (OS mitigation that randomizes memory locations of executables at load time, making buffer overflow exploitation harder. [Memory Safety](concepts/memory-safety.md).)

**AS** — Authentication Service. (Component of Kerberos KDC; verifies user credentials and issues the TGT. [Kerberos](concepts/kerberos.md).)

**ASHRAE** — American Society of Heating, Refrigeration, and Air-Conditioning Engineers. (Sets data center temperature/humidity guidelines (TC 9.9). [Facility Security](concepts/facility-security.md).)

**ATM (1)** — Automated Teller Machine. (Physical device susceptible to skimming attacks. [Physical Security](concepts/physical-security.md).)

**ATM (2)** — Asynchronous Transfer Mode. (Legacy WAN technology; connection-oriented; fixed 53-byte cells; high speed.)

**AV** — Asset Value. (Monetary value of an asset used in quantitative risk calculations. [Risk Management](concepts/risk-management.md).)

## B

**BCM** — Business Continuity Management. (The overarching management function that creates and maintains BCP and DRP plans. [BCP/DRP Operations](concepts/bcp-drp-operations.md).)

**BCP** — Business Continuity Plan / Business Continuity Planning. (Strategic plan to keep critical business processes operating during and after a disaster. [BCP/DRP](concepts/bcp-drp.md), [BCP/DRP Operations](concepts/bcp-drp-operations.md).)

**BGP** — Border Gateway Protocol. (Routing protocol used between autonomous systems on the internet (inter-domain routing).)

**BIA** — Business Impact Analysis. (Process that identifies critical functions, quantifies disruption impact, and determines RPO/RTO/WRT/MTD. [BCP/DRP Operations](concepts/bcp-drp-operations.md).)

**BIS** — Bureau of Industry and Security. (US Commerce Dept. agency responsible for administering EAR.)

**BLP** — Bell–LaPadula Model. (Lattice-based confidentiality security model; "no read up, no write down." [Bell–LaPadula](concepts/bell-lapadula.md).)

**BSIMM** — Building Security In Maturity Model. (Descriptive/benchmarking model of real-world software security practices; contrasts with prescriptive SAMM. [CMMI/SAMM](standards/cmm-samm.md).)

**BYOD** — Bring Your Own Device. (Policy allowing employees to use personal devices for work; requires MDM/MAM for corporate data protection. [Mobile Security](concepts/mobile-security.md).)

## C

**C2** — Command and Control. (Infrastructure used by attackers to remotely control compromised systems in a botnet. [Malware Analysis](concepts/malware-analysis.md).)

**CA** — Certificate Authority. (Issues and manages digital certificates; root of trust in PKI. [PKI](concepts/pki.md).)

**CAB** — Change Advisory Board. (Cross-functional committee that reviews and approves normal and significant changes; security must be represented. [Change Management](concepts/change-management.md), [ITIL v4](standards/itil-v4.md).)

**CaaS** — Containers as a Service. (Cloud service model automating container deployment and management; between PaaS and IaaS in responsibility split.)

**CAPTCHA** — Completely Automated Public Turing test to tell Computers and Humans Apart. (Challenge mechanism preventing automated bot attacks on web forms and login pages.)

**CASB** — Cloud Access Security Broker. (Security policy enforcement point between cloud service consumers and providers; extends DLP-like controls to cloud.)

**CBC** — Cipher Block Chaining. (Block cipher mode that XORs each block with the previous ciphertext; uses IV; good for email. [Cryptography Fundamentals](concepts/cryptography-fundamentals.md).)

**CC** — Common Criteria. (ISO/IEC 15408; international standard for IT security product evaluation; assigns EAL 1–7 ratings. [Common Criteria](concepts/common-criteria.md), [Common Criteria Standard](standards/common-criteria-standard.md).)

**CCB** — Change Control Board. (Alternative name for CAB. [Change Management](concepts/change-management.md).)

**CCMP** — Counter-Mode CBC-MAC Protocol. (AES-based encryption and integrity protocol used in WPA2 and WPA3. [Wireless Security](concepts/wireless-security.md).)

**CCPA** — California Consumer Privacy Act. (California privacy law similar in scope to GDPR.)

**CDI** — Constrained Data Item. (High-integrity data protected by the Clark–Wilson model; access must go through TPs. [Clark–Wilson](concepts/clark-wilson.md).)

**CER** — Crossover Error Rate. (The point where FAR and FRR are equal; lower CER = more accurate biometric system. [Biometrics](concepts/biometrics.md).)

**CFB** — Cipher Feedback. (Block cipher mode; stream-like; uses IV; good for email.)

**CHAP** — Challenge Handshake Authentication Protocol. (Improved authentication over PAP; encrypts password; sends periodic re-challenges to detect session hijacking.)

**CIA** — Confidentiality, Integrity, Availability. (The three core pillars of the security triad. [CIA Triad](concepts/cia-triad.md).)

**CI/CD** — Continuous Integration / Continuous Delivery (or Deployment). (Pipeline automating code commit, test, and release. [CI/CD Security](concepts/ci-cd-security.md).)

**CIDR** — Classless Inter-Domain Routing. (IP addressing scheme that allows flexible subnet sizing using prefix notation (e.g., /24) rather than fixed class boundaries. [IP Addressing](concepts/ip-addressing.md).)

**CMDB** — Configuration Management Database. (Authoritative repository of all hardware and software assets, their configurations, ownership, and relationships. [Configuration Management](concepts/configuration-management.md).)

**CMMI** — Capability Maturity Model Integration. (Six-level (0–5) process maturity model for software development. [CMMI/SAMM](standards/cmm-samm.md).)

**COBIT** — Control Objectives for Information and Related Technologies. (ISACA IT governance framework aligning IT to business objectives; useful for audits and gap assessments. [COBIT](standards/cobit.md).)

**COPE** — Corporate-Owned, Personally Enabled. (Mobile device ownership model; organization owns and controls device but allows personal use. [Mobile Security](concepts/mobile-security.md).)

**COPPA** — Children's Online Privacy Protection Act. (US law protecting privacy of children under 13 online.)

**COPRA** — Consumer Online Privacy Rights Act. (US federal consumer privacy legislation.)

**COSO** — Committee of Sponsoring Organizations. (Enterprise risk management framework providing ERM components, principles, and guidance.)

**COTS** — Commercial Off-the-Shelf. (Pre-built software sold commercially; source code typically not available. [Software Acquisition Security](concepts/software-acquisition-security.md).)

**CRC** — Cyclical Redundancy Check / Cyclic Redundancy Check. (Mathematical checksum used to verify data integrity in storage and transmission; simple and susceptible to collisions. [Backup Strategies](concepts/backup-strategies.md).)

**CRL** — Certificate Revocation List. (Full list of revoked certificate serial numbers downloaded from CA; older revocation method. [PKI](concepts/pki.md).)

**CSP** — Cloud Service Provider. (Third-party provider of cloud computing services; customer retains accountability for data.)

**CSRF** — Cross-Site Request Forgery. (Web attack that exploits persistent browser cookies to execute unauthorized server-side actions; target is the web server.)

**CSIRT** — Computer Security Incident Response Team. (Dedicated team responsible for handling security incidents. [Incident Management](concepts/incident-management.md).)

**CSR** — Certificate Signing Request. (PKCS#10-formatted request submitted to a CA during certificate enrollment. [PKI](concepts/pki.md), [PKCS Standards](standards/pkcs-standards.md).)

**CTR** — Counter mode. (Block cipher mode using an incrementing counter as IV; fastest and most commonly used. [Cryptography Fundamentals](concepts/cryptography-fundamentals.md).)

**CVE** — Common Vulnerabilities and Exposures. (Dictionary/clearinghouse assigning unique IDs to publicly known vulnerabilities. Used with CVSS for vulnerability prioritization. [Vulnerability Assessment](concepts/vulnerability-assessment.md).)

**CVSS** — Common Vulnerability Scoring System. (Framework scoring vulnerability severity 0.0–10.0; includes base, temporal, and environmental components. Critical ≥ 9.0, High 7.0–8.9, Medium 4.0–6.9, Low 0.1–3.9. [CVSS](standards/cvss.md).)

**CYOD** — Choose Your Own Device. (Mobile ownership model; employee selects device from approved list; organization owns and controls it.)

**CCTV** — Closed-Circuit Television. (Primarily a detective physical security control. [Physical Security](concepts/physical-security.md).)

**CPTED** — Crime Prevention Through Environmental Design. (Design methodology that uses the built environment to deter crime. [Physical Security](concepts/physical-security.md).)

## D

**DAC** — Discretionary Access Control. (Access control where the asset owner decides who may access the asset. [Access Control Models](concepts/access-control-models.md).)

**DAST** — Dynamic Application Security Testing. (Black-box testing of a running application; focuses on runtime behavior, not source code. [Security Testing Types](concepts/security-testing-types.md).)

**DDTC** — Directorate of Defense Trade Controls. (US State Dept. office administering ITAR.)

**DDoS** — Distributed Denial of Service. (DoS attack using multiple compromised hosts acting in unison. [Network Attacks](concepts/network-attacks.md).)

**DEA** — Data Encryption Algorithm. (The underlying algorithm used by DES.)

**DEP** — Data Execution Prevention. (OS feature marking memory regions as non-executable; mitigates buffer overflow code injection. [Memory Safety](concepts/memory-safety.md).)

**DES** — Data Encryption Standard. (56-bit key, 16 rounds, 64-bit block; deprecated. [Symmetric Crypto](concepts/symmetric-crypto.md).)

**DevSecOps** — Development, Security, and Operations. (Integration of security into the DevOps process from inception. [DevSecOps](concepts/devsecops.md).)

**DH / DHE** — Diffie–Hellman (Ephemeral). (Key exchange protocol using discrete logarithms; ephemeral variant provides forward secrecy. [Asymmetric Crypto](concepts/asymmetric-crypto.md).)

**DHCP** — Dynamic Host Configuration Protocol. (Automatically assigns IP addresses to devices on a network. Rogue DHCP is a MITM enabler. [Network Protocols](concepts/network-protocols.md).)

**DCS** — Distributed Control System. (ICS type; process control for large industrial facilities; typically locally controlled, unlike SCADA. [ICS/SCADA Security](concepts/ics-scada-security.md).)

**DFS** — Distributed File System. (Files hosted across multiple hosts and shared across a network; presents as a single central location to users.)

**DLP** — Data Loss Prevention. (System for identifying, monitoring, and protecting data in use, in transit, and at rest via deep content inspection. [DLP](concepts/dlp.md).)

**DMZ** — Demilitarized Zone. (A screened subnet between the internet and internal network where public-facing services are hosted. [Secure Network Design](concepts/secure-network-design.md).)

**DMCA** — Digital Millennium Copyright Act. (US law implementing WIPO copyright treaties; governs digital copyright; provides legal recourse for DRM violations.)

**DN** — Distinguished Name. (Unique identifier for an entry in an LDAP/X.500 directory, e.g., `CN=Alice,OU=Finance,DC=example,DC=com`. [Directory Services](concepts/directory-services.md).)

**DNSSEC** — Domain Name System Security Extensions. (Cryptographically signs DNS records to prevent cache poisoning.)

**DNS** — Domain Name System. (Resolves hostnames to IP addresses. Port 53 TCP/UDP. Vulnerable to cache poisoning; DNSSEC is the countermeasure. [Network Protocols](concepts/network-protocols.md).)

**DoS** — Denial of Service. (Attack that impedes or denies functionality of a system using one machine. [Network Attacks](concepts/network-attacks.md).)

**DPIA** — Data Protection Impact Assessment. (GDPR-mandated privacy risk assessment for high-risk data processing; see PIA. [Privacy](concepts/privacy.md).)

**DRM** — Digital Rights Management. (IT system for distributing and controlling intellectual property and its rights; protects assets and rights of owners. [Data Security Controls](concepts/data-security-controls.md).)

**DRP (1)** — Disaster Recovery Plan / Disaster Recovery Planning. (IT-focused plan to restore systems and infrastructure after a disruption. [BCP/DRP](concepts/bcp-drp.md), [BCP/DRP Operations](concepts/bcp-drp-operations.md).)

**DRP (2)** — Digital Rights Protection. (Subset of DRM applied operationally; limits specific user actions on data; often paired with DLP for in-use data protection.)

**DSA** — Digital Signature Algorithm. (FIPS-standardized signing algorithm; uses discrete logarithm; signing only. [Digital Signatures](concepts/digital-signatures.md).)

## E

**EAL** — Evaluation Assurance Level. (Common Criteria rating 1–7 assigned after product evaluation; EAL 4 is the commercial sweet spot. [Common Criteria](concepts/common-criteria.md).)

**EAP** — Extensible Authentication Protocol. (Flexible authentication framework supporting smart cards, certificates, and other methods. Variants: EAP-TLS, PEAP, EAP-TTLS, LEAP, EAP-MD5. [Network Access Control](concepts/network-access-control.md).)

**EAR** — Export Administration Regulations. (US Commerce Dept. rules controlling export of commercial-use items with military potential.)

**ECAB** — Emergency Change Advisory Board. (Subset of CAB authorized to approve emergency changes outside the normal meeting schedule. [ITIL v4](standards/itil-v4.md).)

**ECB** — Electronic Codebook. (Block cipher mode without IV; fastest but insecure for repeating data. [Cryptography Fundamentals](concepts/cryptography-fundamentals.md).)

**ECC** — Elliptic Curve Cryptography. (Asymmetric algorithm using discrete logarithm; smaller keys than RSA for equivalent strength. [Asymmetric Crypto](concepts/asymmetric-crypto.md).)

**ECDSA** — Elliptic Curve Digital Signature Algorithm. (ECC variant of DSA; used in TLS certificates and Bitcoin. [Digital Signatures](concepts/digital-signatures.md).)

**EF** — Exposure Factor. (Percentage of asset value expected to be lost in a single risk event; 0–100%. [Risk Management](concepts/risk-management.md).)

**EK** — Endorsement Key. (Unique RSA key burned into each TPM at manufacture; never leaves the chip; used for authentication. [TPM](concepts/tpm.md).)

**EOL** — End of Life. (Point at which a manufacturer no longer sells or develops a product. Asset inventory must track EOL status. [Information Lifecycle Management](concepts/information-lifecycle-management.md).)

**EOS** — End of Support. (Point at which a manufacturer no longer provides security patches or technical support; represents higher security risk than EOL. [Information Lifecycle Management](concepts/information-lifecycle-management.md).)

**ESP** — Encapsulating Security Payload. (IPsec sub-protocol providing integrity, authentication, replay protection, and **payload encryption (confidentiality)**. [VPN](concepts/vpn.md).)

## F

**FaaS** — Function as a Service. (Cloud model; serverless computing; pay per function invocation; no server provisioning needed. Example: AWS Lambda.)

**FAR** — False Acceptance Rate. (Biometric Type 2 error rate: probability an unauthorized user is falsely accepted. More dangerous than FRR. [Biometrics](concepts/biometrics.md).)

**FCoE** — Fibre Channel over Ethernet. (Encapsulates Fibre Channel storage traffic over Ethernet networks; a converged protocol.)

**FedRAMP** — Federal Risk and Authorization Management Program. (US federal cloud security authorization framework; required for cloud services holding federal data.)

**FERPA** — Family Educational Rights and Privacy Act. (US law protecting privacy of student educational records.)

**FIM (1)** — Federated Identity Management. (Protocols and standards enabling identity portability and trust relationships across organizations. [Federation](concepts/federation.md).)

**FIM (2)** — File Integrity Monitoring. (Tool that hashes key files and alerts when hashes change unexpectedly; detects unauthorized modifications. [Malware Analysis](concepts/malware-analysis.md).)

**FIPS** — Federal Information Processing Standard. (NIST standards for US federal systems; FIPS 140 covers cryptographic modules, FIPS 197 is AES. [FIPS 140](standards/fips-140.md), [AES Standard](standards/aes-standard.md).)

**FISMA** — Federal Information Security Modernization Act / Federal Information Security Management Act. (US law requiring federal agencies to secure information systems and develop agency-wide security programs.)

**FRR** — False Rejection Rate. (Biometric Type 1 error rate: probability a legitimate user is falsely rejected. Less dangerous than FAR. [Biometrics](concepts/biometrics.md).)

**FTP** — File Transfer Protocol. (Ports 20/21 TCP. Plaintext — insecure. Replace with SFTP. [Secure Protocols](concepts/secure-protocols.md).)

## G

**GCMP** — Galois Counter Mode Protocol. (Encryption/integrity protocol used in WPA3. [Wireless Security](concepts/wireless-security.md).)

**GCM** — Galois/Counter Mode. (AEAD block cipher mode; authenticated encryption; used in TLS 1.3.)

**GDPR** — General Data Protection Regulation. (EU privacy law; 7 principles; 72-hour breach reporting; considered global benchmark. [Privacy](concepts/privacy.md).)

**GLBA** — Gramm-Leach-Bliley Act. (US law governing privacy of financial information held by financial institutions.)

**GPS** — Global Positioning System. (Used for high-value mobile asset tracking; highest cost of common labeling approaches. [Data Classification](concepts/data-classification.md).)

**GRE** — Generic Routing Encapsulation. (Tunneling protocol that encapsulates various protocols. Provides **no encryption** — not a VPN by itself. [VPN](concepts/vpn.md).)

## H

**HIDS** — Host-based Intrusion Detection System. (IDS agent running on a specific host/server. [IDS and IPS](concepts/ids-ips.md).)

**HIPAA** — Health Insurance Portability and Accountability Act. (US law protecting personal health information (PHI) in healthcare.)

**HMAC** — Hash-based Message Authentication Code. (Hash combined with a symmetric key; provides integrity + authentication. [Hash Functions](concepts/hash-functions.md).)

**HSM** — Hardware Security Module. (Network-connected device that stores and manages encryption keys for an organization. [Key Management](concepts/key-management.md), [FIPS 140](standards/fips-140.md).)

**HTTP** — Hypertext Transfer Protocol. (Port 80 TCP. Plaintext web communication.)

**HTTPS** — HTTP Secure. (Port 443 TCP. HTTP over TLS. [Secure Protocols](concepts/secure-protocols.md).)

**HVAC** — Heating, Ventilation, and Air Conditioning. (System providing temperature, humidity, and air quality control. [Facility Security](concepts/facility-security.md).)

## I

**IaaS** — Infrastructure as a Service. (Cloud model; customer manages OS and above; provider manages physical/hypervisor. Virtual data center. [Cloud Security Models](concepts/cloud-security-models.md).)

**IAST** — Interactive Application Security Testing. (Agent-based; instruments the running application from inside; combines SAST and DAST advantages. [Security Testing Types](concepts/security-testing-types.md).)

**IAM** — Identity and Access Management. (Processes and technologies for managing user identities and their access to resources.)

**IDaaS** — Identity as a Service. (Cloud-based IAM solutions including SSO, MFA, directory services, and federation. [Federation](concepts/federation.md).)

**IdP** — Identity Provider. (In federation, the entity that owns the user identity and performs authentication. [Federation](concepts/federation.md).)

**ICS** — Industrial Control System. (General term for hardware/software controlling industrial processes and critical infrastructure; subset of OT. [ICS/SCADA Security](concepts/ics-scada-security.md).)

**IDEA** — International Data Encryption Algorithm. (128-bit symmetric algorithm; first to use 128-bit key; used in early PGP.)

**IDS** — Intrusion Detection System. (Passive monitoring device connected to a mirror/span port; detects, logs, and alerts on suspicious activity. [IDS and IPS](concepts/ids-ips.md).)

**IKE** — Internet Key Exchange. (Key management protocol for IPsec; based on Diffie-Hellman; generates the symmetric session key at both VPN endpoints. [VPN](concepts/vpn.md).)

**ILM** — Information Lifecycle Management. (Policies and processes for managing information from creation through destruction. [Information Lifecycle Management](concepts/information-lifecycle-management.md).)

**IMAP** — Internet Message Access Protocol. (Email retrieval. Port 143 TCP.)

**IOC** — Indicator of Compromise. (Artifact (IP, hash, registry key, domain) indicating a system has been or is being compromised. [Malware Analysis](concepts/malware-analysis.md).)

**IoT** — Internet of Things. (Network-connected everyday devices with minimal built-in security. [IoT Security](concepts/iot-security.md).)

**IPS** — Intrusion Prevention System. (Inline device that detects, alerts, and actively blocks malicious traffic. [IDS and IPS](concepts/ids-ips.md).)

**IPsec** — Internet Protocol Security. (Layer 3 suite of protocols (AH + ESP) for VPNs; natively embedded in IPv6; supports transport and tunnel modes. [VPN](concepts/vpn.md).)

**IPT** — Integrated Product Team. (Multidisciplinary team of development, operations, and QA — essentially DevOps in CISSP exam language. [DevSecOps](concepts/devsecops.md).)

**IRM** — Information Rights Management. (Subset of DRM applied to internal sensitive documents within an organization. [Data Security Controls](concepts/data-security-controls.md).)

**ISAC** — Information Sharing and Analysis Center. (Sector-specific organizations that share threat intelligence among members. [Logging and Monitoring](concepts/logging-monitoring.md).)

**ISAE 3402** — International Standard on Assurance Engagements No. 3402. (International equivalent of SSAE 18; governs third-party audits outside the US. [SOC Reports](standards/soc-reports.md).)

**iSCSI** — Internet Small Computer Systems Interface. (SCSI storage commands over IP networks; a converged protocol.)

**ITAR** — International Traffic in Arms Regulations. (US State Dept. law controlling export of defense items on USML.)

**ITIL** — Information Technology Infrastructure Library. (Best-practice framework for IT service management, covering change, incident, and problem management. [ITIL v4](standards/itil-v4.md).)

**ITSM** — IT Service Management. (Discipline for managing IT services; ITIL is the primary ITSM framework. [ITIL v4](standards/itil-v4.md).)

**IV** — Initialization Vector. (Random value combined with key to prevent ciphertext patterns; used in CBC, CFB, OFB, CTR, GCM modes. [Cryptography Fundamentals](concepts/cryptography-fundamentals.md).)

**IVP** — Integrity Verification Procedure. (Clark–Wilson mechanism that verifies CDIs conform to integrity constraints. [Clark–Wilson](concepts/clark-wilson.md).)

## J

**JIT** — Just-in-Time. (Access model where elevated privileges are granted temporarily for a specific task, then revoked. [Privileged Access Management](concepts/privileged-access-management.md).)

## K

**KDC** — Key Distribution Center. (Kerberos component containing both the Authentication Service (AS) and Ticket Granting Service (TGS). [Kerberos](concepts/kerberos.md).)

**KEK** — Key Encrypting Key. (A key used to wrap (encrypt) other keys for distribution. [Key Management](concepts/key-management.md).)

**KPI** — Key Performance Indicator. (Backward-looking metric measuring whether performance targets were achieved. [Security Metrics](concepts/security-metrics.md).)

**KRI** — Key Risk Indicator. (Forward-looking metric indicating current risk exposure and predicting future risk conditions. [Security Metrics](concepts/security-metrics.md).)

## L

**L2TP** — Layer 2 Tunneling Protocol. (Layer 2 tunneling protocol with no native encryption — pairs with IPsec for VPNs. [VPN](concepts/vpn.md).)

**LDAP** — Lightweight Directory Access Protocol. (Open protocol for authenticating to and querying directory services like Active Directory. Port 389 TCP; use LDAPS (port 636) for security. [Directory Services](concepts/directory-services.md).)

## M

**MAC (1)** — Mandatory Access Control. (Access control where the system decides based on object classification labels and subject clearance levels. Used in government/military. [Access Control Models](concepts/access-control-models.md).)

**MAC (2)** — Message Authentication Code. (Generic term for a keyed hash; HMAC is a specific type. [Hash Functions](concepts/hash-functions.md).)

**MAD** — Maximum Allowable Downtime. (Synonym for MTD. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**MAM** — Mobile Application Management. (Software managing and securing specific applications on mobile devices; preferred for BYOD. [Mobile Security](concepts/mobile-security.md).)

**MASTG** — Mobile Application Security Testing Guide. (OWASP guide for mobile app security testing and reverse engineering.)

**MASVS** — Mobile Application Security Verification Standard. (OWASP standard for secure mobile application development and testing.)

**MD5** — Message Digest 5. (128-bit hash algorithm; broken — do not use for security. [Hash Functions](concepts/hash-functions.md).)

**MDM** — Mobile Device Management. (Software managing and securing entire mobile devices; enables remote wipe, policy enforcement, encryption. [Mobile Security](concepts/mobile-security.md).)

**MFA** — Multi-Factor Authentication. (Authentication using two or more *different* factor families: knowledge, ownership, characteristic. [Authentication Factors](concepts/authentication-factors.md).)

**MIC** — Message Integrity Control/Check. (Mechanisms to verify that a message has not been altered in transit. [Hash Functions](concepts/hash-functions.md).)

**MIME** — Multipurpose Internet Mail Extensions. (Email format extension enabling multimedia; no security features — contrast with S/MIME.)

**MITM** — Man-in-the-Middle. (Attack where adversary inserts themselves between two communicating parties. [Network Attacks](concepts/network-attacks.md), [Cryptanalysis Attacks](concepts/cryptanalysis-attacks.md).)

**MOM** — Motive, Opportunity, Means. (Framework for focusing a forensic investigation. [Chain of Custody](concepts/chain-of-custody.md).)

**MPLS** — Multiprotocol Label Switching. (Current WAN technology; uses labels to guarantee traffic isolation within provider network; still requires data encryption for confidentiality. [SDN and NFV](concepts/sdn-nfv.md).)

**MSP** — Managed Service Provider. (External IT provider; must be assessed for security using SOC reports, site visits, reference checks. [Software Acquisition Security](concepts/software-acquisition-security.md).)

**MTBF** — Mean Time Between Failures. (Average time between failures of a component or system; higher = more reliable. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**MTD** — Maximum Tolerable Downtime. (Maximum time a critical process can be disrupted before causing irrecoverable harm; MTD = RTO + WRT; RTO must be ≤ MTD. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**MTTD** — Mean Time to Detect. (Average elapsed time from incident occurrence to detection. Lower is better. [Security Metrics](concepts/security-metrics.md).)

**MTTR** — Mean Time to Repair / Recovery / Respond. (Average time to repair a failed component or respond to an incident; shorter = faster recovery. [Security Metrics](concepts/security-metrics.md), [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

## N

**NAC** — Network Access Control. (Technology restricting network access based on endpoint health and user identity. [Network Access Control](concepts/network-access-control.md).)

**NAT** — Network Address Translation. (Translates private IP addresses to public ones for internet communication; hides internal IP topology. [IP Addressing](concepts/ip-addressing.md).)

**NCA** — Noncompete Agreement. (Contract preventing employees from joining competitors for a defined period after leaving.)

**NDA** — Nondisclosure Agreement. (Contract obligating parties not to disclose specified confidential information.)

**NFV** — Network Functions Virtualization. (Virtualizes network functions (firewalls, routers, IDS/IPS) as software. [SDN and NFV](concepts/sdn-nfv.md).)

**NGFW** — Next-Generation Firewall. (Combines stateful inspection, application awareness, IPS, and threat intelligence. [Firewalls](concepts/firewalls.md).)

**NIDS** — Network-based Intrusion Detection System. (IDS monitoring a network segment. [IDS and IPS](concepts/ids-ips.md).)

**NTP** — Network Time Protocol. (Protocol synchronizing system clocks across a network. Critical for log timestamp consistency and incident correlation. [Logging and Monitoring](concepts/logging-monitoring.md).)

**NVD** — National Vulnerability Database. (NIST-maintained database distributing CVE records and associated CVSS scores. [CVSS](standards/cvss.md).)

## O

**OCSP** — Online Certificate Status Protocol. (Query CA for status of a specific certificate (yes/no revocation response); better than CRL. [PKI](concepts/pki.md).)

**OECD** — Organization for Economic Cooperation and Development. (International organization providing voluntary privacy guidelines used as a global baseline.)

**OFB** — Output Feedback. (Block cipher mode; uses IV; stream-like; good for email.)

**OIDC** — OpenID Connect. (Authentication identity layer built on top of OAuth 2.0; provides user identity via ID token. [OAuth 2.0 and OIDC](concepts/oauth-oidc.md).)

**OSPF** — Open Shortest Path First. (Interior gateway routing protocol with built-in security features.)

**OT** — Operational Technology. (Broad term for hardware/software monitoring and controlling physical processes and industrial systems; ICS is a subset. [ICS/SCADA Security](concepts/ics-scada-security.md).)

**OU** — Organizational Unit. (Container within LDAP/AD directory hierarchy. [Directory Services](concepts/directory-services.md).)

**OWASP** — Open Web Application Security Project. (Community-driven organization producing security guidance; publishes Top 10, Mobile Top 10, SAMM. [OWASP](standards/owasp.md).)

## P

**PaaS** — Platform as a Service. (Cloud model; provider manages infrastructure/runtime; customer manages applications and data. [Cloud Security Models](concepts/cloud-security-models.md).)

**PAM** — Privileged Access Management. (Policies and tools for controlling, monitoring, and protecting privileged accounts. [Privileged Access Management](concepts/privileged-access-management.md).)

**PAP** — Password Authentication Protocol. (Weakest PPP authentication; passwords sent in plaintext; static password. [Network Access Control](concepts/network-access-control.md).)

**PAT** — Port Address Translation. (Extension of NAT that uses unique port numbers to map multiple internal hosts to one public IP. [IP Addressing](concepts/ip-addressing.md).)

**PbD** — Privacy by Design. (Design philosophy embedding privacy from inception, not as an afterthought; seven foundational principles. [Privacy](concepts/privacy.md).)

**PBX** — Private Branch Exchange. (Private internal telephone network for an organization. [VoIP Security](concepts/voip-security.md).)

**PCI DSS** — Payment Card Industry Data Security Standard. (Standard for organizations handling credit/debit card transactions.)

**PDP** — Policy Decision Point. (Centralized component that evaluates authorization requests against policy rules and returns allow/deny decisions. [Access Control Models](concepts/access-control-models.md).)

**PEAP** — Protected Extensible Authentication Protocol. (Encapsulates EAP within a TLS tunnel; supported by Cisco, RSA, Microsoft. [Wireless Security](concepts/wireless-security.md).)

**PEP** — Policy Enforcement Point. (Application component that receives access requests, forwards to PDP, and enforces the PDP's decisions. [Access Control Models](concepts/access-control-models.md).)

**PGP** — Pretty Good Privacy. (Popular cryptosystem supporting all five cryptographic services; used IDEA initially; now uses various algorithms.)

**PHI** — Protected Health Information. (Personal health data protected under HIPAA.)

**PIA** — Privacy Impact Assessment. (Systematic process to evaluate whether personal data is protected appropriately and minimize risks. [Privacy](concepts/privacy.md).)

**PII** — Personally Identifiable Information. (Information that can identify an individual, alone or in combination. Core DLP and privacy target. [Privacy and PII](concepts/privacy-pii.md).)

**PIR** — Passive Infrared (device). (Motion sensor detecting infrared light (heat) differences; must recalibrate to ambient temperature. [Physical Security](concepts/physical-security.md).)

**PIPEDA** — Personal Information Protection and Electronic Documents Act. (Canadian privacy law governing personal information in commercial activity.)

**PKI** — Public Key Infrastructure. (Complete system for distributing and verifying public keys via certificate authorities. [PKI](concepts/pki.md).)

**PKCS** — Public-Key Cryptography Standards. (RSA Laboratories-originated family of standards for PKI operations. [PKCS Standards](standards/pkcs-standards.md).)

**PLC** — Programmable Logic Controller. (ICS type; industrial computer for manufacturing process control; high reliability. [ICS/SCADA Security](concepts/ics-scada-security.md).)

**POA&M** — Plan of Actions and Milestones. (Document tracking remaining weaknesses/deficiencies after system authorization in NIST RMF Step 6.)

**POP3** — Post Office Protocol v3. (Email retrieval. Port 110 TCP.)

**PP** — Protection Profile. (Common Criteria component; defines security requirements for a category of products, e.g., all firewalls. [Common Criteria](concepts/common-criteria.md).)

**PPTP** — Point-to-Point Tunneling Protocol. (Layer 2 tunneling protocol. **Deprecated** — avoid. [VPN](concepts/vpn.md).)

**PSK (1)** — Pre-Shared Key. (Symmetric key established between parties before communication begins; out-of-band distribution. [Key Management](concepts/key-management.md).)

**PSK (2)** — Pre-Shared Key (Wi-Fi). (Shared secret used for WPA/WPA2 Personal authentication (home/small office). [Wireless Security](concepts/wireless-security.md).)

**PSTN** — Public Switched Telephone Network. (Traditional copper-wire telephone network. [VoIP Security](concepts/voip-security.md).)

## Q

*(No entries at this time.)*

## R

**RA** — Registration Authority. (Performs identity proofing on behalf of a CA. [PKI](concepts/pki.md).)

**RAID** — Redundant Array of Independent Disks. (Multiple drives used together for speed (RAID 0), availability (RAID 1), or both (RAID 5, RAID 10). [Disaster Recovery Sites](concepts/disaster-recovery-sites.md).)

**RAM** — Random-Access Memory. (Primary/volatile storage; fast but temporary; lost when power is cut.)

**RARP** — Reverse Address Resolution Protocol. (Maps MAC addresses to IP addresses (opposite of ARP). [Network Protocols](concepts/network-protocols.md).)

**RASP** — Runtime Application Self-Protection. (Security agent embedded within the application runtime that detects and blocks attacks in real time. [Security Testing Types](concepts/security-testing-types.md).)

**RBAC** — Role-Based Access Control. (Authorization model granting access based on job function/role. Best practice for most organizations; reduces administrative overhead. [Access Control Models](concepts/access-control-models.md).)

**RC4** — Rivest Cipher 4. (Stream cipher; deprecated due to key scheduling weaknesses; WEP's broken RC4 use is an implementation attack.)

**RDBMS** — Relational Database Management System. (Database system storing data in related tables linked by primary and foreign keys. [Database Security](concepts/database-security.md).)

**RDP** — Remote Desktop Protocol. (Port 3389 TCP. Windows remote desktop. [Network Protocols](concepts/network-protocols.md).)

**RFC (1)** — Request for Comments. (IETF standard document series; e.g., RFC 1918 defines private IPv4 ranges; RFC 3711 defines SRTP.)

**RFC (2)** — Request for Change. (Formal submission initiating the change management process. [Change Management](concepts/change-management.md).)

**RFID (1)** — Radio Frequency Identification. (Labeling technology for asset tracking; moderate cost; useful for warehouse inventory without individual item scanning. [Data Classification](concepts/data-classification.md).)

**RFID (2)** — Radio Frequency Identification (physical access). (Technology used in proximity cards for physical access control. [Physical Security](concepts/physical-security.md).)

**RIP** — Routing Information Protocol. (Older, simpler routing protocol; less secure than OSPF.)

**RMC** — Reference Monitor Concept. (Concept: subjects access objects through rule-based mediation that is logged. Implemented by the security kernel. [Trusted Computing Base](concepts/trusted-computing-base.md).)

**RP** — Relying Party. (In federation, the service provider that relies on the IdP's authentication assertion. [Federation](concepts/federation.md).)

**RPO** — Recovery Point Objective. (Maximum acceptable data loss expressed as a duration of time. Drives backup frequency and replication strategy. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**RSA** — Rivest, Shamir, Adleman. (Most common asymmetric algorithm; uses factoring of large primes. [Asymmetric Crypto](concepts/asymmetric-crypto.md).)

**RSN** — Robust Security Network. (Security architecture defined in 802.11i / WPA2.)

**RTO** — Recovery Time Objective. (Maximum tolerable time to restore systems to a defined service level after a disaster. Drives recovery site type selection. RTO must be ≤ MTD. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**RUM** — Real User Monitoring. (Passive operational testing technique monitoring actual user interactions with a live application in real time. [Log Management and SIEM](concepts/log-management-siem.md).)

## S

**SA (1)** — Supervisory Authority. (Independent data protection authority in each EU member state; investigates GDPR complaints. [Privacy](concepts/privacy.md).)

**SA (2)** — Security Association. (One-way set of IPsec parameters. Full IPsec connection with AH + ESP requires **4 SAs** (2 directions × 2 components). [VPN](concepts/vpn.md).)

**SaaS** — Software as a Service. (Cloud model; provider manages everything including the application; customer manages users and access only. [Cloud Security Models](concepts/cloud-security-models.md).)

**SABSA** — Sherwood Applied Business Security Architecture. (Enterprise security architecture framework, 1995; risk-driven; open source.)

**SAMM** — Software Assurance Maturity Model. (OWASP's three-level maturity model for software security; five business functions. [CMMI/SAMM](standards/cmm-samm.md).)

**SAML** — Security Assertion Markup Language. (XML-based OASIS standard for federated authentication/authorization via security tokens and assertions. [SAML](concepts/saml.md).)

**SASE** — Secure Access Service Edge. (Converged network security + WAN in a cloud-based service; data and services delivered close to end users with robust security.)

**SAST** — Static Application Security Testing. (White-box analysis of source code while the application is not running; finds code-level vulnerabilities before deployment. [Security Testing Types](concepts/security-testing-types.md).)

**SBOM** — Software Bill of Materials. (Machine-readable inventory of all components and dependencies in a software product; enables rapid CVE impact assessment. [CI/CD Security](concepts/ci-cd-security.md), [Software Acquisition Security](concepts/software-acquisition-security.md).)

**SBU** — Sensitive But Unclassified. (Government classification level between Confidential and Unclassified. [Data Classification](concepts/data-classification.md).)

**SCA** — Software Composition Analysis. (Tool for scanning open-source dependencies for known vulnerabilities (CVEs). [CI/CD Security](concepts/ci-cd-security.md).)

**SCADA** — Supervisory Control and Data Acquisition. (ICS type; system architecture for managing industrial, infrastructure, and facility processes; includes remote management. [ICS/SCADA Security](concepts/ics-scada-security.md).)

**SCM** — Software Configuration Management. (Managing changes to code, configuration, and documentation throughout the SDLC; includes versioning, baseline management, build management. [Change Management (Software)](concepts/change-management-software.md).)

**SCRM** — Supply Chain Risk Management. (Application of risk management principles to all vendors, suppliers, and service providers. [SCRM](concepts/scrm.md).)

**SDK** — Software Development Kit. (Collection of development tools, APIs, libraries, and documentation for building applications for a specific platform or OS.)

**SDLC** — Software Development Life Cycle. (Structured process for planning, creating, testing, and deploying software; security must be embedded at every phase. [SDLC](concepts/sdlc.md).)

**SDN** — Software-Defined Networking. (Network architecture that decouples the control plane from the data plane for software-driven management. [SDN and NFV](concepts/sdn-nfv.md).)

**SESAME** — Secure European System for Applications in a Multi-Vendor Environment. (Improved Kerberos variant supporting asymmetric crypto and multiple tickets; rarely adopted. [Kerberos](concepts/kerberos.md).)

**SFTP** — Secure File Transfer Protocol. (FTP over SSH. Port 22. [Secure Protocols](concepts/secure-protocols.md).)

**SHA** — Secure Hash Algorithm. (Family of NIST hash algorithms: SHA-1 (160-bit, deprecated), SHA-2 (224/256/384/512-bit), SHA-3 (224/256/384/512-bit). [Hash Functions](concepts/hash-functions.md).)

**SIEM** — Security Information and Event Management. (Platform that aggregates, normalizes, correlates, and alerts on log data from across the environment. [Log Management and SIEM](concepts/log-management-siem.md), [Logging and Monitoring](concepts/logging-monitoring.md).)

**SIP** — Session Initiation Protocol. (VoIP signaling protocol; initiates, maintains, and terminates voice/video sessions. [VoIP Security](concepts/voip-security.md).)

**SLA** — Service Level Agreement. (Contractual addendum defining expected service performance, security controls, and compliance requirements with a vendor. [Software Acquisition Security](concepts/software-acquisition-security.md).)

**SLC** — System Life Cycle. (Continuation of SDLC from production deployment through operations, maintenance, and disposal. [SDLC](concepts/sdlc.md).)

**SLE** — Single Loss Expectancy. (Expected cost of a single risk event: AV × EF. [Risk Management](concepts/risk-management.md).)

**SLR** — Service Level Requirements. (Pre-contract document defining service, security, and performance requirements used during vendor evaluation.)

**SMART** — Specific, Measurable, Achievable, Relevant, Timely. (Framework for defining high-quality security metrics. [Security Metrics](concepts/security-metrics.md).)

**SMTP** — Simple Mail Transfer Protocol. (Email sending. Port 25 TCP.)

**SNMP** — Simple Network Management Protocol. (Network device management. Ports 161/162 UDP. Use v3 only. [Network Protocols](concepts/network-protocols.md).)

**SOAR** — Security Orchestration, Automation, and Response. (Platform automating threat response workflows using inputs from SIEM and other tools. [Logging and Monitoring](concepts/logging-monitoring.md).)

**SOC** — System and Organization Controls. (AICPA-defined audit report types (SOC 1/2/3) produced under SSAE 18. [SOC Reports](standards/soc-reports.md).)

**SOCKS** — Socket Secure. (Layer 5 proxy/tunneling protocol with encryption. [VPN](concepts/vpn.md).)

**SoD** — Segregation of Duties (also Separation of Duties). (No single person controls all steps of a critical function; reduces fraud and error.)

**SOX** — Sarbanes-Oxley Act. (US law governing accuracy of financial reporting for public companies; enacted post-Enron to prevent financial fraud.)

**SP** — Service Provider. (In SAML/federation, the SP/RP consumes IdP assertions. [Federation](concepts/federation.md).)

**SPML** — Services Provisioning Markup Language. (Deprecated XML-based OASIS standard for provisioning users to cloud-based services; predecessor to SAML/OAuth.)

**SQL** — Structured Query Language. (Language used to communicate with and query databases; SQL injection exploits inadequate input validation. [Database Security](concepts/database-security.md).)

**SRTP** — Secure Real-time Transport Protocol. (Secure version of RTP; provides encryption, authentication, integrity, and replay protection for VoIP. [VoIP Security](concepts/voip-security.md).)

**SSD** — Solid State Drive. (Flash-memory-based storage; standard overwriting does not reliably work for sanitization; physical destruction or vendor tools preferred. [Data Security Controls](concepts/data-security-controls.md).)

**SSH** — Secure Shell. (Port 22. Encrypted remote login, file transfer (SFTP), tunneling. Replaces Telnet. [Secure Protocols](concepts/secure-protocols.md).)

**SSL** — Secure Sockets Layer. (Predecessor to TLS. SSLv2 and SSLv3 are broken — do not use. [TLS Standard](standards/tls-standard.md).)

**SSAE** — Statements on Standards for Attestation Engagements. (AICPA attestation standards; current version is SSAE 18. Governs SOC audits. [SOC Reports](standards/soc-reports.md).)

**SSO** — Single Sign-On. (One authentication event granting access to multiple systems within a single organization. [SSO](concepts/sso.md).)

**SSRF** — Server-Side Request Forgery. (Attack where server fetches an attacker-controlled URL, potentially reaching internal resources. OWASP Top 10 A10. [OWASP Top 10](concepts/owasp-top-10.md).)

**SSDLC** — Secure Software Development Life Cycle. (SDLC with security embedded at every phase from inception. [SDLC](concepts/sdlc.md).)

**ST** — Security Targets. (Common Criteria component; vendor's written statement explaining how their product meets the PP requirements. [Common Criteria](concepts/common-criteria.md).)

**STP** — Spanning Tree Protocol. (Prevents network loops in Ethernet by disabling redundant paths. [Network Devices](concepts/network-devices.md).)

**SYN** — Synchronize. (TCP flag used in the three-way handshake. Exploited in SYN floods. [TCP/IP Model](concepts/tcp-ip-model.md).)

## T

**TACACS+** — Terminal Access Controller Access Control System Plus. (TCP-based AAA protocol; encrypts entire packet; more secure than RADIUS. Cisco proprietary. [Network Access Control](concepts/network-access-control.md).)

**TCB** — Trusted Computing Base. (The totality of all protection mechanisms within an architecture — hardware, firmware, software, and processes. [Trusted Computing Base](concepts/trusted-computing-base.md).)

**TCSEC** — Trusted Computer System Evaluation Criteria. (First evaluation criteria system; "Orange Book"; DoD 1980s; measures confidentiality only. [TCSEC](standards/tcsec.md).)

**TFTP** — Trivial File Transfer Protocol. (Port 69 UDP. No authentication — highly insecure; disable. [Network Protocols](concepts/network-protocols.md).)

**TGS** — Ticket Granting Service. (Kerberos KDC component that issues service tickets to users presenting a valid TGT. [Kerberos](concepts/kerberos.md).)

**TGT** — Ticket Granting Ticket. (Kerberos token issued by the AS; user presents to TGS to obtain service tickets. Encrypted with TGS key — user cannot decrypt it. [Kerberos](concepts/kerberos.md).)

**TKIP** — Temporal Key Integrity Protocol. (WPA encryption using RC4 with per-packet key mixing. Stopgap for WEP; deprecated — superseded by CCMP/AES. [Wireless Security](concepts/wireless-security.md).)

**TLS** — Transport Layer Security. (Successor to SSL. Current versions: 1.2 (acceptable), 1.3 (preferred). [TLS Standard](standards/tls-standard.md).)

**TOCTOU (1)** — Time-of-Check Time-of-Use (security kernels / ICS). (Race condition; a window of time between when access is checked and when it is used, exploitable by attackers. [Trusted Computing Base](concepts/trusted-computing-base.md).)

**TOCTOU (2)** — Time-Of-Check Time-Of-Use (Kerberos). (Attack exploiting the gap between when access is verified and when it is used; Kerberos is vulnerable due to single-ticket model. [Kerberos](concepts/kerberos.md).)

**TOCTOU (3)** — Time-of-Check Time-of-Use (software). (Race condition vulnerability in code; gap between when a value is validated and when it is used, allowing malicious substitution. [Secure Coding Practices](concepts/secure-coding-practices.md).)

**TOE** — Target of Evaluation. (Common Criteria component; the specific vendor product being evaluated. [Common Criteria](concepts/common-criteria.md).)

**TOGAF** — The Open Group Architecture Framework. (Enterprise architecture framework; focuses on resource efficiency and modular structure.)

**TOR** — The Onion Router. (Onion network providing anonymity and multi-layer encryption for data in transit. [Data States](concepts/data-states.md).)

**TOTP** — Time-based One-Time Password. (Synchronous OTP variant where the token changes on a time interval. [Authentication Factors](concepts/authentication-factors.md).)

**TP** — Transformation Procedure. (Clark–Wilson mechanism; authorized program that mediates subject access to CDIs. [Clark–Wilson](concepts/clark-wilson.md).)

**TPM** — Trusted Platform Module. (Hardware chip implementing ISO/IEC 11889; performs cryptographic operations; provides platform integrity via binding and sealing; every TPM has a unique Endorsement Key. [TPM](concepts/tpm.md), [Key Management](concepts/key-management.md).)

**TSC** — Trust Services Criteria. (Five criteria evaluated in a SOC 2 report: Security, Availability, Confidentiality, Processing Integrity, Privacy. [SOC Reports](standards/soc-reports.md).)

## U

**UBA / UEBA** — User (and Entity) Behavior Analytics. (ML-based system that baselines normal behavior and detects anomalies; used for insider threat and compromised account detection. [Logging and Monitoring](concepts/logging-monitoring.md).)

**UDI** — Unconstrained Data Item. (Clark–Wilson: untrusted external data entering the system; must be transformed to CDI via TP before use. [Clark–Wilson](concepts/clark-wilson.md).)

**UDP** — User Datagram Protocol. (Connectionless, unreliable, fast Layer 4 transport. Used for DNS, VoIP, streaming. [TCP/IP Model](concepts/tcp-ip-model.md).)

**UPS** — Uninterruptible Power Supply. (Battery backup providing short-term power and power conditioning. [Facility Security](concepts/facility-security.md).)

**USML** — United States Munitions List. (List of defense articles and services controlled under ITAR.)

**UTP** — Unshielded Twisted Pair. (Common copper cabling; susceptible to electromagnetic interference.)

## V

**VESDA** — Very Early Smoke Detection Apparatus. (Most advanced and most expensive fire detection system; detects fire at incipient/smoldering stage. [Facility Security](concepts/facility-security.md).)

**VLAN** — Virtual Local Area Network. (Logical network segmentation via 802.1Q tagging. [Network Devices](concepts/network-devices.md).)

**VMM** — Virtual Machine Manager/Monitor. (Another name for a hypervisor; software managing virtual machines on physical hardware. [Virtualization Security](concepts/virtualization-security.md).)

**VoIP** — Voice over Internet Protocol. (Voice communications over IP data networks. [VoIP Security](concepts/voip-security.md).)

**VPC** — Virtual Private Cloud. (Logically isolated portion of a public cloud provider's infrastructure. [SDN and NFV](concepts/sdn-nfv.md).)

**VPN** — Virtual Private Network. (Tunneling + encryption for secure communication over untrusted networks. [VPN](concepts/vpn.md).)

**VRF** — Virtual Routing and Forwarding. (Multiple independent routing tables on a single router/switch; enables logical network isolation.)

## W

**WAF** — Web Application Firewall. (Application-layer firewall specialized for HTTP/S traffic; inspects and blocks common web attacks including SQL injection and XSS. [Firewalls](concepts/firewalls.md).)

**WAN** — Wide Area Network. (Network connecting geographically dispersed LANs.)

**WEP** — Wired Equivalent Privacy. (Original 802.11 security — **broken**; RC4 implemented with short, repeating IVs — classic implementation attack. Use WPA2/WPA3. [Wireless Security](concepts/wireless-security.md).)

**WPA** — Wi-Fi Protected Access. (Interim wireless security using TKIP; deprecated. [Wireless Security](concepts/wireless-security.md).)

**WPA2** — Wi-Fi Protected Access 2. (Current wireless security standard (802.11i); uses CCMP/AES. [Wireless Security](concepts/wireless-security.md).)

**WPA3** — Wi-Fi Protected Access 3. (Latest wireless security standard (2018); uses CCMP or GCMP; not yet widely deployed. [Wireless Security](concepts/wireless-security.md).)

**WRT** — Work Recovery Time. (Time required to verify integrity of systems and data as they return to service; a component of MTD. MTD = RTO + WRT. [RTO/RPO/MTD](concepts/rto-rpo-mtd.md).)

**WSUS** — Windows Server Update Services. (Microsoft tool for managing and distributing patches to Windows systems centrally. [Patch Management](concepts/patch-management.md).)

## X

**XACML** — eXtensible Access Control Markup Language. (Standard defining ABAC policy language, architecture, and processing model. [Access Control Models](concepts/access-control-models.md).)

**XML** — eXtensible Markup Language. (Human and machine-readable markup language; format used by SAML assertions and XACML policies. [SAML](concepts/saml.md).)

**XOR** — Exclusive OR. (Logical operation used in stream ciphers: 0⊕0=0, 0⊕1=1, 1⊕0=1, 1⊕1=0.)

**XSS** — Cross-Site Scripting. (Injection of malicious scripts into web pages viewed by other users; two types: stored (persistent) and reflected (most common); target is the user's browser. [Secure Coding Practices](concepts/secure-coding-practices.md).)

## Y

*(No entries at this time.)*

## Z

*(No entries at this time.)*

## 0–9

**3DES** — Triple Data Encryption Standard. (Applies DES three times; 168-bit nominal / 112-bit effective key length (meet-in-the-middle removes 56 bits); now disallowed by NIST. [Symmetric Crypto](concepts/symmetric-crypto.md).)

**802.11** — IEEE 802.11 Wireless LAN Standard. (Defines wireless communication protocols; 802.11i = WPA2 security standard. [802.11](standards/802-11.md).)

**802.1X** — IEEE 802.1X Port-Based Network Access Control. (Standard for EAP-based authentication to wired and wireless networks. [802.1X](standards/802-1x.md).)
