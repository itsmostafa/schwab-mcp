---
title: "OAuth 2.0 and OpenID Connect (OIDC)"
type: concept
domain: 5
tags: [oauth, oidc, openid, authorization, authentication, federation, access-token]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# OAuth 2.0 and OpenID Connect (OIDC)

## Definitions

**OAuth 2.0:**
- An *authorization* framework (NOT authentication).
- Allows users to grant client applications **delegated access** to resources without sharing credentials.
- Authorizes devices, APIs, servers, and applications using **access tokens** rather than usernames/passwords.
- Operates over HTTPS.

**OpenID Connect (OIDC):**
- An *identity layer* built **on top of OAuth 2.0**.
- Adds **authentication** — verifies who the user is and provides basic profile info via an **ID token**.
- "While OAuth 2.0 is about resource access and sharing, OIDC is about user authentication." (destination-cissp §5.6.1)

**OpenID (original):**
- Separate from OIDC — an older open standard.
- Provides authentication only: lets users use an existing account (e.g., Microsoft account) to identify and authenticate to multiple sites without creating new passwords.
- The user's password is given only to the identity provider, which then confirms identity to third-party sites.

## Exam-Critical Distinction (EXAM CRITICAL)

| Standard | Auth | Authz | Key token |
|---|---|---|---|
| **OAuth 2.0** | **No** | **Yes** | Access token |
| **OpenID (original)** | Yes | No | — |
| **OpenID Connect (OIDC)** | Yes (+ via OAuth) | Yes (via OAuth) | ID token + access token |

> OAuth 2.0 is commonly tested as an **authorization** protocol only. Candidates frequently misidentify it as an authentication protocol. OIDC is the correct answer for authentication built on OAuth.

## How They Work Together

OpenID and OAuth are **complementary** and often deployed together:
- OpenID provides **authentication** (who you are).
- OAuth provides **authorization** (what you can do).
- OIDC unifies both in a single framework built on OAuth 2.0.

## Comparison: SAML vs. OIDC/OAuth

| Feature | SAML | OIDC + OAuth |
|---|---|---|
| Format | XML | JSON / JWT |
| Primary use case | Enterprise SSO, FIM | Consumer web apps, mobile apps, APIs |
| Authentication | Yes | Yes (OIDC) |
| Authorization | Yes | Yes (OAuth) |

## Cross-links

- [SAML](saml.md)
- [Federation](federation.md)
- [SSO](sso.md)

## Sources

- destination-cissp §5.2.14, §5.6.1 (pp. 716-717, 747)
- cissp-exam-outline §5.3, §5.6
