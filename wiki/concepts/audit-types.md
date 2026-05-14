---
title: "Audit Types and SOC Reports"
type: concept
domain: 6
tags: [audit, soc, ssae18, type1, type2, aicpa, isae3402, third-party-audit]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Audit Types and SOC Reports

## Audit Overview

A security audit is a formal, structured examination of controls by a qualified auditor. Unlike a vulnerability assessment, an audit is compliance- and assurance-focused, and findings are typically reported to management and (for third-party audits) to customers. See [Assessment vs Audit](assessment-vs-audit.md).

## Audit Strategies

Three strategies (destination-cissp §6.1.2, §6.5.1):

| Strategy | Who Audits | What |
|---|---|---|
| **Internal** | Organization's own employees | Organization's own systems and controls |
| **External** | Either: internal employees examining a vendor, or an external firm examining the organization | Vendor controls (internal team) or organization's own controls (external firm) |
| **Third-party** | Independent audit firm | Service provider's systems; results shared with the service provider's customers |

Audit plans typically follow these steps (destination-cissp §6.5.1):
1. Define the audit objective.
2. Define the audit scope.
3. Determine business unit leaders and stakeholders to involve.
4. Choose the audit team.
5. Plan the audit.
6. Conduct the audit.
7. Document audit results.
8. Communicate results.

## Audit Roles (destination-cissp §6.5.3)

| Role | Responsibility |
|---|---|
| **Executive/Senior Management** | Sets tone from the top; promotes and supports the audit process. |
| **Audit Committee** | Board members + senior stakeholders; provides oversight and direction to the audit program. |
| **Security Officer (CSO/CISO)** | Advises on security-related risks to evaluate; not the primary audit owner. |
| **Compliance Manager** | Schedules audits; hires and trains auditors; ensures compliance with laws, regulations, and policy. |
| **Internal Auditors** | Company employees; assess effectiveness of corporate internal controls. |
| **External Auditors** | Independent from the organization; provide unbiased, independent audit reports. |

## SOC Report Standards Evolution

Audit standards have evolved (destination-cissp §6.5.2):

> SAS 70 → SSAE 16 → SSAE 18 (current US standard)

- **AICPA** (American Institute of Certified Public Accountants) is the US governing body overseeing these standards.
- **ISAE 3402** is the equivalent international standard (very similar to SSAE 16/18 with minor variations).

## System and Organization Controls (SOC) Reports

SOC reports (previously "Service Organization Controls") are produced under SSAE 18. They are the primary mechanism for a service provider to demonstrate trust to its customers. Three types exist:

### SOC 1

- **Focus**: Financial reporting risks and the controls related to financial reporting.
- **Audience**: Customers who rely on the service provider for financial transaction processing; their auditors.
- **Use case**: Organizations processing payroll, financial transactions, etc.

### SOC 2

- **Focus**: The five **Trust Services Criteria** (TSC):
  1. **Security** — protection against unauthorized access (always included).
  2. **Availability** — system availability for operation and use (always included).
  3. **Confidentiality** — protection of confidential information (always included).
  4. **Processing Integrity** — complete, valid, accurate, timely processing (optional).
  5. **Privacy** — personal information collection, use, retention, and disposal (optional).
- **Audience**: Security professionals, customers evaluating the service provider's security posture.
- **Nature**: Detailed, often confidential. Contains specific control descriptions and system details. Should be protected from unauthorized disclosure.
- **Most relevant SOC report for security professionals** (destination-cissp §6.5.2).

### SOC 3

- **Focus**: A sanitized, public-friendly summary of the SOC 2 report.
- **Audience**: Prospective customers and general public.
- **Use case**: Marketing. Allows a service provider to demonstrate general security assurance without disclosing the confidential details in a SOC 2.
- **Cannot be used for compliance verification** — lacks the detail of SOC 2.

## Type 1 vs Type 2 Reports

Both SOC 1 and SOC 2 can be Type 1 or Type 2:

| | Type 1 | Type 2 |
|---|---|---|
| **Focus** | Design of controls **at a point in time** | Design + **operating effectiveness** over a period of time |
| **What the auditor examines** | Policies, procedures, documentation — "does the control appear properly designed today?" | Everything in Type 1, plus samples of actual control operation over the period (typically ~1 year) |
| **Conclusion** | Controls are appropriately designed | Controls are both designed and operating effectively |
| **Typical use case** | First year of auditing program; identifies gaps before committing to Type 2 | Ongoing annual audit; standard expectation for mature organizations |
| **More comprehensive** | No | **Yes — Type 2 is more comprehensive** |

**Key insight** (destination-cissp §6.5.2): SOC 2, Type 2 is the most desirable report for security professionals because it attests to both control design and operating effectiveness over time.

**Typical progression**: SOC 2 Type 1 in Year 1 → SOC 2 Type 2 from Year 2 onward.

## Cross-links

- [Assessment vs Audit vs Penetration Test](assessment-vs-audit.md)
- [SOC Reports Standard](../standards/soc-reports.md)
- [Security Metrics](security-metrics.md)

## Sources

- destination-cissp §6.5.1 (audit process, roles and strategies)
- destination-cissp §6.5.2 (SOC 1/2/3, Type 1/Type 2, SSAE 18, AICPA, ISAE 3402)
- destination-cissp §6.5.3 (audit roles and responsibilities)
- cissp-exam-outline (domain 6 subtopic list)
