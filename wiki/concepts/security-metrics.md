---
title: "Security Metrics: KPIs and KRIs"
type: concept
domain: 6
tags: [kpi, kri, metrics, smart-metrics, mttd, mttr, security-reporting]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Security Metrics: KPIs and KRIs

## Purpose

Security metrics quantify the effectiveness of security controls and risk posture. They support informed decision-making at all organizational levels — from technical teams to executive management. The key principle: build reports with **"metrics that matter"** for the intended audience (destination-cissp §6.2.8, §6.3.1).

## SMART Metrics Framework (destination-cissp §6.3.1)

Metrics should be **SMART**:

| Letter | Criterion | Question |
|---|---|---|
| **S** | Specific | Is the result clearly stated and easy to understand? |
| **M** | Measurable | Can the result be measured? Is the data available? |
| **A** | Achievable | Can results drive desired outcomes? |
| **R** | Relevant | Is this aligned to business strategy? |
| **T** | Timely | Are results available when needed? |

## KPI vs KRI

The two primary metric types in security (destination-cissp §6.3.1):

| | Key Performance Indicators (KPIs) | Key Risk Indicators (KRIs) |
|---|---|---|
| **Temporal orientation** | **Backward-looking** — historical data | **Forward-looking** — predictive |
| **Purpose** | Did we meet our performance targets? | What is our current risk exposure? |
| **Value** | Reveal what has already occurred; measure achievement | Anticipate future risk shifts and emerging threats |
| **Examples** | Mean time to resolve tickets; patch coverage rate; training completion rate | Unpatched critical vulnerability count; phishing click rate trend; third-party risk scores |

**Memory cue**: KPI = Past (Performance), KRI = Predict (Risk).

## Example Metrics by Area

From destination-cissp §6.3.1 / §6.4 reporting requirements:

### Account Management
- Mean time to resolution (help desk tickets).
- Average response time.
- Number of support emails.
- Last login time, account status, last password change.

### Management Review and Approval
- Time to resolve defects.
- Number of defects identified.
- Defect detection effectiveness.
- Average cost per defect.

### Backup Verification
- Number of backups verified.
- Time between backup verification exercises.
- Amount of data successfully restored in tests.
- Results of full disaster scenario restoration exercises.

### Training and Awareness
- Percentage of employees who completed required training.
- Phishing simulation click rate. *(KPI: historical click rates; KRI: current click rate trend as a predictor of susceptibility)*
- Number of phishing emails reported by employees.

### DR and BC
- Recovery Time Objective (RTO) — maximum acceptable downtime.
- Recovery Point Objective (RPO) — maximum acceptable data loss window.
- Actual time required to restore a critical process.
- Time between plan updates.
- Percentage of critical processes covered by a tested plan.

## Mean Time Metrics

Not covered by name in destination-cissp, but standard CISSP knowledge per cissp-exam-outline:

- **MTTD (Mean Time to Detect)** — average time between an incident occurring and it being detected. Lower is better.
- **MTTR (Mean Time to Respond/Recover)** — average time to respond to and recover from an incident after detection. Lower is better.

These are KPI metrics (backward-looking) that also serve as proxy KRIs when trended over time.

## Audience-Appropriate Reporting (destination-cissp §6.2.8)

Reports must be tailored to the reader:

- **Senior management / CEO** — high-level summary; business impact; key decisions needed. Not interested in technical detail of 1,000 individual test results.
- **Development team** — detailed results; specific vulnerabilities; code locations; remediation steps.
- **Security operations** — technical indicators, alert details, tool output.

Report content should include:
- Objective pass/fail decisions.
- The right level of detail for the right audience.
- Metrics that matter (not every data point).

## Test Output Categories (destination-cissp §6.4.1)

After any security assessment or test, results must address three areas:

| Category | Description |
|---|---|
| **Remediation** | Document specific remediation steps for all identified vulnerabilities. |
| **Exception handling** | Document vulnerabilities that will NOT be remediated, with rationale (cost vs. risk, low asset value, etc.). Exceptions must be formally acknowledged. |
| **Ethical disclosure** | Newly discovered vulnerabilities in widely-used software/hardware must be shared with the broader community (responsible/coordinated disclosure). |

## Cross-links

- [Continuous Monitoring](continuous-monitoring.md)
- [Log Management and SIEM](log-management-siem.md)
- [Audit Types](audit-types.md)

## Sources

- destination-cissp §6.3.1 (KPI vs KRI, SMART metrics, example metric areas)
- destination-cissp §6.2.8 (regression testing, audience-appropriate reporting, metrics that matter)
- destination-cissp §6.4.1 (test output: remediation, exception handling, ethical disclosure)
- cissp-exam-outline (MTTD/MTTR)
