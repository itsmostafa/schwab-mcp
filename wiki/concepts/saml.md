---
title: "SAML — Security Assertion Markup Language"
type: concept
domain: 5
tags: [saml, federation, authentication, authorization, xml, sso, fim]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# SAML — Security Assertion Markup Language

## Definition

SAML (Security Assertion Markup Language) is an **XML-based open standard** for federated identity. It provides both authentication and authorization in Federated Identity Management (FIM) solutions.

**Key technology marker:** SAML assertions are written in **XML** (eXtensible Markup Language) — machine and human-readable.

## How SAML Works

1. **User (Principal)** requests access to a resource at the **Service Provider (SP / Relying Party)**.
2. If the user is not authenticated, the SP redirects them to the **Identity Provider (IdP)**.
3. The IdP authenticates the user (identification + authentication).
4. The IdP issues a **SAML Assertion Ticket/Token** to the user.
   - Critical: The assertion ticket does **NOT** contain the user's username and password.
   - It contains *assertion statements*: identity claims, role, level of access, attributes.
5. The user presents the SAML assertion to the SP.
6. The SP reads the assertion statements and makes an **authorization decision**.

## SAML Assertion Types

| Assertion Type | Content |
|---|---|
| Authentication assertion | Confirms the user was authenticated by the IdP |
| Attribute assertion | User attributes (name, email, role, department) |
| Authorization assertion | What the user is permitted to do |

## Four Key SAML Components

| Component | Function |
|---|---|
| **Assertion** | Authentication, authorization, and attribute statements |
| **Protocol** | Defines how entities request and respond to requests |
| **Bindings** | Mapping SAML onto standard communication protocols (e.g., HTTP) |
| **Profiles** | Define how SAML is used for specific business cases (e.g., Web SSO, LDAP) |

## SAML vs. OAuth/OIDC

| Feature | SAML | OAuth 2.0 | OpenID Connect |
|---|---|---|---|
| Purpose | Auth + Authz (federation) | Authorization only | Authentication (+ Authz via OAuth) |
| Format | XML | JSON/tokens | JSON/JWT |
| Use case | Enterprise SSO / FIM | API access delegation | User login across sites |
| Provides identity? | Yes (via assertions) | No | Yes (ID token) |

> **EXAM CRITICAL:** OAuth 2.0 is **NOT** an authentication protocol — it is an *authorization* framework. OIDC adds the authentication layer on top of OAuth. SAML handles both in enterprise contexts.

## Cross-links

- [Federation](federation.md)
- [OAuth/OIDC](oauth-oidc.md)
- [SSO](sso.md)
- [Kerberos](kerberos.md) — comparison: both use ticket/token model

## Sources

- destination-cissp §5.2.14, §5.2.13 (pp. 715-720)
- cissp-exam-outline §5.3, §5.6
