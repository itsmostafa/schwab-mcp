---
title: "Identity Lifecycle"
type: concept
domain: 5
tags: [identity-lifecycle, provisioning, deprovisioning, access-review, joiner-mover-leaver, orphaned-accounts]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Identity Lifecycle

## Definition

The **Identity Lifecycle** (also called identity and access provisioning lifecycle) describes the full arc of a user account's existence: creation, periodic review, and eventual revocation. Destination-cissp §5.5.2 describes three stages:

```
Provisioning → Review → Revocation
```

## Stage 1: Provisioning

**When:** Upon hire of a new employee, or when an employee changes roles.

**Activities include:**
- Background checks
- Identity proofing / registration (confirming the person is who they claim to be)
- Issuing credentials (account, badge, etc.)
- Granting access to required systems, based on job role

**Identity proofing (§5.2.11):** Before issuing credentials, organizations verify identity using government-issued ID, driver's license, or other unique identifiers. Analogous to a Registration Authority (RA) proofing a certificate applicant before a CA issues a certificate.

## Stage 2: Review (User Access Review / Recertification)

**When:** Periodically, based on risk — not on a fixed schedule for all accounts equally.

**Who reviews:** The **asset or system owner** — they are in the best position to confirm whether continued access is appropriate.

**Frequency guidance (destination-cissp §5.5.3):**
- Standard user access: at minimum annually.
- **Privileged/admin/root accounts: much more frequently** — potentially weekly.
- When a user changes roles: immediately at time of change.
- When a user leaves: immediately upon separation.

**Purpose:** Prevents **privilege creep** (accumulation of access rights beyond what's needed). When a user changes roles, old access should be revoked and new access granted fresh — not simply added to existing permissions.

## Stage 3: Revocation (Deprovisioning)

**When:**
- Employee leaves (voluntary or involuntary termination).
- Employee changes roles (access should be re-evaluated; revoke and re-provision rather than accumulate).

**Risk of skipping revocation:** **Orphaned accounts** — active accounts for departed users that can be exploited for unauthorized access.

**Accounts to deprovision include:** user accounts, remote access, physical access cards, email, system accounts.

## Access Creep / Privilege Creep

When users accumulate permissions over time (especially after role changes) without proper revocation:
- Results in users having more access than their current role requires.
- Violates least privilege principle.
- User access reviews are the primary mitigating control.

## Joiner / Mover / Leaver Model

| Event | Lifecycle Phase | Action |
|---|---|---|
| **Joiner** (new hire) | Provisioning | Create account, grant role-appropriate access |
| **Mover** (role change) | Revocation + Provisioning | Revoke prior access, provision new access |
| **Leaver** (departure) | Revocation | Revoke all access promptly |

## Cross-links

- [Privileged Access Management](privileged-access-management.md)
- [AAA](aaa.md)
- [Access Control Models](access-control-models.md)
- [Directory Services](directory-services.md)

## Sources

- destination-cissp §5.5.1–§5.5.5 (pp. 740-746)
- destination-cissp §5.2.11 (pp. 710)
- cissp-exam-outline §5.5
