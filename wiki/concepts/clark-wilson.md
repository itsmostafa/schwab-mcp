---
title: "Clark–Wilson Model"
type: concept
domain: 3
tags: [security-models, integrity, rule-based, clark-wilson, well-formed-transactions, separation-of-duties]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Clark–Wilson Model

Clark–Wilson is a **rule-based security model** focused exclusively on **integrity**, designed with commercial (non-military) environments in mind. Unlike Biba (which only prevents unauthorized subjects from making any changes), Clark–Wilson addresses all three goals of integrity and is considered superior to Biba for real-world integrity protection.

## Definition

Clark–Wilson formalizes integrity through a triple of subject–program–object and a set of rules that govern how access and transactions must occur. Access is never direct; it always flows through a constrained program.

## Three Goals of Integrity

| Goal | Addresses |
|---|---|
| 1. Prevent unauthorized subjects from making any changes | Biba also addresses this |
| 2. Prevent authorized subjects from making **bad** changes | Clark–Wilson only |
| 3. Maintain **consistency** of the system | Clark–Wilson only |

Biba only addresses goal #1. Clark–Wilson addresses all three.

## Three Rules (Mechanisms)

| Rule | Description |
|---|---|
| **Well-Formed Transactions** | Operations must be performed in a consistent, validated manner that does not compromise object integrity. Data must remain good and consistent. |
| **Separation of Duties** | No single person can perform all steps of a critical function. Multiple parties are required for high-stakes operations. |
| **Access Triple (Subject–Program–Object)** | A subject cannot directly access an object. Access must go through an authorized program (transformation procedure) that enforces access rules. |

## Key Terms

- **CDI** — Constrained Data Item: high-integrity data that Clark–Wilson protects
- **UDI** — Unconstrained Data Item: external/untrusted data entering the system
- **TP** — Transformation Procedure: authorized program through which access to CDIs must occur
- **IVP** — Integrity Verification Procedure: verifies that CDIs conform to integrity constraints

## The Access Triple

```
Subject → (via TP) → CDI
```

Subjects cannot touch CDIs directly. All access is mediated by a TP (Transformation Procedure). This is the Clark–Wilson equivalent of the Reference Monitor Concept.

## Comparison: Clark–Wilson vs. Biba

| | Biba | Clark–Wilson |
|---|---|---|
| **Type** | Lattice-based | Rule-based |
| **Focus** | Integrity (prevent unauthorized changes) | Integrity (all three goals) |
| **Commercial fit** | Limited | Designed for commercial use |
| **Mechanism** | Security level labels | Access triples + TPs + IVPs |

## Exam-Relevant Nuance

- Clark–Wilson is the **commercial** integrity model. The CISSP exam often tests its commercial context vs. Biba.
- The three integrity goals and three rules are both frequently tested — know them separately.
- Separation of duties in Clark–Wilson is a mechanism, not just a principle.
- UDIs entering the system must be transformed into CDIs via a TP — this is how untrusted data is sanitized.

## Cross-links

- [Security Models](security-models.md) — model classification hub
- [Biba](biba.md) — simpler integrity model that Clark–Wilson improves upon
- [Bell–LaPadula](bell-lapadula.md) — confidentiality model

## Sources

- destination-cissp §3.2.4 (pp. 253–254)
- cissp-exam-outline §3.2
