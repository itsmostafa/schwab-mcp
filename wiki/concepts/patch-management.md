---
title: "Patch Management"
type: concept
domain: 7
tags: [patch-management, vulnerability-management, patching, WSUS, EOL, hotfix]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Patch Management

## Definition

**Patch management** is a proactive process to create a consistently configured environment that is secure against **known vulnerabilities**. Patches fix security flaws, bugs, and vulnerabilities in software and firmware; they can also improve performance and add functionality.

> **Critical limitation:** Patching only protects against *known* vulnerabilities. Zero-day vulnerabilities are not addressed by patches until a fix is developed and released.

---

## Patch Terminology

| Term | Meaning |
|---|---|
| **Patch** | A code update that fixes one or more vulnerabilities or bugs. |
| **Hotfix / Quick Fix** | A small, urgent fix for a specific issue, often released outside a normal release cycle. Often applied without full testing. |
| **Service pack** | A collection of patches, hotfixes, and sometimes new features rolled up into a single release. |
| **Firmware update** | A patch to device firmware (BIOS, network device OS, embedded systems). |
| **End-of-life (EOL)** | Software or hardware the vendor no longer supports — no more patches. EOL systems carry permanent, unmitigable vulnerability windows. |

---

## Patch Management Lifecycle

```
Identify need → Assess → Test → Approve (change mgmt) → Deploy → Verify → Document
```

1. **Identify** — threat intelligence feeds, vendor notifications, vulnerability scans reveal missing patches.
2. **Assess** — evaluate criticality (CVSS score), impact on the environment, whether the vulnerable asset is exploitable.
3. **Test** — test patch in a non-production environment before deploying to production.
4. **Approve** — route through the change management process. Emergency patches may use an expedited approval path.
5. **Deploy** — apply patches (manual or automated).
6. **Verify** — confirm patches were successfully applied (re-scan, agent check).
7. **Document** — update CMDB/asset inventory with new patch level.

*Source: destination-cissp §7.8.1 (Fig. 7-8)*

---

## Determining Patch Levels

Three methods for assessing patch level of systems:

| Method | How It Works |
|---|---|
| **Agent-based** | A lightweight agent installed on each host monitors installed software/patches and compares to a master database. Typically triggers automatic updates. |
| **Agentless** | Monitoring software remotely connects to each host to check patch levels (no local install). |
| **Passive detection** | Monitors network traffic and fingerprints OS/application versions to infer patch levels. Less intrusive but less accurate. |

*Source: destination-cissp §7.8.1 (Table 7-8)*

---

## Deploying Patches

| Method | When to Use |
|---|---|
| **Automated** (e.g., WSUS for Windows, SCCM, Intune) | Standard/low-criticality systems — ensures consistent deployment at scale. |
| **Manual** | High-value, high-priority production systems — patching sometimes breaks things; manual gives better control and allows immediate rollback. |

**Windows Server Update Services (WSUS)** — Microsoft tool for managing and distributing Windows patches centrally.

**SCCM (System Center Configuration Manager)** — broader Microsoft endpoint management tool that includes patching.

---

## Emergency Patches

When a critical vulnerability is actively exploited (zero-day going mainstream), the **normal change management window** may be too slow. Organizations should define an **emergency change procedure** that allows accelerated approval while still maintaining documentation and rollback plans.

---

## Patch vs. Vulnerability Window

The **vulnerability window** is the time between a vulnerability becoming known (or a patch being released) and the patch being successfully deployed. Attackers actively exploit this window. Minimizing the window (shorter assessment-to-deployment cycle) reduces risk.

---

## End-of-Life Systems

EOL systems receive no vendor patches. Organizations that cannot immediately decommission them should:
- **Compensating controls**: network segmentation, host-based firewalls, enhanced monitoring
- **Formal exception documentation** with accepted risk
- **Active decommissioning plan** with target date

---

## Exam-Relevant Nuance

- Patch management is part of a **vulnerability management** program — patching is one response to vulnerabilities; others include configuration changes, compensating controls, or accepting risk.
- Change management must be part of every patch management program — an uncontrolled patch can cause outages.
- The exam may ask about the *relationship* between patch and change management: each patch deployment should be submitted as a change request.
- Passive patch detection is the *least accurate* but most non-intrusive method.

---

## Cross-Links

- [Change Management](change-management.md) — all patch deployments go through change control
- [Configuration Management](configuration-management.md) — patch level is a configuration attribute
- [Vulnerability Management Operations](vulnerability-management-ops.md) — vulnerability scanning identifies missing patches

## Sources

- destination-cissp §7.8–7.8.1 (pp. 0879–0882)
- cissp-exam-outline (Domain 7: patch and vulnerability management)
