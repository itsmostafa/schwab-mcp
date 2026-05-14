---
title: "Domain 7 — Security Operations"
type: domain
domain: 7
tags: [incident-response, forensics, logging, siem, dr, bcp, physical-security, change-management, patch-management, malware]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 7 — Security Operations

Exam weight: **13%**. Covers investigations, logging, monitoring, incident management,
disaster recovery, business continuity, and physical security.

## Subtopics (from CISSP Exam Outline)

### 7.1 Understand and comply with investigations

Evidence collection and handling begins with securing the scene — no touching, no mouse movement — to preserve volatile and physical evidence from contamination. The forensic investigation process follows a defined sequence: secure the scene, collect evidence, examine, analyze, and report. Throughout the process, the **chain of custody** must be established from the first moment of collection and maintained until evidence is potentially presented in court. The **five rules of evidence** (authentic, accurate, complete, convincing, admissible) and **Locard's Exchange Principle** (every contact leaves a trace) are foundational concepts. Investigations can be criminal, civil, regulatory, or administrative, and the appropriate response process differs for each.

- Evidence collection and handling → [Chain of Custody](../concepts/chain-of-custody.md)
- Reporting and documentation
- Investigative techniques → [Forensics](../concepts/forensics.md)

### 7.2 Conduct logging and monitoring activities

Logging and monitoring is the detection backbone of security operations. A **SIEM** aggregates and correlates log events from disparate sources (firewalls, IDS, servers, applications) to surface patterns that individual systems would miss — the classic example being the same IP address appearing simultaneously in two different users' authentication logs. **UEBA** adds machine learning to baseline normal behavior and flag anomalies, making it the primary tool for detecting insider threats and compromised accounts. **SOAR** extends this capability by automating response workflows, reducing analyst fatigue and response time. All logging infrastructure requires synchronized clocks (NTP) to enable meaningful cross-system correlation.

- Intrusion detection and prevention (IDPS) *(see Domain 4)*
- Security information and event management (SIEM) → [Logging & Monitoring](../concepts/logging-monitoring.md)
- Continuous monitoring and tuning
- Egress monitoring *(see Domain 4)*
- Log management *(see Domain 6)*
- Threat intelligence (threat feeds, threat hunting)
- User and entity behavior analytics (UEBA)

### 7.3 Perform configuration management (CM)

