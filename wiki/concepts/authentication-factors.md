---
title: "Authentication Factors"
type: concept
domain: 5
tags: [authentication, mfa, otp, totp, hotp, password, passkey, factors]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Authentication Factors

## Definition

Destination-cissp (§5.2.6) identifies **three** factor families for authentication:

| Factor | Label | Examples |
|---|---|---|
| **Authentication by Knowledge** | Something you know | Password, passphrase, security/cognitive questions |
| **Authentication by Ownership** | Something you have | OTP token (hard/soft), smart card, memory card, passkey device |
| **Authentication by Characteristic** | Something you are | Biometrics (fingerprint, iris, retina, voice, gait) |

> **Note on extended factors:** Some exam materials list "somewhere you are" (location) and "something you do" (behavioral biometrics) as separate factors. Destination-cissp §5.2.6 treats these within the three families above. Behavioral biometrics (gait, keystroke dynamics) fall under "Characteristic."

## Single-Factor vs. Multi-Factor Authentication (MFA)

- **Single-factor authentication (SFA):** Uses one factor family, regardless of how many objects. Example: password + security question = still SFA (both are "knowledge").
- **Multi-factor authentication (MFA):** Uses two or more *different* factor families. Example: password (knowledge) + RSA token (ownership) = MFA.

**Exam trap:** An RSA ID key + a Microsoft token is still SFA (both are "ownership/have"). The factor *type* matters, not the count of authentication objects.

## One-Time Passwords (OTP)

OTPs can be generated via:

| Generation Method | Description |
|---|---|
| **Synchronous** | Time-based (TOTP) or counter-based (HOTP); simpler, less expensive |
| **Asynchronous** | Challenge-response; more complex and robust, requires back-and-forth sync with the auth server |

- **Soft token:** Software app generating OTPs (e.g., Google Authenticator).
- **Hard token:** Dedicated hardware device (e.g., RSA SecureID).

## Password-less Authentication

Source (§5.2.6, §5.2.7 context): Password-less options include biometrics, mobile devices, or hardware security tokens. **Passkeys** let users authenticate via device PIN or biometrics; they resist phishing but:
- Still have biometric CER limitations (see [Biometrics](biometrics.md)).
- If device/token lost, user may be locked out.
- Hardware token implementation costs are higher.

## Smart Cards vs. Memory Cards

| Type | Storage | Unique per transaction? | Risk |
|---|---|---|---|
| **Smart card** | Embedded IC chip | Yes — chip computes unique auth data | More secure |
| **Memory card** | Magnetic strip | No — same data read every time | Vulnerable to skimming |

Modern debit/credit cards combine both (chip + magnetic strip).

Smart card communication:
- **Contact:** Chip must physically touch the reader.
- **Contactless:** Reader powers and communicates with chip wirelessly.

## Cross-links

- [AAA](aaa.md)
- [Biometrics](biometrics.md)
- [Kerberos](kerberos.md)
- [SSO](sso.md)

## Sources

- destination-cissp §5.2.3, §5.2.4, §5.2.5, §5.2.6
- cissp-exam-outline §5.2
