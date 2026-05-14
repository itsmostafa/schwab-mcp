---
title: "Bell–LaPadula Model"
type: concept
domain: 3
tags: [security-models, confidentiality, lattice-based, bell-lapadula, no-read-up, no-write-down]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Bell–LaPadula Model

Bell–LaPadula (BLP) is a **lattice-based security model** developed in the early 1970s for the US Department of Defense. It addresses one property of the CIA triad exclusively: **confidentiality**. It is still conceptually valid today because the fundamental rules of confidentiality have not changed.

## Definition

BLP defines rules that govern which subjects can read from and write to objects at various security levels in a layered hierarchy. The goal is to prevent unauthorized disclosure of classified information.

## The Three Properties

| Property | Nickname | Rule | Governs |
|---|---|---|---|
| **Simple Security Property** | "No Read Up" | A subject at level L may NOT read an object at a higher level | Reading |
| **Star (*) Property** | "No Write Down" | A subject at level L may NOT write to an object at a lower level | Writing |
| **Strong Star (*) Property** | Read and write at own level only | A subject may ONLY read and write at their own level | Both |

**Memory aid:** BLP = "No Read Up, No Write Down" — preventing sensitive data from flowing to lower-classified subjects by either reading up or writing it down.

## Why These Rules Protect Confidentiality

- **No Read Up** prevents a lower-clearance subject from accessing higher-classified material.
- **No Write Down** prevents a higher-clearance subject from writing classified information into a lower-classified object (which lower-clearance subjects could then read).
- Together, these rules ensure information only flows upward or stays at the same level — it cannot leak downward.

## Brewer–Nash Comparison

Both BLP and Brewer–Nash (the Chinese Wall model) address **confidentiality**. The difference:
- BLP uses fixed hierarchical security labels.
- Brewer–Nash prevents conflicts of interest: access is denied when it would create a conflict between two data sets a subject has previously accessed (e.g., competitor company data).

## Limitations

- BLP addresses **only** confidentiality; it does not address integrity.
- If both confidentiality and integrity are needed, use the **Lipner Implementation**, which combines BLP with Biba.
- BLP is designed for lattice-based (hierarchical) environments; it is not a good fit for all architectures.

## Exam-Relevant Nuance

- The Star (*) property is also called the "Confinement Property."
- The Strong Star property is sometimes tested separately — remember it restricts both reads and writes to the subject's own level.
- Covert channels (timing and storage) can bypass BLP rules — see [Security Models](security-models.md) for covert channel discussion.
- BLP is lattice-based, not rule-based.

## Cross-links

- [Security Models](security-models.md) — model classification hub
- [Biba](biba.md) — the integrity counterpart; inverse rules
- [Clark–Wilson](clark-wilson.md) — commercial integrity model
- [Trusted Computing Base](trusted-computing-base.md) — implementation context

## Sources

- destination-cissp §3.2.3 (pp. 245–249)
- cissp-exam-outline §3.2
