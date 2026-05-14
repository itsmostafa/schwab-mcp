---
title: "Domain 1 — Security and Risk Management"
type: domain
domain: 1
tags: [governance, compliance, risk, ethics, law, policy]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 1 — Security and Risk Management

Exam weight: **16%** (heaviest single domain). Covers the foundation of security strategy,
governance, risk, legal/regulatory compliance, and personnel security.

## Subtopics (from CISSP Exam Outline)

### 1.1 Understand, adhere to, and promote professional ethics

The ISC2 Code of Ethics governs all CISSP holders globally. The Preamble establishes that
adherence is a **condition of certification**. The four Canons must be memorized in order:
(1) protect society, (2) act honorably, (3) serve principals diligently, (4) advance the profession.
Canon order determines priority when Canons conflict — Canon 1 always wins. Organizational ethics
are codified into **policies** (corporate laws) driven top-down by senior management.

- **1.1.1 ISC2 Code of Professional Ethics**
- **1.1.2 Organizational code of ethics**

See: [Professional Ethics](../concepts/professional-ethics.md)

### 1.2 Understand and apply security concepts

Security's role is to **support the business** in achieving its goals and objectives, and to
**increase organizational value**. The CIA Triad (Confidentiality, Integrity, Availability) forms
the foundational model. CISSP extends this to five pillars by adding Authenticity and
Nonrepudiation. All five goals apply to **all assets**, not just information.

- **1.2.1 Confidentiality, Integrity, Availability, Authenticity, and Nonrepudiation**

See: [CIA Triad](../concepts/cia-triad.md)

### 1.3 Evaluate and apply security governance principles

Governance = directing an organization to increase its value. **Security governance must be
top-down**: Board → CEO → senior management → security function. Security governance aligns
with corporate governance; the security function is an **enabler**, not a gatekeeper. Key
sub-concepts: alignment to business strategy, scoping/tailoring of controls, accountability vs.
responsibility (accountability can never be delegated), and organizational roles.

- **Alignment of the security function to business strategy, goals, mission, and objectives**
- **Organizational processes** (e.g., acquisitions, divestitures, governance committees)
- **Organizational roles and responsibilities** — Owners are accountable; custodians and users are responsible; security enables both.
- **Security control frameworks** (ISO, NIST, COBIT, SABSA, PCI, FedRAMP) — deferred to Domain 3 (§3.3.1) in destination-cissp; see concept stubs below.
- **Due care / due diligence** — Due care = doing the right thing; Due diligence = proving you did.

See: [Governance](../concepts/governance.md) | [Accountability vs. Responsibility](../concepts/accountability-vs-responsibility.md) | [Due Care and Due Diligence](../concepts/due-care-due-diligence.md)

### 1.4 Understand legal, regulatory, and compliance issues

Organizations face compliance obligations across laws, regulations, and industry standards that
vary by jurisdiction, asset type, and industry. Security works with Legal and Compliance to
implement the required controls. Key privacy framework: GDPR (7 principles; 72-hour breach
notification; data subjects have rights). OECD Privacy Guidelines provide a voluntary baseline.
PIAs/DPIAs are systematic assessments of privacy risks triggered by new or changed data processing.

- **Cybercrimes and data breaches** — understand threat landscape; implement cost-effective defenses.
- **Licensing and Intellectual Property** — trade secrets, patents, copyrights, trademarks; each protects different things.
- **Import/export controls** — Wassenaar Arrangement (cryptography); ITAR (munitions, State Dept.); EAR (commercial items, Commerce Dept.).
- **Transborder data flow** — many countries require data residency within their borders; GDPR is the leading example.
- **Privacy regulations** — GDPR, CCPA, PIPL, POPIA; privacy is impossible without security.
- **Contractual, legal, industry standards, and regulatory requirements** — security works with Legal to identify requirements, then implements controls.

See: [Privacy](../concepts/privacy.md) | [Compliance Requirements](../concepts/compliance-requirements.md) | [Intellectual Property](../concepts/intellectual-property.md)

### 1.5 Understand requirements for investigation types

Investigation types (administrative, criminal, civil, regulatory, industry standards) are covered
in depth in **Domain 7** (§7.1.7). At the Domain 1 level, understand that different investigation
types have different standards of evidence, scope, and legal authority.

