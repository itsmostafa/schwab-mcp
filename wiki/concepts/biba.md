---
title: "Biba Model"
type: concept
domain: 3
tags: [security-models, integrity, lattice-based, biba, no-read-down, no-write-up]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Biba Model

The Biba model is a **lattice-based security model** that addresses one property of the CIA triad exclusively: **integrity**. It is the direct counterpart to Bell–LaPadula: where BLP protects confidentiality (preventing data from leaking down), Biba protects integrity (preventing untrusted data from flowing up).

## Definition

Biba defines rules that govern which subjects can read from and write to objects at various integrity levels. The goal is to prevent corruption of high-integrity data by lower-integrity subjects or inputs.

## The Three Properties

| Property | Nickname | Rule | Governs |
|---|---|---|---|
| **Simple Integrity Property** | "No Read Down" | A subject at integrity level L may NOT read an object at a lower integrity level | Reading |
| **Star (*) Integrity Property** | "No Write Up" | A subject at integrity level L may NOT write to an object at a higher integrity level | Writing |
| **Invocation Property** | — | A subject cannot send information to a subject at a higher integrity level | Messaging/invocation |

**Memory aid:** Biba = "No Read Down, No Write Up" — preventing low-integrity data from contaminating high-integrity objects.

## Comparison: Biba vs. Bell–LaPadula

| | Bell–LaPadula | Biba |
|---|---|---|
| **Protects** | Confidentiality | Integrity |
| **No Read** | Up (can't read higher) | Down (can't read lower) |
| **No Write** | Down (can't write lower) | Up (can't write higher) |
| **Direction of concern** | Data leaking down | Corruption flowing up |

The rules are mirror images of each other.

## Biba vs. Clark–Wilson on Integrity

Biba only prevents unauthorized subjects from making any changes (one of the three goals of integrity). Clark–Wilson goes further:

1. Prevent unauthorized subjects from making changes — Biba addresses this; Clark–Wilson also addresses it.
2. Prevent **authorized** subjects from making **bad** changes — Clark–Wilson only.
3. Maintain consistency of the system — Clark–Wilson only.

Therefore, Clark–Wilson provides more complete integrity protection than Biba.

## Lipner Implementation

When both confidentiality (BLP) and integrity (Biba) are required, the **Lipner Implementation** combines the best features of both. Lipner separates objects into data and programs and applies sensitivity levels and job categories to subjects. Lipner is not a model — it is an implementation.

## Exam-Relevant Nuance

- Biba is lattice-based, not rule-based.
- The Invocation Property is a third property often forgotten on exams — know all three.
- Integrity in Biba means the data is "accurate, relevant, or meaningful."
- Biba does not address confidentiality at all.

## Cross-links

- [Security Models](security-models.md) — model classification hub
- [Bell–LaPadula](bell-lapadula.md) — confidentiality counterpart; inverse direction rules
- [Clark–Wilson](clark-wilson.md) — more comprehensive integrity model

## Sources

- destination-cissp §3.2.3 (pp. 247–249)
- cissp-exam-outline §3.2
