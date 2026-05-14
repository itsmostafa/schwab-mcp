---
title: "Logging and Monitoring"
type: concept
domain: 7
tags: [SIEM, logging, monitoring, UEBA, SOAR, threat-intelligence, continuous-monitoring, correlation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Logging and Monitoring

## Definition

Security logging and monitoring is the practice of collecting, storing, correlating, and analyzing event data from systems and networks to detect security incidents, support investigations, and demonstrate compliance. Effective monitoring requires technology, well-tuned processes, and trained personnel.

---

## What to Log

High-priority events for security logging:
- **Authentication events** — logins, logouts, failed attempts, account lockouts
- **Privilege changes** — privilege escalation, group membership changes, admin account usage
- **Policy violations** — access control violations, firewall deny events
- **System events** — startup/shutdown, service start/stop, patch installation
- **Object access** — file/folder access, database queries, registry changes
- **Network events** — firewall logs, IDS/IPS alerts, DNS queries, web proxy logs
- **Application events** — errors, configuration changes, user actions in critical apps

Logs should be protected from tampering (stored in read-only format in the SIEM) and retained per regulatory requirements.

---

## Security Information and Event Management (SIEM)

A **SIEM system** ingests logs from disparate sources across the organization, normalizes and correlates them, and alerts analysts to suspicious patterns that individual systems would miss.

### SIEM Capabilities

| Capability | Description |
|---|---|
| **Aggregation** | Collects events from firewalls, IDS/IPS, servers, apps, DLP, etc. under one umbrella. |
| **Normalization** | Converts different log formats (12-hour vs 24-hour clocks, date formats) into a common schema. Includes deduplication. |
| **Correlation** | Links events across systems and time to reveal patterns (e.g., same IP logging into two different accounts simultaneously). |
| **Secure storage** | Retains all log events, ideally read-only, for extended periods. |
| **Analysis & reporting** | Rules-based alerting; dashboards; compliance reporting. |
| **Threat intelligence** | Many SIEMs offer threat feed subscriptions (STIX/TAXII), US-CERT alerts, ISAC feeds. |

### Key Exam Point

SIEM systems are **technology + process + people**. Without trained analysts who understand escalation procedures, even a perfectly configured SIEM fails to prevent breaches. The example in source: two teams detected a breach, but the alert was not properly escalated → ignored.

*Source: destination-cissp §7.2.1*

---

## User and Entity Behavior Analytics (UEBA)

**UEBA** (also called UBA) uses machine learning to:
1. Establish a **behavioral baseline** for each user and entity (device, application).
2. Detect **anomalies** — deviations from normal patterns.
3. Trigger alerts when anomalies cross a threshold.

Use cases: detecting insider threats, compromised privileged accounts, brute-force attacks, data exfiltration.

UEBA is typically bundled with SIEM or available as a subscription add-on.

*Source: destination-cissp §7.2.1*

---

## Security Orchestration, Automation, and Response (SOAR)

**SOAR** takes inputs from disparate sources (SIEM, email, manual reports) and applies workflows to orchestrate an automated or semi-automated response.

Three focus areas:
1. **Threat and vulnerability management**
2. **Incident response automation**
3. **Security operations automation**

SOAR tools include: incident/threat intelligence management, reporting, data analytics, and ML-powered assistance for SOC analysts. Goal: consistent, faster response while reducing analyst fatigue.

*Source: destination-cissp §7.2.3*

---

## Continuous Monitoring

Setting up a SIEM is not a one-time task. **Continuous monitoring** requires:
- Updating rules as new threats emerge
- Adding monitoring for new assets as the environment changes
- Tuning to reduce false positives without creating false negatives
- Reviewing and updating after each incident

The continuous monitoring lifecycle: Define → Establish → Implement → Analyze/Report → Respond → Review/Update.

---

## NTP and Time Synchronization

Accurate time across all systems is critical for log correlation. All systems must use **NTP (Network Time Protocol)** to synchronize clocks. Without it, log timestamps may not align, making attack reconstruction impossible.

---

## Threat Intelligence

Sources for actionable threat intelligence:
- Internal: internal security assessments, previous incidents
- Vendor trend reports
- US-CERT, national CERTs
- **ISACs** (Information Sharing and Analysis Centers) — sector-specific threat sharing
- SIEM threat intelligence subscription feeds (STIX/TAXII format)
- Open-source intelligence (OSINT)

---

## Insider Threat Detection

UEBA and SIEM correlation are primary tools for detecting insider threats because insiders use legitimate credentials. Indicators:
- Unusual data access patterns (accessing files outside normal role)
- Access at unusual hours or from unusual locations
- Large-scale data downloads or file transfers
- Accessing sensitive systems during off-hours
- Anomalous login patterns (same IP for two users — destination-cissp example)

---

## Cross-Domain Notes

- **IDS/IPS** and **egress monitoring** are covered in Domain 4. SIEM ingests alerts from IDS/IPS.
- **Log management** basics (what to log, log retention) are introduced in Domain 6. This page covers the SIEM/UEBA layer on top.

---

## Exam-Relevant Nuance

- SIEM **correlation** is the key differentiator from simple log collection — it connects events across systems.
- **False positives** (legitimate activity flagged as malicious) and **false negatives** (malicious activity not flagged) must be continuously balanced.
- SIEM data should be captured based on **organizational risk, budget, and regulatory obligations** — not everything needs to go into the SIEM.
- Log aggregation sources: security appliances, network devices, DLP, data activity monitors, applications, OS, servers, IPS/IDS.

---

## Cross-Links

- [Incident Management](incident-management.md) — SIEM/monitoring feeds the detection phase of IR
- [Forensics](forensics.md) — log files are key forensic evidence sources
- [Configuration Management](configuration-management.md) — configuration changes should be logged
- [Vulnerability Management Operations](vulnerability-management-ops.md) — vulnerability scans feed SIEM

## Sources

- destination-cissp §7.2.1–7.2.3 (pp. 0844–0856)
- cissp-exam-outline (Domain 7: SIEM, continuous monitoring, UEBA, egress monitoring, threat intelligence)
