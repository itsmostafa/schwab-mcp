---
title: Facility Security
type: concept
domain: 3
tags: [facility-security, data-center, HVAC, fire-suppression, power, UPS, VESDA, halon]
sources: [destination-cissp]
updated: 2026-05-13
---

# Facility Security

## Definition

Facility security addresses the physical infrastructure controls needed to ensure availability and protect equipment in buildings and data centers. It encompasses power management, HVAC (temperature/humidity/air quality), and fire prevention, detection, and suppression. Sometimes summarized as: **power, ping, and pipe** (power, internet, and cooling water).

## Power

### Goal
Provide **clean, steady, uninterrupted** electrical power to critical systems.

### Power Protection Systems

| System | Function | Duration |
|---|---|---|
| **UPS (Uninterruptible Power Supply)** | Battery backup + power conditioning (removes sags/dips) | Short-term (minutes) |
| **Generator** | Diesel-powered alternator; long-term power | Long-term (hours/days, fuel-dependent) |

Typical workflow: power fails → UPS provides immediate bridge (seconds to minutes) → generator spins up (~1 minute) → generator takes over.

### Types of Power Disruption

| | Short Duration (milliseconds) | Long Duration (seconds–hours–days) |
|---|---|---|
| **No power** | Fault | Blackout |
| **Insufficient power (low voltage)** | Sag / Dip | Brownout |
| **Excessive power (high voltage)** | Spike | Surge |

## HVAC — Temperature, Humidity, and Air Quality

### ASHRAE TC 9.9 Guidelines for Data Centers

| Parameter | Minimum | Maximum |
|---|---|---|
| Temperature | 64.4°F / 18°C | 80.6°F / 27°C |
| Humidity | 40% RH | 60% RH |

### Why These Matter

- **Too hot**: equipment overheats, failures.
- **Too cold**: condensation risk, equipment stress.
- **Low humidity (<40%)**: static electricity → electrical shorts.
- **High humidity (>60%)**: condensation → corrosion, shorts.
- **Poor air quality**: dust accumulates in servers → overheating, failure.

### Positive Pressurization

Air is pumped into server rooms at **slightly above ambient pressure**. When a door opens, clean air flows *out* rather than dirty air flowing *in*. Cracks, gaps, or open windows allow clean air to escape rather than admitting contaminants. Keeps data center air clean.

### Hot/Cold Aisle Containment

Data center racks arranged so cool air enters from the front (cold aisle) and hot exhaust exits at the rear (hot aisle). Prevents hot and cold air mixing, improving cooling efficiency.

## Fire

### The Fire Triangle

All three elements are required to sustain a fire:
- **Fuel** — combustible material
- **Oxygen** — air supply
- **Heat** — ignition and sustaining temperature

Remove any one → fire goes out. This drives fire suppression strategy choices.

### Fire Prevention

Keep combustible materials away from valuable assets. In data centers: unbox equipment on the loading dock; bring only bare metal into the room.

### Fire Detection Systems

| Type | Method | Best For | Limitations |
|---|---|---|---|
| **Flame detector** | UV/infrared camera detects flame light | — | Late-stage detection; fire is advanced by the time flames are visible |
| **Smoke — Ionization** | Radioactive material ionizes particles; smoke disrupts | Fast/flaming fires | Less effective for smoldering fires |
| **Smoke — Photoelectric / Optical** | Off-angle light refracted by smoke particles into sensor | Smoldering fires | Less effective for fast/flaming fires |
| **Smoke — Dual** | Both ionization + photoelectric | Both fire types | Most common modern design |
| **VESDA** (Very Early Smoke Detection Apparatus) | Samples air continuously; detects incipient/smoldering stage | Data centers, high-value areas | Most expensive; best early warning |
| **Heat (rate-of-rise)** | Temperature sensor detects rapid temperature increase | — | Late-stage detection; fire is well-advanced |

**Best early detection**: VESDA (detects fire at the **incipient** stage before visible smoke).

### Water-Based Fire Suppression Systems

Water removes **heat** from the fire triangle. Cheaper than gas systems but can destroy equipment.

| System | How It Works | Key Characteristics |
|---|---|---|
| **Wet Pipe** | Pipes always filled with pressurized water | Simplest; always ready; risk of leaks; pipes burst if frozen |
| **Dry Pipe** | Pipes filled with pressurized gas (air/nitrogen); water added on activation | Better for freezing environments; combined with pre-action |
| **Pre-action** | Requires fire detection signal before water releases into pipes; only activated zones release water | Reduces false-activation water damage; best water-based option for data centers |
| **Deluge** | All sprinkler heads open; massive water release | Only where immediate total suppression needed (explosives, fireworks factories) |

### Gas-Based Fire Suppression Systems

Gas removes **oxygen** or **interrupts the chemical process** of combustion. Preferred for data centers (no water damage). More expensive than water systems. **Warning**: gas can be lethal — systems sound an alarm and provide evacuation time before discharge.

| Gas | Properties | Notes |
|---|---|---|
| **INERGEN** | Reduces oxygen concentration | Replaces Halon; still allows breathing during discharge |
| **Argonite** | Argon + nitrogen mixture | Safe for people and environment; Halon replacement |
| **FM-200** | Clean agent; removes heat or interrupts combustion | No residue; safe for equipment and people; Halon replacement |
| **Aero-K** | Ultrafine potassium aerosol | Suppresses fire fast; prevents reignition; no residue; environmentally friendly |
| **~~Halon~~** | ~~Interrupts chemical process~~ | **ILLEGAL** — destroys ozone layer; must not be used |
| **CO2** | Removes oxygen | Effective and non-corrosive; **lethal to people in enclosed spaces** |

### Fire Extinguisher Classes

| Class | Fuel Type | Suppression Agent |
|---|---|---|
| **A** | Common combustibles (wood, paper) | Water, foam, dry chemicals |
| **B** | Liquids (gasoline, oil) | CO2, foam, gas, dry chemicals |
| **C** | Electrical | CO2, gas, dry chemicals (NOT water) |
| **D** | Combustible metals | Dry powder |
| **K** | Commercial kitchen (grease, oil) | Wet chemicals |

## Exam Traps

- **VESDA** = best / most expensive / detects at incipient (smoldering) stage.
- **Halon is ILLEGAL** — always wrong answer for new fire suppression.
- **CO2** is effective but dangerous to people in confined spaces.
- **Pre-action** is the best water-based system for data centers (detection required before release).
- **Deluge** is only for explosives/fireworks — never for data centers.
- Low humidity → static electricity. High humidity → condensation/corrosion.
- **Positive pressurization**: clean air flows *out* when door opens, not dirty air *in*.
- UPS = short-term (minutes); Generator = long-term (hours).
- Fault/sag/spike = short duration; Blackout/brownout/surge = long duration.

## Cross-References

- [Physical Security](physical-security.md) — perimeter and access controls
- Domain 3 §3.9.12–3.9.14

## Sources

- destination-cissp §3.9.12–3.9.14 (pages 0476–0485)
