---
title: "Chain of Custody"
type: concept
domain: 7
tags: [chain-of-custody, evidence, forensics, legal, admissibility, Locard]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Chain of Custody

## Definition

The **chain of custody** is the chronological, documented record of who handled what evidence, when, and where. Its primary focus is **control** — ensuring that evidence integrity is preserved from the moment of collection through its potential presentation in court, which may be years later.

> "Tag it, bag it, carry it." — destination-cissp §7.1.5

---

## Purpose

- Establishes and proves that evidence was not tampered with, contaminated, or altered.
- Makes evidence **admissible** in legal proceedings (though maintaining chain of custody does not *guarantee* admissibility).
- Protects the investigator from accusations of planting or altering evidence.
- Enables prosecutors, defense counsel, judges, and juries to trace the history of every piece of evidence.

---

## Locard's Exchange Principle

Formulated by 19th-century French criminologist **Dr. Edmond Locard**:

> *Every time two objects interact, some type of transfer occurs — something is taken and something is left behind.*

In digital forensics: every intrusion leaves artifacts (logs, file changes, registry entries) and the attacker carries away something (data, credentials). This principle motivates meticulous scene documentation and artifact collection.

*Source: destination-cissp §7.1.3*

---

## Establishing Chain of Custody

### Steps

1. **Identify** the scene and potential evidence locations.
2. **Secure the scene** — prevent access, contamination, or alteration.
3. **Document** via photographs, diagrams, and written notes before touching anything.
4. **Tag** each item of evidence: what it is, where found, date/time, name of the collector.
5. **Bag** the evidence in tamper-evident containers (e.g., sealed evidence bags, anti-static bags for electronics).
6. **Transport** to a secured evidence storage location (e.g., locked evidence room, evidence locker).
7. **Log every transfer** — each person who receives or handles the evidence signs off with date and time.

### For Digital Evidence

- Use a **write blocker** before connecting to any storage media.
- Create **bit-for-bit forensic copies** and verify with hash values (MD5, SHA-256).
- Seal the original in a tamper-evident evidence bag; note hash value on the bag.
- Document the chain for every copy, every analysis workstation, every analyst.

---

## Legal Admissibility Requirements

For evidence to be admitted in court, it should satisfy the **Five Rules of Evidence**:

| Rule | Meaning |
|---|---|
| **Authentic** | Evidence is genuine, not fabricated or planted. Proven through scene photos, hash values, chain-of-custody log. |
| **Accurate** | Evidence has not been changed or modified — it has integrity. |
| **Complete** | All parts of the evidence must be presented, whether or not they help the case (no cherry-picking). |
| **Convincing / Reliable** | Evidence can be understood by non-technical audiences (judges, juries). Must demonstrate a high degree of truth/veracity. |
| **Admissible** | Accepted by the court. Chain of custody helps but does not guarantee admission. |

*Source: destination-cissp §7.1.6*

---

## Types of Evidence

| Type | Description |
|---|---|
| **Real evidence** | Tangible physical objects (hard drives, USB drives). Can be held and inspected. |
| **Direct evidence** | Speaks for itself — no inference needed (eyewitness accounts, video footage of a crime). |
| **Circumstantial (indirect)** | Suggests a fact by inference (witness near the scene, suspicious log entries). |
| **Corroborative** | Supports and confirms other evidence; powerful when combined with direct evidence. |
| **Hearsay** | Testimony from witnesses not present at the incident. Usually inadmissible (exceptions exist). |
| **Best evidence** | The original document or media is preferred over copies. |
| **Secondary evidence** | A copy or substitute (e.g., a printout of log files) — admitted when originals no longer exist. |

*Source: destination-cissp §7.1.2*

---

## MOM Framework

Investigators use **MOM** to focus their inquiry:
- **M**otive — why would the suspect commit the act?
- **O**pportunity — did the suspect have access and timing to commit the act?
- **M**eans — did the suspect have the skills, tools, and knowledge required?

---

## Types of Investigations

| Type | Focus | Who Drives |
|---|---|---|
| **Criminal** | Crimes with potential jail time / criminal record | Law enforcement (local, state, federal) |
| **Civil** | Disputes between individuals/organizations; monetary penalties | Organizations and their attorneys |
| **Regulatory** | Violations of regulated activities | Associated regulatory body |
| **Administrative** | Internal policy violations, employee misconduct | The organization |

Once criminal activity is identified, the organization hands the investigation to law enforcement.

*Source: destination-cissp §7.1.7*

---

## Exam-Relevant Nuance

- Chain of custody starts at the **moment of collection**, not later.
- **Hash values** prove data integrity; chain-of-custody documentation proves handling integrity — both are needed.
- Original evidence should **never be analyzed directly** — always work from a forensic copy.
- Even if chain of custody is impeccable, a judge still has discretion over admissibility.
- Digital evidence can include **artifacts** (IP addresses, registry keys, log entries) — each artifact must also be documented and handled properly.

---

## Cross-Links

- [Forensics](forensics.md) — forensic investigation process that generates evidence
- [Incident Management](incident-management.md) — chain of custody applies during IR investigations
- [Logging & Monitoring](logging-monitoring.md) — logs are frequently secondary evidence

## Sources

- destination-cissp §7.1.1–7.1.7 (pp. 0824–0843)
- cissp-exam-outline (Domain 7: evidence collection and handling, chain of custody)
