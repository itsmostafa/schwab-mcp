---
title: "Change Management"
type: concept
domain: 7
tags: [change-management, CAB, RFC, ITIL, change-control, rollback, baseline]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Change Management

## Definition

**Change management** is the formal process that ensures all changes to systems, configurations, and processes are analyzed, approved, built, tested, implemented, and documented in a controlled manner. It reduces the risk that unauthorized or poorly planned changes introduce vulnerabilities, outages, or compliance gaps.

> "The discipline and rigor that a company places upon change management is directly proportional to how well a company operates." — destination-cissp §7.9.1

---

## Change Management Process Steps

| Step | Description |
|---|---|
| **1. Change Request (RFC)** | A **Request for Change** can originate from any part of the organization — IT, business owners, security, compliance, or automated vulnerability tools. Changes are tracked in change management software. |
| **2. Assess Impact** | Evaluate the scope and risk of the proposed change. How big is it? Who does it affect? What systems and processes are impacted? |
| **3. Approval** | Based on impact, determine the appropriate level of review. Minor changes → minimal approvals; significant/costly/multi-stakeholder changes → higher levels including the **CAB**. Emergency changes use an expedited approval path. |
| **4. Build and Test** | Develop and test the change in a non-production environment. Includes regression testing (does existing functionality still work?) and validation of new functionality. |
| **5. Notification** | Notify key stakeholders **before** implementation so they can prepare, adjust SLAs, or escalate concerns. |
| **6. Implement** | Apply the change to production. |
| **7. Validation** | Notify senior management and stakeholders to confirm the change works as expected. |
| **8. Version and Baseline** | Update all documentation, CMDB entries, and baselines to reflect the new state. This step is critical for CM integrity. |

*Source: destination-cissp §7.9.1 (Fig. 7-9, Table 7-10)*

---

## Change Advisory Board (CAB)

The **Change Advisory Board (CAB)** — also called the Change Control Board (CCB) — is a cross-functional committee that reviews and approves significant changes. Membership includes key stakeholders from throughout the organization (IT, security, operations, business owners, legal) to ensure all perspectives are considered.

For minor changes, CAB review may not be required. For major changes, CAB approval is mandatory before implementation.

---

## Emergency Change Procedures

When a critical vulnerability or system failure demands an immediate change (e.g., emergency security patch):
- An **expedited approval** path bypasses the normal CAB meeting cycle.
- The change still requires documentation and (often) post-implementation review by the CAB.
- Rollback plans are especially important for emergency changes, since testing time is compressed.
- Emergency patches approved in emergency change management are an exception — not the norm.

---

## Rollback Plans

Every change must have a documented **rollback plan** — a tested procedure for reverting to the prior state if the change causes problems. Without a rollback plan, a failed change can cause extended outages.

---

## Relationship to Patch Management

Every patch deployment is a change and must flow through the change management process. This ensures:
- Impact is assessed before patching production systems.
- Rollback plans exist if a patch causes issues.
- Documentation is updated after successful patching.

---

## ITIL Change Management

ITIL (IT Infrastructure Library) provides a widely-used service management framework that formalizes change management. ITIL distinguishes:
- **Standard changes** — low-risk, pre-approved, well-understood (e.g., routine patching).
- **Normal changes** — require full CAB review.
- **Emergency changes** — urgent; expedited approval with post-implementation review.

See [ITIL v4](../standards/itil-v4.md) for more detail.

---

## Exam-Relevant Nuance

- Change management is about **deliberate, documented, approved** change — not preventing change.
- **Too little** change management → chaotic, reactive, uncontrolled environment.
- **Too much** change management → paralysis; people circumvent the process, also leading to chaos.
- The exam will ask: who can approve an emergency change? → Senior management / designated authority. Not a low-level IT admin.
- Configuration management and change management are a pair: CM documents what the baseline *is*; change management controls *how it changes*.
- Documentation at the **version and baseline** step is often the overlooked step that trips organizations up.

---

## Cross-Links

- [Configuration Management](configuration-management.md) — CM baselines that change management protects
- [Patch Management](patch-management.md) — patch deployments are changes
- [ITIL v4](../standards/itil-v4.md) — formal framework for change management

## Sources

- destination-cissp §7.9–7.9.1 (pp. 0883–0885)
- cissp-exam-outline (Domain 7: change management processes)
