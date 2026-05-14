---
title: "Practice Questions — Domain 08: Software Development Security"
type: practice
domain: 08
tags: [practice]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 08: Software Development Security

## Questions

### Q1: During which SDLC phase should threat modeling be performed?

**Answer:** The Architecture and Design phase.

**Why:** Threat modeling identifies what threats exist against the application and what security controls should be built into the design. If it is performed earlier (during requirements), the design is not yet defined. If later (during development or testing), the architecture is already fixed and changes are expensive. The design phase is the right point because security elements can still be incorporated into the architecture before coding begins.

*Subtopic: 8.1 — [SDLC](../domains/08-software-development-security.md)*

---

### Q2: A security team wants to test a production application without access to its source code, by sending various inputs and observing outputs. Which type of testing is this?

**Answer:** Dynamic Application Security Testing (DAST) — specifically black-box testing.

**Why:** DAST tests the running application without examining source code. The tester observes how the application behaves in response to inputs, mimicking an attacker's perspective. This contrasts with SAST (white-box; source code examined; application not running) and IAST (running application + source code visible via agent). Fuzz testing is a subtype of dynamic/black-box testing that uses random/malformed inputs specifically.

*Subtopic: 8.5 — [Security Testing Types](../concepts/security-testing-types.md)*

---

### Q3: A web application concatenates user-supplied input directly into SQL queries without validation. An attacker enters `' OR '1'='1` into a login field and gains unauthorized access. Which OWASP Top 10 (2021) category does this represent, and what is the primary defense?

**Answer:** A03 — Injection. The primary defense is parameterized queries (prepared statements).

**Why:** SQL injection occurs when user input is interpreted as SQL code rather than data. Parameterized queries separate the query structure from the data — the user input is always treated as a literal value and can never alter the query's logic. Input filtering alone is insufficient because it can be bypassed. Least-privilege database accounts reduce the damage but do not prevent the injection itself.

*Subtopic: 8.5 — [Database Security](../concepts/database-security.md); [OWASP Top 10](../concepts/owasp-top-10.md)*

---

### Q4: An organization purchases commercial software and is concerned that if the vendor goes out of business, they will lose the ability to maintain the application. What control addresses this risk?

**Answer:** Software escrow.

**Why:** Software escrow involves a neutral third party holding the source code. Predefined trigger conditions (vendor bankruptcy, failure to support the product) release the source code to the purchasing organization, allowing them to maintain, modify, or migrate the application independently. This is a standard mitigation for COTS software where the vendor does not provide source code to customers.

*Subtopic: 8.4 — [Software Acquisition Security](../concepts/software-acquisition-security.md)*

---

### Q5: Which SDLC testing activity can overlap with the development phase, beginning before all code is complete?

**Answer:** Static Application Security Testing (SAST) / code review.

**Why:** SAST examines source code without running the application, so it can be applied incrementally as individual units of code are completed. This allows developers to get feedback on security issues while the code is still fresh and changes are inexpensive. DAST requires a running application and therefore must wait until at least a testable build exists.

*Subtopic: 8.1 / 8.5 — [SDLC](../concepts/sdlc.md); [Security Testing Types](../concepts/security-testing-types.md)*

---

### Q6: A DevOps team wants to incorporate security testing into their pipeline so that every code commit is automatically checked. Which concept describes this approach, and what is the primary value?

**Answer:** DevSecOps (or SecDevOps); the value is automating security testing so it keeps pace with rapid development iteration.

**Why:** Traditional security techniques (manual penetration tests, periodic security reviews) are too slow for DevOps velocity. By automating SAST and DAST within the CI/CD pipeline, security checks run on every commit — making security continuous rather than periodic. This is the "shift left" approach: catching security issues early in the pipeline when they are cheapest to fix.

*Subtopic: 8.1 — [DevSecOps](../concepts/devsecops.md); [CI/CD Security](../concepts/ci-cd-security.md)*

---

### Q7: A military database system uses a technique that allows the same named data record to exist as two separate entries — one visible to users with a lower clearance level and one visible only to those with a higher clearance. What technique is this, and what attack does it prevent?

**Answer:** Polyinstantiation. It prevents unauthorized inference.

**Why:** Without polyinstantiation, a lower-clearance user attempting to add a record that already exists at a higher classification level would receive an error like "record already exists" — inadvertently revealing that a classified record exists. Polyinstantiation creates a separate instance of the same named object at each classification level, so users at lower levels see their own version and receive no signal about higher-level data.

