---
title: "Security Models"
type: concept
domain: 3
tags: [security-models, bell-lapadula, biba, clark-wilson, brewer-nash, lattice-based, rule-based]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Security Models

A security model is a formal representation of what security should look like in an architecture — essentially, the rules that must be implemented to achieve a desired security property. Security models provide the conceptual foundation; actual system implementations then realize those rules.

## Model Types

All security models are classified as either:

| Category | Description | Examples |
|---|---|---|
| **Lattice-based (layer-based)** | Rules are confined to hierarchical security layers | Bell–LaPadula, Biba |
| **Rule-based** | Access is mediated by a general set of rules | Clark–Wilson, Brewer–Nash, Information Flow, Graham–Denning, Harrison–Ruzzo–Ullman |

### Why Not Just "Rule-Based" for Lattice Models?
Lattice-based models do include rules, but those rules operate within defined layers (the lattice). The term "lattice-based" is more specific and is the correct exam classification.

## Individual Model Summary

| Model | Property | Direction Rules | Notes |
|---|---|---|---|
| Bell–LaPadula | Confidentiality | No read up; no write down | Military origin; lattice-based |
| Biba | Integrity | No read down; no write up | Inverse of BLP; lattice-based |
| Lipner | Confidentiality + Integrity | Combines BLP + Biba | Not a model — an implementation |
| Clark–Wilson | Integrity | Well-formed transactions, SoD, Access Triple | Commercial focus; rule-based |
| Brewer–Nash (Chinese Wall) | Confidentiality | No cross-conflict access | Prevents conflicts of interest; rule-based |
| Graham–Denning | Access rights | 8 rules for subject/object access | Lesser known; rule-based |
| Harrison–Ruzzo–Ullman | Access rights integrity | Finite rule set for editing access rights | Adds generic rights to groups |
| Information Flow | All | Tracks info through life cycle | Basis for BLP and Biba |

## Enterprise Security Architecture Frameworks

These are not security models but frameworks for implementing enterprise security architecture:

- **Zachman** (1970s) — classification framework; answers who/what/when/where/why/how for each organizational stakeholder group. Older model focused on classification.
- **SABSA** (Sherwood Applied Business Security Architecture, 1995) — risk-driven, open source, scalable. Embeds security in IT functions and facilitates compliance.
- **TOGAF** (The Open Group Architecture Framework) — focuses on resource efficiency and cost minimization; modular structure.

## Detailed Model Pages

- [Bell–LaPadula](bell-lapadula.md) — confidentiality model
- [Biba](biba.md) — integrity model
- [Clark–Wilson](clark-wilson.md) — commercial integrity model
- Brewer–Nash is covered in this hub page and [bell-lapadula.md](bell-lapadula.md) comparison

## Exam-Relevant Nuance

- Bell–LaPadula dates to the early 1970s but is still conceptually valid. Technology changes; fundamental confidentiality rules do not.
- Most exam questions test direction of access (up/down) and which property (confidentiality vs. integrity) each model addresses.
- Brewer–Nash primarily addresses **confidentiality**, not just conflict of interest — it is listed under confidentiality models.
- Lipner is not a model; calling it one is a common mistake.

## Cross-links

- [Bell–LaPadula](bell-lapadula.md)
- [Biba](biba.md)
- [Clark–Wilson](clark-wilson.md)
- [Trusted Computing Base](trusted-computing-base.md) — how models are implemented
- [Common Criteria](common-criteria.md) — evaluation of security implementations

## Sources

- destination-cissp §3.2.1–3.2.4 (pp. 241–256)
- cissp-exam-outline §3.2
