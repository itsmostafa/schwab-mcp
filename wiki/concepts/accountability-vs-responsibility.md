---
title: "Accountability vs. Responsibility"
type: concept
domain: 1
tags: [accountability, responsibility, governance, ownership, delegation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Accountability vs. Responsibility

A foundational governance concept and classic exam trap. These two words are **not** synonymous.

## Definitions (Table 1-4)

| Accountability | Responsibility |
|---|---|
| "Where the buck stops" | "The doer" |
| Ultimate ownership and liability | In charge of a task or process |
| Only **ONE** person or group can be accountable | **Multiple** people can be responsible |
| Sets rules and policies | Develops plans and implements controls |
| **Cannot be delegated** | **Can be delegated** |

## Key Rules

1. **Accountability can never be delegated.** The CEO/Board can delegate tasks (responsibility) but they remain accountable for the outcome.
2. **Responsibility can be delegated** — but the person who delegated it remains accountable.
3. **Security is everyone's responsibility** — but accountability for security remains with asset owners and senior management.

## Application Examples

- A CFO is accountable for the accuracy of financial reports. They can delegate the responsibility to an accounting team, but if reports are fraudulent, the CFO remains accountable.
- A data owner is accountable for their data even when a Cloud Service Provider (CSP) is responsible for processing it. If there is a breach of that data, the data owner is accountable — not just the CSP.
- If the question asks "who is ultimately accountable for the finance system?" the answer is the **VP of Finance** (or the next most senior person if the VP isn't listed).
- On the CISSP exam: if the question includes the CEO or Board, they are the last line of accountability when no other owner is identified.

## Exam-Relevant Nuance

- Exam trap: "The cloud provider is responsible for security, so the customer isn't accountable." **FALSE.** Accountability always remains with the data owner.
- Data Custodians are **responsible** for protecting assets in their custody but not **accountable** — the asset owner remains accountable.
- If a custodian fails to protect an asset: the **custodian** is responsible for the failure; the **owner** is accountable.

## Cross-links

- [Governance](governance.md)
- [SCRM](scrm.md)
- [Personnel Security](personnel-security.md)
- [Privacy](privacy.md)

## Sources

- destination-cissp §1.3.2, §1.3.3 (Tables 1-4, 1-5)