*Subtopic: 8.5 — [Secure Coding Practices](../concepts/secure-coding-practices.md); [Database Security](../concepts/database-security.md)*

---

### Q8: An organization's CMMI assessment places them at Level 2. What does this tell you about the organization's software development process, and what must they achieve to reach Level 3?

**Answer:** Level 2 (Managed) means projects are planned and managed with metrics; work gets done but the organization is still primarily reactive. To reach Level 3 (Defined), the organization must become proactive — establishing standardized processes across the entire organization (not just project-by-project) that guide all programs, portfolios, and projects.

**Why:** The key distinction between Level 2 and Level 3 is the shift from reactive to proactive. Level 3 requires organization-wide standards — not just individual project management. Levels 4 and 5 then add quantitative measurement and continuous optimization on top of that foundation.

*Subtopic: 8.1 — [CMMI/SAMM](../standards/cmm-samm.md)*

---

### Q9: Which OWASP Top 10 (2021) category specifically addresses the risk that an application uses libraries or frameworks with known security vulnerabilities, and what is the most effective mitigation?

**Answer:** A06 — Vulnerable and Outdated Components. The most effective mitigation is maintaining an inventory of all dependencies (SBOM), subscribing to CVE notifications, and regularly updating/patching components.

**Why:** A06 directly targets supply chain risk from third-party components. A Software Bill of Materials (SBOM) enables rapid identification of which products are affected when a new CVE is disclosed (e.g., Log4Shell affected any product using Log4j). Software Composition Analysis (SCA) tools automate dependency scanning in the CI/CD pipeline.

*Subtopic: 8.5 — [OWASP Top 10](../concepts/owasp-top-10.md); [Software Acquisition Security](../concepts/software-acquisition-security.md)*

---

### Q10: A developer writes code where Module A calls Module B, Module B calls Module C, and Module C calls Module A — a tight circular dependency. What software design problem does this represent, and what is the preferred state?

**Answer:** This represents high coupling — modules are highly interdependent and cannot function independently. The preferred state is low coupling (modules can stand alone) combined with high cohesion (each module's internal code is highly related to a single purpose).

**Why:** High coupling makes code difficult to test, maintain, and secure — a change in one module may break others unpredictably. Low coupling enables modular development, isolated testing, and reduces the blast radius of bugs. Low coupling + high cohesion is the hallmark of well-designed, maintainable, and secure code.

*Subtopic: 8.5 — [Secure Coding Practices](../concepts/secure-coding-practices.md)*

---

### Q11: An attacker sends an input string of 500 characters to a field designed to hold 50 characters. The excess data overwrites adjacent memory, eventually allowing the attacker to redirect code execution. What vulnerability is this, and what OS-level mitigation specifically randomizes memory layouts to make this attack harder to execute reliably?

**Answer:** Buffer overflow vulnerability. The OS-level mitigation is Address Space Layout Randomization (ASLR).

**Why:** Buffer overflows occur when input exceeds the allocated buffer, allowing overflow data to overwrite adjacent memory (including return addresses on the stack). ASLR defeats deterministic exploitation by randomizing where executables, stack, heap, and libraries are loaded each time — an attacker cannot hard-code target addresses into their exploit. Bounds/parameter checking prevents the overflow from occurring; ASLR makes exploitation harder if overflow does occur.

*Subtopic: 8.5 — [Memory Safety](../concepts/memory-safety.md)*

---

### Q12: An organization is using an open source library in their production application. A critical vulnerability (CVE) is discovered in that library. Which OWASP Top 10 (2021) category does this fall under, and what organizational control would have given the earliest warning of the exposure?

**Answer:** A06 — Vulnerable and Outdated Components. A Software Bill of Materials (SBOM) maintained for all applications would provide immediate visibility into which applications use the affected library, enabling rapid assessment and remediation.

**Why:** Without an SBOM or dependency inventory, organizations must manually search all codebases to determine exposure — an error-prone, time-consuming process. Log4Shell demonstrated this at massive scale; organizations with SBOMs could identify and prioritize affected systems in hours rather than days. Software Composition Analysis (SCA) tools integrated into CI/CD pipelines provide continuous monitoring of dependencies against known CVEs.

*Subtopic: 8.4 / 8.5 — [Software Acquisition Security](../concepts/software-acquisition-security.md); [OWASP Top 10](../concepts/owasp-top-10.md); [CI/CD Security](../concepts/ci-cd-security.md)*
