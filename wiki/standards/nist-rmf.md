---
title: "NIST Risk Management Framework (RMF) — SP 800-37"
type: standard
domain: 1
tags: [nist, rmf, sp800-37, risk-management-framework, authorization, categorize, monitor]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# NIST Risk Management Framework (RMF) — SP 800-37

The NIST RMF (NIST SP 800-37 Rev. 2) provides comprehensive guidance for applying risk management
to federal information systems and organizations. It is widely used in both government and private
sector contexts, and is a core framework for CISSP candidates.

## The 7 Steps of NIST SP 800-37 Rev. 2 (Table 1-26)

| Step | Name | Key Activities |
|---|---|---|
| **1** | **Prepare** | Establish context and organizational risk management strategy; roles, responsibilities, risk tolerance |
| **2** | **Categorize** | Classify information systems by sensitivity and potential adverse impact to CIA; answer "What do we have?" and "How sensitive is it?" |
| **3** | **Select** | Choose, tailor, and document security controls based on risk assessment; controls protect CIA of systems and information |
| **4** | **Implement** | Deploy selected controls; document specific, baseline implementation details so controls are understood in organizational context |
| **5** | **Assess** | Determine if controls are implemented correctly, operating as intended, and meeting security/privacy requirements; formulate and approve a comprehensive assessment plan |
| **6** | **Authorize** | Senior management decides whether to authorize system operation given potential risk, controls, and residual risk; approval typically tied to Plan of Actions & Milestones (POA&M) |
| **7** | **Monitor** | Continuous monitoring of control effectiveness over time; adapts to changing threats, vulnerabilities, technologies; near-real-time capability with automated tools |

Mnemonic: **P-C-S-I-A-A-M** → "**P**lease **C**all **S**teve **I**f **A**nybody **A**sks **M**e"

## Key Facts

- **NIST SP 800-37 Rev. 2** is the focus for the CISSP exam (not earlier revisions).
- The RMF is **cyclical** — the Monitor step feeds back continuously.
- **POA&M** (Plan of Actions & Milestones): documents remaining weaknesses and deficiencies after authorization; tracks and monitors failed controls.
- The **Authorize** step requires explicit senior management sign-off; this is the risk acceptance decision point.
- Automated monitoring tools are supported but **not required**.
- Configuration drift is a key concern addressed by the Monitor step.

## Related Risk Frameworks (Table 1-25)

| Framework | Description |
|---|---|
| **NIST SP 800-37 (RMF)** | Seven-step risk management framework for information systems |
| **ISO 31000** | International family of standards for risk management; provides structure and guidance for all organizations |
| **COSO** | Enterprise risk management framework; defines ERM components and provides direction for enterprise risk management |
| **ISACA Risk IT** | Guidelines for risk optimization, security, and business value; aligns with COBIT |

## Exam-Relevant Nuance

- The Authorize step (Step 6) = **senior management makes the risk acceptance decision** — not the security team.
- The Assess step (Step 5) determines *if* controls are working; Monitor (Step 7) ensures they *keep* working.
- "Categorize" (Step 2) = understand what you have and how sensitive it is — feeds everything downstream.
- NIST SP 800-37 Rev. 2 added the "Prepare" step (Step 1) — distinguishing Rev. 2 from earlier versions.

## Cross-links

- [Risk Management](../concepts/risk-management.md)
- [Governance](../concepts/governance.md)

## Sources

- destination-cissp §1.9.10 (Tables 1-25, 1-26)
