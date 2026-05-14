---
title: "Compliance Requirements"
type: concept
domain: 1
tags: [compliance, laws, regulations, industry-standards, hipaa, gdpr, pci-dss, fisma, glba]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Compliance Requirements

Organizations face a matrix of compliance obligations based on industry, jurisdiction, asset types,
and business activities. Controls must align with applicable requirements.

## Categories of Compliance (Table 1-13)

| Category | Description | Examples |
|---|---|---|
| **Laws** | Legally binding requirements based on jurisdiction or industry | HIPAA, GLBA, COPPA, FERPA, GDPR, FISMA, DMCA, CCPA, SOX |
| **Regulations** | Government-imposed rules often stemming from laws | ITAR (State Dept.), EAR (Commerce), Encryption Export Controls |
| **Industry Standards** | Procedural and technical rules specific to an industry | NERC CIP (energy), NIST frameworks, ISO standards, PCI-DSS |
| **Import/Export Controls** | Country-based rules on technology movement | Wassenaar Arrangement, ITAR, EAR |
| **Transborder Data Flow** | Rules restricting data movement across national borders | GDPR (EU data must stay in EU), PIPEDA (Canada) |

## Key Laws and Regulations

| Law/Regulation | Jurisdiction | Covers |
|---|---|---|
| **HIPAA** | US | Health information (PHI) |
| **GLBA** | US | Financial institutions; privacy of financial information |
| **SOX** | US | Financial reporting accuracy |
| **COPPA** | US | Children's online privacy |
| **CCPA / CPRA** | California | Consumer privacy (GDPR-like) |
| **FERPA** | US | Student educational records |
| **FISMA** | US Federal | Federal information systems security |
| **DMCA** | US | Digital copyright protection |
| **GDPR** | EU | Personal data of EU citizens |
| **PIPEDA** | Canada | Personal information in commercial activity |

## Compliance Function Interplay

- **Legal function** — identifies applicable laws and regulations.
- **Compliance function** — monitors ongoing compliance.
- **Security function** — implements and enforces controls.
- **Privacy function** — addresses personal data requirements.

Best practice: security works with legal to identify requirements, then designs and implements the controls.

## Transborder Data Flow

- Many countries require data to remain within their physical borders (data residency / data localization).
- GDPR is the leading example: personal data of EU citizens must remain within EU borders.
- Challenge: cloud services and global SaaS make compliance complex for multinational organizations.

## Exam-Relevant Nuance

- The exam does NOT expect you to be a compliance expert in any specific law — know the major laws by name and general purpose.
- GDPR is tested most deeply (7 principles, 72-hour notification, data subject rights).
- The exam may ask which functional group (Legal, Compliance, Security, Privacy) is responsible for which compliance activity — see the roles described above.

## Cross-links

- [Privacy](privacy.md)
- [Governance](governance.md)
- [Intellectual Property](intellectual-property.md)
- [SCRM](scrm.md)

## Sources

- destination-cissp §1.4.4, §1.4.5, §1.4.6 (Tables 1-9, 1-10, 1-13)
