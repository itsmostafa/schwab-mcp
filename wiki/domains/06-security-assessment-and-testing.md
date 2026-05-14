---
title: "Domain 6 — Security Assessment and Testing"
type: domain
domain: 6
tags: [assessment, testing, audit, penetration-testing, soc, metrics, sast, dast, cvss, cve, log-management]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 6 — Security Assessment and Testing

Exam weight: **12%**. Covers designing and executing security assessments, penetration tests,
audits, and continuous monitoring programs.

## Subtopics (from CISSP Exam Outline)

### 6.1 Design and validate assessment, test, and audit strategies

- Internal (within organization control)
- External (outside organization control)
- Third-party (outside enterprise control)

**Purpose**: Security assessment and testing provides *assurance* to stakeholders that security controls are defined, tested, and operating effectively. Two pillars support every security control: functional (it does the job) and assurance (we can prove it does the job). Domain 6 focuses on the assurance pillar. Testing applies to the full system lifecycle — development, operation, and retirement — not just at initial deployment.

**Validation vs Verification**: Two foundational concepts (destination-cissp §6.1.1):
- **Validation** — "Are we building the *right* product?" Begins before development; ensures requirements are correctly understood.
- **Verification** — "Are we building the product *correctly*?" Confirms the product meets its requirements. Grounded in the 3 Cs: **Completeness**, **Correctness**, **Consistency**.

**Effort proportionality**: Testing investment should be proportional to the value the asset or application represents to the organization.

**Strategies** — Internal, external, and third-party (see [Assessment vs Audit vs Penetration Test](../concepts/assessment-vs-audit.md)):
- **Internal**: employees audit the organization's own controls.
- **External**: internal employees examine a vendor, or an external firm independently assesses the organization.
- **Third-party**: independent audit firm examines a service provider; results shared with the service provider's customers (dominant model in cloud computing).

**Location dimension**: On-premises, cloud, or hybrid — separate from who conducts the audit.

**Role of security professional**: Identify risk; advise testing processes; support stakeholders. The security team does not test alone.

### 6.2 Conduct security control testing

- Vulnerability assessment
- Penetration testing (red, blue, and/or purple team exercises)
- Log reviews
- Synthetic transactions / benchmarks
- Code review and testing
- Misuse case testing
- Interface testing (UI, network interface, API)
- Breach attack simulations
- Compliance checks

**Software testing progression** (destination-cissp §6.2.0): Unit testing → Interface testing → Integration testing → System testing. Testing is required at every SDLC phase: Planning, Design, Develop, Deploy, Operate, Retire.

**Vulnerability Assessment vs Penetration Test** — the critical distinction:

| | Vulnerability Assessment | Penetration Test |
|---|---|---|
| Exploitation step | **No** | **Yes** (the key differentiator) |
| Primarily automated | Yes | No (manual-driven) |
| Duration | Minutes–days | Several days |

Both share: defined scope, activity schedule, formal approval requirement, and a reporting phase. See [Vulnerability Assessment](../concepts/vulnerability-assessment.md) and [Penetration Testing](../concepts/penetration-testing.md).

**Pen test phases** (destination-cissp §6.2.2): Reconnaissance (passive) → Enumeration (active) → Vulnerability Analysis → Exploitation → Reporting. The fork: vulnerability assessments stop at vulnerability analysis; pen tests continue to exploitation.

**Red/Blue/Purple Teams**:
- **Red team**: offensive; simulates attacks using pen testing, social engineering, threat intelligence.
- **Blue team**: defensive; monitors, responds to incidents, evaluates posture, recommends mitigations.
- **Purple team**: collaboration between red and blue; promotes information sharing; not a separate team.

**Testing techniques** (destination-cissp §6.2.1):
- **SAST** — Static; white box; app not running; source code visible. See [Security Testing Types](../concepts/security-testing-types.md).
- **DAST** — Dynamic; black box; app running; behavior-focused.
- **Fuzz testing** — Dynamic; chaos-based; mutation (dumb) vs generation (intelligent/smart) fuzzers.
- **IAST / RASP** — Not in destination-cissp source; see cissp-exam-outline.

**Test types**: Positive (normal usage), Negative (normal errors), Misuse (attacker perspective).

**Test efficiency techniques**: Boundary value analysis (testing at behavioral boundaries); Equivalence partitioning (group inputs by behavior, test a representative value per partition).

**Code review**: White box (source code access) vs black box (behavior only). See [Code Review and Security](../concepts/code-review-security.md).

**Vulnerability scanning**:
- Credentialed scans: deeper; fewer false positives; baseline compliance checks.
- Non-credentialed scans: attacker-perspective; more false positives.
- CVE identifies the vulnerability; CVSS scores its severity (0–10, Critical ≥ 9.0). See [CVSS](../standards/cvss.md).
- False positives: scanner reports vuln that doesn't exist. False negatives: scanner misses a real vuln — **worse**.

**Log management** (destination-cissp §6.2.5–6.2.6):
- Log what is relevant, review it, identify errors and anomalies.
- NTP synchronizes clocks for timestamp consistency across systems — critical for breach correlation.
- **Circular overwrite**: auto-deletes oldest entries when size limit reached. May destroy breach evidence.
- **Clipping levels**: threshold-based logging; preserves relevant events; does not delete data. Better for security.
See [Log Management and SIEM](../concepts/log-management-siem.md).

