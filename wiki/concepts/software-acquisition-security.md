---
title: "Software Acquisition Security"
type: concept
domain: 8
tags: [cots, open-source, third-party, managed-services, cloud, software-escrow, acquisition, vendor-assessment, sdlc]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Software Acquisition Security

## Definition

**Software acquisition security** is the application of security rigor to the process of obtaining software from sources other than in-house development. Regardless of how software is acquired — purchased, licensed, open source, contracted, or cloud-hosted — the security organization must be involved throughout the acquisition process.

> **Core principle (destination-cissp §8.4.1):** Acquiring software must be treated with the same seriousness as developing it. Security should be part of every step.

---

## Software Assurance Phases for Acquisition

*(source: destination-cissp §8.4.1)*

| Phase | Key Activities |
|---|---|
| **Planning / Requirements** | Define security requirements for the software. What does it need to do securely? What regulatory/compliance constraints apply? |
| **Contracting** | Include security requirements in contracts; define SLAs, right-to-audit, breach notification obligations, liability, and source code escrow terms. |
| **Acceptance** | Test and evaluate the software against security requirements before accepting delivery. Apply SDLC testing to the extent possible. |
| **Monitoring and Follow-On** | Ongoing security assessment of the vendor relationship; track vulnerabilities and patches; periodic re-evaluation. |

> **Exam note:** When acquiring ANY software, the first step should always be planning/requirements — define what is needed security-wise before engaging vendors.

---

## Acquisition Methods

### Commercial-Off-the-Shelf (COTS)

Pre-built, commercially available software sold to the general public. Examples: Microsoft 365, antivirus software, ERP systems, CRM tools.

**Pros:**
- Functionality can be evaluated and compared
- Third-party evaluations available (e.g., Common Criteria)
- Existing customer base to consult
- Vendor-provided patches and updates

**Cons:**
- Source code typically not available → **no white-box testing possible**
- Vendor may go out of business (mitigated by software escrow)
- Missing features/functionality may require workarounds
- Widely-used COTS is a larger target — more attackers probe it for vulnerabilities

**Security approach:** Apply as much of the SDLC process as possible (black-box testing, acceptance testing, vendor assessment). If source code is unavailable, request it be placed in **software escrow**.

---

### Open Source Software

Software with publicly available source code that anyone can inspect, modify, and distribute.

**Advantages over COTS:**
- **Control**: Can inspect the code to verify security; can modify or remove risky components
- **Training**: Can use as a learning/testing tool
- **Security**: Community-driven vulnerability discovery often identifies and fixes issues quickly
- **Longevity**: Not dependent on a single vendor; community provides ongoing care

**Risks:**
- **Supply chain risk**: Community contributions can introduce vulnerabilities (accidental or malicious). Classic example: **Heartbleed** — a simple programming error in OpenSSL affected millions of devices globally because OpenSSL was so widely used.
- Code quality and security practices vary widely across projects.

**Security approach:** Before using in production, treat as if internally developed. Apply the SDLC (or equivalent SSDLC approach). Inspect the code. Confirm safety for use.

---

### Third-Party Development

Code developed by external programmers (contractors or software development companies) not employed by the organization.

**Security approach:** Same scrutiny as vendor software or internal development. Apply SDLC to the extent possible. Contractual requirements for secure development, testing, and code review. Access to source code enables white-box testing.

---

### Managed Services (e.g., Enterprise Applications via MSP)

Managed Service Providers (MSPs) offer IT infrastructure and support, including enterprise applications. Common for small-to-medium businesses where running IT in-house is cost-prohibitive.

**Assessment approach (destination-cissp §8.4.1):**
- Evaluate **SOC reports** (SOC 1, SOC 2) prepared by independent auditors
- **Site visits** by senior management and key stakeholders
- **Reference checks** with existing MSP customers
- Industry and regulatory assessments specific to organizational needs (e.g., HIPAA for healthcare)

---

### Cloud Services (SaaS, IaaS, PaaS)

Cloud service acquisition requires security involvement in vendor selection, data migration, and ongoing oversight.

**Key principle:** The **shared responsibility model** — the customer is always *accountable* but shares *control* and *responsibility* with the cloud provider. Both parties must understand their specific responsibilities.

- **SaaS**: Provider manages everything; customer responsible for data and access management.
- **PaaS**: Provider manages infrastructure/runtime; customer responsible for applications and data.
- **IaaS**: Provider manages physical infrastructure; customer responsible for OS, applications, and data.

Due diligence is required whenever data or applications are moved to the cloud. Responsibilities and risks must be documented in the service agreement.

---

## Software Escrow

When source code is not provided to the purchaser (common with COTS), the purchasing organization may negotiate a **software escrow** arrangement:
- Source code is held by a neutral third-party escrow agent.
- The purchasing organization gains access to the source code if predefined trigger conditions occur (vendor bankruptcy, vendor ceases support, failure to maintain the product).
- Ensures the organization can maintain, modify, or migrate the software even if the vendor disappears.

---

## SBOM (Software Bill of Materials)

> **Coverage gap:** destination-cissp does not cover SBOM. The following is from cissp-exam-outline scope.

A **Software Bill of Materials (SBOM)** is a formal, machine-readable inventory of all components, libraries, and dependencies included in a software product — analogous to an ingredient list on food packaging.

**Security value:**
- Enables rapid identification of which products are affected when a vulnerability is disclosed in a specific component (e.g., Log4Shell in Log4j)
- Supports supply chain risk management
- Increasingly required by government contracts (Executive Order 14028, 2021)
- Applicable to COTS, open source, and custom software

---

## Exam-Relevant Nuance

- **COTS**: No white-box testing possible. Black-box and acceptance testing required. Software escrow mitigates vendor-gone-out-of-business risk.
- **Open source**: Heartbleed is the canonical example of open source supply chain risk.
- **The SDLC must be applied to all acquisition types** — even when purchasing from a vendor, the organization must still conduct its own security assessment.
- **Security should be part of contracting** — SLA security requirements, right-to-audit, breach notification.
- **Shared responsibility model**: Customer is always ultimately accountable in cloud arrangements.

---

## Cross-Links

- [SDLC](sdlc.md) — SDLC applied to acquisition phases
- [CI/CD Security](ci-cd-security.md) — SBOM in pipeline context; supply chain attack mitigations
- [OWASP Top 10](owasp-top-10.md) — A06 Vulnerable & Outdated Components — directly tied to software acquisition risk

## Sources

- destination-cissp §8.4–8.4.1 (pp. 0961–0964)
- cissp-exam-outline (Domain 8.4: Assessing security impact of acquired software)
- SBOM coverage gap noted — cissp-exam-outline scope
