---
title: "Security Policies, Standards, Procedures, Baselines, and Guidelines"
type: concept
domain: 1
tags: [policy, standard, procedure, baseline, guideline, governance, compliance]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Security Policies, Standards, Procedures, Baselines, and Guidelines

The security document hierarchy is how management's intent flows into operational reality. Each
level has a distinct definition and purpose. The exam frequently asks you to classify a described
document into the correct category.

## The Hierarchy (Table 1-14)

### Policies
- **Corporate laws** — communicate management's goals and objectives.
- Provide authority to security activity; define scope, elements, and function of security.
- Must be approved and communicated by **management** (top-down).
- The **overarching security policy** should come from the CEO or Board of Directors.
- Policies do not need to be reviewed as frequently as standards/procedures.
- Example: "All systems must use anti-malware software."

### Standards
- Specific hardware/software solutions, mechanisms, and products.
- Define *what* must be used (vendor, version, technology).
- Example: "McAfee Endpoint Security v10 is the required anti-malware product."
- Note: Published guidelines adopted by an organization (e.g., ISO 27001) can become an organizational standard.

### Procedures
- **Step-by-step descriptions** of how to perform a task; **mandatory** actions.
- Example: "To install anti-malware: Step 1 — download installer from internal server; Step 2 — run as admin..."
- Other examples: user registration, incident response procedures, secure contracting process.

### Baselines
- Defined **minimal implementation levels** for security mechanisms.
- Configuration-based (e.g., minimum system hardening settings).
- Example: "All IDS systems must be configured with baseline ruleset XYZ."

### Guidelines
- **Recommended or suggested** actions — **not mandatory**.
- Allows an organization to suggest best practices without creating a hard compliance requirement.
- Exam trap: guidelines do **not** cause audit findings if not followed (because they're optional).
- Example: "Consider enabling heuristic scanning in anti-malware software where possible."

## Document Hierarchy Diagram

```
Overarching Security Policy (Board/CEO)
  └── Functional Security Policies
        ├── Standards (what products/tools)
        ├── Procedures (how to do it, step-by-step)
        ├── Baselines (minimum config levels)
        └── Guidelines (suggestions/recommendations)
```

## Exam-Relevant Nuance

- On the exam, the key differentiator between **guidelines** and **procedures** is mandatory vs. optional. Procedures must be followed; guidelines are recommended.
- **Baselines** vs. **standards**: standards specify *what* to use; baselines specify *minimum configuration levels* for what's in place.
- The overarching policy sets tone from the top — if CEO/Board won't lead, security culture fails.
- Administrative controls = policies, procedures, standards, guidelines, and baselines collectively.

## Cross-links

- [Governance](governance.md)
- [Due Care and Due Diligence](due-care-due-diligence.md)
- [Security Controls Types](security-controls-types.md)
- [Compliance Requirements](compliance-requirements.md)

## Sources

- destination-cissp §1.6 (Table 1-14, Figure 1-3)