**Operational testing** (destination-cissp §6.2.7):
- **RUM (Real User Monitoring)**: passive; monitors actual user interactions.
- **Synthetic performance monitoring**: scripted ("fake") transactions; tests functionality and performance under load.

**Regression testing** (destination-cissp §6.2.8): Verifies software still works after patches or enhancements. Reports must use "metrics that matter" appropriate to the audience (CEO vs. dev team).

**Compliance checks** (destination-cissp §6.2.9): Ongoing confirmation that controls align with documented security requirements and organizational policies/standards/baselines.

### 6.3 Collect security process data (technical and administrative)

- Account management
- Management review and approval
- Key performance and risk indicators
- Backup verification data
- Training and awareness
- Disaster recovery (DR) and Business Continuity (BC)

**KPIs vs KRIs** (destination-cissp §6.3.1):
- **KPI (Key Performance Indicator)** — backward-looking; measures achievement of past performance targets.
- **KRI (Key Risk Indicator)** — forward-looking; predicts risk exposure and emerging risk conditions.

**SMART metrics**: Specific, Measurable, Achievable, Relevant, Timely.

See [Security Metrics](../concepts/security-metrics.md) for example metrics by area (account management, backup verification, training and awareness, DR/BC).

### 6.4 Analyze test output and generate report

- Remediation
- Exception handling
- Ethical disclosure
- Coverage analysis

Three required outputs from any security assessment (destination-cissp §6.4.1):
1. **Remediation** — documented remediation steps for all identified vulnerabilities.
2. **Exception handling** — formal documentation of vulnerabilities that will NOT be remediated, with rationale.
3. **Ethical disclosure** — newly discovered vulnerabilities in widely-used software must be shared with the broader community (responsible/coordinated disclosure).

**Coverage analysis**: percentage of source code exercised by tests. Formula: covered lines / total lines.

### 6.5 Conduct or facilitate security audits

- Internal (within organization control)
- External (outside organization control)
- Third-party (outside of enterprise control)
- Location (on-premises, cloud, hybrid)

**Audit process** (destination-cissp §6.5.1): Define objective → Define scope → Involve stakeholders → Choose team → Plan → Conduct → Document results → Communicate.

**SOC Reports** (destination-cissp §6.5.2): Produced under SSAE 18 (US standard; AICPA governing body). International equivalent: ISAE 3402.

| | SOC 1 | SOC 2 | SOC 3 |
|---|---|---|---|
| Focus | Financial controls | 5 Trust Services Criteria | SOC 2 summary (public) |
| Distribution | Restricted | Restricted | Public |
| Has Type 1/2? | Yes | Yes | No |

**SOC 2 Trust Services Criteria**: Security (required), Availability (required), Confidentiality (required), Processing Integrity (optional), Privacy (optional).

**Type 1 vs Type 2**:
- Type 1 = design of controls at a point in time.
- Type 2 = design + operating effectiveness over a period of time (~1 year). **More comprehensive. SOC 2 Type 2 is the gold standard.**

Typical progression: SOC 2 Type 1 (Year 1) → SOC 2 Type 2 (Year 2+).

**Audit roles** (destination-cissp §6.5.3): Executive management (sets tone), Audit Committee (oversight), Security Officer (security risk advisory), Compliance Manager (scheduling, hiring auditors), Internal Auditors (employees), External Auditors (independent).

See [Audit Types](../concepts/audit-types.md) and [SOC Reports](../standards/soc-reports.md).

## Key concepts

- [Assessment vs Audit vs Penetration Test](../concepts/assessment-vs-audit.md) — strategies, validation/verification, security professional role
- [Vulnerability Assessment](../concepts/vulnerability-assessment.md) — scanning, credentialed/uncredentialed, CVE/CVSS, false positives/negatives, vuln management lifecycle
- [Penetration Testing](../concepts/penetration-testing.md) — phases, red/blue/purple teams, perspective/approach/knowledge variables
- [Security Testing Types](../concepts/security-testing-types.md) — SAST vs DAST vs fuzz, manual vs automated, SDLC phases, coverage analysis
- [Code Review and Security](../concepts/code-review-security.md) — white box vs black box code review, SAST, fuzz testing
- [Log Management and SIEM](../concepts/log-management-siem.md) — log review principles, NTP, circular overwrite vs clipping levels, RUM vs synthetic monitoring
- [Audit Types](../concepts/audit-types.md) — internal/external/third-party, SOC 1/2/3, Type 1/Type 2, audit roles, SSAE 18
- [Security Metrics](../concepts/security-metrics.md) — KPI vs KRI, SMART metrics, MTTD/MTTR, test output (remediation/exceptions/disclosure)
- [Continuous Monitoring](../concepts/continuous-monitoring.md) — NIST SP 800-137, regression testing, compliance checks, configuration drift
- [CVSS](../standards/cvss.md) — v3.1 score ranges, Base/Temporal/Environmental metrics
- [SOC Reports](../standards/soc-reports.md) — SSAE 18, AICPA, Trust Services Criteria, Type 1/Type 2 comparison table

## Sources

- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §6.1–6.5 (full domain coverage, ingested 2026-05-13)
