---
title: Information System Lifecycle (IS Lifecycle)
type: concept
domain: 3
tags: [IS-lifecycle, SDLC, system-lifecycle, security-engineering]
sources: [destination-cissp]
updated: 2026-05-13
---

# Information System Lifecycle (IS Lifecycle)

## Definition

The Information System Lifecycle (IS lifecycle) is the complete lifespan of an information system, from initial conceptualization through eventual decommissioning. Security must be integrated at every phase — not added as an afterthought post-deployment.

## Relationship to SDLC

The IS lifecycle in Domain 3 (Security Architecture and Engineering) is functionally equivalent to the **System Development Lifecycle (SDLC)** in Domain 8 (Software Development Security). Minor naming differences exist:
- Integration testing falls within the Development phase in the IS lifecycle.
- Verification and Validation are distinct sub-processes within the overall testing process.

For exam purposes, both refer to the same concept: structured phases for building and retiring information systems with security embedded throughout.

## Exam-Relevant Notes

- Security controls must be considered at **every phase**, not retrofitted after design.
- The cost of fixing security issues increases dramatically as the lifecycle progresses — it is cheapest to address security during the **requirements/design phases**.
- Decommissioning must include secure data destruction (see [Key Management — crypto shredding](key-management.md)).

## Cross-References

- [Domain 8 — Software Development Security](../domains/08-software-development-security.md) — SDLC in depth
- [Key Management](key-management.md) — crypto shredding for secure decommissioning

## Sources

- destination-cissp §3.10 (page 0485)
