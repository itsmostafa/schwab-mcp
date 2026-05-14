---
title: "Single Sign-On (SSO)"
type: concept
domain: 5
tags: [sso, kerberos, saml, authentication, centralized, federation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Single Sign-On (SSO)

## Definition

**Single Sign-On (SSO):** A user authenticates *once* and gains access to *multiple systems* within a single organization. Contrast with [Federated Identity Management (FIM)](federation.md), which extends this across organizational boundaries.

## How SSO Works (high-level flow)

1. User sends a login request to an application.
2. Application checks: "Is the user authenticated?" If not, redirects to the authentication server.
3. User authenticates (knowledge / ownership / characteristic, or combination).
4. User receives a **ticket or token**.
5. User presents the ticket to the application.
6. Application grants access based on the ticket.

## Pros and Cons

| Pros | Cons |
|---|---|
| Better user experience | **Single point of failure** — if compromised, attacker gets everything; if down, users get nothing |
| Users more likely to use one stronger password | Difficult to integrate unique/legacy systems |
| Centralized enforcement of timeout and attempt thresholds | Centralized administration = higher-value target |
| Simplified administration | |

## SSO Protocols

The major SSO protocol is **Kerberos** (see [Kerberos](kerberos.md)). Other SSO/federated standards include [SAML](saml.md), [OpenID/OAuth](oauth-oidc.md), and WS-Federation.

### SESAME

**SESAME** (Secure European System for Applications in a Multi-Vendor Environment) is an improved but rarely-used alternative to Kerberos:

| Feature | Kerberos | SESAME |
|---|---|---|
| Cryptography | Symmetric only | Symmetric + asymmetric |
| Tickets | Single ticket (TGT) | Multiple tickets |
| TOCTOU vulnerability | Yes | Mitigated (multiple tickets) |
| Adoption | Very widespread (built into Windows, macOS, Linux) | Low — Kerberos's ubiquity won |

> Windows requires **Active Directory** to use Kerberos.

## Relationship to Centralized Administration

SSO implies centralized administration. Destination-cissp §5.1.2 notes centralized administration has the single point of failure problem. This is a direct trade-off in SSO design.

## Cross-links

- [Kerberos](kerberos.md)
- [Federation](federation.md)
- [SAML](saml.md)
- [OAuth/OIDC](oauth-oidc.md)
- [Directory Services](directory-services.md)

## Sources

- destination-cissp §5.2.8 (pp. 697-705)
- cissp-exam-outline §5.6
