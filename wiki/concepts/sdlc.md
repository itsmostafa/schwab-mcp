---
title: "SDLC — Software Development Life Cycle"
type: concept
domain: 8
tags: [sdlc, slc, waterfall, agile, spiral, cmmi, ssdlc, threat-modeling, certification, accreditation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# SDLC — Software Development Life Cycle

## Definition

The **Software Development Life Cycle (SDLC)** is a structured process for planning, creating, testing, and deploying software. Security must be embedded at every phase — from inception through disposal — because retrofitting security after development is both more costly and less effective than designing it in from the start.

The **System Life Cycle (SLC)** picks up where the SDLC leaves off: beginning when software enters production and continuing through operations, maintenance, decommissioning, and disposal. Together, SDLC + SLC cover the entire application life.

---

## SDLC Phases and Security Activities

| Phase | What Happens | Security Activities |
|---|---|---|
| **Initiation / Planning** | Define project scope, feasibility, resource allocation. | Identify security requirements, initial asset classification, assign data owners. |
| **Requirements** | Elicit and document functional and non-functional requirements from stakeholders. | Conduct **risk analysis** specific to the application; identify regulatory/compliance constraints. |
| **Architecture & Design** | Translate requirements into technical specifications; design data flows and system components. | Perform **threat modeling** (STRIDE) to drive security architecture decisions; select controls. |
| **Development / Coding** | Programmers implement the design. | Apply secure coding guidelines; conduct **code review** and unit-level **SAST** as units are completed. |
| **Testing** | Validate that the system meets requirements and is free from defects. | Comprehensive SAST, DAST, and fuzz (dynamic) testing; penetration testing; certification. |
| **Release / Deployment** | Move software to production. | **Accreditation** (management sign-off); change-controlled deployment; baseline configuration recorded. |
| **Operations & Maintenance** | Ongoing production use; monitoring, patching, change requests. | Continuous monitoring; patch management; change management with security review; periodic re-assessment. |
| **Disposal** | Decommission the application and retire associated data. | Secure media sanitization (see Domain 2); archival; ensure no residual sensitive data. |

> **Exam note:** Testing and development phases can overlap — specifically, static code review (SAST) can begin while coding is still underway. The bulk of dynamic testing (DAST, fuzz) occurs in the formal testing phase.

---

## Key Security Focus Points by Phase

- **Requirements phase**: This is the right time for risk analysis. If risks are not defined here, controls cannot be designed.
- **Design phase**: Threat modeling is performed *here*, not later. The design determines which security elements are built in.
- **Testing phase**: Must include normal, error-prone, *and malicious* usage scenarios. Static (SAST) + Dynamic (DAST) + Fuzz is the comprehensive combination.
- **Deployment transition**: Certification (technical evaluation) and Accreditation (management approval) are performed here.
- **Operations phase**: Change management must include security review/approval. Many organizations fail to include security on change control committees.
- **Disposal**: Often neglected. Security must oversee data destruction and archival just as it does at inception.

---

## Development Methodologies

All methodologies are reflections or adaptations of the waterfall model. Security must be embedded regardless of which is used.

| Methodology | Key Characteristics |
|---|---|
| **Waterfall** | Strictly phased; must complete each phase before moving to the next. No going back. Sign-offs mark each transition. Simple but inflexible. |
| **Structured Programming** | Logical approach emphasizing structured control flow; foundational to OOP. |
| **Agile** | Iterative, rapid sprints; heavy customer involvement throughout; led by Scrum Master. More responsive to change than waterfall. |
| **Scaled Agile Framework (SAFe)** | Agile adapted for large organizations with many teams collaborating. |
| **Spiral** | Risk-driven; iterative like agile but includes defined phases similar to waterfall. The process "spirals" back through phases to address issues from prior iterations. |
| **Cleanroom** | Focus on defect *prevention* (not detection); aims for certifiable reliability through statistical quality control. |

**Waterfall vs. Agile:**  
Waterfall is phased and sequential (best when requirements are stable). Agile uses small skilled teams and produces working software in short iterations (best when requirements may evolve). Both require security at every step.

**Agile Scrum Master** role: understands how all team efforts fit together; shields developers from external interruptions; enforces scrum principles; removes barriers; facilitates close cooperation.

**Combining methodologies**: Organizations often mix approaches (e.g., agile + structured), and DevSecOps can be incorporated with any combination.

---

## Certification and Accreditation

- **Certification**: Technical and comprehensive evaluation of a system against security requirements. Produces a formal finding.
- **Accreditation**: Management's official decision and sign-off to authorize operation of the system. Accountability rests with management.

Both occur at the transition between testing and deployment (SDLC) or during the Implementation phase (SLC).

---

## Secure SDLC (SSDLC)

The SSDLC is simply the SDLC with security embedded at every phase from the outset — not added later. Key practices include:

- Security requirements captured during requirements phase.
- Threat modeling during design.
- Secure coding standards enforced during development.
- SAST/DAST/fuzz testing in testing phase.
- Security sign-off as part of accreditation/deployment.
- Ongoing monitoring and change management during operations.
- Secure disposal.

---

## Operation and Maintenance

Successful and secure software requires ongoing care:

- **Monitoring**: Proactive detection of security problems before they become widespread.
- **Periodic evaluation**: Deeper dives confirming that coding techniques and components remain secure.
- **Patching**: Most commonly addresses security vulnerabilities, but also adds functionality. Change management governs all patches.

---

## Cross-Links

- [DevSecOps](./devsecops.md) — DevOps with security integrated; relationship to SDLC
- [Secure Coding Practices](secure-coding-practices.md) — what to apply during the development phase
- [Software Testing Types](./security-testing-types.md) — SAST/DAST/fuzz; unit, integration, system testing
- [CMMI/SAMM](../standards/cmm-samm.md) — maturity models applied to the SDLC
- [Change Management (Software)](change-management-software.md) — governs changes during operations phase

## Sources

- destination-cissp §8.1–8.1.3 (pp. 0925–0934)
- cissp-exam-outline (Domain 8 subtopic list)
