---
title: "Risk Management"
type: concept
domain: 1
tags: [risk, threat, vulnerability, risk-treatment, quantitative, qualitative, ale, sle, aro, risk-analysis]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Risk Management

Risk management is the **identification, assessment, and prioritization of risks** and the
cost-efficient application of resources to minimize the probability and/or impact of those risks.

## Core Definitions (Table 1-20)

| Term | Definition |
|---|---|
| **Threat Agent** | Entity with potential to cause damage (external attackers, disgruntled employees) |
| **Threat** | Any potential danger to an asset |
| **Vulnerability** | A weakness that could be exploited |
| **Attack** | A harmful action exploiting a vulnerability |
| **Risk** | Significant exposure where a threat can exploit a vulnerability |
| **Asset** | Anything valued by the organization |
| **Exposure / Impact** | Negative consequences if a risk materializes |
| **Countermeasures / Safeguards** | Controls that reduce threats, vulnerabilities, and impact |
| **Residual Risk** | Risk remaining after controls are implemented |

## Risk Management Process (Table 1-17)

### Step 1: Value (Asset Valuation)
Identify and rank assets from most to least valuable. Two methods:
- **Qualitative**: Relative ranking ("Low/Medium/High", 1-5 scale). No monetary values. Faster, simpler.
- **Quantitative**: Assigns objective monetary values. Preferred but difficult to perform accurately.

### Step 2: Risk Analysis
For each asset, identify: **Threats + Vulnerabilities + Impact + Probability/Likelihood**.

**Risk types (examples):**
| Risk Type | Threat Example | Vulnerability Example |
|---|---|---|
| Natural/Environmental | Flood | Building on a floodplain |
| Human | Hacker | Untrained employees susceptible to SE |
| Operational/Process | Check fraud | No segregation of duties |
| Technical | Malware | Unpatched software |
| Physical | Power outage | No backup power |

### Step 3: Treatment
Four risk treatment options (see [Risk Treatment section below](#risk-treatment-options)):
- **Avoid** — stop doing the risky activity
- **Transfer** — purchase cyber insurance
- **Mitigate** — implement controls
- **Accept** — the asset owner accepts the risk

## ALE Calculation (Quantitative Risk — Table 1-20)

**Formula: ALE = SLE × ARO**  
**SLE = AV × EF**

| Component | Definition | Example (CCTV) |
|---|---|---|
| **AV** — Asset Value | Monetary cost of the asset | $2,000 |
| **EF** — Exposure Factor | % of asset value lost per incident (0–100%) | 10% (3 of 30 cameras = $200) |
| **SLE** — Single Loss Expectancy | Cost if the risk occurs once: AV × EF | $2,000 × 10% = $200 |
| **ARO** — Annualized Rate of Occurrence | Expected number of times per year | 3 |
| **ALE** — Annualized Loss Expectancy | Annual cost of the risk: SLE × ARO | $200 × 3 = **$600** |

**Key rule**: Do not implement a control that costs more than the ALE. If a control costs $5,000 to address a $600/year risk, the best business decision is to **accept** the risk.

## Risk Treatment Options

- **Avoid**: Stop the activity that creates the risk. Incurs **opportunity cost** (you also lose the potential upside).
- **Transfer**: Share risk with an insurer (cyber insurance). The organization retains **accountability** even after transfer — only responsibility is transferred.
- **Mitigate**: Implement controls to reduce risk to an acceptable level. Risk is never eliminated to zero; **residual risk** always remains.
- **Accept**: Owner decides the cost of control exceeds the benefit. Must be made by the **asset owner or senior management** — not security staff.
- **Ignore**: NOT a valid approach. Constitutes failure of due care and due diligence.

## Qualitative vs. Quantitative Analysis (Table 1-18)

| Qualitative | Quantitative |
|---|---|
| No monetary values | Assigns objective monetary values |
| Relative ranking (Low/Med/High) | Fully numerical (ALE, SLE, ARO) |
| Fast, simple | Difficult, time-consuming |
| Based on professional judgment | Preferred when accuracy is needed |

## Continuous Improvement — Deming Cycle (PDCA)

| Step | Action |
|---|---|
| **Plan** | Determine which controls to implement based on identified risks |
| **Do** | Implement the controls |
| **Check** | Monitor/assess whether controls are operating effectively |
| **Act** | Take corrective action based on Check findings; feed back to Plan |

Risk analysis frequency: "as often as necessary" — triggered by new assets, new threats, new vulnerabilities, regulatory changes, or changes in asset value.

## Cross-links

- [Security Controls Types](security-controls-types.md)
- [Threat Modeling](threat-modeling.md)
- [NIST RMF](../standards/nist-rmf.md)
- [Governance](governance.md)
- [SCRM](scrm.md)

## Sources

- destination-cissp §1.9.1–§1.9.10 (Tables 1-17 through 1-26, Figures 1-4 through 1-11)