Configuration management ensures every asset is deployed in a known, secure, hardened state and tracked in a **CMDB** (Configuration Management Database) with an assigned owner. Secure provisioning means never deploying with default credentials or settings — hardening guides such as CIS Benchmarks define the approved baseline. **Configuration drift** (when a system's running state diverges from its approved baseline) is detected through credentialed vulnerability scans and file integrity monitoring. Any configuration change must flow through the change management process to maintain a consistent and auditable environment.

- Provisioning, baselining, automation → [Configuration Management](../concepts/configuration-management.md)

### 7.4 Apply foundational security operations concepts

The operational security principles in this subtopic work together as a system: **need-to-know** restricts access to data, while **least privilege** restricts what actions a user can take. **Segregation of Duties (SoD)** ensures no single person can complete a sensitive transaction alone, and **job rotation** provides a detective control by periodically placing another employee in a role where fraud would be discovered. **Privileged account management (PAM)** wraps these principles around the accounts with greatest power — requiring MFA, minimizing standing access, and enabling enhanced logging on every privileged session. SLAs formalize operational commitments between the organization and its vendors.

- Need-to-know / least privilege
- Segregation of Duties (SoD) and responsibilities
- Privileged account management
- Job rotation
- Service-level agreements (SLA)

### 7.5 Apply resource protection

Data protection depends on recognizing that media degrades over time — even the most reliable storage technology has a **MTBF** (Mean Time Between Failures) — and that data's encryption must be periodically re-evaluated as cryptographic standards evolve. Media management includes policies for labeling, access control, transport, sanitization, and end-of-life disposal. For very long-retention data, organizations must plan not just for physical media rotation but also for format migration to ensure future readability. Hardware and software assets require current inventories, assigned ownership, and active license management.

- Media management
- Media protection techniques
- Data at rest / data in transit

### 7.6 Conduct incident management

Incident management begins with a crucial distinction: **events** happen continuously; **incidents** are the adverse events that trigger a formal response. Detection uses a combination of automated tools (IPS/IDS, DLP, SIEM, anti-malware) and manual processes (administrative review, physical security). The response process — whether framed as NIST SP 800-61's four phases or the eight-step operational model — moves from preparation through detection, containment, reporting, recovery, remediation, and lessons learned. A single spokesperson should handle all external communications during an incident so the response team stays focused. When the incident impact assessment shows that MTD will be exceeded, a **disaster** must be declared and the DRP activated.

- Detection
- Response
- Mitigation
- Reporting
- Recovery
- Remediation
- Lessons learned
- Digital forensics tools, tactics, and procedures → [Forensics](../concepts/forensics.md), [Incident Management](../concepts/incident-management.md)
- Artifacts (data, computer, network, mobile device)

### 7.7 Operate and maintain detection and preventative measures

This subtopic covers the operational layer of security controls. **Malware** is the most varied threat: viruses require user action to trigger; worms self-propagate; ransomware encrypts and extorts; rootkits conceal attacker tools; zero-days have no existing signatures. Anti-malware defenses include signature-based detection (fast, misses zero-days), heuristic/behavioral detection (catches novel malware, generates more false positives), and ML/AI-assisted classification. **Patch management** is the primary mechanism for closing known vulnerabilities — patches must be tested before deploying to production, and all deployments must flow through change management. **Change management** (covered in 7.9 of the source) ensures changes are deliberate, approved, tested, and documented; the CAB/CCB provides cross-functional oversight for significant changes.

- Firewalls (next-generation, web application, network) *(see Domain 4)*
- Intrusion detection systems (IDS) and intrusion prevention systems (IPS) *(see Domain 4)*
- Whitelisting / blacklisting *(see Domain 4)*
- Third-party provided security services
- Patch and vulnerability management → [Patch Management](../concepts/patch-management.md), [Vulnerability Management Ops](../concepts/vulnerability-management-ops.md)
- Change management processes → [Change Management](../concepts/change-management.md)
- Malware types and anti-malware → [Malware Analysis](../concepts/malware-analysis.md)

### 7.10 Implement recovery strategies

Recovery strategies span from component-level redundancy up to full alternate site strategies. At the component level, **fail modes** define what happens when something breaks: fail-safe prioritizes human safety (doors unlock), fail-secure prioritizes security (firewall blocks all traffic), and fail-soft prioritizes availability (everything passes). **Spare parts** strategies (cold/warm/hot spares) and **RAID** configurations (striping for speed, mirroring for availability, parity for balance) address component and disk-level resilience. **Backup strategies** (full, differential, incremental, mirror) determine RPO: the key distinction is that differential does not reset the archive bit (grows larger, restores faster) while incremental does (stays small, restores slower). Recovery **sites** (cold/warm/hot/mobile/redundant) address facility-level disasters, with cost inversely proportional to recovery time.

- Backup storage strategies (cloud, onsite, offsite) → [Backup Strategies](../concepts/backup-strategies.md)
- Recovery site strategies (cold vs. hot, resource capacity agreements) → [Disaster Recovery Sites](../concepts/disaster-recovery-sites.md)
- Failure modes → [Failure Modes](../concepts/failure-modes.md)
- RAID levels, clustering, redundancy

### 7.11 Implement disaster recovery (DR) processes

Disaster recovery is the tactical execution of the DRP when a disaster is declared. A **disaster** is declared when the impact assessment shows that MTD will be exceeded — the authority to declare rests with the CEO or a designated BCM Board. **BCM** creates the structure; **BCP** addresses business process continuity (strategic); **DRP** addresses technology recovery (tactical). The **BIA** is the foundational step that identifies critical functions and determines their RPO, RTO, WRT, and MTD values. The formula **MTD = RTO + WRT** is the central constraint: recovery plus verification must fit within the maximum tolerable downtime. **Restoration order** at the DR site follows BIA priority (most critical first); restoration back to the primary site reverses the order (least critical first) to validate the rebuilt environment.

- Response, personnel, communications, assessment → [BCP/DRP Operations](../concepts/bcp-drp-operations.md)
- RPO, RTO, WRT, MTD → [RTO/RPO/MTD](../concepts/rto-rpo-mtd.md)

### 7.12 Test disaster recovery plans (DRP)

DRP testing validates that plans actually work and that personnel know how to execute them. Tests progress from paper-based to live, in order of increasing risk and realism: **read-through** (checklist review by author), **walkthrough** (tabletop discussion by all stakeholders), **simulation** (facilitated disaster scenario, paper-based), **parallel** (scenario response using backup systems only, not production), and **full-interruption** (production systems are involved — most realistic, highest risk). Full-interruption tests require successful completion of all prior test types and explicit management approval because production disruption is possible. All relevant stakeholders — not just IT — must participate.

- Read-through / tabletop
- Walkthrough
- Simulation
- Parallel
- Full-interruption / full-scale

### 7.13 Participate in Business Continuity (BC) planning and exercises

BCM encompasses both BCP and DRP and has three goals in strict priority order: (1) **safety of people** — always first; (2) **minimization of damage** to facilities and assets; (3) **survival of the business** and its critical functions. The BCP process begins with a contingency planning policy, flows through the BIA to identify critical functions, and results in documented plans that are regularly tested, trained, and maintained as living documents. An organization must remain compliant with legal and regulatory obligations even during a declared disaster.

### 7.14 Implement and manage physical security

All physical security concepts are covered in Domain 3, including perimeter controls (fencing, guards, lighting, CCTV) and internal controls (locks, badges, mantraps). Domain 7 references these controls as operational elements that must be maintained and tested as part of the overall security operations program.

- Perimeter security controls *(see Domain 3)*
- Internal security controls *(see Domain 3)*

### 7.15 Address personnel safety and security concerns

Personnel safety extends beyond the physical campus. Employees traveling internationally — especially to high-risk regions — remain the organization's responsibility; the organization should provide medical coverage, travel insurance, 24/7 emergency assistance access, and pre-travel security briefings. **Duress** situations (coercion to take actions against one's will) require employees to be trained in non-reactive responses — code words for requesting help, silent alarm triggers, and de-escalation techniques. Security awareness training must address insider threats, social engineering, and MFA fatigue attacks.

