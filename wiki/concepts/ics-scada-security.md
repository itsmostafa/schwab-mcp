---
title: "ICS/SCADA Security"
type: concept
domain: 3
tags: [ics, scada, ot, operational-technology, critical-infrastructure, air-gap, plc, dcs]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# ICS/SCADA Security

Industrial Control Systems (ICS) and SCADA systems are the computerized hardware and software that control physical processes in critical infrastructure — power grids, water treatment, nuclear plants, manufacturing, and similar environments. They represent a unique security challenge: they are mission-critical, difficult to patch, often running legacy software, and increasingly connected to corporate and internet-facing networks.

## Key Terminology

| Term | Full Name | Definition |
|---|---|---|
| **OT** | Operational Technology | Broad term: hardware and software that monitors and controls physical processes and industrial systems |
| **ICS** | Industrial Control System | Subset of OT; focuses on control and automation of industrial processes |
| **SCADA** | Supervisory Control and Data Acquisition | System architecture combining computers, networking, and proprietary devices for management of industrial, infrastructure, and facility processes — includes local and remote management |
| **DCS** | Distributed Control System | Process control system monitoring, controlling, and gathering data from components in large processing facilities — typically **locally** controlled (no remote) |
| **PLC** | Programmable Logic Controller | Industrial computer for manufacturing process control; high reliability; often networked with other PLCs and SCADA |

**Hierarchy**: OT > ICS > SCADA/DCS/PLC

## Why ICS Security is Difficult

1. **Legacy software and hardware**: ICS systems are often designed to run for 20–30 years. A system designed for Windows XP cannot be upgraded to Windows 10 without risking breaking the custom control software.
2. **Fear of patching**: Patching may cause unintended downtime or malfunction of critical infrastructure — operators often avoid patching even when patches are available.
3. **Specialized software**: ICS uses highly customized, proprietary software that vendors may not support with security updates.
4. **Convergence of IT and OT**: Modern networks increasingly connect corporate IT networks (and the internet) to OT/ICS networks, dramatically expanding the attack surface.

## Primary Risk Reduction: Air Gapping

The **best** way to protect ICS is to keep them **offline** — this is called an **air gap**:
- ICS devices can communicate with each other on their own isolated network
- The ICS network has **no connection** to the internet or corporate network
- Even an attacker who compromises the corporate network cannot reach the ICS

When air gapping is not feasible, compensating controls include:
- VLANs and network zoning to limit lateral movement
- Nonstop logging, monitoring, and anomaly detection
- Regular vulnerability assessments with focus on: internet connections, rogue devices, plaintext authentication
- Privileged Access Management (PAM) tools for managing legacy system access

## When Patching is Required

If patching is necessary despite the risks:
- Maintain strong configuration management and patch management processes
- Create backups and archives before patching
- Test patches in a staging environment before production deployment
- Implement additional monitoring during and after patch windows

## SCADA vs. DCS vs. PLC

| | SCADA | DCS | PLC |
|---|---|---|---|
| **Scope** | Large-scale; local + remote management | Typically local; large processing facilities | Single manufacturing process; device-level |
| **Focus** | System-wide supervision and control | Process control for industrial facilities | Highly reliable discrete control |
| **Networking** | Networked; interfaces with enterprise systems | Typically more isolated | Often networked with SCADA |

## Attack Surface Considerations

- ICS connectivity to corporate networks is the most common attack vector
- ICS protocols (Modbus, DNP3, OPC) were designed without security in mind — no authentication, no encryption
- Smart/IoT technology is blurring the line between IT and OT, further expanding ICS attack surfaces

## Exam-Relevant Nuance

- **Air gap** = the best ICS protection. Know this answer.
- OT is broader than ICS; ICS is a subset of OT.
- DCS is locally controlled; SCADA includes remote management — this distinction is testable.
- ICS patching avoidance is intentional (not negligence) — critical infrastructure uptime outweighs patch cycles.
- Stuxnet (2010) is the canonical real-world ICS attack — attacked uranium enrichment PLCs via Windows zero-days.

## Cross-links

- [IoT Security](iot-security.md) — IoT devices increasingly connected to OT networks
- [Cloud Security Models](cloud-security-models.md) — OT/cloud convergence considerations
- [Secure Design Principles](secure-design-principles.md) — defense in depth for ICS

## Sources

- destination-cissp §3.5.7 (pp. 316–318)
- cissp-exam-outline §3.5
