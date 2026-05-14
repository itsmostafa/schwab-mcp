---
title: "Disaster Recovery Sites"
type: concept
domain: 7
tags: [DR-sites, cold-site, warm-site, hot-site, mobile-site, redundant-site, recovery, RTO, geographic-disparity]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Disaster Recovery Sites

## Purpose

When a primary site becomes unavailable (due to a disaster), the organization must operate from an alternate location. Recovery site strategies balance the **cost** of maintaining the alternate site against the **RTO** (Recovery Time Objective) required by the business.

The core tradeoff: **the faster the recovery needed, the more expensive the solution.**

---

## Recovery Site Types Comparison

| Site Type | Infrastructure / HVAC | Basic Equipment (racks, cabling) | Computer Hardware | Data | People On-site | Recovery Time | Cost |
|---|---|---|---|---|---|---|---|
| **Cold** | Yes | Maybe | No | No | No | Weeks | $ (cheapest) |
| **Warm** | Yes | Yes | No | No | No | Days | $$ |
| **Hot** | Yes | Yes | Yes | No (must be restored) | No | Hours | $$$ |
| **Mobile** | Built-in | Built-in | Yes | No (must be restored) | No | Days to hours | $$$ |
| **Redundant** | Yes | Yes | Yes | Yes (real-time sync) | Yes | Instant / seconds | $$$$ (most expensive) |

*Source: destination-cissp §7.10.6 (Table 7-16, Fig. 7-16)*

---

## Site Descriptions

### Cold Site
An **empty building shell** with basic infrastructure (power, HVAC) but no computer or networking equipment. Everything must be procured, installed, and configured after the disaster is declared.

- **Advantage:** Cheapest option.
- **Disadvantage:** Takes **weeks** to become operational. Longest RTO.
- Use case: Organizations with high MTD tolerance and tight budgets.

### Warm Site
In addition to the building shell, **basic equipment is pre-installed** (racks, cabling, some hardware) but servers, networking gear, data, and people are not present.

- **Advantage:** Faster than cold — can be operational in **days**.
- **Disadvantage:** Still requires significant setup time; more expensive than cold.

### Hot Site
A **fully equipped facility** with servers and networking hardware already in place. Only data (most recent backup) and people are needed to begin operations.

- **Advantage:** Can be operational in **hours**.
- **Disadvantage:** Significantly more expensive; data must still be restored and verified.
- Most common choice for organizations with strict RTO requirements.

### Mobile Site
A **hot site on wheels** — a mini data center built inside a shipping container or trailer. Can be transported to wherever it's needed after a disaster.

- **Advantage:** Deployable to disaster location; comparable speed to hot site (hours once moved).
- **Disadvantage:** Transit time is the primary delay variable; expensive.
- Widely used by government agencies for hurricane and disaster response.

### Redundant Site
A **fully operational mirror** of the primary site with data synchronized in real time and people already present.

- **Advantage:** Can fail over **instantly or in seconds**. Highest availability.
- **Disadvantage:** Cost equals or exceeds the primary site — essentially two full data centers.
- Use case: Organizations with near-zero RTO (banks, stock exchanges, critical infrastructure).

---

## Geographic Disparity

Recovery sites must be **geographically remote** from the primary site — far enough that a regional disaster (hurricane, earthquake, power grid failure) affecting the primary site does not simultaneously affect the recovery site.

General guideline: East Coast primary → Midwest or West Coast recovery (for North American organizations).

---

## Internal vs. External Recovery Sites

| Type | Owner | Example |
|---|---|---|
| **Internal** | The organization | Organization's own secondary data center |
| **External** | Third-party service provider | Sungard, Iron Mountain, cloud-based DR providers |

**Reciprocal agreements** — two organizations agree to host each other's recovery operations if either suffers an outage. In practice, reciprocal agreements are rare in private enterprise due to capacity, confidentiality, and compatibility concerns.

**Resource capacity agreements** — agreements with vendors guaranteeing access to specific resources (hardware, bandwidth, cloud capacity) during a declared disaster.

**Multiple processing sites** — multiple active locations running key business functions simultaneously, with geographically dispersed redundancy (common in credit card processing, financial services). Expensive to architect but very reliable.

---

## Connecting Site Type to RTO/RPO

| RTO Required | Site Type to Consider |
|---|---|
| Weeks | Cold site |
| Days | Warm site or mobile site |
| Hours | Hot site |
| Minutes/Seconds | Redundant site |

Lower RTO → more expensive site type.

Higher RPO (more data loss tolerable) → less real-time data sync needed → less expensive.

---

## Spare Parts (component-level equivalent)

For individual components within a system:

| Spare Type | Location | Activation Time |
|---|---|---|
| **Cold spare** | On a shelf (storage room) | Longest — system offline until retrieved and installed |
| **Warm spare** | Installed but not powered | System goes offline briefly; failover by switching to spare |
| **Hot spare** | Installed, powered, synchronized | Instantly takes over when primary fails — no downtime |

*Source: destination-cissp §7.10.3*

---

## RAID and High Availability (Component Level)

RAID and clustering/redundancy complement the recovery site strategy at the component and system level:

| RAID Level | Mechanism | Min Drives | Key Benefit |
|---|---|---|---|
| RAID 0 (Striping) | Splits data across drives | 2 | Speed (no redundancy) |
| RAID 1 (Mirroring) | Writes identical data to multiple drives | 2 | Availability/redundancy |
| RAID 5 (Parity) | Striping + parity bit | 3 | Balance of speed and redundancy |
| RAID 10 (Mirror + Stripe) | Combines RAID 0 and RAID 1 | 4 | Speed + redundancy (expensive) |

*Source: destination-cissp §7.10.4*

---

## Exam-Relevant Nuance

- **Cold site = cheapest, longest RTO; redundant = most expensive, shortest RTO.** This linear relationship is heavily tested.
- **Data is NOT present** at cold, warm, or hot sites until backups are restored — only the redundant site has live synchronized data.
- The exam may describe a scenario and ask which site type is appropriate — match RTO requirement to site type.
- **Fail-safe** (doors unlock in a fire) is distinct from **fail-secure** (system locks down if it fails) — separate concept from DR sites but tested in the same domain.

---

## Cross-Links

- [RTO/RPO/MTD](rto-rpo-mtd.md) — RTO drives site selection
- [Backup Strategies](backup-strategies.md) — backup data must be transported to or staged at the recovery site
- [BCP/DRP Operations](bcp-drp-operations.md) — DR site is activated as part of the DRP

## Sources

- destination-cissp §7.10.3–7.10.6 (pp. 0892–0901)
- cissp-exam-outline (Domain 7: recovery site strategies)
