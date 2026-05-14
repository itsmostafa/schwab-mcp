---
title: "Change Management (Software Context)"
type: concept
domain: 8
tags: [change-management, version-control, scm, release-management, configuration-management, patch-management]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Change Management (Software Context)

## Definition

**Change management** ensures that the costs and benefits of changes are analyzed and that changes are implemented in a controlled manner to reduce risk. In the software development context, change management applies both during the SDLC (controlling changes to software during development) and during the operations phase (controlling updates and patches to production software).

> **Cross-reference:** Change management is covered in depth in [Domain 7 — Security Operations](../domains/07-security-operations.md) §7.9.1. This page focuses on the software development-specific applications.

---

## Change Management in Software Development

*(source: destination-cissp §8.1 — Operation and Maintenance review)*

A **Request for Change (RFC)** may arise from:
- A **service request** (user-initiated feature or enhancement)
- The **incident management process** (change needed to resolve an incident)
- A **Service Level Agreement (SLA)** requirement
- A security vulnerability requiring a patch

**Software-specific components of change management:**
- **Configuration management**: Tracking the approved state (baseline) of software and its components. Any deviation is a change requiring approval.
- **Release management**: Managing the packaging and deployment of approved changes into production in a controlled, safe manner.
- **Security review and approval**: Security must be part of the change approval process — many organizations fail to include security on change control committees (Change Advisory Boards / CABs).

---

## Software Configuration Management (SCM)

*(source: destination-cissp §8.2.1)*

SCM specifically focuses on managing changes in software artifacts (source code, configuration files, documentation) throughout the SDLC:

| SCM Activity | Description |
|---|---|
| **Baseline establishment** | Define and record the approved state of software at a point in time. All changes measured against this baseline. |
| **Revision control** | Track all changes to code and configuration files — who changed what, when, and why. Enables rollback. |
| **Build management** | Consistent, repeatable build procedures ensure that code compiles identically each time. Prevents "works on my machine" problems. |
| **Process management** | Defines workflows for how changes move from development through testing to production. |
| **Teamwork facilitation** | Prevents conflicting changes (merge conflicts); enables parallel development; supports code review workflows. |

---

## Version Control Security

Version control systems (Git, SVN) are critical security infrastructure:

| Control | Description |
|---|---|
| **Access controls** | Apply least privilege: not all developers need write access to all branches. Separate read from write from admin permissions. |
| **Branch protection** | Protect main/production branches: require pull requests, mandatory code reviews, and passing CI checks before merge. Prevent direct pushes. |
| **Signed commits** | Developers sign commits with their GPG/SSH key to prove authorship. Enables detection of impersonated or tampered commits. |
| **Audit logs** | All repository actions (commits, merges, force pushes, permission changes) logged. Essential for forensic investigation of supply chain incidents. |
| **Secrets scanning** | Automated scanning of commits for accidentally included credentials, API keys, or certificates. |

---

## Patch Management

**Patch management** is a subset of change management focused on applying security fixes and updates:
- Patches most commonly address security vulnerabilities but may also add functionality.
- Must go through the change management process (test in staging → approve → deploy to production).
- Failure to patch is one of the most common causes of breach.
- Monitoring and periodic evaluation help identify when patching is needed (destination-cissp §8.1 — Operation and Maintenance).

---

## Security on Change Control Committees

Security must be a member of the **Change Advisory Board (CAB)** or equivalent change approval body. This is commonly absent in organizations, leading to:
- Security-impacting changes approved without security review
- Emergency changes bypassing security controls
- Security-relevant patches delayed indefinitely

> **Exam note:** The source explicitly calls out that many organizations fail to include security as part of change control committees. Security's role is to ensure changes meet security requirements before approval.

---

## Cross-Links

- [CI/CD Security](ci-cd-security.md) — automated implementation of change in DevSecOps pipelines
- [SDLC](sdlc.md) — change management during operations phase of SDLC/SLC
- [Domain 7 — Security Operations](../domains/07-security-operations.md) — change management in operations context

## Sources

- destination-cissp §8.1 (Change Management review, Operation and Maintenance) (pp. 0940–0941)
- destination-cissp §8.2.1 (SCM, code repositories) (pp. 0948)
- cissp-exam-outline (Domain 8.1: Change management, software configuration management)
