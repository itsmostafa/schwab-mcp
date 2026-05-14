---
title: "Log Management and SIEM"
type: concept
domain: 6
tags: [log-management, siem, ntp, log-review, clipping-levels, circular-overwrite, log-retention]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Log Management and SIEM

## Purpose

Log review and analysis is a best practice that every organization should implement. Logs provide the audit trail to detect errors, unauthorized modifications, and breaches. Without logs, there is no post-incident forensic trail and no ability to correlate attacker activity across multiple systems.

## Log Review Principles (destination-cissp §6.2.5)

Three core principles govern effective log review:

| Principle | Explanation |
|---|---|
| **Log what is relevant** | Not all system output needs to be logged. Risk management drives relevance — log events that would indicate a risk has materialized. |
| **Review the logs** | Logs must actually be reviewed — automated tools (SIEM) are essential when events number in the thousands or millions. |
| **Identify errors and anomalies** | Focus on: unexpected errors (system malfunction), unauthorized modifications (significant red flag, potential breach), and confirmed breach activity. |

## Time Synchronization — NTP

All systems generating logs must use **synchronized clocks**. Without consistent timestamps, correlating events as an attacker moves across multiple systems becomes extremely difficult (or impossible) during incident response.

**Network Time Protocol (NTP)** is used to synchronize clocks. Typically:
1. One or more network devices are synchronized to a publicly available atomic/nuclear clock (e.g., NIST).
2. All other network devices synchronize from those primary NTP sources.
3. Redundancy is recommended (at least two NTP sources).

(destination-cissp §6.2.5)

## Log File Management — Controlling Size

### Circular Overwrite

When the log file reaches its configured maximum size (or entry count), **circular overwrite** begins writing from the beginning, overwriting the oldest entries first. Result: log file never exceeds the configured limit.

- **Advantage**: Prevents disk exhaustion; guaranteed bounded storage use.
- **Disadvantage**: Old entries are deleted. If a breach occurred in the early part of the log window, that evidence may be overwritten. **Not recommended where forensic preservation matters.**

### Clipping Levels

**Clipping levels** define a threshold before an event is logged. Events below the threshold are not recorded; only events that cross the threshold trigger a log entry.

Example: Log a failed login only after 15 consecutive failures (threshold). Single failed logins are noise; 15+ in a row indicate a potential brute-force attack or credential-stuffing attempt.

- **Advantage**: Logs remain focused on meaningful, anomalous events. Log file size is limited by relevance, not brute truncation. Unlike circular overwrite, clipping levels do **not** delete data.
- **Disadvantage**: May miss early-stage reconnaissance activity that never reaches the threshold.

**Which is better for security?** Clipping levels — relevant breach-related entries are preserved; circular overwrite may destroy them. (destination-cissp §6.2.6)

## Log Data Lifecycle

The source notes (destination-cissp §6.2.5) that logging and monitoring includes these stages (covered further in Domain 7):

> Generation → Transmission → Collection → Normalization → Analysis → Retention → Disposal

## SIEM — Security Information and Event Management

A SIEM is the operational tool for log aggregation, correlation, alerting, and dashboarding at scale. It is not covered in depth by this source (destination-cissp Domain 6) — Domain 7 addresses SIEM and monitoring further.

Key SIEM capabilities (general CISSP knowledge):
- **Centralized log collection** from heterogeneous sources.
- **Normalization** — converts diverse log formats into a common schema.
- **Correlation rules** — detects patterns across multiple events (e.g., failed login + privilege escalation + data exfiltration).
- **Alerting** — triggers notifications or automated responses.
- **Dashboards** — provides visibility into security posture.
- **Log retention** — stores log history for compliance and forensics.

## Real User Monitoring (RUM) vs Synthetic Transactions (destination-cissp §6.2.7)

Two operational testing techniques that relate to log monitoring:

| Technique | Type | Description |
|---|---|---|
| **Real User Monitoring (RUM)** | Passive | Monitors actual user interactions and transactions with a live app in real time. Often consumes log streams. |
| **Synthetic Performance Monitoring** | Active | Runs scripted ("synthetic/fake") transactions against a system to test functionality and performance under controlled or load conditions. Best run in production (after full QA in lower environments). |

## Cross-links

- [Security Metrics](security-metrics.md) — KPI/KRI metrics including MTTD/MTTR
- [Continuous Monitoring](continuous-monitoring.md)
- [Assessment vs Audit](assessment-vs-audit.md)

## Sources

- destination-cissp §6.2.5 (log review principles, NTP time synchronization)
- destination-cissp §6.2.6 (circular overwrite vs clipping levels)
- destination-cissp §6.2.7 (RUM vs synthetic performance monitoring)
- cissp-exam-outline (domain 6 subtopics)
