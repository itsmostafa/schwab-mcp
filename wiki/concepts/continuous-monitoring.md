---
title: "Continuous Monitoring"
type: concept
domain: 6
tags: [continuous-monitoring, nist-sp-800-137, siem, vulnerability-management, configuration-management, compliance]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Continuous Monitoring

## Definition

Continuous monitoring is the ongoing, real-time (or near-real-time) observation of an organization's security posture, with the goal of detecting security events, configuration drift, compliance deviations, and new vulnerabilities as quickly as possible after they arise. It is the operational layer that bridges point-in-time assessments and real-world security posture.

**Standard reference**: **NIST SP 800-137** — *Information Security Continuous Monitoring (ISCM) for Federal Information Systems and Organizations* — defines continuous monitoring for federal environments. While the destination-cissp source does not cite 800-137 by name, it is the authoritative NIST reference for this topic per cissp-exam-outline.

## Why Continuous Monitoring

Point-in-time assessments (vulnerability scans, audits) capture a snapshot. Between assessments:
- New vulnerabilities emerge (CVE disclosures happen daily).
- Configurations drift from approved baselines.
- New assets are added that may not yet be inventoried.
- Threats evolve and new attack techniques emerge.

Continuous monitoring closes these gaps by automating detection across the assessment cycle.

## Core Components

### SIEM (Security Information and Event Management)

Aggregates and correlates log data from across the environment. See [Log Management and SIEM](log-management-siem.md) for full detail. Key role in continuous monitoring:
- Real-time correlation of events across systems.
- Alert generation for anomalous or policy-violating activity.
- Dashboard visibility into security posture.

### Vulnerability Scanners

Automated tools (Nessus, Qualys, InsightVM, OpenVAS) run on a recurring schedule against assets. More frequent scanning for high-sensitivity assets. Feeds the vulnerability management lifecycle. See [Vulnerability Assessment](vulnerability-assessment.md).

### Configuration Management and Drift Detection

Continuous monitoring includes verifying that systems remain in their approved, baseline-compliant configuration. Configuration drift — when a system deviates from its baseline — can introduce vulnerabilities or indicate unauthorized change.

The source notes (destination-cissp §6.2.0, Operate phase): "Configuration management reviews can be performed to ensure the product is working as intended without its security being compromised."

### SOAR (Security Orchestration, Automation, and Response)

Not covered in destination-cissp Domain 6 source material. Covered in cissp-exam-outline / Domain 7 scope:
- Automates repetitive security operations tasks.
- Orchestrates responses across multiple tools (SIEM, ticketing, firewalls).
- Reduces mean time to respond (MTTR).

## Assessment Frequency

The appropriate frequency of assessments should be proportional to asset sensitivity and criticality (destination-cissp §6.1.2). High-value, high-sensitivity assets warrant more frequent assessment, tighter monitoring thresholds, and faster response SLAs.

## Regression Testing and Continuous Monitoring

**Regression testing** is the process of verifying that previously working software continues to work after updates (patches, enhancements). It is a component of continuous assurance: every change should trigger regression tests to confirm no controls were broken (destination-cissp §6.2.8).

After regression testing, results must be reported in a manner appropriate to the audience (see [Security Metrics](security-metrics.md)).

## Compliance Checks (destination-cissp §6.2.9)

Compliance checking is an integral, ongoing component of security control testing. It confirms:
- Implemented controls align with documented security requirements.
- Controls align with organizational policies, standards, procedures, and baselines.

Compliance checks provide the linkage between technical security testing and policy/regulatory obligations.

## Cross-links

- [Log Management and SIEM](log-management-siem.md)
- [Vulnerability Assessment](vulnerability-assessment.md)
- [Security Metrics](security-metrics.md)
- [Assessment vs Audit](assessment-vs-audit.md)

## Sources

- destination-cissp §6.2.0 (software testing lifecycle phases — Operate phase)
- destination-cissp §6.2.8 (regression testing)
- destination-cissp §6.2.9 (compliance checks)
- destination-cissp §6.1.2 (effort proportional to asset value)
- cissp-exam-outline (NIST SP 800-137, SOAR, continuous monitoring)
