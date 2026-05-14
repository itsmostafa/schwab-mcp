---
title: "BCP/DRP Operations"
type: concept
domain: 7
tags: [BCP, DRP, BCM, business-continuity, disaster-recovery, crisis-communications, alternate-sites, BIA]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# BCP/DRP Operations

## Hierarchy of Terms

| Term | What It Is | Focus |
|---|---|---|
| **BCM** (Business Continuity Management) | The overarching function; creates and maintains BCP and DRP | Structure, policy, oversight |
| **BCP** (Business Continuity Plan/Planning) | Strategic plan for keeping critical business processes operating | Business processes; survival of the organization |
| **DRP** (Disaster Recovery Plan/Planning) | Tactical plan for restoring IT infrastructure and systems | Technology; returning to BAU |

> "BCP focuses on the processes; DRP focuses on the systems." — destination-cissp §7.11.1

A **disaster** = any sudden event that creates an inability to support critical business functions for a predetermined period. When MTD is at risk of being exceeded, a **disaster should be declared** and the DRP activated.

---

## BCM: Three Primary Goals (in priority order)

1. **Safety of people** — always the first priority.
2. **Minimization of damage** to facilities and the business.
3. **Survival of the business** and its critical functions.

*Source: destination-cissp §7.13.1*

---

## BCP/DRP Development Steps

| Step | Activity |
|---|---|
| 1. Develop contingency planning policy | Formal policy providing authority and guidance for BCP/DRP. |
| 2. Conduct BIA | Identify and prioritize critical functions, systems, and components; determine RPO, RTO, WRT, MTD. |
| 3. Identify preventive controls | Reduce the probability or impact of disruptions. |
| 4. Create contingency strategies | Ensure systems can be recovered quickly and effectively. |
| 5. Develop contingency plan | Write the actual BCP and DRP documents. |
| 6. Test, train, and exercise | Testing validates recovery capabilities; training prepares personnel; exercises identify gaps. |
| 7. Maintenance | Plans are living documents — update them regularly as systems and the organization change. |

*Source: destination-cissp §7.11.1 (Table 7-19)*

---

## Business Impact Analysis (BIA)

The **BIA** is the most important step in the BCP process. It:
- Identifies and prioritizes critical business processes, functions, and systems.
- Determines the impact (financial, operational, reputational) of disruption.
- Establishes **RPO, RTO, WRT, and MTD** for each critical function.
- Identifies resource requirements and dependencies (people, equipment, data, facilities).
- Establishes restoration priority order.

The BIA output drives all subsequent recovery strategy decisions.

*Source: destination-cissp §7.11.3*

---

## Declaring a Disaster

**Who declares:** CEO, BCM Board, or another authoritative executive entity.
**When to declare:** When MTD is expected to be exceeded by the incident at hand.

Process:
1. Incident response process is followed.
2. Impact assessment includes MTD evaluation.
3. If MTD will be exceeded → declare disaster → activate DRP.
4. Emergency response team is activated.

---

## Emergency Response Team Composition

Personnel should represent all major functional areas:
- Executive / senior management
- IT
- Legal
- HR
- PR / communications
- Security

---

## Crisis Communications

Communications must reach all relevant stakeholders:

**Internal:** Senior management, board, business owners, legal, HR, PR/comms team.

**External:** Regulators, law enforcement, customers, media, business partners.

One designated **spokesperson** ensures message consistency and allows the response team to remain focused on technical recovery.

---

## Restoration Order

When moving to the DR site: **most critical systems first** (BIA-driven priority).

When returning to the primary site after it is rebuilt: **least critical systems first** — to test that the rebuilt site works correctly before trusting it with the most critical workloads.

Dependency charts map out which underlying components (load balancers, databases, web servers) must be online before a given service can be restored.

*Source: destination-cissp §7.11.5*

---

## External Dependencies

BCP must account for external dependencies (e.g., fuel suppliers for generators, cloud providers, third-party logistics). If a disaster disrupts an external dependency, the plan must include alternatives.

---

## Exam-Relevant Nuance

- Not all business functions are critical or essential — BIA helps identify which ones are.
- Even during a disaster, the organization must still comply with laws, regulations, and privacy requirements.
- BCP is **strategic** (survival); DRP is **tactical** (technical recovery).
- MTD = RTO + WRT. If this equation is violated (RTO alone exceeds MTD), the organization cannot survive the disaster through normal recovery.
- Safety of people is **always** the top BCM priority — it will appear on the exam.

---

## Cross-Links

- [RTO/RPO/MTD](rto-rpo-mtd.md) — key time measurements that drive BCP/DRP decisions
- [Disaster Recovery Sites](disaster-recovery-sites.md) — alternate site types and their RTOs
- [Backup Strategies](backup-strategies.md) — backup drives RPO decisions
- [DRP Test Types](../domains/07-security-operations.md#712-test-disaster-recovery-plans-drp) — see subtopic 7.12 in domain page

## Sources

- destination-cissp §7.11–7.13 (pp. 0903–0917)
- cissp-exam-outline (Domain 7: disaster recovery, business continuity planning and exercises)
