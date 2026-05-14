---
title: "Failure Modes"
type: concept
domain: 7
tags: [failure-modes, fail-safe, fail-secure, fail-open, high-availability, resilience]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Failure Modes

## Definition

**Failure modes** define the state a system enters when it fails. The appropriate failure mode depends on what the system must protect — security, availability, or human safety.

---

## Three Failure Modes

| Mode | Also Called | What Happens on Failure | Priority | Example |
|---|---|---|---|---|
| **Fail-soft** | Fail-open | System fails into a **less secure** state | Availability | Firewall fails → allows all traffic through |
| **Fail-secure** | Fail-closed | System fails into a **same or more secure** state | Security | Firewall fails → blocks all traffic |
| **Fail-safe** | — | System fails in a way that **prioritizes human safety** | Human safety | Door lock fails → doors unlock automatically (fire egress) |

*Source: destination-cissp §7.10.1*

---

## Key Distinctions

- **Fail-soft (fail-open)**: maximizes availability at the cost of security. Acceptable when continuous operation is more critical than security (e.g., a hospital patient monitoring system that cannot afford to lock staff out during a failure).
- **Fail-secure (fail-closed)**: maximizes security at the cost of availability. Preferred for security controls (firewalls, access control systems). Example: an access door that defaults to locked if the electronic lock fails.
- **Fail-safe**: the override priority is always human safety. A fail-safe mechanism may reduce security (unlocking doors) to protect lives — safety takes precedence over security.

---

## Design Considerations

- Security devices (firewalls, IDS) should generally be **fail-secure**.
- Physical access controls in buildings must be **fail-safe** for fire egress (required by building codes).
- High-availability systems may use **fail-soft** when downtime is unacceptable, with compensating controls.
- The appropriate failure mode should be documented in system design and verified during disaster recovery testing.

---

## Exam-Relevant Nuance

- **Fail-safe ≠ fail-secure.** Students often confuse these. Fail-safe protects people; fail-secure protects data/systems.
- A fire door that unlocks automatically is **fail-safe** — it would be fail-secure if it locked, but that is illegal in most jurisdictions for egress routes.
- The exam may present a scenario and ask which failure mode applies — match the priority (people vs. security) to the mode.

---

## Cross-Links

- [Disaster Recovery Sites](disaster-recovery-sites.md) — high availability concepts
- [BCP/DRP Operations](bcp-drp-operations.md) — BCM's top priority is human safety (aligns with fail-safe)

## Sources

- destination-cissp §7.10.1 (p. 0886)
- cissp-exam-outline (Domain 7: implement recovery strategies)
