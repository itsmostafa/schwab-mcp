---
title: "Access Control Models"
type: concept
domain: 5
tags: [dac, mac, rbac, abac, rule-based, access-control, xacml, pep, pdp, non-discretionary]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Access Control Models

## Overview

Access control models define *how* authorization decisions are made. Destination-cissp §5.4 organizes them under three broad categories:

1. **Discretionary** — owner decides access
2. **Mandatory** — system decides based on labels
3. **Non-discretionary** — neither owner nor system; someone other than the owner decides

## Quick Comparison Table (EXAM CRITICAL)

| Model | Who decides access | Basis for decision | Use case | Admin overhead |
|---|---|---|---|---|
| **DAC** | Asset owner | Owner's discretion | General enterprise | Medium |
| **Rule-Based** | Owner (via rules) | Explicit rule set | Granular control needs | High |
| **RBAC** | Organization (via role definitions) | Job function / role | Organizations with defined roles | Low–Medium |
| **ABAC** | Policy engine | User + environment attributes | Cloud, complex dynamic environments | Medium (policy-driven) |
| **MAC** | System | Clearance label vs. classification label | Government / military | Very high setup |
| **Non-Discretionary** | Someone other than asset owner | IT admin or delegated party | (Anti-pattern; avoid) | Variable |
| **Risk-Based** | System | Risk profile of connection | Adaptive/zero-trust systems | Low (automated) |

---

## Discretionary Access Control (DAC)

**Defining characteristic:** The **asset owner** decides who can access the asset, at their discretion. Best practice because owners are accountable and best positioned to judge access needs.

**Three subtypes of DAC:**

### Rule-Based Access Control

Access is governed by explicit rules assigned to users (e.g., an Access Control List / ACL).
- Very granular — can define read/write/execute per user per resource.
- High administrative overhead as rule tables grow.
- Example: Alice has read-only access to Bob's directory; read+write to her own.

### Role-Based Access Control (RBAC)

Access based on **job function / role**. Users are assigned to roles; roles carry permissions.
- Mirrors organizational chart.
- Significant reduction in per-user administration.
- Considered a "best practice" for most organizations.
- **Exam trap:** Full-RBAC across an entire complex organization can result in *more roles than employees*, increasing admin burden. Most organizations use **Limited or Hybrid RBAC**.

### Attribute-Based Access Control (ABAC)

Access based on **attributes** of the user, the resource, and the environment:
- Job function, device type, OS/browser version, IP address, time of day, asset classification.
- Most **flexible** and **granular** model.
- Enables fine-grained cloud access control where traditional network perimeters don't apply.
- Standard enabler: **XACML** (eXtensible Access Control Markup Language) — defines the ABAC policy language, architecture, and processing model.

---

## Mandatory Access Control (MAC)

**Defining characteristic:** The **system** determines access based on labels — not the owner.

- Every **object** gets a **classification label** (e.g., Public, Secret, Top Secret).
- Every **subject** (user) gets a **clearance level**.
- System grants/denies access: user clearance must meet or exceed object classification.
- Primarily protects **confidentiality**.
- Rare in private companies (requires all assets classified + all users cleared).
- Common in **government / military** contexts.

---

## Non-Discretionary Access Control

**Defining characteristic:** Someone **other than the owner** makes access decisions (e.g., IT department).

- Not a security best practice — avoid if possible.
- Occurs when: no identified owner exists, or owner delegates access decisions without retaining input.
- Can lead to over-provisioning ("grant everything just in case").

---

## Risk-Based Access Control

Evaluates elements of a connection request (IP address, time of access, geolocation) to compute a **risk profile**. Based on the result, may challenge the user with additional authentication before granting access.

- Supports adaptive authentication / zero-trust architectures.
- Related to context-based access control (considers whether connection is internal or external, typically enforced via firewall rules).

---

## Access Policy Enforcement: PEP and PDP

| Component | Full Name | Role |
|---|---|---|
| **PEP** | Policy Enforcement Point | Receives authorization requests; enforces PDP decisions by granting or denying access. Placed at all access points in an application. |
| **PDP** | Policy Decision Point | Evaluates authorization requests against pre-defined policy rules. Centralized. |

**Flow:** Request arrives at PEP → PEP sends to PDP → PDP evaluates rules → PDP returns decision → PEP enforces.

---

## Reference Monitor Concept (RMC)

All access control implementations are conceptually based on the **Reference Monitor Concept (RMC)**: a rules-based decision-making tool placed between subjects and objects to mediate access, with all activity logged for accountability. Any implementation of the RMC is called a **security kernel**.

---

## Cross-links

- [AAA](aaa.md)
- [Privileged Access Management](privileged-access-management.md)
- [Identity Lifecycle](identity-lifecycle.md)
- [Federation](federation.md) — ABAC + cloud
- [Domain 5 — IAM](../../domains/05-identity-and-access-management.md)

## Sources

- destination-cissp §5.4.1–§5.4.4 (pp. 725-740)
- destination-cissp §5.1.1 (RMC, pp. 667-669)
- cissp-exam-outline §5.4
