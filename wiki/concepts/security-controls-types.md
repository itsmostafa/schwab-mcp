---
title: "Security Control Types and Categories"
type: concept
domain: 1
tags: [controls, preventive, detective, corrective, directive, deterrent, recovery, compensating, administrative, technical, physical, safeguards, countermeasures]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Security Control Types and Categories

Security controls are organized by **type** (what they do) and **category** (how they are implemented).

## Seven Types of Controls (Table 1-21)

| Type | Description | Example |
|---|---|---|
| **Directive** | Directs/confines actions to encourage compliance | Fire exit sign |
| **Deterrent** | Discourages policy violations | "Trespassers will be prosecuted" sign |
| **Preventive** | Prevents undesired actions/events | Fence, locked door, firewall rule |
| **Detective** | Identifies that a risk has occurred (**after** the event) | Smoke alarm, IDS alert, CCTV |
| **Corrective** | Minimizes damage after a risk occurs | Fire suppression system activating |
| **Recovery** | Returns system to normal operation after an incident | Data backup/restore process |
| **Compensating** | Supports/supplements other controls; can replace a control | HIPS deployed in addition to NIPS |

### Timing of controls
- **Before an incident**: Directive, Deterrent, Preventive, Compensating (safeguards).
- **After an incident**: Detective, Corrective, Recovery (countermeasures).

## Complete Control Concept

A **complete control** = at minimum, **preventive + detective + corrective** controls together.

Rationale: No preventive control is perfect. Therefore, always pair a preventive control with detective
controls (to catch what gets through) and corrective controls (to minimize damage when it does).
This should be applied at every layer of defense (defense-in-depth).

## Safeguards vs. Countermeasures

| Safeguards | Countermeasures |
|---|---|
| **Proactive** — before risk occurs | **Reactive** — after risk occurs |
| Directive, Deterrent, Preventive, Compensating | Detective, Corrective, Recovery |

## Three Categories of Controls (Table 1-22)

| Category | Examples |
|---|---|
| **Administrative** | Policies, procedures, baselines, guidelines, background checks, acceptable use policies, onboarding/offboarding, job rotation |
| **Logical/Technical** | Firewalls, IDS/IPS, anti-malware, proxies, login mechanisms, OS restrictions |
| **Physical** | Doors, fences, gates, bollards, mantraps, guards, CCTV, RF ID badges |

Note: Logical vs. Technical is a subtle distinction — logical = software component; technical = hardware component. Usually used interchangeably on the exam.

## Functional and Assurance Aspects (Table 1-23)

Every properly designed control should have:
- **Functional**: The control does what it was designed to do (e.g., firewall filters traffic).
- **Assurance**: The control can be proven to be working on an ongoing basis (e.g., via testing, logging, monitoring, assessments).

## Control Selection Criteria

Controls must be:
1. Aligned to organizational goals and objectives.
2. Cost-effective (control cost must not exceed the ALE of the risk it mitigates).
3. Structured as a complete control (preventive + detective + corrective).
4. Both functional and providing assurance.

## Exam-Relevant Nuance

- Exam trap: "CCTV" — it is both **detective** (records events) and can be **deterrent** (deters would-be attackers). Context determines the answer.
- "Compensating" controls can **replace** a primary control (e.g., camera monitoring instead of a lock in a certain zone) or **supplement** it.
- Always remember: detective, corrective, recovery controls trigger **after** a risk event. If the question describes actions before harm — it's preventive/deterrent/directive.

## Cross-links

- [Risk Management](risk-management.md)
- [Due Care and Due Diligence](due-care-due-diligence.md)
- [Security Policies](security-policies.md)

## Sources

- destination-cissp §1.9.6, §1.9.7, §1.9.8, §1.9.9 (Tables 1-21 through 1-23)
