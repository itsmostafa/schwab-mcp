---
title: "ITIL v4 — IT Infrastructure Library"
type: standard
domain: 7
tags: [ITIL, ITSM, change-management, incident-management, problem-management, service-management]
sources: [cissp-exam-outline]
updated: 2026-05-13
---

# ITIL v4 — IT Infrastructure Library

## Bibliographic Information

- **Full title:** Information Technology Infrastructure Library (ITIL), version 4
- **Issuing body:** AXELOS (a joint venture; originally developed by the UK government's CCTA)
- **Current version:** ITIL 4 (2019)
- **Purpose:** A comprehensive framework of best practices for IT service management (ITSM). Describes how IT services should be planned, delivered, and continuously improved.
- **CISSP relevance:** Domain 7 references ITIL change management and incident/problem management processes.

---

## ITIL Service Management Framework

ITIL v4 is organized around the **Service Value System (SVS)**, which includes:
- **Guiding principles** (e.g., focus on value, collaborate, keep it simple)
- **Governance**
- **Service value chain** — activities that create value: plan, improve, engage, design & transition, obtain & build, deliver & support
- **34 management practices** (formerly "processes" in earlier ITIL versions)
- **Continual improvement**

For the CISSP, the key ITIL practices are **change management**, **incident management**, and **problem management**.

---

## ITIL Change Management (Change Enablement in ITIL 4)

ITIL defines three change types:

| Change Type | Description | Authorization |
|---|---|---|
| **Standard change** | Low-risk, pre-approved, frequently performed. Follows a documented procedure. | Pre-authorized (no CAB needed) |
| **Normal change** | Requires assessment and approval via the Change Advisory Board (CAB). Timeline varies by urgency. | CAB |
| **Emergency change** | Must be implemented immediately due to a critical incident or vulnerability. Expedited approval. Post-implementation review still required. | Emergency CAB (ECAB) or delegated authority |

The ITIL **Change Advisory Board (CAB)** is cross-functional and includes representatives from affected business units, IT operations, and security. The CAB meets on a defined schedule to review and approve normal changes.

Key ITIL change management steps (closely mirrors destination-cissp's model):
1. Request for change (RFC)
2. Assessment (impact, risk, cost, rollback plan)
3. Authorization (CAB, standard pre-auth, or ECAB)
4. Planning and scheduling
5. Implementation
6. Review and closure (update CMDB, close RFC)

---

## ITIL Incident Management

ITIL distinguishes:
- **Event** — any change of state significant enough to be managed.
- **Alert** — a notification that a threshold has been crossed.
- **Incident** — an unplanned interruption to, or reduction in quality of, a service.

ITIL's incident management goal: **restore normal service as quickly as possible** with minimum disruption to the business.

Steps:
1. Detect and log
2. Classify and prioritize (by impact + urgency)
3. Diagnose and investigate
4. Resolve and recover
5. Close

Note: ITIL incident management focuses on **service restoration**, while NIST SP 800-61 focuses on **security incident handling**. The CISSP exam tests both but in their respective contexts.

---

## ITIL Problem Management

ITIL separates **incidents** (symptoms) from **problems** (root causes):
- **Incident management** = restore service quickly (reactive)
- **Problem management** = find and eliminate root causes to prevent future incidents (proactive + reactive)

A **Known Error** is a problem whose root cause is identified but not yet resolved. Known errors are documented in the Known Error Database (KEDB) and used to speed up incident diagnosis.

---

## ITIL vs. NIST 800-61 for CISSP

| Aspect | ITIL | NIST SP 800-61 |
|---|---|---|
| Focus | IT service management (all services) | Security incidents specifically |
| Incident framing | Service disruption | Security event causing adverse impact |
| Change management | Detailed, formal (RFC → CAB → ECAB) | Referenced but not the focus |
| Best for CISSP | Change management questions | Incident response phase questions |

---

## Exam-Relevant Nuance

- ITIL's **three change types** (standard, normal, emergency) and the **CAB** are high-yield for Domain 7 change management questions.
- Distinguish **incident** (restore service) from **problem** (eliminate root cause) — the CISSP exam may test this distinction.
- The **ECAB (Emergency CAB)** handles emergency changes — not a regular CAB meeting, not a single person's unilateral decision.
- ITIL **does not define security-specific practices** — for security incident handling, NIST 800-61 is the authoritative standard.

---

## Cross-Links

- [Change Management](../concepts/change-management.md) — concept page with the operational process
- [Incident Management](../concepts/incident-management.md) — security incident handling vs. ITIL service incidents
- [NIST SP 800-61](nist-sp-800-61.md) — security-specific incident handling standard

## Sources

- ITIL v4 (AXELOS, 2019) — primary source
- cissp-exam-outline (Domain 7: change management processes)
