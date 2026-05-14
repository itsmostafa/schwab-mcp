---
title: "Mobile Security"
type: concept
domain: 3
tags: [mobile, mdm, mam, byod, cope, cyod, containerization, owasp-mobile]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Mobile Security

Mobile devices — smartphones, tablets, and similar small-form-factor computers — present unique security challenges because of their power, portability, and tendency to be lost or stolen. They can store and access vast amounts of sensitive data while being carried outside the protected perimeter of an organization.

## Key Risks

- **Physical loss/theft**: Small, powerful, portable — frequently lost or stolen
- **Expanded attack surface**: Connects to public Wi-Fi, personal apps, personal cloud services
- **BYOD risk**: Personal devices with corporate data have no device-level corporate control unless MDM is deployed
- **App-based attacks**: Malicious apps can exfiltrate data, intercept credentials, or install malware

## Mobile Device Management (MDM) vs. Mobile Application Management (MAM)

| | MDM | MAM |
|---|---|---|
| **Scope** | Manages the entire device | Manages specific applications only |
| **Controls** | Remote wipe, encryption enforcement, device policy, VPN | Application whitelisting, app-level encryption, corporate data containers |
| **Use case** | Corporate-owned devices (COPE) | BYOD environments where full device control is inappropriate |
| **Data isolation** | Full device | Application-level containerization |

MDM and MAM are often combined in a single platform. Together, they enforce:
- Device encryption
- VPN requirements for remote access
- Application whitelisting
- Remote wipe capability

**Important caveat on remote wipe**: Remote wipe requires the device to be connected to the internet. A sophisticated attacker can prevent remote wipe by keeping the device offline or in airplane mode.

## Deployment Models (Ownership)

| Model | Full Name | Who owns device | Who controls it | Best for |
|---|---|---|---|---|
| **BYOD** | Bring Your Own Device | Employee | Employee (MDM/MAM enrollment voluntary or required) | Cost savings; employee flexibility |
| **COPE** | Corporate-Owned, Personally Enabled | Organization | Organization | High-security environments; full device control |
| **CYOD** | Choose Your Own Device | Organization | Organization | Balance of user choice and corporate control |

## Security Controls for Mobile

| Control Category | Examples |
|---|---|
| **Remote access** | VPN + 2FA required by default for any corporate access |
| **Endpoint security** | Antivirus/malware, DLP, MDM-provisioned software |
| **Application control** | Application whitelisting — only approved apps may be installed |
| **Hardening** | Minimize attack surface; disable unnecessary features |
| **Policy** | Acceptable Use Policy, BYOD Policy, Lost Device Policy |
| **Containerization** | Corporate data in an encrypted, isolated container separate from personal data |

## OWASP Mobile Top 10

The Open Web Application Security Project (OWASP) publishes mobile-specific vulnerability rankings:

| Rank | Risk |
|---|---|
| M1 | Improper Credential Usage |
| M2 | Inadequate Supply Chain Security |
| M3 | Insecure Authentication/Authorization |
| M4 | Insufficient Input/Output Validation |
| M5 | Insecure Communication |
| M6 | Inadequate Privacy Controls |
| M7 | Insufficient Binary Protections |
| M8 | Security Misconfiguration |
| M9 | Insecure Data Storage |
| M10 | Insufficient Cryptography |

Related OWASP tools:
- **MASTG** (Mobile Application Security Testing Guide) — manual for mobile app security testing and reverse engineering
- **MASVS** (Mobile Application Security Verification Standard) — standard for secure mobile app development and testing

## Exam-Relevant Nuance

- **MDM vs. MAM**: MDM = full device management; MAM = application-level management. MAM is preferred for BYOD because it doesn't require control of the entire personal device.
- Remote wipe is only effective if the device is online.
- VPN + 2FA is the best remote access posture for mobile.
- Application whitelisting is a key MDM control — only pre-approved apps allowed.
- BYOD introduces personal apps and personal cloud services as uncontrolled data channels.

## Cross-links

- [IoT Security](iot-security.md) — mobile shares some IoT constraints (embedded, mass-produced)
- [Cloud Security Models](cloud-security-models.md) — mobile apps frequently use SaaS/cloud backends
- [Secure Design Principles](secure-design-principles.md) — least privilege, secure defaults apply to mobile

## Sources

- destination-cissp §3.5.3–3.5.4 (pp. 306–309)
- cissp-exam-outline §3.5
