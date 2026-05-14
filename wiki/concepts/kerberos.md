---
title: "Kerberos"
type: concept
domain: 5
tags: [kerberos, sso, kdc, tgt, authentication, symmetric, toctou, sesame]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Kerberos

## Definition

Kerberos is a network authentication protocol that provides **Single Sign-On (SSO)** services. The name comes from the three-headed Greek dog Cerberus, guarding the gates of Hell — fitting for guarding access to resources. Kerberos provides three primary functionalities:

1. **Accounting**
2. **Authentication**
3. **Auditing**

Kerberos is built into Windows (requires Active Directory), macOS, and many Linux/Unix distributions.

## Key Components

| Component | Full Name | Role |
|---|---|---|
| **KDC** | Key Distribution Center | Container for AS and TGS; the central authority |
| **AS** | Authentication Service | Verifies user identity; issues the TGT |
| **TGS** | Ticket Granting Service | Issues service tickets to authenticated users |
| **TGT** | Ticket Granting Ticket | Issued by AS; user presents to TGS (encrypted with TGS key — user cannot decrypt it) |
| **Service Ticket** | (Service Ticket) | Issued by TGS; grants access to a specific service |

## Authentication Flow (simplified)

1. **Alice → AS:** Sends initial authentication messages.
2. **AS → Alice:** Returns two messages — one encrypted with Alice's password (she decrypts to prove knowledge), and the **TGT** (encrypted with TGS key — Alice cannot decrypt).
3. **Alice → TGS:** Sends new tickets + the encrypted TGT.
4. **TGS → Alice:** Performs verifications; returns a **Service Ticket** (encrypted with the service's key).
5. **Alice → Service:** Sends new messages + the encrypted Service Ticket.
6. **Service → Alice:** Final verification passes; access granted.

> **Key insight:** Kerberos verifies passwords *without transmitting them* — the user decrypts a message with their password as the key. If they know it, they can proceed.

## Weaknesses (source: destination-cissp §5.2.8)

| Weakness | Detail |
|---|---|
| **Symmetric encryption only** | Supports RC4, DES, AES — no asymmetric crypto. Implies symmetric key distribution challenges. |
| **TOCTOU vulnerability** | Time-Of-Check Time-Of-Use — only one major ticket issued; session hijacking possible. Mitigated by increasing re-authentication frequency. |

> **Note:** The commonly cited "5-minute clock skew" requirement and "Golden Ticket attack" are **not** mentioned in destination-cissp. They appear in general CISSP exam prep materials. The source's named weaknesses are TOCTOU and symmetric-only crypto.

## TOCTOU Mitigation

Increase re-authentication frequency (expire tickets more often). Trade-off: for high-value systems, this makes sense. Forcing all users to re-auth frequently for a low-value system creates unnecessary burden. Isolate high-value systems and use shorter ticket lifetimes for those.

## SESAME: The Improved Alternative

See [SSO — SESAME section](sso.md#sesame). SESAME adds asymmetric cryptography and multiple tickets, solving both Kerberos weaknesses — but Kerberos's ubiquity has kept it dominant.

## Cross-links

- [SSO](sso.md)
- [Directory Services](directory-services.md)
- Session management (TOCTOU mitigation: increase re-authentication frequency — see §5.2.10 in domain page)

## Sources

- destination-cissp §5.2.8 (pp. 700-705)
- cissp-exam-outline §5.6
