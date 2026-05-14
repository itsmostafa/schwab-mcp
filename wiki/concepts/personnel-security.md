---
title: "Personnel Security Policies and Procedures"
type: concept
domain: 1
tags: [personnel, hiring, onboarding, termination, separation-of-duties, job-rotation, least-privilege, need-to-know, nda, duress]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Personnel Security Policies and Procedures

Personnel represent one of the most significant security risks in any organization. Personnel
security policies address the full employee lifecycle: hiring → employment → termination.

## Candidate Screening and Hiring (§1.8.1)

New personnel represent inherent risk. Controls to mitigate:
- Background checks (criminal, employment history, education verification)
- Reference checks
- Access badges, ID cards
- Acceptable use policies and code of conduct signed before access is granted

## Employment Agreements and Controls (§1.8.2)

**At onboarding:**
- Security policies and acceptable use agreements must be reviewed and signed **before** issuing credentials or badges.
- Identity proofing verifies the person's identity.
- Access provisioning based on **least privilege** and **need-to-know**.

**During employment — key controls (Table 1-15):**

| Control | Purpose |
|---|---|
| **Job Rotation** | Detects/prevents fraud; builds personnel redundancy; ensures no single person has prolonged unchecked access |
| **Mandatory Vacation** | Forces another employee to cover the role — uncovers hidden fraud or misconfigurations |
| **Separation of Duties** | Requires ≥2 people to complete critical tasks; prevents single-person fraud |
| **Least Privilege** | Grants only minimum permissions needed for the job |
| **Need-to-Know** | Access to sensitive assets only for those who require it to do their job |

## Termination / Offboarding (Table 1-16)

- **Timely removal of access** — disable accounts immediately (or simultaneously with termination notification).
- **Voluntary vs. involuntary termination:**
  - Voluntary: lower security risk; standard offboarding.
  - **Involuntary**: high security risk — potentially hostile employee may steal/tamper with data. May require physical security escort from building.
- Credentials, badges, and physical keys must be recovered.
- Relevant internal parties must be notified.

## Employee Duress

An employee acting under duress (coercion) may perform actions under threat. Organizations use
**duress codes/keywords** — a word or phrase given in an otherwise normal communication that
signals the employee is under threat. Training is critical so employees remain calm and signal correctly.

## Third-Party Controls

Personnel security extends to **contractors, consultants, and vendors**:
- Governed by contracts, SLAs, and NDAs.
- NDAs prevent disclosure of sensitive information accessed in the course of work.
- NCAs (Noncompete Agreements) prevent employees from joining competitors immediately after leaving.
- Attestation and audit confirm third-party compliance with organizational policies.

## Exam-Relevant Nuance

- **Separation of duties** vs. **least privilege**: Separation of duties = requires multiple people for a task; Least privilege = limits individual permissions. They address different threats.
- **Job rotation** vs. **mandatory vacation**: Both detect fraud but job rotation also provides cross-training/redundancy. Mandatory vacation specifically forces someone else to cover the role for an extended period.
- Involuntary termination is the **higher-risk** event on the exam — always requires more immediate and thorough offboarding steps.
- Even when vendors are responsible for processing your data, **your organization remains accountable** for that data.

## Cross-links

- [Governance](governance.md)
- [Security Policies](security-policies.md)
- [Security Awareness](security-awareness.md)
- [SCRM](scrm.md)

## Sources

- destination-cissp §1.8.1, §1.8.2 (Tables 1-15, 1-16)
