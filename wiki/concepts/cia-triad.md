---
title: "CIA Triad and the Five Pillars of Information Security"
type: concept
domain: 1
tags: [cia, confidentiality, integrity, availability, authenticity, nonrepudiation, foundations]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# CIA Triad and the Five Pillars of Information Security

The CIA triad is the foundational model of security. CISSP extends it to five pillars by adding
authenticity and nonrepudiation.

## Key Facts

- **Confidentiality** — Protects assets from unauthorized disclosure. Implemented via need-to-know and least privilege.
- **Integrity** — Prevents unauthorized or accidental changes to assets; makes data more accurate, timely, and meaningful.
- **Availability** — Ensures organizational assets are accessible when required by stakeholders.
- **Authenticity** — Proves an asset's source and origin ("proof of origin"); verifies assets are legitimate and trusted.
- **Nonrepudiation** — The inability to deny having done something; assurance that a party cannot dispute the validity of a transaction or action.

Together, CIA+AN = **the Five Pillars of Information Security**.

## Exam-Relevant Nuance

- The source explicitly frames the CIA goals as "goals of **asset** security," not merely "goals of information security." Assets include people, buildings, processes — not just data.
- Authenticity is sometimes called "proof of origin" — both terms may appear on the exam.
- Nonrepudiation is achieved through mechanisms like digital signatures and audit logs.
- Confidentiality relies on **need-to-know** (limiting who has access to what) and **least privilege** (limiting the permissions granted). These are distinct controls; both can appear on exams.
- The triad is a design model: it guides how an organization structures and implements its security function.

## Cross-links

- [Due Care and Due Diligence](due-care-due-diligence.md)
- [Risk Management](risk-management.md)
- [Security Controls Types](security-controls-types.md)
- [Governance](governance.md)

## Sources

- destination-cissp §1.2.1 (Tables 1-2, 1-3)
