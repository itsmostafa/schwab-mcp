---
title: "Identity Federation"
type: concept
domain: 5
tags: [federation, fim, identity-provider, relying-party, trust, idaas, saml, openid]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Identity Federation

## Definition

**Federated Identity Management (FIM):** Allows a user to authenticate *once* and gain access to multiple systems *across organizational boundaries*. Extends [SSO](sso.md) beyond a single organization.

**Key distinction:**
- SSO = one authentication → multiple systems *within one organization*
- FIM = one authentication → multiple systems *across different organizations*

**Real-world example:** Logging into Pinterest using your Google account. Google authenticates you; Pinterest trusts Google's assertion.

## Three Components of Any Federated Access System

| Component | Also known as | Role |
|---|---|---|
| **Principal / User** | End user | The person who wants access |
| **Identity Provider (IdP)** | — | Owns the identity; performs authentication |
| **Relying Party (RP)** | Service Provider (SP) | Trusts the IdP's assertion; grants access |

The system depends on a **trust relationship** between IdP and RP/SP. In the Google/Pinterest example: Google = IdP, Pinterest = RP.

## Federated Access Standards

| Standard | Auth | Authz | Notes |
|---|---|---|---|
| **SAML** | Yes | Yes | XML-based; see [SAML](saml.md) |
| **WS-Federation** | Yes | Yes | Created by IBM/Microsoft/Verisign; standardized by OASIS |
| **OpenID** | Yes | No | Authentication only; user uses existing account to identify across sites |
| **OAuth** | No | Yes | Authorization only; see [OAuth/OIDC](oauth-oidc.md) |
| **OpenID Connect (OIDC)** | Yes | Yes | Authentication layer built on top of OAuth 2.0 |

> **Exam critical:** OpenID = authentication. OAuth = authorization. OIDC = authentication + authorization (combines both).

## Identity as a Service (IDaaS)

IDaaS is FIM implemented in the cloud. The IdP lives in a cloud service, and the organization's identity infrastructure is managed externally.

### IDaaS capabilities

- Provisioning, Administration, SSO, MFA, Directory Services
- On-premises and cloud support

### IDaaS identity types (destination-cissp §5.3.1)

| Type | Account stored | Authenticated by |
|---|---|---|
| **Cloud Identity** | Created/managed in cloud | Cloud service |
| **Synced Identity** | Created in local store (e.g., AD), synced to cloud | Either local or cloud |
| **Linked Identities** | Two separate accounts linked (local + cloud) | Either local or cloud |
| **Federated Identity** | Managed by IdP | Identity Provider |

### IAM deployment models

| Model | Characteristics |
|---|---|
| **On-Premises** | Controlled by organization; no internet dependency; typically very secure |
| **Cloud** | Provided by CSP; uses federated protocols; availability risk + multitenant security risks |
| **Hybrid** | Combines on-premises + cloud; most flexibility for growing organizations |

### IDaaS Risks

1. **Availability:** CSP outage = users locked out.
2. **Protection of identity data:** PII and sensitive data controlled by the CSP.
3. **Third-party data exposure:** Organization details can be inferred from shared identity data.

## Trust Models

Federation requires establishing trust between the IdP and SP/RP. This is analogous to airport security checkpoints in different airports trusting each other's screening — the trust enables transitive access without re-screening.

## Cross-links

- [SSO](sso.md)
- [SAML](saml.md)
- [OAuth/OIDC](oauth-oidc.md)
- [Directory Services](directory-services.md)
- [Identity Lifecycle](identity-lifecycle.md)

## Sources

- destination-cissp §5.2.12, §5.2.13, §5.2.14, §5.3.1 (pp. 712-725)
- cissp-exam-outline §5.3, §5.6
