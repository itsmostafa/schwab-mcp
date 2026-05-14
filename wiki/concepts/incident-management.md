---
title: "Incident Management"
type: concept
domain: 7
tags: [incident-response, IR, detection, mitigation, recovery, breach-notification, NIST]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Incident Management

## Definition

Incident management is the process used to detect, respond to, and reduce the impact of security incidents. It encompasses the people, processes, and technology that enable an organization to identify adverse events, contain their damage, and restore normal operations—while learning from each event to strengthen future defenses.

**Key distinction:**
- **Event** — any observable occurrence (the vast majority are benign).
- **Incident** — an adverse event that triggers a formal response. Not all events are incidents.
- **Breach** — a confirmed unauthorized disclosure, acquisition, or loss of protected data (often triggers legal notification requirements).

---

## Incident Response Models

### destination-cissp 8-Step Model

The source uses an eight-step cycle. Each step builds on the prior:

| Step | Focus |
|---|---|
| Preparation | Develop IR process, assign team roles, deploy detection tools, train staff. |
| Detection | Identify that an adverse event is occurring (via IDS/IPS, SIEM, DLP, guards, etc.). |
| Response (IR Team) | Activate IR team; conduct initial impact assessment (scope, duration, stakeholders). |
| Mitigation (Containment) | Limit further damage — isolate affected systems, extinguish the fire, stop the bleeding. Goal is NOT to fix the root cause yet. |
| Reporting | Formal communication to stakeholders (management, legal, HR, PR, regulators, media). One designated spokesperson keeps the message consistent. |
| Recovery | Return operations to normal (replace hardware, restore data, bring systems back online). |
| Remediation | Implement fixes and process improvements to prevent recurrence. |
| Lessons Learned | Holistic review of what happened, what worked, what to improve, what to add or eliminate. |

*Source: destination-cissp §7.6.1*

### NIST SP 800-61 4-Phase Model

NIST's *Computer Security Incident Handling Guide* (SP 800-61 Rev. 2) organizes incident response into four higher-level phases:

| Phase | Key Activities |
|---|---|
| **1. Preparation** | Build IR capability, develop policies/procedures, acquire tools, train personnel. |
| **2. Detection & Analysis** | Identify signs of incidents (precursors, indicators); prioritize by impact; document initial findings. |
| **3. Containment, Eradication & Recovery** | Contain the incident, eradicate the root cause (malware, compromised accounts), recover affected systems. |
| **4. Post-Incident Activity** | Lessons learned meeting, evidence retention, reporting, metric collection for improvement. |

*Source: NIST SP 800-61 Rev. 2 (standard CISSP knowledge; see [nist-sp-800-61.md](../standards/nist-sp-800-61.md)).*

> **Exam note:** Both models are testable. The NIST 4-phase model is authoritative; the 8-step model maps the same content at finer granularity. Preparation appears at the start of both.

---

## IR Team Roles

A typical Computer Security Incident Response Team (CSIRT) includes:

- **Incident manager / team lead** — declares incidents, manages escalation, interfaces with management.
- **Technical analysts** — forensic investigation, malware analysis, log review.
- **Legal / compliance** — advises on breach notification, law enforcement liaison.
- **HR** — when insider threat is involved.
- **PR / communications** — manages external messaging.
- **Senior management / executive sponsor** — authorizes major containment actions (e.g., taking production systems offline).

One person should be designated as the **single point of communication** to stakeholders so the IR team can stay focused.

---

## Escalation

Escalation is triggered when:
- The initial impact assessment shows damage is larger than expected.
- MTD (Maximum Tolerable Downtime) is at risk of being exceeded → declare a **disaster**, activate DRP.
- Legal thresholds are triggered (e.g., personal data exposed).
- Criminal activity is suspected → notify law enforcement.

Failure to escalate properly (even with good detection technology) has caused significant secondary breaches at real organizations (*destination-cissp* recounts a case where two teams detected a breach but alerts were not escalated).

---

## Detection Tools

Organizations should use a combination of automated and manual detection:
- IPS/IDS
- DLP (Data Loss Prevention)
- Anti-malware
- SIEM (with UEBA)
- Administrative review
- Physical controls (motion sensors, cameras, guards)

---

## Breach Notification

When an incident constitutes a **data breach**, breach notification laws may require organizations to notify:
- **Affected individuals** (e.g., GDPR, CCPA, HIPAA Breach Notification Rule)
- **Regulators** (e.g., supervisory authorities under GDPR within 72 hours)
- **Law enforcement** (especially when criminal activity is suspected)

Notification triggers, timelines, and thresholds vary by jurisdiction and regulation. Legal counsel must be involved early.

---

## Exam-Relevant Nuance

- **Containment ≠ remediation.** Containment stops further damage; remediation fixes the root cause. The exam distinguishes them.
- **Lessons learned** is a discrete step, not optional. It feeds back into Preparation.
- **Reporting** happens *throughout* the process (not just at the end), but formal stakeholder reporting occurs after mitigation.
- Incident response and disaster recovery are related but distinct: IR handles incidents; when MTD will be exceeded, IR escalates to DR.
- A **computer security incident** includes malware, hacker attacks, insider attacks, employee errors, system errors, data corruption, and workplace injuries.

---

## Cross-Links

- [NIST SP 800-61](../standards/nist-sp-800-61.md) — detailed incident handling standard
- [Forensics](forensics.md) — digital forensics used during incident investigation
- [Chain of Custody](chain-of-custody.md) — evidence handling during IR
- [Logging & Monitoring](logging-monitoring.md) — detection capability
- [BCP/DRP Operations](bcp-drp-operations.md) — when IR escalates to disaster
- [RTO/RPO/MTD](rto-rpo-mtd.md) — MTD triggers disaster declaration

## Sources

- destination-cissp §7.6–7.6.1 (pp. 0867–0872)
- cissp-exam-outline (Domain 7 subtopic list)
- NIST SP 800-61 Rev. 2 (Computer Security Incident Handling Guide)
