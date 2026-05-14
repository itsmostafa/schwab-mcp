---
title: "Domain 5 — Identity and Access Management (IAM)"
type: domain
domain: 5
tags: [iam, authentication, authorization, aaa, sso, federation, rbac, mfa, kerberos, saml, oauth, biometrics]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 5 — Identity and Access Management (IAM)

Exam weight: **13%**. Covers controlling access to systems and data through authentication,
authorization, identity federation, and access provisioning lifecycle management.

## Subtopics (from CISSP Exam Outline)

### 5.1 Control physical and logical access to assets

Access control is the collection of mechanisms that protect organizational assets while allowing controlled access to authorized subjects. It applies to all asset types: facilities, systems/devices, information, personnel, and applications.

**Fundamental access control principles:**

| Principle | Meaning |
|---|---|
| **Need to know** | Access granted only to assets required for the job — nothing beyond |
| **Least privilege** | Minimum level of access needed; nothing more |
| **Separation of duties** | No single person controls all aspects of a process; prevents fraud and errors |

**Groups vs. Roles:**
- *Groups:* Collections of users, not necessarily tied to a specific job (e.g., a BCM team).
- *Roles:* Sets of permissions tied to a specific job function (e.g., call center agent). Roles are more focused on function and permissions.

**Access control administration approaches:**
- *Centralized:* One system controls access to all; single point of failure.
- *Decentralized:* Control at the resource level; peer-to-peer; separate credentials per resource; lacks standardization.
- *Hybrid:* Most common; combines both approaches, often due to legacy systems.

**Reference Monitor Concept (RMC):** A rules-based mediator placed between subjects and objects to control access; all activity logged. Any implementation is called a **security kernel**.

### 5.2 Design identification and authentication strategy

**Seven Laws of Identity** (Kim Cameron): User control and consent; minimal disclosure; justifiable parties; directed identity (omni-directional for public entities, uni-directional for private); pluralism of operators/technologies; human integration; consistent experience across contexts.

**Access Control Services sequence:**

```
Identification → Authentication → Authorization → Accounting
```

All four must be in place to achieve the **Principle of Access Control (= Accountability)**. Shared accounts undermine this principle.

**Identification (§5.2.2):** The unique assertion of identity (username, access card, biometric scan). Must be: unique, non-descriptive of role, issued and used securely.

**Authentication by Knowledge (§5.2.3):** Something you know — password, passphrase, cognitive/security questions. Answers to security questions don't need to be true (which makes them unguessable).

**Authentication by Ownership (§5.2.4):** Something you have — OTP (hard/soft tokens), smart cards, memory cards, passkeys.
- *OTP:* Synchronous (simpler) or asynchronous (more robust, more complex) generation.
- *Smart card:* Contains IC chip; computes unique data per transaction. Contact or contactless communication.
- *Memory card:* Magnetic strip; same data read every time — vulnerable to skimming.

**Authentication by Characteristics (§5.2.5):** [Biometrics](../concepts/biometrics.md) — physiological (fingerprint, hand geometry, vascular pattern, facial, iris, retina) and behavioral (voice, signature, keystroke, gait). Retina = most accurate but invasive.

**Factors of Authentication (§5.2.6):** Destination-cissp defines **three factor families**:
1. Something you know (knowledge)
2. Something you have (ownership)
3. Something you are (characteristic)

**MFA requires two or more *different* factor families.** Using two "knowledge" factors is still single-factor authentication.

**Password-less authentication:** Passkeys, biometrics, hardware tokens — reduce password friction and phishing risk, but introduce device-loss risk and higher implementation cost.

**Credential Management Systems (§5.2.7):** Manage credentials at scale using strong two-factor authentication and PKI. Includes **password vaults** (password managers) — generate/store/sync passwords behind one master password. Advantage: strong unique passwords per account. Disadvantage: single point of failure.

**Single Sign-On / Kerberos (§5.2.8):** See [SSO](../concepts/sso.md) and [Kerberos](../concepts/kerberos.md). Kerberos uses KDC (AS + TGS), TGT, and service tickets. Weaknesses: symmetric-only encryption, TOCTOU vulnerability. **SESAME** improves Kerberos with asymmetric crypto + multiple tickets but is rarely adopted.

**CAPTCHA (§5.2.9):** Completely Automated Public Turing test to tell Computers and Humans Apart. Prevents bot-based account creation, spam, and brute-force attacks.

**Session Management (§5.2.10):** Sessions created after successful identification + authentication + authorization. Session hijacking risk is mitigated by: frequent re-authentication, schedule limitations, login limitations, timeouts, screensavers.

**Registration and Proofing of Identity (§5.2.11):** Identity proofing confirms a person is who they claim before issuing credentials. Uses government-issued ID. Analogous to RA proofing before a CA issues a certificate.

**Authenticator Assurance Levels / AAL (§5.2.12):** Defined in [NIST SP 800-63B](../standards/nist-sp-800-63b.md). AAL1 = SFA; AAL2 = MFA + crypto; AAL3 = MFA + hard cryptographic authenticator + impersonation resistance.

**Federated Identity Management (§5.2.13):** See [Federation](../concepts/federation.md). FIM = SSO across organizational boundaries. Three components: Principal (user), Identity Provider (IdP), Relying Party (SP).

