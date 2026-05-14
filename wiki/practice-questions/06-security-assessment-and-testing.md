---
title: "Practice Questions — Domain 06: Security Assessment And Testing"
type: practice
domain: 06
tags: [practice, domain-6, assessment, audit, penetration-testing, soc, sast, dast, cvss, log-management, metrics]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 06: Security Assessment And Testing

## Questions

### Q1: What is the single step that differentiates a penetration test from a vulnerability assessment?

**Answer:** Exploitation. A penetration test actively attempts to exploit identified vulnerabilities to confirm they are real (true positives); a vulnerability assessment stops at identifying and reporting vulnerabilities without attempting exploitation.

**Why:** Both processes begin identically — reconnaissance, enumeration, vulnerability analysis — but the pen test goes one critical step further by attempting to breach the system. This is the exam's most-tested distinction in Domain 6.

*Subtopic: 6.2 — [Penetration Testing](../concepts/penetration-testing.md)*

---

### Q2: A company engages an independent audit firm to audit its cloud provider's controls. The audit firm's report is then shared with the company's customers. What type of audit is this?

**Answer:** Third-party audit. Three parties are involved: the customer, the service provider (cloud provider), and the independent audit firm. The service provider commissions and pays for the audit, then shares the report with customers as evidence of security assurance.

**Why:** The third-party audit model is prevalent in cloud computing. It is distinct from an internal audit (employees auditing their own org) and from an external audit (external firm auditing the requesting org directly). The three-party structure — customer / vendor / independent auditor — defines third-party.

*Subtopic: 6.5 — [Audit Types](../concepts/audit-types.md)*

---

### Q3: A security team wants to assess the operating effectiveness of a cloud provider's security controls over the past year, not just their design. Which SOC report type should they request?

**Answer:** SOC 2, Type 2.

**Why:** SOC 2 focuses on the five Trust Services Criteria (security, availability, confidentiality, processing integrity, privacy) — directly relevant to security professionals. Type 2 examines controls over a period of time (typically ~1 year) and attests to both design AND operating effectiveness. Type 1 only attests to design at a point in time. SOC 3 is a public summary insufficient for compliance. SOC 1 covers financial controls only.

*Subtopic: 6.5 — [SOC Reports](../standards/soc-reports.md)*

---

### Q4: Which type of SOC 2 report would a new cloud provider most likely commission in their first year of undergoing third-party audits, and why?

**Answer:** SOC 2, Type 1.

**Why:** A Type 1 report evaluates control design at a point in time — it does not require demonstrating sustained operating effectiveness. For a provider just beginning their audit program, controls are likely still maturing and gaps may exist. A Type 1 lets the auditor identify design gaps for remediation before committing to the more rigorous Type 2 (which covers a full year of operational evidence). The typical progression is Type 1 in Year 1, Type 2 from Year 2 onward.

*Subtopic: 6.5 — [SOC Reports](../standards/soc-reports.md)*

---

### Q5: A developer runs an automated tool that analyzes source code for SQL injection vulnerabilities without executing the application. What type of testing is this?

**Answer:** Static Application Security Testing (SAST) — white box testing.

**Why:** SAST examines source code while the application is not running (static = not executing). Because the source code is visible, it is white box testing. It finds code-level vulnerabilities (SQL injection, buffer overflows, hardcoded credentials) before runtime. DAST, by contrast, requires a running application and does not have access to source code (black box).

*Subtopic: 6.2 — [Security Testing Types](../concepts/security-testing-types.md)*

---

### Q6: During a vulnerability scan, the tool reports a critical vulnerability in a web server that has already been patched. What type of result is this, and which type is considered more dangerous?

**Answer:** This is a **false positive** (the vulnerability is reported but does not exist). False **negatives** are more dangerous — when a scanner reports no vulnerability but one actually exists, the organization remains exposed without knowing it.

**Why:** False positives create noise and administrative overhead but don't leave systems vulnerable. False negatives are far more serious: they give a false sense of security while real vulnerabilities remain undetected and unpatched. Credentialed scans help reduce false positives.

*Subtopic: 6.2 — [Vulnerability Assessment](../concepts/vulnerability-assessment.md)*

---

### Q7: A security team is reviewing logs after a breach and finds that timestamps across servers don't match, making it impossible to sequence the attacker's movements. What protocol should have been implemented to prevent this?

**Answer:** Network Time Protocol (NTP).

**Why:** NTP synchronizes clocks across all network devices to ensure consistent timestamps in log entries. Without synchronized timestamps, correlating events across multiple systems during incident response becomes extremely difficult or impossible. At least one device should sync to a government-maintained atomic clock (e.g., NIST), with all other devices syncing from that source. Redundancy (two+ NTP sources) is recommended.

*Subtopic: 6.2 — [Log Management and SIEM](../concepts/log-management-siem.md)*

---

### Q8: A CISO wants to implement a log management approach that preserves evidence of security breaches while limiting log file size. Should they use circular overwrite or clipping levels?

**Answer:** Clipping levels.

**Why:** Circular overwrite limits file size by deleting the oldest log entries as new ones arrive — breach-related evidence in the early portion of the log window may be overwritten and permanently lost. Clipping levels limit log size by only recording events that exceed a defined threshold (e.g., only log failed logins after 15 consecutive attempts). Unlike circular overwrite, clipping levels do not delete data — relevant events are preserved for forensic investigation.

*Subtopic: 6.2 — [Log Management and SIEM](../concepts/log-management-siem.md)*

---

### Q9: A penetration tester is given only the company name and told the internal security team has not been notified of the test. What approach and knowledge level describes this engagement?

**Answer:** **Double-blind** approach with **zero knowledge** (black box).

**Why:** Double-blind means both the tester has minimal information AND the internal IT/security team does not know a test is coming — only senior management (who commissioned it) knows. This tests both the tester's skill in discovering vulnerabilities and the team's ability to detect and respond to an actual attack. Zero knowledge / black box means the tester starts with no information about the target network and must use reconnaissance and enumeration to gain visibility.

*Subtopic: 6.2 — [Penetration Testing](../concepts/penetration-testing.md)*

---

### Q10: What is the difference between a KPI and a KRI? Give an example of each in a security context.

**Answer:** A **KPI (Key Performance Indicator)** is backward-looking — it measures whether past performance targets were achieved. A **KRI (Key Risk Indicator)** is forward-looking — it predicts current and emerging risk exposure.

Examples:
- **KPI**: "85% of employees completed security awareness training last quarter" — historical achievement.
- **KRI**: "The phishing simulation click rate has increased 3% month-over-month" — predicts growing susceptibility to phishing attacks (a future risk).

**Why:** The backward/forward distinction is the exam's key test point. KPIs tell you what already happened; KRIs tell you where risk is heading. Both should be part of a security metrics program using SMART criteria (Specific, Measurable, Achievable, Relevant, Timely).

*Subtopic: 6.3 — [Security Metrics](../concepts/security-metrics.md)*