### 1.6 Develop, document, and implement security policies, standards, procedures, baselines, and guidelines

The security document hierarchy flows from the overarching policy (Board/CEO) down to functional
policies, then to standards, procedures, baselines, and guidelines. **Policies are corporate laws**
and must be communicated top-down. Standards specify what to use; procedures are mandatory
step-by-step instructions; baselines define minimum configuration levels; guidelines are
recommendations only (no audit findings for non-compliance).

- **Policies** — corporate laws; communicate management's intent; must come from senior management.
- **Standards** — specific hardware/software solutions (e.g., specific AV product).
- **Procedures** — mandatory step-by-step instructions.
- **Baselines** — minimum implementation/configuration levels.
- **Guidelines** — recommendations; not mandatory.

See: [Security Policies](../concepts/security-policies.md)

### 1.7 Identify, analyze, and prioritize business continuity (BC) requirements

The Business Impact Analysis (BIA) is the foundation of business continuity planning. It analyzes
consequences of disruption, identifies critical functions, and establishes recovery priorities.
External dependencies (suppliers, vendors, utilities) must be mapped as part of the BIA. Full
BCP/DRP coverage is in **Domain 7** (§7.11).

- **1.7.1 Business Impact Analysis (BIA)** — analyzes disruption consequences; establishes recovery priorities; maps critical processes and assets.
- **1.7.2 External dependencies** — identify third parties that affect critical functions.

See: [BCP and DRP](../concepts/bcp-drp.md)

### 1.8 Contribute to and enforce personnel security policies and procedures

New employees represent security risk. Personnel security controls span the full employment
lifecycle: pre-hire screening → onboarding agreements → ongoing controls (job rotation, separation
of duties, mandatory vacation, least privilege, need-to-know) → termination/offboarding.
**Involuntary termination** requires immediate and more rigorous offboarding than voluntary
termination. Third-party personnel (contractors, vendors) are governed through NDAs, contracts, SLAs.

- **Candidate screening and hiring** — background checks, reference checks, identity proofing.
- **Employment agreements and policy-driven requirements** — sign policies before credentials issued; NDA, NCA.
- **Onboarding, transfers, and termination processes** — timely access removal; involuntary = high risk.
- **Vendor, consultant, and contractor agreements and controls** — SLAs, NDAs, attestation, audits.

See: [Personnel Security](../concepts/personnel-security.md)

### 1.9 Understand and apply risk management concepts

Risk management is the identification, assessment, and prioritization of risks and the
cost-efficient application of resources to minimize probability and/or impact. The three-step
process: (1) Value assets — qualitative or quantitative analysis; (2) Risk Analysis — identify
threats, vulnerabilities, impact, probability; (3) Treatment — Avoid/Transfer/Mitigate/Accept.
Key formula: **ALE = SLE × ARO**, where SLE = AV × EF. Do not implement a control costing
more than the ALE. Risk is continuous — use the PDCA/Deming Cycle for ongoing improvement.

- **Threat and vulnerability identification**
- **Risk analysis, assessment, and scope** — qualitative vs. quantitative; ALE formula.
- **Risk response and treatment** — Avoid, Transfer (cyber insurance), Mitigate, Accept. Risk ignorance is NOT valid.
- **Applicable types of controls** — 7 types (directive, deterrent, preventive, detective, corrective, recovery, compensating) × 3 categories (administrative, logical/technical, physical).
- **Control assessments** — functional + assurance aspects; both required.
- **Continuous monitoring and measurement** — metrics must match the audience.
- **Reporting** — tailored to internal (senior management) and external (regulators, customers) audiences.
- **Risk frameworks** — NIST SP 800-37, ISO 31000, COSO, ISACA Risk IT.
- **Continuous improvement** — Deming Cycle (Plan-Do-Check-Act).

See: [Risk Management](../concepts/risk-management.md) | [Security Controls Types](../concepts/security-controls-types.md) | [NIST RMF](../standards/nist-rmf.md)

### 1.10 Understand and apply threat modeling concepts and methodologies

