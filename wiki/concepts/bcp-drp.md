---
title: "Business Continuity Planning and Disaster Recovery"
type: concept
domain: 1
tags: [bcp, drp, bia, rto, rpo, mtd, continuity, recovery, disaster]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Business Continuity Planning and Disaster Recovery

BCP and DRP are overlapping but distinct disciplines. The CISSP exam tests the distinction between
them and the key metrics: MTD, RTO, and RPO.

## BCP vs. DRP

| | BCP (Business Continuity Planning) | DRP (Disaster Recovery Planning) |
|---|---|---|
| **Focus** | Keeping critical business functions running during a disruption | Restoring IT systems and infrastructure after a disruption |
| **Scope** | Organization-wide; includes non-IT processes | Primarily IT/technical recovery |
| **Timing** | During the event (keeping things running) | After the event (restoring to normal) |
| **Ownership** | Business unit leaders, senior management | IT and security operations |

BCP encompasses DRP — DRP is a subset of the overall BCP program.

## Business Impact Analysis (BIA)

The BIA is the foundational activity for both BCP and DRP. It:
- Analyzes the **consequences** of a disaster to the organization.
- Identifies critical processes, functions, and assets.
- Establishes **priorities** for recovery.
- Gathers information needed to develop recovery strategies.
- Maps **external dependencies** (suppliers, vendors, utilities) that affect critical functions.

## Key Metrics

| Metric | Definition | Exam Tip |
|---|---|---|
| **MTD** — Maximum Tolerable Downtime | The maximum time a business process can be disrupted before causing unacceptable harm | Sets the upper bound; RTO must be ≤ MTD |
| **RTO** — Recovery Time Objective | Target time within which systems/processes must be restored after a disruption | Operational target for recovery teams |
| **RPO** — Recovery Point Objective | Maximum acceptable amount of data loss measured in time (how far back can you restore from?) | Drives backup frequency decisions |

**Relationship**: RPO ≤ RTO ≤ MTD

## BCP/DRP Test Types (ordered by comprehensiveness)

| Test Type | Description | Disruption |
|---|---|---|
| **Tabletop Exercise** | Discussion-based; team walks through scenarios verbally | None — paper-based only |
| **Walkthrough / Structured Walk-through** | More detailed than tabletop; step-by-step review of the plan | Minimal |
| **Simulation** | Simulates a disaster scenario without activating actual recovery systems | Low |
| **Parallel Test** | Activates the alternate site/recovery systems while primary site continues to operate | Low — both run simultaneously |
| **Full Interruption / Cutover Test** | Primary site is taken offline; full failover to recovery site | High — actual disruption |

## Exam-Relevant Nuance

- **BCP** = keeping business running; **DRP** = restoring systems after disaster. On a scenario question, if the focus is IT restoration — it's DRP. If it's keeping people working during the event — it's BCP.
- The source (destination-cissp) covers BIA briefly at §1.7 and defers full BCP/DRP coverage to Domain 7 (§7.11). Expect detailed exam questions on BCP/DRP to draw from Domain 7 content.
- **MTD** sets the constraint; **RTO** is the goal. An organization with an MTD of 24 hours must have an RTO ≤ 24 hours.
- Parallel test is safer than full interruption — use when you can't risk the disruption of a full cutover.

## Cross-links

- [Risk Management](risk-management.md)
- [Governance](governance.md)
- [SCRM](scrm.md)

## Sources

- destination-cissp §1.7 (BIA overview, external dependencies)
- Domain 7 (§7.11) for full BCP/DRP coverage (not yet ingested)
