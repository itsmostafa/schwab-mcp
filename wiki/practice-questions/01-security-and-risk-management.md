---
title: "Practice Questions — Domain 1: Security and Risk Management"
type: practice
domain: 1
tags: [governance, compliance, risk, ethics, law, policy, practice]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Practice Questions — Domain 1: Security and Risk Management

## Questions

### Q1: A server valued at $50,000 has a 20% exposure factor and is expected to experience a hardware failure twice per year. What is the ALE?

**Answer:** $20,000

**Why:** SLE = AV × EF = $50,000 × 20% = $10,000. ALE = SLE × ARO = $10,000 × 2 = $20,000. If a proposed control costs more than $20,000, the best business decision is to accept the risk instead. This is a direct application of the quantitative risk formula.

*Subtopic: 1.9 — Risk Management Concepts*

---

### Q2: A CISSP discovers their employer is committing financial fraud that harms thousands of customers. According to the ISC2 Code of Ethics, what should the CISSP do?

**Answer:** Report the fraud to appropriate authorities, even if it means acting against the employer's interests.

**Why:** Canon 1 (Protect society, the common good, public trust, and the infrastructure) takes precedence over Canon 3 (Provide diligent service to principals). When canons conflict, they are applied in numbered order. The employer is the "principal" in Canon 3, but society's welfare in Canon 1 overrides that obligation.

*Subtopic: 1.1 — Professional Ethics*

---

### Q3: An organization performs a penetration test and fixes all identified vulnerabilities. The CISO provides a report to the Board documenting the remediation. Which concept does the CISO's report to the Board represent?

**Answer:** Due diligence

**Why:** Due care = performing the penetration test and fixing the vulnerabilities (doing the right thing). Due diligence = providing the Board with proof that due care was exercised. The report is the evidence of due care, which is the definition of due diligence. Exam trap: both concepts are needed, but the act of reporting/proving is always due diligence.

*Subtopic: 1.3 — Security Governance Principles*

---

### Q4: A cloud service provider (CSP) experiences a data breach affecting customer data stored under their service. Who is ultimately accountable for the breach?

**Answer:** The customer organization (data owner), not the CSP.

**Why:** Accountability can never be outsourced or transferred. Even though the CSP is responsible for protecting the data under the SLA, the data owner remains accountable to regulators, customers, and shareholders. The CSP may be contractually liable for breach of the SLA, but ultimate accountability lies with the owner of the data.

*Subtopic: 1.3 — Organizational Roles and Responsibilities*

---

### Q5: Which risk treatment option should an organization choose when the cost of implementing a security control exceeds the Annualized Loss Expectancy of the risk?

**Answer:** Accept the risk.

**Why:** Risk acceptance is appropriate when the cost of control exceeds the benefit (i.e., the ALE). This is a sound business decision. The decision to accept risk must always be made by the **asset owner or senior management** — not the security team. Risk ignorance (not acknowledging the risk at all) is not a valid option and violates due care.

*Subtopic: 1.9 — Risk Response/Treatment*

---

### Q6: A document specifies exactly which version of anti-malware software all corporate laptops must run. What type of security document is this?

**Answer:** Standard

**Why:** A standard specifies specific hardware or software solutions, mechanisms, or products — in this case, a specific AV product and version. A policy would state the requirement (e.g., "all laptops must have anti-malware software"). A procedure would give step-by-step installation instructions. A baseline would specify minimum configuration settings. A guideline would suggest (but not mandate) best practices.

*Subtopic: 1.6 — Security Policies, Standards, Procedures, Baselines, and Guidelines*

---

### Q7: An attacker successfully follows an employee through a secured door without using any badge or credential. What type of attack is this?

**Answer:** Piggybacking

**Why:** Piggybacking = following an authorized person through a restricted door without any badge. Tailgating = the same physical act, but the attacker possesses a fake (but realistic-looking) badge. The distinction is the presence/absence of a fraudulent credential. Both are physical social engineering attacks best mitigated by employee awareness training and physical controls (mantraps, turnstiles).

*Subtopic: 1.10 — Threat Modeling / Social Engineering*

---

### Q8: In the STRIDE threat model, which threat category maps to a violation of Availability?

**Answer:** Denial of Service (D in STRIDE)

**Why:** STRIDE maps to security violations: Spoofing → Authentication, Tampering → Integrity, Repudiation → Nonrepudiation, Information Disclosure → Confidentiality, **Denial of Service → Availability**, Elevation of Privilege → Authorization. Knowing this mapping is essential because exam questions may describe a threat and ask for its STRIDE category or the CIA pillar violated.

*Subtopic: 1.10 — Threat Modeling Concepts*

---

### Q9: A security manager identifies a risk but is told by their manager to "just ignore it for now." Why is this problematic?

**Answer:** Risk ignorance violates due care and due diligence, and is not a valid risk treatment option.

**Why:** The four valid risk treatment options are Avoid, Transfer, Mitigate, and Accept. Ignoring a risk (failing to make any documented decision about it) is not a valid option. It exposes the organization to liability, violates due care (not protecting assets appropriately) and due diligence (not demonstrating responsible governance). The asset owner or senior management must formally acknowledge and decide on the risk.

*Subtopic: 1.9 — Risk Response/Treatment*

---

### Q10: What is the PRIMARY distinction between a Business Continuity Plan (BCP) and a Disaster Recovery Plan (DRP)?

**Answer:** BCP focuses on keeping critical business functions running during a disruption; DRP focuses on restoring IT systems after a disruption.

**Why:** BCP is broader and organization-wide — it addresses how the business continues operating during a crisis. DRP is a subset focused specifically on IT/technical recovery after the fact. Both are informed by the Business Impact Analysis (BIA). On the exam, if the scenario involves keeping people working during an outage → BCP; if it's about restoring servers after an event → DRP.

*Subtopic: 1.7 — Business Continuity Requirements*

---

### Q11: PASTA threat modeling is described as "attacker-focused and risk-centric." What does this mean in practice compared to STRIDE?

**Answer:** PASTA analyzes threats from the attacker's perspective across 7 stages, incorporating business risk context, while STRIDE categorizes threats by the type of violation (spoofing, tampering, etc.) without the full strategic/business context.

**Why:** STRIDE is a checklist-style methodology that categorizes threats against CIA pillars — useful, but tactical. PASTA is a 7-stage methodology that starts with business objectives (Stage 1) and risk context before drilling into technical threat analysis, making it more thorough and strategic. When the exam asks which is "more comprehensive" or "more strategic," the answer is PASTA.

*Subtopic: 1.10 — Threat Modeling Concepts*

---

### Q12: An organization is implementing a firewall. According to the "complete control" concept, what additional controls should also be implemented?

**Answer:** At minimum, a detective control (e.g., IDS/SIEM alerts) and a corrective control (e.g., incident response procedure to isolate compromised systems).

**Why:** A complete control = preventive + detective + corrective controls at minimum. A firewall is preventive — it prevents unauthorized traffic. But no preventive control is perfect, so a detective control is needed to identify when something gets through, and a corrective control minimizes damage when it does. This complete control concept should be applied at every layer of a defense-in-depth architecture.

*Subtopic: 1.9 — Applicable Types of Controls*