- Travel
- Security training and awareness (insider threat, social media, 2FA fatigue)
- Emergency management
- Duress

---

## Key concepts

- [Incident Management](../concepts/incident-management.md) — event vs. incident vs. breach; 8-step model vs. NIST 4-phase; escalation; breach notification
- [Forensics](../concepts/forensics.md) — order of volatility; forensic imaging; write blockers; live vs. dead acquisition; artifacts
- [Chain of Custody](../concepts/chain-of-custody.md) — evidence handling; five rules of evidence; evidence types; Locard's principle; MOM; investigation types
- [Logging & Monitoring](../concepts/logging-monitoring.md) — SIEM capabilities; UEBA; SOAR; continuous monitoring; NTP; threat intelligence; insider threat detection
- [Configuration Management](../concepts/configuration-management.md) — CMDB; baseline; hardening; CIS Benchmarks; configuration drift
- [Patch Management](../concepts/patch-management.md) — patch lifecycle; agent/agentless/passive detection; WSUS; emergency patches; EOL systems
- [Change Management](../concepts/change-management.md) — RFC; CAB; emergency change; rollback plans; ITIL change types
- [BCP/DRP Operations](../concepts/bcp-drp-operations.md) — BCM hierarchy; BIA; three BCM goals; declaring a disaster; restoration order
- [RTO/RPO/MTD](../concepts/rto-rpo-mtd.md) — MTD = RTO + WRT; cost relationships; MTBF/MTTR
- [Backup Strategies](../concepts/backup-strategies.md) — full/incremental/differential/mirror; archive bit; 3-2-1 rule; tape rotation; CRC
- [Disaster Recovery Sites](../concepts/disaster-recovery-sites.md) — cold/warm/hot/mobile/redundant; geographic disparity; spare parts; RAID 0/1/5/10
- [Vulnerability Management Operations](../concepts/vulnerability-management-ops.md) — CVSS; prioritization; exception management; continuous scanning
- [Malware Analysis](../concepts/malware-analysis.md) — malware types; IOCs; signature vs. heuristic detection; sandboxing; static vs. dynamic analysis
- [Failure Modes](../concepts/failure-modes.md) — fail-safe vs. fail-secure vs. fail-soft

### Standards
- [NIST SP 800-61](../standards/nist-sp-800-61.md) — Computer Security Incident Handling Guide (4 phases, IR team models)
- [ITIL v4](../standards/itil-v4.md) — Change management (standard/normal/emergency); incident vs. problem management

## Sources

- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §7.1–7.15 (pp. 0824–0917)
