---
title: "Privacy and Data Protection"
type: concept
domain: 1
tags: [privacy, gdpr, pii, phi, personal-data, oecd, pia, dpia, data-protection]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Privacy and Data Protection

Privacy is "the state or condition of being free from being observed or disturbed by other people."
In security, privacy is achieved through security controls — **privacy is impossible without security**.

## Personal Data Categories

- **PI** — Personal Information
- **PII** — Personally Identifiable Information
- **SPI** — Sensitive Personal Information
- **PHI** — Protected Health Information

### Identifiers (Table 1-8)

| Direct Identifiers | Indirect Identifiers | Online Identifiers |
|---|---|---|
| Name, phone, government ID, account numbers, biometric data | Age, gender, ethnicity, city, state, ZIP code | Email address, IP address, cookies |

- **Direct identifiers** uniquely identify an individual on their own.
- **Indirect identifiers** identify individuals when combined with other data.

## GDPR (General Data Protection Regulation)

Enacted May 2018. Considered the global bellwether for privacy law. Key elements:
- Applies to personal data of **EU citizens**, regardless of where the data is processed.
- Each EU state has an independent **Supervisory Authority (SA)** to investigate complaints.
- Data subjects have the right to lodge complaints with the SA.
- **Privacy breach reporting within 72 hours** of discovery.

### GDPR's 7 Principles of Lawful Processing

1. **Lawfulness, fairness, and transparency**
2. **Purpose limitation** — data collected for a specific purpose; used only for that purpose
3. **Data minimization** — collect only what is needed
4. **Accuracy** — keep data accurate and up-to-date
5. **Storage limitation** — don't keep data longer than necessary
6. **Integrity and confidentiality (security)** — protect data with appropriate controls
7. **Accountability** — data controller is responsible for demonstrating compliance

## OECD Privacy Guidelines (Table 1-11)

The OECD provides voluntary (non-mandatory) guidelines as a common baseline for privacy.
Organizations can use them as a starting point before consulting applicable local laws.

| Principle | Core Idea |
|---|---|
| **Collection Limitation** | Collect only what's needed; obtain lawfully with consent |
| **Data Quality** | Data must be relevant, accurate, complete, up-to-date |
| **Purpose Specification** | Specify purpose at time of collection |
| **Use Limitation** | Use only for stated purpose or with consent |
| **Security Safeguards** | Protect with reasonable security controls |
| **Openness** | Be transparent about how data is used |
| **Individual Participation** | Data subjects can obtain and request removal of their data |
| **Accountability** | Data controller is accountable for compliance with all principles |

## Privacy Impact Assessment (PIA / DPIA)

A PIA is a systematic process to determine whether personal data is being protected appropriately
and to minimize risks. GDPR Article 35 mandates DPIAs in certain situations.

### 8-Step PIA Process (Table 1-12)

1. Identify the need for a DPIA (legal requirements, GDPR Article 35)
2. Describe the data processing (what, where, how, why)
3. Assess necessity and proportionality (is collection justified?)
4. Consult interested parties (DPO, stakeholders, data subjects)
5. Identify and assess risks (storage locations, access controls, retention)
6. Identify measures to mitigate risks (defense-in-depth, data minimization)
7. Sign off and record outcomes
8. Monitor and review (ongoing; triggered by changes in processing)

A PIA is not a one-time event — it should be performed whenever risk from personal data processing changes.

## Exam-Relevant Nuance

- GDPR's **72-hour breach notification** requirement is a highly tested fact.
- OECD guidelines are **voluntary** — not a replacement for reviewing specific local laws.
- The exam distinguishes **direct identifiers** (identify someone alone) from **indirect identifiers** (identify someone in combination).
- **Privacy without security is impossible** — this is the source's foundational statement and may appear as a scenario question.
- Data subject = the individual to whom personal data relates.

## Cross-links

- [Compliance Requirements](compliance-requirements.md)
- [Governance](governance.md)
- [Risk Management](risk-management.md)

## Sources

- destination-cissp §1.4.5, §1.4.6 (Tables 1-8 through 1-12)
