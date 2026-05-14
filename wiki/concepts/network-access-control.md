---
title: "Network Access Control"
type: concept
domain: 4
tags: [nac, 802.1x, radius, tacacs, diameter, posture-assessment]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Network Access Control

Network Access Control (NAC) restricts network access based on endpoint health and user identity. It unifies endpoint security, authentication, and network security policy enforcement.

## Key Facts

### 802.1X Port-Based NAC

IEEE 802.1X provides port-based access control. Three roles:
- **Supplicant:** The client device requesting network access.
- **Authenticator:** The network device (switch or AP) that enforces access. Passes credentials to the authentication server.
- **Authentication Server:** Validates credentials (typically RADIUS or TACACS+). Sends access-allow or access-deny to the authenticator.

The supplicant uses **EAP** to authenticate to the authentication server, carried through the authenticator. If authentication fails, the port remains in a blocked or guest VLAN state.

### Remote Authentication Protocols

| Protocol | Transport | Encryption | AAA? | Notes |
|----------|-----------|-----------|------|-------|
| **RADIUS** | UDP | Poorly obfuscates password only | Yes (Authentication, Authorization, Accounting) | Originally for dial-in networking; widely used |
| **TACACS+** | TCP | **Encrypts entire packet** | Yes | Cisco extension; more secure than RADIUS; separates AAA |
| **Diameter** | TCP/SCTP | Yes (EAP support) | Yes | Successor to RADIUS; enhanced security; mobile/4G networks |

**RADIUS vs TACACS+:**
- RADIUS uses UDP; TACACS+ uses TCP (more reliable).
- RADIUS encrypts only the password field (poorly); TACACS+ encrypts the **entire payload**.
- TACACS+ separates authentication, authorization, and accounting into distinct processes; RADIUS combines them.
- Both are used in AAA frameworks (also covered in Domain 5 — Identity and Access Management).

### Posture Assessment

NAC solutions can evaluate an endpoint's **posture** (health check) before granting access:
- Is antivirus installed and up to date?
- Are OS patches current?
- Is the endpoint compliant with security policy?

If posture check fails, the device may be placed in a quarantine VLAN with remediation resources only.

### Endpoint Security and NAC

Modern endpoint security includes: antivirus, device management policies, DLP (Data Loss Prevention), NAC for access restriction, and EDR (Endpoint Detection and Response) platforms.

## Exam Nuance

- 802.1X is the **framework** for port-based NAC; RADIUS or TACACS+ is the **backend authentication server**.
- RADIUS uses **UDP** — this is frequently tested. TACACS+ uses **TCP**.
- TACACS+ encrypts the **full packet**; RADIUS only obfuscates the password — TACACS+ is more secure.
- **Diameter** is the successor to RADIUS and is used in modern mobile/4G networks. It adds EAP and improved security.
- AAA (Authentication, Authorization, Accounting) — RADIUS and TACACS+ both support this triad, but the implementation differs. Full AAA treatment is in Domain 5.

## Cross-links

- [Wireless Security](./wireless-security.md)
- [Secure Protocols](./secure-protocols.md)
- [VPN](./vpn.md)
- [802.1X Standard](../standards/802-1x.md)

## Sources

- destination-cissp §4.2.7 (pp. 0637–0638)
- destination-cissp §4.3.4 (pp. 0656–0658)
