---
title: Physical Security
type: concept
domain: 3
tags: [physical-security, CPTED, defense-in-depth, CCTV, mantrap, bollards, layered-defense]
sources: [destination-cissp]
updated: 2026-05-13
---

# Physical Security

## Definition

Physical security protects an organization from the outside perimeter inward, extending the CIA triad (confidentiality, integrity, availability) to the physical world. It encompasses controls that protect people, property, and equipment from physical threats.

## Primary Goal

> **Safety and protection of human life is the #1 priority of physical security.**

Every physical security decision must prioritize human safety. This principle drives many counterintuitive design choices (e.g., doors that hinge outward for egress even though it's less secure).

## Physical Security Control Types

| Type | Purpose | Examples |
|---|---|---|
| **Deter / Prevent** | Discourage or block intrusion attempts | Fences, warning signs, lighting |
| **Delay** | Slow down an attacker | Locks (cannot prevent; only delay) |
| **Detect** | Identify and alert to an intrusion | CCTV, PIR sensors, alarms, guard dogs |
| **Assess** | Evaluate an alert to determine appropriate response | Security personnel review |
| **Respond** | Take corrective action | Guards, law enforcement, lockdown |

## Layered Defense Model (Defense-in-Depth)

Physical security uses the same defense-in-depth principle as logical security:

```
Outer perimeter (fencing, lighting, landscaping)
    ↓
Building perimeter (walls, doors, CCTV)
    ↓
Interior access points (mantraps, card readers)
    ↓
High-value areas (server rooms, wiring closets, evidence storage)
```

**Optimal door count**: zero (most secure) but building egress safety requires multiple exits. Target: as few as safely possible.

## CPTED (Crime Prevention Through Environmental Design)

A professional discipline providing guidelines for designing buildings and surrounding environments to naturally deter criminal activity. Uses the environment itself as a security control:
- Landscaping that directs pedestrian flow and eliminates hiding spots
- Lighting that deters nighttime threats and improves camera effectiveness
- Grading (ground slope) to prevent flooding of critical facilities
- Building placement and sight-lines that enable natural surveillance

## Key Physical Security Controls

### Perimeter Controls
- **Fencing**: prevents casual entry; force multiplier when combined with lighting/CCTV.
- **Bollards**: concrete/steel barriers preventing vehicle attacks. Often disguised as decorative planters. Required in front of government buildings following the Oklahoma City bombing.
- **Lighting**: deters attacks, improves CCTV effectiveness, reduces assault incidents in parking areas.
- **Grading**: slope ground away from critical facilities to prevent flood damage.
- **Landscaping**: trees and bushes can block camera sightlines if poorly placed; well-designed landscaping directs access and eliminates cover.

### CCTV (Closed-Circuit Television)
- **Primary role**: detective control.
- Also functions as: deterrent, forensic evidence source, security audit tool.
- Design considerations: placement (cover entrances/exits, capture faces), image quality (day and night), transmission media, retention policy (local laws may govern).

### Passive Infrared (PIR) Devices
- Motion detectors that sense infrared light (heat) differences.
- Take rapid successive images and compare; human body (warmer than room) triggers alert.
- Must constantly recalibrate to ambient temperature (Texas heat can exceed body temperature).
- Very sensitive to temperature changes.

### Doors
- **Composition**: door material must be resistant to forced entry.
- **Frame**: most secure door is useless on a weak frame (wood frames can be kicked in).
- **Hinge placement**: hinges on the outside = security risk (pins can be knocked out). However, outward-opening doors (hinges outside) protect people during emergency egress.
- **Mantraps**: double-door airlock or turnstile preventing tailgating/piggybacking. Usually requires two-factor authentication (card + biometric). Unauthorized person trapped in vestibule if detected.

### Locks (Delay Controls Only)
All locks are **delay controls** — given enough time, any lock can be defeated.

| Type | Category | Notes |
|---|---|---|
| Key | Mechanical | Most common; many variants |
| Combination | Mechanical | Security based on combination complexity |
| Magnetic (Maglock) | Mechanical/Electronic | Electromagnet + metal plate; releases on power failure |
| Proximity / RFID | Electronic | Card-based; convenient but transferable |
| Keypad | Electronic | Susceptible to shoulder-surfing; should have shields |
| Biometric | Electronic | Most accurate; privacy concerns; most expensive |

### Windows
- **Major physical vulnerability** in most structures.
- **Shock sensors**: installed on each pane; detect vibration from breaking glass; effective in **noisy** environments.
- **Glass break sensors**: microphone-based; listen for glass-breaking sound; one sensor can cover a whole room; effective in **quiet** environments.

### Walls
- Composition determines breach resistance (wood/steel stud walls are easily defeated).
- Must extend from **true floor to true ceiling** — not to a drop ceiling.
- In data centers with raised floors, walls must span from actual floor to actual ceiling.

## High-Value Areas Requiring Special Protection

| Area | Description |
|---|---|
| Wiring closets | Network cables and equipment (typically one per floor) |
| Media storage | Physical/digital media containing sensitive data |
| Evidence storage | Chain of custody areas |
| Server rooms | Core infrastructure; highest value |
| Restricted work areas | Based on work type, personnel clearance, or data sensitivity |

## Threats to Physical Security

Theft, espionage (targeting IP/trade secrets), dumpster diving (improperly disposed materials), social engineering, shoulder surfing, HVAC compromise (for access or equipment damage).

## Exam Traps

- Primary goal of physical security = **safety and protection of human life**, not data protection.
- Locks are **delay** controls — they do not prevent access.
- CCTV is **primarily a detective** control (with secondary deterrent function).
- Mantraps prevent **tailgating / piggybacking**.
- Shock sensors work best in **noisy** environments; glass break sensors work best in **quiet** environments.
- PIR sensors must recalibrate because ambient temperature can exceed body temperature.
- Fewer access points = better security, but human safety requires sufficient egress.

## Cross-References

- [Facility Security](facility-security.md) — data center-specific controls (power, HVAC, fire)
- Domain 3 §3.8–3.9

## Sources

- destination-cissp §3.8.1–3.9.10 (pages 0461–0476)
