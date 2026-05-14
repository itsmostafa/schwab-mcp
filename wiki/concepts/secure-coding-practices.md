---
title: "Secure Coding Practices"
type: concept
domain: 8
tags: [secure-coding, input-validation, coupling, cohesion, polyinstantiation, inheritance, encapsulation, owasp, toctou, covert-channel]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Secure Coding Practices

## Definition

**Secure coding practices** are steps and techniques applied during software development to minimize or eliminate vulnerabilities that could be exploited. They transform security requirements into concrete coding behaviors, ensuring that the finished application is resistant to attack by design rather than by luck.

---

## Key Secure Coding Practices

*(source: destination-cissp §8.5.4)*

| Practice | Description |
|---|---|
| **Input validation** | Check all input before accepting it. Validate type, length, format, and range. Inadequate input validation is one of the leading causes of web application attacks. |
| **Authentication & password management** | Implement strong authentication; store passwords using approved hashing algorithms (bcrypt, Argon2); enforce complexity and rotation policies. |
| **Session management** | Use secure, random session tokens; enforce timeouts; invalidate sessions on logout; protect tokens from XSS/CSRF. |
| **Cryptographic practices** | Use vetted, approved algorithms and libraries; never implement custom crypto; protect keys. |
| **Error handling and logging** | Return generic error messages to users (do not expose stack traces, file paths, or internals); log detailed errors server-side for diagnostics. |
| **System configuration** | Secure defaults; disable unnecessary features; principle of least privilege for service accounts and database users. |
| **File/database security** | Parameterized queries; avoid dynamic SQL; restrict file access. |
| **Memory management** | Bounds checking; safe functions; avoid unsafe string operations. See [Memory Safety](memory-safety.md). |

---

## OOP Security Terms (Exam-Relevant)

Modern programming languages use object-oriented concepts that have security implications:

| Term | Definition | Security Relevance |
|---|---|---|
| **Inheritance** | A new object automatically inherits characteristics (methods, properties) of a parent object. | Starting point must be secure — if parent has weak security, all child objects inherit it. |
| **Encapsulation** | An object (piece of code) is wrapped to hide internal details. Data and methods are bundled together; access is controlled via defined interfaces. | Protects internals; limits attack surface. |
| **Polymorphism** | Code that can change behavior based on requirements or the environment ("smart code"). | Similar to polymorphic malware in concept — code that adapts. Unlike malware, used for legitimate adaptation. |
| **Polyinstantiation** | A single data object exists as multiple separate instances, each visible only at an appropriate classification level. | Prevents unauthorized inference (see below). |

---

## Coupling and Cohesion

**Coupling** = the degree of interdependence between separate units of code.
**Cohesion** = the degree to which the code *within* a single unit is related to a single purpose.

| Metric | Optimal State | Why |
|---|---|---|
| **Coupling** | **Low** — units can stand alone without depending on other units | Reduces blast radius: a change/bug in one module doesn't cascade to others |
| **Cohesion** | **High** — all code within a module does one thing well | Easier to reason about; easier to test; fewer unexpected side effects |

> **Exam rule:** Low coupling + high cohesion = well-written, maintainable, secure code. High coupling + low cohesion = poorly written code that is difficult to secure, test, and maintain.

---

## Polyinstantiation (Deep Dive)

Polyinstantiation allows the *same named data* to exist as multiple independent instances at different classification levels. This prevents **unauthorized inference** — where a lower-privileged user discovers the existence of classified data simply by observing an error message.

**Example (destination-cissp §8.5.4):**  
A military system tracks units. An operator at a lower clearance level tries to add "Charlie Company 6" and receives: *"Unable to add unit Charlie Company 6."* This error reveals that the unit already exists at a higher classification level — an unauthorized inference.

With polyinstantiation: the system recognizes Charlie Company 6 exists at a higher level, and silently creates a separate, lower-classification instance for the operator's view. No inference is possible.

**Applies to:** Database systems handling multi-level security (MLS) environments.

---

## Source-Level Vulnerabilities (Common Weaknesses)

*(source: destination-cissp §8.5.1)*

| Vulnerability | Description |
|---|---|
| **Covert channels** | Unintentional communications paths that can disclose sensitive information. Two types: **timing** (information encoded in timing patterns) and **storage** (shared storage used as a side channel). |
| **Buffer overflow** | Input exceeds allocated buffer size; can allow attacker to inject and execute malicious code. See [Memory Safety](memory-safety.md). |
| **Memory/object reuse** | Residual data in memory from a previous operation can be read by another process. Data should be zeroed/overwritten before reuse. |
| **Executable mobile code** | Code downloaded and executed locally (e.g., via a clicked link). If malicious, can cause serious harm. Mitigation: sandbox environment testing before execution. |
| **TOCTOU (race condition)** | Time-of-Check Time-of-Use gap: a value is checked/enforced at one time but used at another. The gap allows malicious code to swap a legitimate resource for a malicious one. |
| **Backdoors / trapdoors** | Intentionally placed by developers for maintenance access; dangerous if not removed post-development. Also called **maintenance hooks**. Often only discoverable via source code review. |
| **Malformed input** | Input that does not meet validation criteria. Inadequate input validation is one of the top causes of web application attacks (see OWASP Top 10). |
| **Citizen developers** | Non-technical users given access to powerful programming tools (e.g., SQL query tools) without commensurate security knowledge. Mitigated through policies, security awareness, and training. |

---

## Software-Defined Security

Security implemented, controlled, and managed entirely by software rather than dedicated hardware. Grown alongside cloud and virtualization. Can implement firewalls, IDS/IPS, access control, and other security functions as software services. Policy-driven; cost-effective; dynamic.

---

## Exam-Relevant Nuance

- **Polyinstantiation** prevents inference; does NOT prevent access — it controls what different security levels *see*.
- **TOCTOU** = race condition = time gap exploit. Mitigation: atomic operations, locking, minimize gap between check and use.
- **Covert channels** cannot always be eliminated; they can be reduced (bandwidth limiting) or monitored.
- **Backdoors** are a software supply chain risk — present in developed, COTS, and open source software.
- **Error handling**: Never return detailed errors to the user; always log them server-side.

---

## Cross-Links

- [Memory Safety](memory-safety.md) — buffer overflow, ASLR, DEP, stack canaries
- [Database Security](database-security.md) — SQL injection, polyinstantiation in DB context, parameterized queries
- [OWASP Top 10](owasp-top-10.md) — mapping of vulnerabilities to OWASP categories
- [SDLC](sdlc.md) — which SDLC phase secure coding applies to
- [Software Testing Types](./security-testing-types.md) — SAST/DAST for finding secure coding violations

## Sources

- destination-cissp §8.5.1, §8.5.4, §8.5.5 (pp. 0965–0972)
- cissp-exam-outline (Domain 8.5: Secure coding guidelines and standards)
