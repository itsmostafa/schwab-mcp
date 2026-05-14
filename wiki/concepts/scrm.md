---
title: "Supply Chain Risk Management (SCRM)"
type: concept
domain: 1
tags: [scrm, supply-chain, third-party, vendor, sla, slr, tampering, counterfeits, implants]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Supply Chain Risk Management (SCRM)

Risk management must extend to **all external entities**: vendors, suppliers, cloud service
providers, contractors, and customers. Accountability for data and security can never be
outsourced, even when responsibility is transferred to a third party.

## Key Risks in the Supply Chain (Table 1.31)

| Risk | Definition | Example |
|---|---|---|
| **Product Tampering** | Unauthorized alteration of a product after manufacturing but before delivery | Attacker intercepts a keyboard shipment, solders a keylogger inside, repackages it |
| **Counterfeits** | Unauthorized replicas intended to deceive consumers | Fake networking components with reduced performance or increased vulnerabilities |
| **Implants** | Hardware/software stealthily inserted into products to perform unauthorized activities (espionage, data theft) | Backdoor chip in server hardware; unauthorized firmware |

## Risk Mitigations (Table 1.32)

| Mitigation | Description |
|---|---|
| **Third-party Assessment and Monitoring** | Continuously evaluate vendors' security practices and performance |
| **Minimum Security Requirements** | Predefined baseline security standards vendors must meet |
| **Service-level Requirements (SLR)** | Specifications for expected performance, availability, responsiveness |
| **Silicon Root of Trust** | Secure cryptographic identity embedded in hardware; ensures trusted boot and genuine firmware |
| **Physically Unclonable Function (PUF)** | Uses unique physical characteristics of semiconductors to generate cryptographic keys — each device has an unclonable identity; mitigates counterfeit risk |
| **Software Bill of Materials (SBOM)** | Comprehensive list of components, libraries, modules, versions, sources, dependencies; enables detection of backdoors and hidden vulnerabilities |

## Procurement Process: SLR → SLA → Service Level Reports

### Service Level Requirements (SLR)
- Pre-contract document that defines **detailed service descriptions**, **service level targets**, and **mutual responsibilities**.
- Used during procurement to evaluate competing vendors against documented security requirements.
- Informs requirements in the eventual SLA.

### Service Level Agreement (SLA)
- Contract addendum — **legally enforceable**.
- Documents: service performance levels, governance (who is responsible for what), security controls, compliance requirements, liability/indemnification terms.
- Even when an SLA exists, **the customer (data owner) remains accountable** for all customer data.

### Service Level Reports
- Issued by vendor/service provider to the client on an ongoing basis.
- Compares agreed SLA targets vs. actual performance.
- May include: metrics achievement, issue identification, third-party SOC 2 Type 2 audit reports.
- **SOC 2 Type 2** = third-party assurance report — used when a customer cannot audit the provider directly.

## SCRM Assessment Activities

When evaluating suppliers, organizations should include:
- Governance review
- Site security review
- Formal security audit
- Penetration testing
- Adherence to security baselines
- Evaluation of hardware and software
- Adherence to security policies

## Exam-Relevant Nuance

- **Accountability cannot be outsourced.** Even with a CSP handling data, the owning organization remains accountable to regulators and customers.
- SBOM is a relatively new supply chain control — appearing more frequently on modern exam questions.
- Silicon Root of Trust and PUF are hardware-level supply chain controls — know their purpose at a conceptual level.
- The SLR comes **before** the contract; the SLA is **in** the contract.

## Cross-links

- [Risk Management](risk-management.md)
- [Governance](governance.md)
- [Personnel Security](personnel-security.md)
- [Compliance Requirements](compliance-requirements.md)

## Sources

- destination-cissp §1.11, §1.11.1, §1.11.2, and SLR/SLA section (Tables 1.31, 1.32)