Threat modeling systematically identifies, enumerates, and prioritizes threats related to an
asset. Three major methodologies: **STRIDE** (threat-focused; 6 threat categories mapped to CIA
violations; Microsoft-origin), **PASTA** (7-stage; attacker-focused, risk-centric; strategic and
technical), **DREAD** (scoring model; 5 factors scored 1–10; used with STRIDE to rank threats).
Social engineering exploits human psychology; best defense is awareness/training/education.

See: [Threat Modeling](../concepts/threat-modeling.md) | [Social Engineering](../concepts/social-engineering.md)

### 1.11 Apply supply chain risk management (SCRM) concepts

Every supplier, vendor, and service provider introduces risk. Accountability for data and
security cannot be outsourced. Key supply chain risks: product tampering, counterfeits, implants
(hardware/software backdoors). Key mitigations: third-party assessments, minimum security
requirements, SLRs, SLAs, service level reports, Silicon Root of Trust, Physically Unclonable
Functions, and Software Bill of Materials (SBOM).

- **Risks associated with acquisition from suppliers/providers** — tampering, counterfeits, implants.
- **Risk mitigations** — third-party assessment/monitoring, minimum security requirements, SLR/SLA, SBOM, Silicon Root of Trust, PUF.

See: [SCRM](../concepts/scrm.md)

### 1.12 Establish and maintain a security awareness, education, and training program

Everyone is responsible for security but must know what to do. Three levels: **Awareness**
(cultural sensitivity; organization-wide), **Training** (specific skills; role-focused), **Education**
(concepts and decision-making). Topics must align with the risk register. Programs must evolve
with the threat landscape. Effectiveness measured via simulated phishing, assessments, and
reporting rates.

- **Methods and techniques** — live sessions, gamification, security champions, phishing simulations.
- **Periodic content reviews** — emerging topics (AI, blockchain, cloud); triggered by threat landscape changes.

See: [Security Awareness](../concepts/security-awareness.md) | [Social Engineering](../concepts/social-engineering.md)

## Key Concepts

- [CIA Triad](../concepts/cia-triad.md) — Confidentiality, Integrity, Availability (+ Authenticity, Nonrepudiation)
- [Professional Ethics](../concepts/professional-ethics.md) — ISC2 Code of Ethics, 4 canons in order
- [Governance](../concepts/governance.md) — security governance aligned top-down to corporate governance
- [Accountability vs. Responsibility](../concepts/accountability-vs-responsibility.md) — accountability cannot be delegated
- [Due Care and Due Diligence](../concepts/due-care-due-diligence.md) — doing vs. proving
- [Security Policies](../concepts/security-policies.md) — policy/standard/procedure/baseline/guideline hierarchy
- [Compliance Requirements](../concepts/compliance-requirements.md) — laws, regulations, industry standards
- [Privacy](../concepts/privacy.md) — GDPR, OECD, PIA/DPIA
- [Intellectual Property](../concepts/intellectual-property.md) — trade secret, patent, copyright, trademark
- [Risk Management](../concepts/risk-management.md) — ALE formula, risk treatment, PDCA
- [Security Controls Types](../concepts/security-controls-types.md) — 7 types, 3 categories, complete control
- [Threat Modeling](../concepts/threat-modeling.md) — STRIDE, PASTA, DREAD
- [Social Engineering](../concepts/social-engineering.md) — phishing variants, tailgating vs. piggybacking
- [BCP and DRP](../concepts/bcp-drp.md) — BIA, MTD/RTO/RPO, test types
- [Personnel Security](../concepts/personnel-security.md) — hiring, controls, termination
- [SCRM](../concepts/scrm.md) — supply chain risks and mitigations
- [Security Awareness](../concepts/security-awareness.md) — awareness vs. training vs. education
- [NIST RMF](../standards/nist-rmf.md) — SP 800-37, 7 steps
- [NIST CSF](../standards/nist-csf.md) — 5 functions (stub; D3 source pending)
- [ISO/IEC 27001](../standards/iso-27001.md) — ISMS requirements (stub; D3 source pending)
- [ISO/IEC 27002](../standards/iso-27002.md) — Controls guidance (stub; D3 source pending)
- [COBIT](../standards/cobit.md) — IT governance (stub; D3 source pending)

## Sources

- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §1.1–1.12
