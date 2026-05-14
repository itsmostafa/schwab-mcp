---
title: "Privileged Access Management (PAM)"
type: concept
domain: 5
tags: [pam, least-privilege, sudo, privilege-escalation, jit, service-accounts]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Privileged Access Management (PAM)

## Definition

Privileged Access Management (PAM) encompasses the policies, processes, and tools used to control, monitor, and protect privileged accounts — accounts with elevated rights (admin, root, superuser) beyond standard user permissions.

## Core Principle: Two-Account Model

Destination-cissp §5.5.4 recommends that privileged users maintain **two separate accounts**:

1. **Standard user account** — for regular business tasks (email, meetings, browsing).
2. **Privileged account** — used *only* when performing administrative tasks requiring elevated access.

**Rationale:** High-risk activities (email, web browsing) that expose accounts to compromise should be performed under the lower-privilege account. The privileged account is only surfaced when strictly needed.

**Examples:**
- Unix/Linux: `sudo` ("superuser do") — run specific commands with elevated privileges while logged in as a standard user.
- Windows: `RunAs` command — equivalent to sudo.

**Auditing:** Use of sudo/RunAs should be audited. Logs of privileged command execution are essential for accountability.

## Just-in-Time (JIT) Access

**Just-in-Time (JIT) access:** Temporarily elevate user privileges for a specific task window, then revoke the elevation. Mitigates the risk of standing (permanent) elevated privilege.

**Example:** A user needs to access a sensitive database partition once a month to run a report. JIT grants the elevated access during that window, then automatically removes it — rather than leaving the user permanently privileged.

Benefits:
- Reduces the window of exposure for privileged access.
- Can be automated, reducing administrative overhead.
- Minimizes potential for privilege abuse or account compromise.

## Privileged Account Review Frequency

Per destination-cissp §5.5.3: privileged/admin/root/superuser accounts should be reviewed **more frequently** than standard user accounts — potentially as often as **weekly**. Standard user access may be reviewed annually.

## Service Account Management

**Service accounts:** Accounts used by services, workloads, or applications — not by humans directly.

Best practices:
- Limit service accounts to **single purposes**.
- Apply **least privilege** — minimum permissions needed to function.
- Maintain human oversight: monitor and audit service account usage.
- Risks if poorly managed: privilege escalation attacks, spoofing.

## Vendor Access

Third-party vendor access should be provisioned with **equal or greater care** than employee access. Vendor access provisioning may include:
- Security review of the vendor organization.
- Onsite inspection of vendor facilities and systems.
- Same lifecycle controls: provisioning → review → revocation.

## Cross-links

- [Identity Lifecycle](identity-lifecycle.md)
- [Access Control Models](access-control-models.md)
- [AAA](aaa.md)

## Sources

- destination-cissp §5.5.1, §5.5.4, §5.5.5 (pp. 740-746)
- destination-cissp §5.2.16 — JIT Access (pp. 721-722)
- cissp-exam-outline §5.5
