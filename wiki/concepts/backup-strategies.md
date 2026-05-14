---
title: "Backup Strategies"
type: concept
domain: 7
tags: [backup, incremental, differential, full, mirror, archive-bit, 3-2-1, CRC, tape-rotation, cloud-backup]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Backup Strategies

## Purpose

Backup strategies ensure that data can be recovered if the primary data store is lost, corrupted, stolen, or otherwise compromised. The strategy chosen is driven by organizational goals: how quickly must data be restored (RTO) and how much data can be lost (RPO)?

---

## The Archive Bit

The **archive bit** is metadata attached to every file in a filesystem:
- **0** = file does not need to be backed up (no changes since last backup, or backup was just performed)
- **1** = file needs to be backed up (new file, or modified since last backup)

Different backup strategies treat the archive bit differently — this is the key technical differentiator between incremental and differential backups.

*Source: destination-cissp §7.10.2*

---

## Backup Types Comparison

| Type | Data Backed Up | Archive Bit After Backup | Backup Time | Restore Time | Storage |
|---|---|---|---|---|---|
| **Full** | ALL data, regardless of archive bit | Reset to 0 | Slowest | Fast (1 tape set) | High |
| **Incremental** | Changes since the **last backup** (any type) | Reset to 0 | Fast (small sets at end of week) | Slow (full + every incremental) | Lowest |
| **Differential** | Changes since the **last full backup** | NOT reset (stays at 1) | Moderate → slow by end of week | Moderate (full + latest differential only) | Moderate |
| **Mirror** | Exact copy, no compression | N/A | Fastest | Fastest | Highest |

*Source: destination-cissp §7.10.2 (Table 7-12)*

### Key distinction: Incremental vs. Differential

- **Incremental**: resets archive bit to 0 after each backup → each incremental only captures *that day's* changes. Restoration requires the full backup + every incremental tape.
- **Differential**: does NOT reset the archive bit → each differential captures *all changes since the last full*. Restoration requires only 2 tapes: full + latest differential.

---

## Backup Storage Locations

| Location | Description | Trade-off |
|---|---|---|
| **Onsite** | Same location as primary data | Fast access; same disaster risk |
| **Offsite** | Geographically separate location | Protected against site-level disaster; slower access |
| **Cloud** | Stored with cloud provider | High availability, scalable cost, inherently offsite |

**Geographically remote** = far enough from the primary site that a regional disaster (hurricane, earthquake) would not affect both locations simultaneously.

### 3-2-1 Rule

The widely-cited backup best practice:
- **3** copies of data (1 primary + 2 backups)
- **2** different storage media types
- **1** offsite copy

The source does not explicitly state "3-2-1" — this is standard CISSP knowledge per cissp-exam-outline.

---

## Tape Rotation Strategies

Tape rotation determines when a tape is used, how long it is retained, and when it is reused:
- **FIFO (First In, First Out)** — simplest; oldest tape is overwritten first.
- **Grandfather-Father-Son (GFS)** — hierarchical: daily (son), weekly (father), monthly (grandfather) tapes. Allows point-in-time recovery for different timeframes.
- **Tower of Hanoi** — a mathematical rotation scheme that balances media wear and retention period efficiency. More complex to manage.

---

## Electronic Vaulting

**Electronic vaulting** typically refers to an automated tape management system (tape jukebox) with robotic arms that manage tape insertion, removal, and scheduling. Used in large enterprises for high-volume backup automation.

---

## Backup Verification: CRC

A **Cyclic Redundancy Check (CRC)** — also called a checksum — is used to verify data integrity after backup. The CRC is computed on the backed-up data and stored alongside it. When the data is later restored, the CRC is recomputed and compared to the stored value. If they match, data integrity is confirmed.

CRC can be applied to data at rest (disk), data in motion (network), and in RAM.

*Source: destination-cissp §7.10.2*

---

## Backup Testing

Backups have no value if they cannot be restored. Organizations must regularly test restoration:
- Perform test restores to verify data can be recovered.
- Document restoration procedures so personnel can execute them under pressure.
- Include backup restoration in DRP exercises.

---

## Media Considerations

- **MTBF** of media determines reliability — no media is reliable forever; data must be migrated to new media periodically.
- File formats should be updated to maintain compatibility with current applications.
- Cryptographic protection applied at backup time should use current standards; old ciphers may be broken by the time the data is needed.

---

## Exam-Relevant Nuance

- **Incremental: faster to backup, slower to restore. Differential: slower to backup, faster to restore.** The exam loves this distinction.
- The archive bit behavior (reset vs. not reset) is the technical root of the incremental vs. differential difference.
- Backup selection is driven by RPO and RTO: low RPO → frequent backups or continuous replication; short RTO → fast restore strategy (mirror or differential preferred over incremental).
- Cloud backup is **not a substitute** for testing — verify that cloud backups can actually be restored within RTO.

---

## Cross-Links

- [RTO/RPO/MTD](rto-rpo-mtd.md) — RPO drives backup frequency; RTO drives restore strategy
- [Disaster Recovery Sites](disaster-recovery-sites.md) — recovery sites house the backup data
- [BCP/DRP Operations](bcp-drp-operations.md) — backup is a key element of the overall BCM strategy

## Sources

- destination-cissp §7.10.2 (pp. 0887–0892)
- cissp-exam-outline (Domain 7: backup storage strategies)
