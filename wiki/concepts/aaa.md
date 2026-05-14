---
title: "AAA — Authentication, Authorization, Accounting"
type: concept
domain: 5
tags: [aaa, identification, authentication, authorization, accounting, access-control]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# AAA — Authentication, Authorization, Accounting

## Definition

The AAA framework describes the four core services of Access Control. Together they form the **Principle of Access Control** (also called accountability). The source uses the term "Access Control Services" and presents four components, though AAA conventionally groups the last three.

| Component | What it means | Key point |
|---|---|---|
| **Identification** | Claiming an identity (asserting who you are) | Username, access card, biometric scan *before* proof |
| **Authentication** | Proving the identity claim | Knowledge / Ownership / Characteristic |
| **Authorization** | What the authenticated identity is allowed to do | Permissions, roles, access levels |
| **Accounting** | Logging and monitoring all actions | Enables individual accountability |

## Exam-critical distinctions (EXAM CRITICAL)

- **Identification ≠ Authentication.** Identification is the *claim*; authentication is the *proof*. You type your username = identification. You type your password = authentication.
- **Authorization** answers: "What can you do?" — comes *after* authentication.
- **Accounting** (the Principle of Access Control) requires all three preceding steps to be in place. Without unique identification, you cannot achieve meaningful accountability.
- Shared accounts undermine the Principle of Access Control because actions can no longer be traced to a single individual.

## Accountability as the Principle of Access Control

Destination-cissp §5.2.15 states explicitly: **Accountability = the Principle of Access Control.** The four requirements:

1. Users must be uniquely identified.
2. Users must be properly authenticated.
3. Users must be properly authorized.
4. All actions must be logged and monitored.

Only when all four conditions hold can the Principle of Access Control be achieved.

## Cross-links

- [Authentication Factors](authentication-factors.md) — the three factor families
- [Access Control Models](access-control-models.md) — how authorization is implemented
- [Identity Lifecycle](identity-lifecycle.md) — provisioning to revocation
- [Domain 5 — IAM](../../domains/05-identity-and-access-management.md)

## Sources

- destination-cissp §5.2.1, §5.2.2, §5.2.15
- cissp-exam-outline §5.2
