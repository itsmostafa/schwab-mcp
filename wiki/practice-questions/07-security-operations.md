---
title: "Practice Questions — Domain 07: Security Operations"
type: practice
domain: 07
tags: [practice]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 07: Security Operations

## Questions

### Q1: An investigator responds to a suspected breach. The suspect's computer is running. What should the investigator collect FIRST?

**Answer:** Contents of RAM (volatile memory) — specifically the running processes, network connections, and any data in RAM before powering off the system.

**Why:** Volatile data (RAM, cache, CPU registers) is lost the moment the system is powered down. The order of volatility mandates collecting the most volatile data first: registers/cache → RAM → swap → disk → remote logs → archives. Starting with the disk (a common mistake) would forfeit the volatile evidence.

*Subtopic: 7.1 / [Forensics](../concepts/forensics.md)*

---

### Q2: A company's nightly backup completes in three hours and can restore data from the previous night's backup. An earthquake destroys the primary data center at 3 PM on a Tuesday. What is the RPO implied by this backup strategy?

**Answer:** Approximately 24 hours (one business day's worth of data).

**Why:** The Recovery Point Objective is the maximum amount of data loss the organization can tolerate, measured in time. If backups run nightly and the disaster strikes mid-afternoon, the organization could lose up to the transactions that occurred since the previous night's backup — roughly 15–24 hours of data. RPO drives backup frequency: to achieve a lower RPO, the organization would need more frequent backups or real-time replication.

*Subtopic: 7.10–7.11 / [RTO/RPO/MTD](../concepts/rto-rpo-mtd.md)*

---

### Q3: An organization's RTO is 6 hours and its WRT is 2 hours. Its MTD is 7 hours. Is this a safe configuration?

**Answer:** No — this is unsafe. The sum of RTO + WRT (6 + 2 = 8 hours) exceeds the MTD of 7 hours.

**Why:** MTD = RTO + WRT. The total time to restore systems (RTO) and verify they work correctly (WRT) must fit within the MTD. If it doesn't, the organization will inevitably exceed its maximum tolerable downtime before fully recovering — meaning the business may fail. The golden rule: RTO must always be less than MTD, and MTD must always be ≥ RTO + WRT.

*Subtopic: 7.11 / [RTO/RPO/MTD](../concepts/rto-rpo-mtd.md)*

---

### Q4: An organization needs to recover from a disaster within four hours. Which recovery site type is MOST appropriate?

**Answer:** A hot site.

**Why:** Recovery site types and their approximate RTOs: cold site (weeks), warm site (days), hot site (hours), redundant site (instant). A four-hour RTO requirement aligns with a hot site, which has all hardware pre-installed and only needs data restored and personnel on site. A warm site would take days and a redundant site would be far more expensive than needed.

*Subtopic: 7.10 / [Disaster Recovery Sites](../concepts/disaster-recovery-sites.md)*

---

### Q5: A disaster has been declared. At the recovery site, which systems should be brought online FIRST?

**Answer:** The most critical systems, as determined by the Business Impact Analysis (BIA).

**Why:** Recovery resources are limited during a disaster. The BIA prioritizes systems by their criticality to the organization's mission. The most critical systems are recovered first so the organization can resume minimum viable operations. When the primary site is later rebuilt and systems are moved back, the order reverses — least critical systems are restored first to validate the new site is working correctly before trusting it with the most critical workloads.

*Subtopic: 7.11 / [BCP/DRP Operations](../concepts/bcp-drp-operations.md)*

---

### Q6: During a forensic investigation, why should an analyst NEVER perform analysis directly on the original hard drive?

**Answer:** To preserve the integrity of the original evidence, which must remain unmodified to be legally admissible and to allow re-analysis if needed.

**Why:** The forensic process requires creating two bit-for-bit copies of the original drive and verifying them with cryptographic hashes. Analysis is performed on the second copy (the working copy). The original is sealed and stored as evidence — never touched. If analysis is performed on the original, any changes (even benign read artifacts on some file systems) could compromise evidence integrity and violate the chain of custody, potentially making the evidence inadmissible.

*Subtopic: 7.1 / [Forensics](../concepts/forensics.md), [Chain of Custody](../concepts/chain-of-custody.md)*

---

### Q7: A security analyst is reviewing evidence collected during an incident. A witness who was not present at the time of the incident provides testimony about what someone told them. What type of evidence is this?

**Answer:** Hearsay evidence.

**Why:** Hearsay evidence is testimony from a witness who was not present at the incident and is reporting what they were told by someone else. Hearsay is generally inadmissible in court unless an exception applies. In contrast, direct evidence (a witness present at the scene), circumstantial evidence (implies a fact by inference), and corroborative evidence (supports other evidence) are distinct categories. Chain of custody helps make evidence admissible, but hearsay has specific admissibility rules.

*Subtopic: 7.1 / [Chain of Custody](../concepts/chain-of-custody.md)*

---

### Q8: An organization performs a weekly full backup every Sunday and daily backups on weekdays. On Friday, a system fails and must be restored. Which backup strategy results in FEWER restore tapes: incremental or differential?

**Answer:** Differential requires fewer tapes (at most 2: Sunday's full + Friday's differential). Incremental requires more tapes (Sunday's full + Monday + Tuesday + Wednesday + Thursday + Friday incrementals = up to 6 tapes).

**Why:** Differential backups capture all changes since the last full backup and do NOT reset the archive bit, so each differential grows larger but only one differential is needed at restore time. Incremental backups capture only changes since the last backup of any kind, reset the archive bit, and thus stay small — but every incremental tape is needed to reconstruct the complete picture. Trade-off: incremental is faster/smaller to run; differential is faster to restore.

*Subtopic: 7.10 / [Backup Strategies](../concepts/backup-strategies.md)*

---

### Q9: Which type of DRP test is the MOST realistic but also carries the HIGHEST risk?

**Answer:** Full-interruption (full-scale) test.

**Why:** A full-interruption test involves both backup/parallel systems and production systems. It is the most realistic test because it validates whether the DRP actually works end-to-end, including production system failover and recovery. However, it carries the highest risk because production systems can be disrupted. It should only be conducted after all other tests (read-through, walkthrough, simulation, parallel) have been successfully completed, and only with explicit management approval.

*Subtopic: 7.12 / [BCP/DRP Operations](../concepts/bcp-drp-operations.md)*

---

### Q10: An employee's workstation is suspected of being infected with malware. The security team wants to understand what the malware does without letting it connect to the internet. What analysis technique should they use?

**Answer:** Dynamic malware analysis in a sandbox (also called behavioral/heuristic analysis).

**Why:** A sandbox is an isolated execution environment where the malware can run and its behavior (file system changes, registry modifications, network connection attempts, process creation) can be observed without affecting production systems or reaching the internet. This is dynamic analysis — as opposed to static analysis, which examines the malware code without running it. Note: sophisticated malware can detect sandbox environments and remain dormant — a limitation of this approach.

*Subtopic: 7.7 / [Malware Analysis](../concepts/malware-analysis.md)*

---

### Q11: An organization needs to make an urgent change to patch a critical zero-day vulnerability being actively exploited. The normal CAB meeting is scheduled for next week. What should the organization do?

**Answer:** Invoke the emergency change management procedure, which allows expedited approval by an Emergency CAB (ECAB) or designated authority, bypassing the normal CAB meeting schedule.

**Why:** Emergency changes exist precisely for situations where a critical security vulnerability requires immediate remediation and waiting for the normal change cycle would create unacceptable risk. The change still requires documentation, a rollback plan, and a post-implementation review — the process is accelerated, not eliminated. Using emergency change procedures improperly (for non-emergencies) undermines the integrity of the change management process.

*Subtopic: 7.9 / [Change Management](../concepts/change-management.md)*

---

### Q12: What is the NUMBER ONE priority of Business Continuity Management (BCM), and what are the other two goals in order?

**Answer:** 1) Safety of people (always first). 2) Minimization of damage to facilities and assets. 3) Survival of the business.

**Why:** Destination-cissp emphasizes this ordering explicitly: human safety overrides all other considerations. An organization may lose data or facilities and survive; loss of human life is irreversible. This is also why fail-safe mechanisms (like doors that unlock during a fire) take precedence over fail-secure mechanisms — safety overrides security when human lives are at stake.

*Subtopic: 7.13 / [BCP/DRP Operations](../concepts/bcp-drp-operations.md)*
