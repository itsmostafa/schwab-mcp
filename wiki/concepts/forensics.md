---
title: "Digital Forensics"
type: concept
domain: 7
tags: [forensics, digital-forensics, live-evidence, forensic-imaging, write-blocker, artifacts, anti-forensics, mobile-forensics]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Digital Forensics

## Definition

Digital forensics is the **scientific examination and analysis of data from storage media** in such a way that the information can be used as part of an investigation to identify a culprit or the root cause of an incident. The process must maintain evidence integrity and the chain of custody so findings can be presented in legal proceedings.

---

## Forensic Investigation Process

| Phase | Activities |
|---|---|
| **Identify & Secure the Scene** | Seal off the area; photograph and document; avoid touching anything; begin chain of custody. |
| **Collect Evidence** | Gather physical and digital evidence using proper tools; preserve integrity; maintain chain of custody. |
| **Examine Evidence** | Media analysis, software analysis, network analysis using automated and manual techniques. |
| **Analyze Evidence** | Focus on most relevant evidence; build hypotheses; attribute actions to a suspect. |
| **Report** | Document every facet of the process for all stakeholders (prosecution, defense, judge/jury, regulators, insurers). |

Documentation must occur at every step, not only at the end.

---

## Live Evidence

**Live evidence** is data stored in a running system: RAM, CPU registers, cache, network connections, running processes, clipboard contents. If the system is powered off or the mouse is moved, live evidence changes or disappears permanently.

**Order of Volatility** (most volatile → least volatile):

| Priority | Location |
|---|---|
| 1 | CPU registers, cache |
| 2 | RAM (including routing tables, ARP cache, process list) |
| 3 | Swap space / paging file |
| 4 | Hard disk / SSD |
| 5 | Remote logging / monitoring data |
| 6 | Archived media (tapes, optical) |

Collect in this order to capture the most ephemeral evidence first. Examining a live system inevitably changes its state; specialist tools minimize contamination.

*Source for live-evidence framing: destination-cissp §7.1.4. The canonical order-of-volatility list is standard CISSP knowledge per RFC 3227.*

---

## Forensic Copies (Dead Acquisition)

When a system is already powered off, the primary evidence source is the **hard drive**:

1. Create **two bit-for-bit (sector-by-sector) copies** of the original drive using a forensic imaging tool.
2. Hash all three (original + two copies) with a cryptographic hash (e.g., MD5, SHA-256). Matching hashes confirm integrity.
3. Place the **original** in a sealed evidence bag — never touched again.
4. Store **Copy 1** as an archival copy — also sealed.
5. Use **Copy 2** as the working copy for analysis.

A **write blocker** (hardware or software) prevents any writes to the original drive during the imaging process, ensuring integrity.

---

## Forensic Imaging and Write Blockers

- **Write blockers** intercept write commands to the source drive; all reads pass through normally. Hardware write blockers are preferred in legal proceedings.
- **Forensic imaging tools**: FTK Imager, dd (Unix), Guymager, EnCase. They create bit-for-bit images and compute hash values.
- The hash value serves as the **digital fingerprint** of the evidence — it proves the copy is identical to the original and that the original was never modified.

---

## Investigative Analysis Techniques

| Technique | Focus |
|---|---|
| **Media analysis** | Hard drives, SSDs, USB drives, tapes. Searches for deleted files (pointer removed, data may remain), file signatures, hidden data. |
| **Software / malware analysis** | Determines how malware works, what it targets, and potentially who created it (attribution via code patterns). See [Malware Analysis](malware-analysis.md). |
| **Network analysis** | Traces how a network was penetrated, how the attacker traversed it, which systems were compromised. Log files are the primary source. |

---

## Forensic Artifacts

**Artifacts** are remnants of a breach or attempted breach — breadcrumbs that point to an attacker's path and actions:

- IP addresses, domain names, URLs
- Cryptographic hashes of known-malicious files (IOCs)
- Registry keys (Windows), file names/types, file sizes
- Browser history, cookies, cache
- Log entries: account changes, privilege escalations, file modifications

Sources: computer systems, web browsers, mobile devices, hard drives, flash drives.

Identifying relevant artifacts is like finding a needle in a haystack — forensic examiners must focus on artifacts that support or refute specific hypotheses.

---

## Mobile Device Forensics

Mobile forensics is more challenging than PC forensics because:
- Manufacturers frequently change OS structure, file systems, and connectors.
- No single method or tool can extract all data from all devices.
- Hibernation and app suspension complicate live acquisition.
- New training is constantly required for examiners.

---

## Anti-Forensics

Attackers use techniques to impede forensic investigation:
- **File deletion / secure wiping** (overwriting data to prevent recovery)
- **Encryption** of data or drives
- **Timestamp manipulation** (changing file metadata)
- **Log tampering / log clearing**
- **Steganography** (hiding data in plain sight, e.g., in image files)
- **Rootkits** that hide files, processes, and network connections from the OS

---

## Exam-Relevant Nuance

- **Live acquisition vs. dead acquisition:** Live acquisition captures volatile data but risks changing the evidence state; dead acquisition (powered-off system) is safer for disk evidence but loses volatile data.
- The **best evidence rule** prefers the original. When the original cannot be presented, a forensic copy (with verified hash) is the next best option.
- Hash verification proves **integrity**, not authenticity — you still need chain of custody documentation to prove the evidence was never tampered with.
- **Five rules of evidence**: authentic, accurate, complete, convincing/reliable, admissible.

---

## Cross-Links

- [Chain of Custody](chain-of-custody.md) — evidence handling, legal admissibility
- [Incident Management](incident-management.md) — forensics is performed during IR
- [Malware Analysis](malware-analysis.md) — software analysis sub-discipline
- [Logging & Monitoring](logging-monitoring.md) — log files as forensic evidence

## Sources

- destination-cissp §7.1.1–7.1.6 (pp. 0824–0841)
- cissp-exam-outline (Domain 7 subtopic: digital forensics tools, tactics, and procedures)
- RFC 3227 — Guidelines for Evidence Collection and Archiving (order of volatility)