**Federated Access Standards (§5.2.14):** See [SAML](../concepts/saml.md) and [OAuth/OIDC](../concepts/oauth-oidc.md).
- SAML: XML-based; authentication + authorization via assertion tickets.
- WS-Federation: authentication + authorization; codified by OASIS.
- OpenID (original): authentication only.
- OAuth: authorization only.
- OIDC: authentication layer on top of OAuth 2.0.

**Accountability = Principle of Access Control (§5.2.15):** Unique ID + proper authentication + proper authorization + logging/monitoring.

**Just-in-Time (JIT) Access (§5.2.16):** Temporary privilege elevation for specific task windows. Reduces standing privilege exposure; often automated. See [Privileged Access Management](../concepts/privileged-access-management.md).

### 5.3 Federated identity with a third-party service

**Identity as a Service (IDaaS) (§5.3.1):** Cloud-based IAM — provisioning, administration, SSO, MFA, directory services. Supports cloud, synced, linked, and federated identity types.

IAM deployment models: On-Premises (secure, no internet dependency), Cloud (availability + multitenant risk), Hybrid (most flexible).

IDaaS risks: availability outages, protection of PII held by CSP, third-party data exposure.

### 5.4 Implement and manage authorization mechanisms

See [Access Control Models](../concepts/access-control-models.md) for full detail.

- **DAC (Discretionary):** Owner decides access. Subtypes: rule-based, RBAC, ABAC.
- **Rule-Based AC:** Explicit rules/ACL per subject-object pair; granular but high admin overhead.
- **RBAC (Role-Based):** Access by job function; best practice; mirrors org chart. Full-RBAC can produce role explosion — most orgs use Hybrid/Limited RBAC.
- **ABAC (Attribute-Based):** Access by user + environment attributes; most flexible; enables cloud/zero-trust. Enabled by XACML standard.
- **MAC (Mandatory):** System decides based on labels (classification) and clearances; protects confidentiality; used in government/military.
- **Non-Discretionary:** Third party (not the owner) decides; anti-pattern.
- **Risk-Based:** System computes risk profile of request; may challenge with additional auth.
- **Context-Based:** Similar to risk-based; typically enforced via firewall rules (internal vs. external context).

**Access Policy Enforcement:**
- **PEP (Policy Enforcement Point):** Gatekeeper at access points; sends requests to PDP; enforces decisions.
- **PDP (Policy Decision Point):** Centralized evaluator; returns allow/deny based on policy rules.

### 5.5 Manage the identity and access provisioning lifecycle

See [Identity Lifecycle](../concepts/identity-lifecycle.md) and [Privileged Access Management](../concepts/privileged-access-management.md).

**Identity Life Cycle (§5.5.2):** Provisioning → Review → Revocation.

**Vendor Access (§5.5.1):** Third-party vendor access requires same or greater care as employee provisioning. May include security review and onsite inspection.

**User Access Review (§5.5.3):** Conducted by asset/system owner. Frequency driven by asset value and risk. Privileged accounts: potentially weekly. Standard accounts: at minimum annually. Role changes and departures trigger immediate review.

**Privilege Escalation / sudo (§5.5.4):** Privileged users maintain two accounts (standard + privileged). Use sudo/RunAs only for admin tasks. Audit sudo use.

**Service Account Management (§5.5.5):** Accounts used by services, not humans. Apply single-purpose and least-privilege principles. Human oversight required to prevent privilege escalation and spoofing.

### 5.6 Implement authentication systems

Authentication systems covered (§5.6.1):
- **OpenID Connect (OIDC) / OAuth 2.0** — authorization + authentication layer. See [OAuth/OIDC](../concepts/oauth-oidc.md).
- **SAML** — enterprise FIM. See [SAML](../concepts/saml.md).
- **Kerberos** — SSO for Windows/AD environments. See [Kerberos](../concepts/kerberos.md).
- **RADIUS / TACACS+** — centralized network authentication protocols. (Also covered in Domain 3.)

## Key concepts

- [AAA — Authentication, Authorization, Accounting](../concepts/aaa.md) — the four services of access control; Principle of Access Control
- [Authentication Factors](../concepts/authentication-factors.md) — three factor families; MFA vs. SFA; OTP; smart/memory cards
- [Biometrics](../concepts/biometrics.md) — CER, FAR, FRR; physiological vs. behavioral; retina most accurate
- [SSO](../concepts/sso.md) — single sign-on; pros/cons; SESAME
- [Kerberos](../concepts/kerberos.md) — KDC, AS, TGS, TGT; TOCTOU weakness; symmetric-only
- [Federation](../concepts/federation.md) — FIM; IdP/SP/Principal; IDaaS; trust relationships
- [SAML](../concepts/saml.md) — XML-based; assertion tickets; auth + authz
- [OAuth/OIDC](../concepts/oauth-oidc.md) — OAuth = authorization only; OIDC = authentication on OAuth
- [Access Control Models](../concepts/access-control-models.md) — DAC/MAC/RBAC/ABAC/Rule-based; PEP/PDP; XACML; RMC
- [Privileged Access Management](../concepts/privileged-access-management.md) — JIT; two-account model; sudo; service accounts
- [Identity Lifecycle](../concepts/identity-lifecycle.md) — provisioning → review → revocation; joiner/mover/leaver; privilege creep
- [Directory Services](../concepts/directory-services.md) — LDAP, Active Directory, DN/OU
- [NIST SP 800-63B](../standards/nist-sp-800-63b.md) — AAL levels

## Sources

- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §5.1–§5.6 (pp. 664–750)
