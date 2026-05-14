---
title: "Information Lifecycle Management"
type: concept
domain: 2
tags: [asset-inventory, asset-management, ilm, asset-classification]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Information Lifecycle Management

Information Lifecycle Management (ILM) encompasses policies, procedures, and technologies for managing information assets from creation through final disposition. The foundation of ILM is knowing what assets you have, classifying them, and protecting them based on classification.

## Asset Inventory

An organization cannot protect what it does not know it has. An accurate and continuously updated **asset inventory** is the starting point for all asset classification and management activities.

Key challenges with asset inventory:
- Assets can be **created, purchased, rented, or acquired** — all methods must be tracked
- **Shadow IT** (e.g., a department manager signing up for a cloud service and forgetting about it) leaves valuable assets unprotected and unknown to security
- Large multinational organizations face enormous scale challenges

Asset inventory must capture:
- All tangible assets (hardware, physical media, servers, endpoints)
- All intangible assets (data, software, intellectual property, licenses, cloud services)
- The **owner** of each asset
- The **classification level** assigned to each asset

## Asset Management and Classification as Ongoing Processes

Asset management is not a one-time exercise:
- Assets are constantly added and removed
- Owners change
- Laws, regulations, and business priorities evolve — all of which can change the value and therefore the classification of an asset
- **Periodic review** of all assets is required

The classification process (from inventory → owner identification → classification → handling → review) is a continuous cycle, not a project with an end date.

## Protection baselines

Once assets are classified, a **baseline of controls** is defined for each classification level. These baselines:
- Apply across all assets at a given classification level
- Are cost-effective: high-value assets warrant more investment; low-value assets may not warrant costly controls
- Must address all three data states (at rest, in transit, in use) — protections may differ per state
- Drive scoping and tailoring decisions (which controls from a standard framework are applicable)

## Scoping and Tailoring

- **Scoping**: Selecting which controls from a larger control framework apply to the organization's specific environment
- **Tailoring**: Adjusting selected controls to fit the specific context (e.g., compensating controls when a standard control is infeasible)

Both are part of how organizations apply classification-driven control baselines from frameworks like NIST SP 800-53.

## End of Life (EOL) and End of Support (EOS)

Exam outline §2.5 specifically addresses asset retention at end of life. Key considerations:
- **EOL** (End of Life): manufacturer no longer sells or develops the product
- **EOS** (End of Support): manufacturer no longer provides security patches or technical support — this is the higher security risk because vulnerabilities will no longer be remediated
- Assets at EOS pose risk and should be inventoried, risk-assessed, and either upgraded/replaced or compensating controls applied

## Cross-links

- [Data Classification](data-classification.md) — classification process that ILM supports
- [Data Lifecycle](data-lifecycle.md) — the lifecycle that ILM manages
- [Data Retention](data-retention.md) — retention requirements as part of ILM
- [Data Roles](data-roles.md) — accountability roles within ILM
- [Domain 2 — Asset Security](../domains/02-asset-security.md)

## Sources

- destination-cissp §2.1.1, §2.1.2, §2.3.1, §2.6 (pages 0185–0188, 0200–0201, 0212)
- cissp-exam-outline §2.3, §2.5 (asset inventory, EOL/EOS)
