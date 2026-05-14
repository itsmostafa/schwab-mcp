---
title: "Biometrics"
type: concept
domain: 5
tags: [biometrics, cer, far, frr, fingerprint, retina, iris, authentication]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Biometrics

## Definition

Biometric authentication ("Authentication by Characteristic") verifies identity using physical or behavioral attributes. Unlike passwords, biometric data cannot be changed if compromised.

## Types of Biometric Attributes

| Category | Types |
|---|---|
| **Physiological** | Fingerprint, hand geometry, vascular pattern, facial features, iris, retina |
| **Behavioral** | Voice (how a person speaks), signature/handwriting, keystroke dynamics, gait |

### Notable individual types

| Type | Notes |
|---|---|
| **Fingerprint** | Most common; used on devices, border crossings; reasonably accurate |
| **Hand geometry** | Rare in practice; often depicted in movies |
| **Vascular pattern** | Scans veins in the hand; used at testing centers (including CISSP exam venues) |
| **Iris** | Scans the colored ring around the eye; high accuracy; non-invasive |
| **Retina** | Scans vein patterns at the *back* of the eye; **most accurate** biometric type; invasive (user must press eye to rubber cup, bright light flashed); controversial because scan results can reveal medical conditions |
| **Facial** | Scans facial features and pattern |
| **Voice** | Behavioral — how a person speaks |
| **Keystroke dynamics** | Behavioral — typing rhythm and patterns |
| **Gait** | Behavioral — walking pattern |

> **Exam point:** Retina = most accurate. Iris = non-contact, more user-friendly. Both scan the eye, but *different parts*.

## Biometric Error Rates (EXAM CRITICAL)

Biometric systems are not binary — they operate on probabilities, not exact matches.

| Error Type | Full Name | Abbrev. | What it means | Severity |
|---|---|---|---|---|
| **Type 1** | False Rejection Rate | **FRR** | A valid/legitimate user is *rejected* by the system | Lower severity — just frustrates users |
| **Type 2** | False Acceptance Rate | **FAR** | An invalid/unauthorized user is *accepted* | **Higher severity** — security breach risk |

> **Source note (§5.2.5, p. 686-687):** The source text contains a typo labeling FAR as "FRR" in the table header. The narrative is correct: Type 2 = False Acceptance = FAR.

### Inverse relationship

FRR and FAR are **inversely related** — tuning the system to reduce one increases the other:
- Tighten sensitivity → fewer false accepts (lower FAR), more false rejects (higher FRR).
- Loosen sensitivity → fewer false rejects (lower FRR), more false accepts (higher FAR).

### Crossover Error Rate (CER)

**CER = the point where FRR and FAR are equal.** It is the intersection of the two error curves.

- CER is the standard metric for comparing biometric system accuracy.
- **Lower CER = more accurate system.**
- No system achieves CER = 0.

```
Error
Rate
  |  FRR
  |    \     FAR
  |     \   /
  |      \ /
  |       X  ← CER
  |      / \
  +----------→ Sensitivity
```

## Biometric Template Storage

Raw biometric data (e.g., a raw fingerprint image) must **never be stored** directly. Instead, a one-way mathematical function produces a **template** — a digital representation of unique features.

| Template use | Mode | Description |
|---|---|---|
| **1:N (one-to-many)** | Identification | System compares new template against the full database to *find* who the person is |
| **1:1 (one-to-one)** | Authentication | System compares new template against the *single stored template* for the claimed user identity |

## Implementation Considerations

Before deploying biometrics, consider:
1. **Processing speed** — biometric systems can be slower than password-based systems.
2. **User acceptance** — invasive or uncomfortable systems (e.g., retina scanners) face resistance.
3. **Protection of biometric data** — exposure is irreversible; users cannot grow new fingers.
4. **Accuracy** — CER determines fitness for the security context.

## Cross-links

- [Authentication Factors](authentication-factors.md)
- [AAA](aaa.md)

## Sources

- destination-cissp §5.2.5, §5.2.6 (pp. 684-692)
- cissp-exam-outline §5.2
