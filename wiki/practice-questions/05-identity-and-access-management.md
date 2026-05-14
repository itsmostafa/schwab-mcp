---
title: "Practice Questions — Domain 05: Identity And Access Management"
type: practice
domain: 05
tags: [practice, iam, authentication, access-control, biometrics, kerberos, saml, mfa]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 05: Identity And Access Management

## Questions

### Q1: A biometric system is tuned to minimize false acceptances. What is the likely effect on false rejections?

**Answer:** False rejections will increase.

**Why:** FAR (False Acceptance Rate) and FRR (False Rejection Rate) are inversely related. Tightening the system's sensitivity reduces FAR (Type 2 errors — bad actors accepted) at the cost of increasing FRR (Type 1 errors — legitimate users rejected). Neither can be eliminated simultaneously without raising the other.

*Subtopic: 5.2.5 — [Biometrics](../concepts/biometrics.md)*

---

### Q2: When evaluating two competing biometric systems for a new authentication deployment, which metric is most useful for comparison?

**Answer:** Crossover Error Rate (CER).

**Why:** CER is the point where FAR and FRR are equal. It provides a single, comparable metric for overall biometric system accuracy. A lower CER means a more accurate system. CER removes the ambiguity of comparing systems that may excel in FAR but suffer in FRR or vice versa.

*Subtopic: 5.2.5 — [Biometrics](../concepts/biometrics.md)*

---

### Q3: Which biometric type is considered most accurate, and why is it rarely deployed despite its accuracy?

**Answer:** Retina scanning is most accurate. It is rarely deployed because it is invasive (requires pressing the eye to a rubber cup with a bright light flash), unpleasant for users, and controversial — retina scans can reveal medical conditions, creating privacy concerns.

**Why:** Accuracy alone is insufficient; user acceptance and data sensitivity must also be considered in biometric system selection.

*Subtopic: 5.2.5 — [Biometrics](../concepts/biometrics.md)*

---

### Q4: A user authenticates with their username/password and then answers a security question. Is this single-factor or multi-factor authentication?

**Answer:** Single-factor authentication (SFA).

**Why:** Both a password and a security question fall under "something you know" (Authentication by Knowledge). MFA requires two or more *different factor families*. Using multiple objects from the same factor family does not constitute MFA.

*Subtopic: 5.2.6 — [Authentication Factors](../concepts/authentication-factors.md)*

---

### Q5: In Kerberos, what is the purpose of the Ticket Granting Ticket (TGT), and why can the user not decrypt it?

**Answer:** The TGT is issued by the Authentication Service (AS) and presented by the user to the Ticket Granting Service (TGS) to obtain service tickets. The user cannot decrypt the TGT because it is encrypted with the TGS's key — not the user's key. This is by design: the TGT is an opaque token that proves the AS has validated the user, without exposing TGS secrets to the user.

**Why:** The separation of encryption domains (user's key for one message, TGS key for TGT) allows Kerberos to verify user identity without transmitting the password over the network.

*Subtopic: 5.2.8 — [Kerberos](../concepts/kerberos.md)*

---

### Q6: What are the two primary weaknesses of Kerberos identified in the source material, and how is each addressed?

**Answer:**
1. **Symmetric encryption only** (RC4, DES, AES) — implies key distribution challenges inherent to symmetric crypto. Addressed by SESAME, which adds asymmetric cryptography support.
2. **TOCTOU (Time-Of-Check Time-Of-Use) vulnerability** — single ticket model means a compromised session can be replayed until the ticket expires. Addressed by increasing re-authentication frequency (shortening ticket lifetimes), especially for high-value systems.

**Why:** Understanding Kerberos weaknesses is important for both exam questions on protocol selection and for recommending compensating controls.

*Subtopic: 5.2.8 — [Kerberos](../concepts/kerberos.md)*

---

### Q7: What is the critical difference between OAuth 2.0 and OpenID Connect (OIDC)?

**Answer:** OAuth 2.0 is an *authorization* framework only — it delegates resource access without authenticating the user. OpenID Connect (OIDC) is an *authentication* layer built on top of OAuth 2.0 — it adds user identity verification (via an ID token).

**Why:** Exam questions frequently test whether candidates confuse OAuth 2.0 (authorization only) with authentication. The correct answer for "authentication across federated systems" is OIDC or SAML — not OAuth 2.0 alone.

*Subtopic: 5.2.14, 5.6.1 — [OAuth/OIDC](../concepts/oauth-oidc.md)*

---

### Q8: An organization needs an access control model for a cloud-based application where access decisions must consider the user's device type, IP address, time of access, and asset classification. Which model is most appropriate?

**Answer:** Attribute-Based Access Control (ABAC).

**Why:** ABAC evaluates multiple attributes of the subject, the resource, and the environment simultaneously. It is the most flexible access control model and is specifically suited to cloud environments where traditional network perimeters do not exist. RBAC would only consider role/job function and could not natively account for dynamic environmental factors.

*Subtopic: 5.4.1 — [Access Control Models](../concepts/access-control-models.md)*

---

### Q9: In the context of Mandatory Access Control (MAC), what determines whether a subject can access an object?

**Answer:** The system compares the subject's **clearance level** against the object's **classification label**. Access is granted only if the subject's clearance meets or exceeds the object's classification. The owner does not make this decision — the system does.

**Why:** The defining feature of MAC is that the system enforces access based on labels, not owner discretion. This distinguishes it from DAC, where the owner decides.

*Subtopic: 5.4.2 — [Access Control Models](../concepts/access-control-models.md)*

---

### Q10: What is the key SAML component that a Service Provider uses to make an authorization decision, and what does it NOT contain?

**Answer:** The **SAML Assertion Ticket/Token** contains assertion statements about the user (identity claims, role, level of access, attributes). It does **NOT** contain the user's username or password.

**Why:** SAML is designed so the SP never needs to see credentials. The IdP authenticates the user and vouches for them via the assertion, which the SP trusts based on the established trust relationship.

*Subtopic: 5.2.13–5.2.14 — [SAML](../concepts/saml.md)*

---

### Q11: An employee moves from the Finance department to the Marketing department. According to identity lifecycle best practices, what should happen to their account access?

**Answer:** The employee's Finance access should be **revoked**, and Marketing access should be **provisioned fresh** — rather than simply adding Marketing permissions on top of existing Finance access.

**Why:** Simply adding new access without revoking old access leads to **privilege creep**, violating least privilege. The correct approach is to treat a role change like a leaver/joiner event: deprovision the old role's access, then provision the new role's access. This prevents accumulation of unauthorized permissions.

*Subtopic: 5.5.2, 5.5.3 — [Identity Lifecycle](../concepts/identity-lifecycle.md)*

---

### Q12: In a federated identity system using SAML, what are the three key components and what role does each play?

**Answer:**
1. **Principal (User):** The person requesting access.
2. **Identity Provider (IdP):** Owns the user's identity; performs authentication and issues the SAML assertion.
3. **Relying Party (Service Provider / SP):** Receives and trusts the SAML assertion; makes the authorization decision based on assertion statements.

**Why:** Understanding these three roles and how trust flows between them (IdP → SP via assertion) is fundamental to federated identity questions on the exam. The SP trusts the IdP based on a pre-established trust relationship, not on the user's credentials.

*Subtopic: 5.2.13 — [Federation](../concepts/federation.md) | [SAML](../concepts/saml.md)*
