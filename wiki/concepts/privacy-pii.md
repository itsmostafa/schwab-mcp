---
title: "Privacy and PII"
type: concept
domain: 2
tags: [pii, privacy, gdpr, data-subject, data-controller, anonymization, pseudonymization]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Privacy and PII

Personally Identifiable Information (PII) is information that can identify a specific individual, either on its own or combined with other data. Protecting PII is a core obligation of asset security and is governed by privacy regulations such as GDPR.

## Key facts

- **PII** is listed as an example classification label in destination-cissp §2.1.3
- **Privacy** is about an individual's right to control their personal information; **security** is about the mechanisms to enforce that control. Privacy is the objective; security is the enabler.
- The **data subject** is the individual to whom personal data pertains (GDPR term; destination-cissp Table 2-3)
- **GDPR** is referenced in destination-cissp as a compliance driver for information obfuscation (§2.6.4) and for data handling obligations generally

## GDPR data roles (brief)

| Role | Definition |
|---|---|
| **Data Controller** | Legal entity that determines the purposes and means of processing personal data |
| **Data Processor** | Entity that processes personal data on behalf of the controller (e.g., a cloud provider) |
| **Data Subject** | Individual whose personal data is being processed |

Note: destination-cissp conflates Data Owner and Data Controller in Table 2-3. In GDPR, these are distinct: the organization is the legal controller; the Data Owner is an internal accountability role. See [Data Roles](data-roles.md) for full discussion.

## Anonymization vs. Pseudonymization

> **Content gap**: destination-cissp does not cover this distinction directly. The exam outline §2.6 lists "data protection methods" and the broader privacy topic. The following is based on exam-outline scope; a future ingest from a source covering GDPR/privacy in depth would strengthen this section.

- **Anonymization**: Irreversible removal of identifying information; re-identification is impossible. Anonymized data is no longer PII under GDPR.
- **Pseudonymization**: Replacing identifying fields with artificial identifiers (pseudonyms); re-identification is *possible* with access to the mapping table. Pseudonymized data is still personal data under GDPR.
- **Tokenization**: Substituting sensitive data (e.g., credit card numbers) with a non-sensitive token; the mapping is stored in a secure token vault. Similar to pseudonymization; widely used in PCI DSS environments.

| Technique | Re-identification possible? | GDPR personal data? |
|---|---|---|
| Anonymization | No | No |
| Pseudonymization | Yes (with mapping table) | Yes |
| Tokenization | Yes (with token vault) | Yes |

## Information obfuscation methods (from source)

Destination-cissp §2.6.4 describes these data-masking techniques relevant to PII protection:

- **Concealing data**: Completely removes access; attribute not visible at all
- **Pruning data**: Removes data value but attribute field remains visible (used in non-production environments)
- **Fabricating data**: Replaces real data with fake data; enables functional testing without real PII exposure
- **Trimming data**: Partially masks an attribute (e.g., showing only last 4 digits of SSN or credit card)
- **Encrypting data**: Creates ciphertext; reversible with key. Used for credit card storage and transit.

See [Information Obfuscation](information-obfuscation.md) for full page.

## Cross-links

- [Data Roles](data-roles.md) — data controller, processor, subject roles
- [Information Obfuscation](information-obfuscation.md) — masking techniques
- [DLP](dlp.md) — DLP often targets PII for monitoring and blocking
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.1.3, §2.3.1, §2.6.4, Table 2-3 (pages 0195, 0200–0202, 0218–0219)
- cissp-exam-outline §2.4, §2.6 (data roles, data protection methods)
- *Anonymization/pseudonymization/tokenization definitions are based on GDPR and general CISSP exam scope — not directly covered in destination-cissp. Flag for future source enrichment.*
