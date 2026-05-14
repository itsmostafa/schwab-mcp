---
title: "Configuration Management"
type: concept
domain: 7
tags: [configuration-management, CM, CMDB, baseline, hardening, provisioning, asset-management, change-control]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Configuration Management

## Definition

**Configuration management (CM)** is the discipline of ensuring that all hardware and software assets are deployed, tracked, and maintained in a known, secure, and consistent state. It encompasses asset inventory, secure provisioning, baseline configuration, and ongoing verification.

---

## Key Components

### 1. Asset Inventory (CMDB)

A **Configuration Management Database (CMDB)** maintains the authoritative inventory of all assets:
- What the asset is (hardware/software)
- Who owns it (designated owner accountable for patching, licensing, configuration)
- Current version, patch level, configuration state

**Without an accurate inventory, assets cannot be patched, scanned, or secured.** Assets represent the organization's attack surface — each undocumented asset is a potential blind spot.

Asset management lifecycle:
1. Plan / procure
2. Securely provision (harden before deployment)
3. Deploy and document
4. Monitor and maintain (patching, config reviews)
5. Decommission and sanitize

*Source: destination-cissp §7.3.1*

### 2. Secure Provisioning and Baseline Configuration

**Provisioning** = how hardware, software, and devices are deployed. Secure provisioning means:
- **Never deploy with default settings** — default credentials and configurations are well-known to attackers.
- Apply **system hardening**: disable unnecessary services, close unused ports, remove default accounts, enable logging.
- Align to a documented **baseline configuration** — the approved security state for a given device type.
- Update the CMDB immediately upon deployment.

**Hardening guides** that establish baseline configurations:
- **CIS Benchmarks** (Center for Internet Security) — prescriptive secure configuration guides for common OS, databases, and applications.
- DISA STIGs (Security Technical Implementation Guides) — DoD hardening standards.
- Vendor-specific hardening guidance.

### 3. Automation

In large environments, **automated provisioning tools** (e.g., Ansible, Puppet, Chef, SCCM, Intune) ensure:
- Consistency — every device of a given type is configured identically.
- Speed — no manual step-by-step configuration.
- Auditability — configuration state is recorded and can be compared to the baseline.

---

## Configuration Drift

**Configuration drift** occurs when a system's actual configuration diverges from its approved baseline over time (unauthorized changes, manual tweaks, failed patches). Drift introduces vulnerabilities.

Detection methods:
- **Credentialed vulnerability scans** — log into systems and check actual configuration vs. expected.
- **File integrity monitoring (FIM)** — hashes key OS files and detects unauthorized changes.
- **Configuration compliance tools** — continuously compare running configs to stored baselines.

---

## CM Process Steps

1. **Identify** all assets to bring under CM control.
2. **Configure** assets to the approved baseline.
3. **Document** the configuration (store in CMDB).
4. **Verify** configuration is accurate and compliant.

Any change to a configuration must go through the **[Change Management](change-management.md)** process.

---

## Change Control Board (CCB)

The **Change Control Board (CCB)** (also called Change Advisory Board / CAB) is a cross-functional group that reviews, approves, and tracks changes to the baseline configuration. Membership typically includes IT, security, business owners, and operations. The CCB prevents unauthorized configuration changes that could introduce vulnerabilities.

---

## Exam-Relevant Nuance

- CM and change management are related but distinct: CM maintains the approved state; change management controls the process for modifying that state.
- **Configuration drift** is a primary root cause of vulnerabilities in real-world environments.
- The **change management process** must include CM updates (version and baseline step) — a change without updating the CM baseline leaves the CMDB stale.
- Software licensing is an explicit CM responsibility: each software asset's license must be current and managed.

---

## Cross-Links

- [Change Management](change-management.md) — process for authorizing configuration changes
- [Patch Management](patch-management.md) — patching is a subset of CM
- [Logging & Monitoring](logging-monitoring.md) — configuration changes should be logged and monitored
- [Vulnerability Management Operations](vulnerability-management-ops.md) — vuln scans detect configuration weaknesses

## Sources

- destination-cissp §7.3.1–7.3.2 (pp. 0857–0861)
- cissp-exam-outline (Domain 7: configuration management — provisioning, baselining, automation)
