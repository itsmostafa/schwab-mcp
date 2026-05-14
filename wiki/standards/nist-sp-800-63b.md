---
title: "NIST SP 800-63B — Digital Identity Guidelines: Authentication and Lifecycle Management"
type: standard
domain: 5
tags: [nist, 800-63, aal, authentication, digital-identity, mfa]
sources: [destination-cissp]
updated: 2026-05-13
---

# NIST SP 800-63B — Digital Identity Guidelines: Authentication and Lifecycle Management

## Overview

**NIST SP 800-63B** is part of the NIST SP 800-63 Digital Identity Guidelines suite. It covers **Authentication and Lifecycle Management**, defining standards for digital authentication strength via **Authenticator Assurance Levels (AAL)**.

- Publication suite URL: https://pages.nist.gov/800-63-3/
- Part of a four-document suite (800-63, 800-63A, 800-63B, 800-63C).

## Authenticator Assurance Levels (AAL)

AALs measure the **robustness of the authentication process**. Higher number = more robust.

| Level | Name | Requirements |
|---|---|---|
| **AAL1** | Some assurance | Single-factor authentication + secure authentication protocol |
| **AAL2** | High confidence | Multi-factor authentication + secure authentication protocol + approved cryptographic techniques |
| **AAL3** | Very high confidence | Multi-factor authentication + secure authentication protocol + "hard" cryptographic authenticator providing proof of possession of key + impersonation resistance |

### AAL3 Key Detail

AAL3 requires a **hardware-based cryptographic authenticator** (a "hard" authenticator) that:
- Proves possession of a cryptographic key.
- Provides **impersonation resistance** (cannot be phished — the authenticator is bound to the specific RP/SP).

## Relationship to Domain 5

- AAL levels directly map to MFA design decisions (§5.2.12).
- AAL1 is acceptable for low-sensitivity systems.
- AAL2 aligns with standard enterprise MFA requirements.
- AAL3 is required for high-value government/sensitive systems.

## Cross-links

- [Authentication Factors](../concepts/authentication-factors.md)
- [AAA](../concepts/aaa.md)
- [Biometrics](../concepts/biometrics.md)

## Sources

- destination-cissp §5.2.12 (p. 711)
- NIST: https://pages.nist.gov/800-63-3/
