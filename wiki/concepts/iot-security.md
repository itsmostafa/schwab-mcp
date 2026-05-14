---
title: "IoT Security"
type: concept
domain: 3
tags: [iot, internet-of-things, embedded, botnet, ddos, firmware, attack-surface]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# IoT Security

The Internet of Things (IoT) refers to the multitude of everyday devices — appliances, vehicles, industrial equipment, medical devices — that have been fitted with cheap network connectivity and connected to the internet. These devices present significant security challenges because they are mass-produced with minimal security, rarely updated, and increasingly present in both home and enterprise networks.

## Definition

IoT devices are physical devices embedded with:
- Cheap, mass-produced processing chips and network interfaces
- Sensors or actuators (physical interaction with the environment)
- Connectivity to the internet (directly or through a gateway)
- Little to no built-in security — authentication is often absent or default credentials are never changed

## The Core Security Problem

Manufacturers optimize IoT devices for cost and connectivity, not security:

1. **Mass-produced, cheap hardware**: The chips and network cards used are designed for volume, not security.
2. **No patching culture**: Users don't think about patching their refrigerators, door locks, or security cameras. Devices run unpatched for their full lifecycle.
3. **Long refresh cycles**: Enterprise IT refreshes every 3–5 years; IoT devices may be in service for 10–20 years.
4. **Default credentials**: Many devices ship with admin/admin or similar defaults that are never changed.
5. **No encryption**: Many IoT protocols lack encryption by default.

## Attack Surface

The attack surface of IoT includes:
- **Default credentials** — easily guessable or publicly known
- **Unpatched firmware** — known vulnerabilities persist for years
- **Unencrypted communications** — traffic can be intercepted
- **Physical access** — many IoT devices are deployed in accessible locations
- **Supply chain** — malicious firmware injected during manufacturing

## Real-World Example: Mirai Botnet (2016)

Security cameras running vulnerable IoT firmware were compromised at massive scale. The attacker:
1. Identified a firmware vulnerability in millions of security cameras
2. Installed malware turning cameras into botnet nodes
3. Used the botnet to launch one of the largest DDoS attacks in history

Irony: security cameras became the attack vector. This illustrates the primary IoT risk: cheap, unpatched devices in high volumes create massive botnets.

## IoT and DDoS

IoT devices are ideal botnet candidates:
- Always on, always connected
- Sufficient processing power to send network traffic
- Virtually never monitored for anomalous behavior
- Distributed globally — hard to block by geography

## Risk Reduction

Reducing IoT risk is difficult but not impossible:

- **Avoid IoT unless necessary** — the best control is not deploying the device
- **Vet before purchase** — choose vendors with a security commitment; check for patch history
- **Change default credentials** immediately upon deployment
- **Network segmentation** — place IoT devices on a dedicated, isolated network segment; prevent lateral movement to corporate resources
- **Keep firmware updated** — actively monitor vendor advisories
- **Vulnerability scanning** of IoT network segment
- **Physical security** — restrict physical access to IoT devices

## IoT vs. Embedded Systems

| | IoT | Embedded Systems |
|---|---|---|
| **Connectivity** | Always network-connected | May or may not be networked |
| **Examples** | Smart thermostat, IP camera, smartwatch | BIOS, car ECU, medical device controllers |
| **Key risk** | Network-facing attack surface | Physical access; limited software update capability |

IoT is a subset of embedded systems that specifically features network connectivity.

## Exam-Relevant Nuance

- IoT = cheap, network-connected devices with minimal security.
- The biggest IoT risk is unpatched firmware on devices with long refresh cycles.
- Network segmentation is the most effective enterprise IoT control after avoiding deployment.
- DDoS via IoT botnet is the canonical IoT attack scenario.
- IoT increases the attack surface of any network it's connected to.

## Cross-links

- [ICS/SCADA Security](ics-scada-security.md) — OT/ICS also involves embedded/connected devices
- [Mobile Security](mobile-security.md) — mobile devices share some IoT characteristics
- [Cloud Security Models](cloud-security-models.md) — cloud-connected IoT architectures
- [Secure Design Principles](secure-design-principles.md) — defense in depth, secure defaults

## Sources

- destination-cissp §3.5.8 (pp. 318–320)
- cissp-exam-outline §3.5
