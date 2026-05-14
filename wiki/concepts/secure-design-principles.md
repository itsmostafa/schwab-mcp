---
title: "Secure Design Principles"
type: concept
domain: 3
tags: [secure-design, least-privilege, defense-in-depth, zero-trust, fail-secure, privacy-by-design, threat-modeling]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Secure Design Principles

Secure design principles are foundational rules that should be applied from the inception of any architecture or system — not retrofitted afterward. The guiding philosophy is **security by design**: embedding protection into the structure of a solution rather than bolting it on post-deployment.

## Key Principles

### Least Privilege
Every subject (user, process, program) should be granted only the minimum permissions required to perform its function, and nothing more. Reduces blast radius of compromise.

### Defense in Depth (Layered Security)
No single control should be the sole protection for a valuable asset. Multiple control layers should be stacked, each comprising a combination of preventive, detective, and corrective controls. Covered in depth under [Trusted Computing Base](trusted-computing-base.md) (§3.4.9).

### Fail Securely (Fail-Safe)
When a system or component fails, it should fail in a state that does not expose the system to attack. Example: an electronic safe lock that fails locked, not unlocked, when power is cut.

### Secure Defaults
All default settings should be as secure as possible. An admin account with no password set by default is a failure of this principle.

### Segregation of Duties (Separation of Duties)
No single person should be able to perform all tasks related to a critical function without oversight from another. Reduces fraud and error. Also relevant in Clark–Wilson integrity model. Cross-reference: [Domain 1 — SoD coverage](./governance.md).

### Need to Know
Access to information should be granted only when the subject has a legitimate operational need for it, independent of clearance level.

### Keep It Simple and Small
Complexity is the enemy of security. Simpler designs have smaller attack surfaces, fewer vulnerabilities, easier testing, and simpler troubleshooting. Unnecessary complexity increases the likelihood of misconfiguration.

### Zero Trust
"Trust nothing" by default — not internal or external entities. Prior to granting access, every user, device, and service must be authenticated and authorized. Key principles:
1. Know your architecture (users, devices, services)
2. Know your identities
3. Know the health of users, devices, services
4. Use policies to authorize requests
5. Authenticate everywhere
6. Focus monitoring on devices and services
7. Don't trust any network, including your own
8. Choose services designed for zero trust

Network micro-segmentation is a primary implementation tactic — it forces re-authentication as users move between segments.

### Trust but Verify
An oxymoron: complete trust implies no verification is needed, but complete distrust is unworkable. In practice, "trust but verify" means implementing strong authentication and authorization controls combined with real-time monitoring. Third-party trust should be verified through audits, SOC reports, SLAs, and ongoing monitoring.

### Privacy by Design (PbD)
Privacy should be embedded into system design from inception, not added as an afterthought. Seven foundational principles:
1. **Proactive, not reactive** — anticipate and prevent privacy shortcomings before they occur
2. **Privacy as the default** — automatic protection of personal data without user action required
3. **Privacy embedded into design** — part of the core functionality, not a layer on top
4. **Full functionality** — "win-win," no unnecessary trade-offs
5. **End-to-end security** — from cradle to grave across the data life cycle
6. **Visibility and transparency** — operations are open to independent verification
7. **Respect for user privacy** — user-centric; strong defaults, appropriate notice, user control

Security professionals note that PbD is largely equivalent to "security by design," since privacy requirements are enabled by security controls.

### Shared Responsibility
As organizations increasingly rely on cloud and third-party services, responsibility for security is shared between internal staff and external providers. The cloud customer is always **accountable** for their data; responsibility for specific controls can be delegated through SLAs. See [Cloud Security Models](cloud-security-models.md).

### Open Design
Security mechanisms should not rely on secrecy of their design (no "security by obscurity"). The strength of the mechanism should reside in the key, not in the algorithm being hidden.

## Exam-Relevant Nuance

- **Fail secure vs. fail safe**: Fail secure = system locks down on failure (e.g., door locks). Fail safe = system allows people to exit safely in an emergency (life-safety context). Context determines which is appropriate.
- CISSP Exam Outline lists threat modeling, least privilege, defense in depth, secure defaults, fail securely, and SoD explicitly for section 3.1.
- Zero trust does not eliminate the need for monitoring — it increases the importance of continuous logging and anomaly detection.
- Privacy by Design's principle of "privacy as default" mirrors the firewall concept of "implicit deny."

## Cross-links

- [Security Models](security-models.md) — models that formalize many of these principles into lattice or rule structures
- [Trusted Computing Base](trusted-computing-base.md) — how these principles are implemented in hardware/software
- [Cloud Security Models](cloud-security-models.md) — shared responsibility detail
- [Trusted Computing Base](trusted-computing-base.md) — defense in depth / layered security detail

## Sources

- destination-cissp §3.1.1–3.1.2 (pp. 225–240)
- cissp-exam-outline §3.1
