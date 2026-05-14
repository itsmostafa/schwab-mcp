---
title: "Social Engineering"
type: concept
domain: 1
tags: [social-engineering, phishing, vishing, smishing, pretexting, baiting, tailgating, piggybacking, whaling, spear-phishing]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Social Engineering

Social engineering is using **deception or intimidation to manipulate people into providing
sensitive information or performing actions** that facilitate fraudulent activities. It exploits
human psychology rather than technical vulnerabilities.

## Core Tactics

- **Intimidation** — Inducing fear to manipulate someone (e.g., blackmail).
- **Deception** — Tricking someone (e.g., lying about identity or context).
- **Rapport** — Building a relationship with the victim over time to exploit later.

## Common Attack Types (Table 1-30)

| Attack | Description |
|---|---|
| **Phishing** | Mass email campaign hoping a recipient clicks a malicious link or opens a malicious file |
| **Spear Phishing** | Targeted phishing — researched and customized for specific individuals or groups (e.g., malicious invoice to AP team) |
| **Whaling** | Spear phishing targeting "big fish" — CEO, COO, CFO |
| **Smishing** | Phishing via SMS/text message; may allow attacker to control the victim's phone |
| **Vishing** | Phishing via voice/VoIP phone calls |
| **Pretexting** | Attacker creates a fabricated scenario (script) that triggers the victim emotionally — e.g., "your bank account has suspicious activity" |
| **Baiting** | Exploits curiosity via physical media (e.g., dropping USB drives in a parking lot) |
| **Tailgating** | Following an authorized person through a restricted door — attacker has a **fake but realistic-looking badge** |
| **Piggybacking** | Same as tailgating, but attacker has **no badge at all** |

## Key Distinction: Tailgating vs. Piggybacking

- **Tailgating**: Attacker has a fake badge.
- **Piggybacking**: Attacker has no badge; relies solely on following someone.
(This distinction is a documented exam trap.)

## Mitigations

- **Awareness, training, and education** — primary mitigation for all social engineering.
- Requiring proof of identity before granting access.
- Callback authorization for voice/text requests (verify out-of-band).
- Contacting organizations via confirmed numbers (not numbers in the suspicious email).
- Strong security policies (acceptable use, clean desk, tailgating awareness).
- Simulated phishing campaigns to measure and improve employee vigilance.

## Exam-Relevant Nuance

- Best defense against social engineering = **user awareness and training** (not technical controls).
- Pretexting may involve impersonating IT support, bank representatives, or colleagues.
- Baiting exploits curiosity — a dropped USB drive is a classic scenario.
- Whaling vs. spear phishing: Whaling targets specific executives; spear phishing targets any specific individual or group.

## Cross-links

- [Security Awareness](security-awareness.md)
- [Personnel Security](personnel-security.md)
- [Threat Modeling](threat-modeling.md)

## Sources

- destination-cissp §1.10 (Social Engineering, Table 1-30) and §1.12.1
