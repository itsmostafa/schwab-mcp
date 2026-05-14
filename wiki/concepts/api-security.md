---
title: "API Security"
type: concept
domain: 8
tags: [api, rest, soap, oauth, tls, api-gateway, rate-limiting, input-validation, web-services]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# API Security

## Definition

An **Application Programming Interface (API)** is a set of standards that allows applications to communicate with each other. APIs act as translators — enabling disparate systems (regardless of underlying programming language) to exchange data and invoke functionality through a defined interface.

> **Analogy (destination-cissp §8.5.3):** A restaurant server takes orders from customers (one system) and relays them to the kitchen (another system). The server is like the API — a defined interface that both sides understand.

---

## Two Primary API Formats

*(source: destination-cissp §8.5.3)*

| Property | REST | SOAP |
|---|---|---|
| **Full name** | Representational State Transfer | Simple Object Access Protocol |
| **Age** | Newer | Older (originally by Microsoft) |
| **Protocol** | HTTP-based | XML-based (protocol-agnostic in theory) |
| **Flexibility** | More flexible and lightweight | More rigid and standardized |
| **Ease of use** | Easy to learn and use; fast to process | Steeper learning curve |
| **Output formats** | CSV, JSON, RSS, XML (multiple options) | XML only |
| **Error handling** | Limited (relies on HTTP status codes) | Strong, built-in error handling |
| **Extensibility** | Lightweight; relies on developer conventions | Extensible via WS-* standards (WS-Security, etc.) |
| **Best for** | Modern web/mobile APIs; microservices | Enterprise integrations requiring strong contracts and security |

**Exam note:** REST is more commonly used today. SOAP is more rigid and older but has stronger built-in error handling and is still common in enterprise/legacy environments.

---

## API Security Controls

*(source: destination-cissp §8.5.3)*

| Control | Description |
|---|---|
| **Authentication and authorization** | Use access tokens (OAuth 2.0), API keys, or mutual TLS (mTLS). OAuth 2.0 is the standard for delegated authorization in modern REST APIs. |
| **Encryption (TLS)** | All API traffic traversing insecure channels (internet) must use TLS 1.2+ to prevent eavesdropping and tampering. |
| **Data validation** | Validate all inputs to the API — type, length, format. APIs are subject to the same injection attacks as web applications. |
| **API gateways** | Centralized point of control for all API traffic; enforces authentication, rate limiting, input validation, logging, and routing policies. |
| **Quotas and throttling (rate limiting)** | Limits the number of requests per client per unit time. Prevents abuse, DoS attacks, and data harvesting. |
| **Testing and validation** | Include APIs in the security testing regimen (DAST, penetration testing); use purpose-built API security scanners. |

---

## OWASP API Security Top 10 (Exam Awareness)

> **Coverage gap:** destination-cissp does not enumerate the OWASP API Security Top 10. The following is drawn from cissp-exam-outline scope and OWASP's published list. Top exam-relevant categories:

| Rank | Category | Core Issue |
|---|---|---|
| API1 | Broken Object Level Authorization | Accessing data belonging to other users by manipulating resource IDs. |
| API2 | Broken Authentication | Weak or missing authentication mechanisms on API endpoints. |
| API3 | Broken Object Property Level Authorization | Returning or accepting more object properties than the user should see/modify (mass assignment). |
| API4 | Unrestricted Resource Consumption | Lack of rate limiting enabling DoS or data harvesting. |
| API5 | Broken Function Level Authorization | Unauthorized access to admin-level API functions by regular users. |
| API6 | Unrestricted Access to Sensitive Business Flows | Exploiting business flows at scale (e.g., bulk account creation, scalper bots). |
| API7 | Server Side Request Forgery (SSRF) | API server makes requests to internal resources based on attacker-controlled input. |
| API8 | Security Misconfiguration | Insecure defaults, verbose error messages, unnecessary HTTP methods enabled. |
| API9 | Improper Inventory Management | Undocumented or deprecated ("shadow") API versions exposed and unpatched. |
| API10 | Unsafe Consumption of APIs | Trusting third-party API responses without validation. |

---

## Exam-Relevant Nuance

- **REST vs. SOAP**: REST = flexible, HTTP-based, multiple output formats. SOAP = rigid, XML-only, stronger error handling. REST is dominant in modern APIs.
- **OAuth** is for *authorization* (delegated access), not authentication. OpenID Connect (OIDC) adds authentication on top of OAuth — see [OAuth/OIDC](./oauth-oidc.md).
- **API gateways** are a key architectural control — centralizing enforcement avoids inconsistent security across individual APIs.
- **Rate limiting** is both a performance control and a security control (prevents DoS and brute-force attacks on API endpoints).
- APIs are subject to all the same vulnerabilities as web applications (injection, broken access control, etc.) plus API-specific issues (broken object-level authorization, mass assignment).

---

## Cross-Links

- [Secure Coding Practices](secure-coding-practices.md) — input validation; underlying principles
- [OWASP Top 10](owasp-top-10.md) — web application security; overlaps with API security
- [OAuth/OIDC](./oauth-oidc.md) — OAuth for API authorization (cross-domain: Domain 5)
- [Software Testing Types](./security-testing-types.md) — DAST for API testing

## Sources

- destination-cissp §8.5.3 (APIs, REST vs. SOAP, API security controls) (pp. 0968–0969)
- cissp-exam-outline (Domain 8.5: Security of APIs)
- OWASP API Security Project (owasp.org/api-security) — coverage gap noted
