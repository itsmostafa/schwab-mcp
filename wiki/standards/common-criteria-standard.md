---
title: "Common Criteria (ISO/IEC 15408)"
type: standard
domain: 3
tags: [common-criteria, iso-15408, eal, evaluation-criteria, certification]
sources: [destination-cissp]
updated: 2026-05-13
---

# Common Criteria — ISO/IEC 15408

## Overview

**Common Criteria** is an international standard (ISO/IEC 15408) for evaluating the security properties of IT products and systems. It emerged from the convergence of TCSEC (US) and ITSEC (Europe) and is now the globally dominant evaluation framework. Products evaluated under CC receive an **Evaluation Assurance Level (EAL)** rating from 1 to 7.

## Standard Details

| Attribute | Value |
|---|---|
| **Formal name** | ISO/IEC 15408 |
| **Common name** | Common Criteria (CC) |
| **Status** | Current (version 3.1 Rev 5 as of publication) |
| **Replaced** | ITSEC (2005); TCSEC (earlier) |
| **Nature** | Voluntary; evaluation performed by CC-licensed independent organizations |

## EAL Ratings

| Level | Description |
|---|---|
| EAL 7 | Formally verified, designed, and tested |
| EAL 6 | Semi-formally verified, designed, and tested |
| EAL 5 | Semi-formally designed and tested |
| EAL 4 | Methodically designed, tested, and reviewed |
| EAL 3 | Methodically tested and checked |
| EAL 2 | Structurally tested |
| EAL 1 | Functionally tested |

- **Commercial sweet spot**: EAL 4 (firewalls); EAL 3 (operating systems)
- **Above EAL 4**: Rarely purchased commercially due to complexity and cost
- **EAL persistence**: Patches do not change the EAL; only major functionality changes trigger re-evaluation

## Key Components

| Component | Purpose |
|---|---|
| **Protection Profile (PP)** | Defines requirements for a category of products (e.g., all firewalls) |
| **Target of Evaluation (TOE)** | The specific product being evaluated |
| **Security Targets (ST)** | Vendor's statement of how their product meets the PP |
| **Evaluation** | Independent testing of functional + assurance requirements |

## Exam Notes

- Know both names: Common Criteria = ISO/IEC 15408
- CC evaluation is **voluntary** — vendors choose to pursue it
- Higher EAL does not always mean more secure in practice — usability and maintainability matter
- Accreditation (management sign-off) is separate from and follows certification (technical evaluation)

## Related Concept Page

For full detail on CC mechanics and comparison to TCSEC/ITSEC, see [Common Criteria concept page](../concepts/common-criteria.md).

## Sources

- destination-cissp §3.2.7 (pp. 267–273)
