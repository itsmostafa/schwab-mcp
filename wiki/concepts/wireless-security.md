---
title: "Wireless Security"
type: concept
domain: 4
tags: [wireless, wep, wpa, wpa2, wpa3, eap, 802.11, 802.1x, tkip, ccmp]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Wireless Security

Wireless networks carry the same data as wired networks but over radio frequencies accessible to anyone within range. Security is not native to 802.11 — it must be explicitly added. Four services are required: access control, authentication, encryption, and integrity protection.

## Key Facts

### Wireless Security Protocol Evolution

| Standard | Released | Access Control | Authentication | Encryption | Integrity | Notes |
|----------|---------|---------------|---------------|-----------|---------|-------|
| WEP | 1997 | 802.1X | EAP methods | RC4 (weak IV) | None | **Broken — do not use** |
| WPA | 2003 | 802.1X or PSK | EAP or PSK | TKIP (RC4, 128-bit per-packet keys) | Michael MIC | Stopgap fix for WEP |
| WPA2 | 2004 | 802.1X or PSK | EAP or PSK | **CCMP (AES)** | CCMP | Current standard |
| WPA3 | 2018 | 802.1X or PSK | EAP or PSK | CCMP or GCMP | CCMP or GCMP | Latest; not yet widely deployed |

**WEP weakness:** Used a weak Initialization Vector (IV). The IV was short and reused, allowing attackers to statistically recover the WEP key. TKIP was created as a temporary fix.

**TKIP:** Designed to replace WEP without requiring hardware replacement. Uses RC4 with 128-bit per-packet key mixing and Michael MIC for integrity. **No longer considered secure — superseded by AES/CCMP.**

**CCMP (Counter-Mode CBC-MAC Protocol):** Uses AES with 128-bit keys. Used in WPA2 and WPA3. The gold standard for wireless encryption.

**GCMP (Galois Counter Mode Protocol):** Used in WPA3 for even stronger encryption.

### EAP Variants

| Type | Client Auth | Server Auth | Security | Notes |
|------|------------|------------|---------|-------|
| EAP-TLS | Certificate | Certificate | **High** | Strongest; mutual auth; no proprietary |
| EAP-TTLS | ID + Password | Certificate | Medium | Funk/Certicom; proprietary |
| PEAP (EAP-PEAP) | ID + Password | Certificate | Medium-High | Cisco/RSA/Microsoft; wraps EAP in TLS tunnel |
| LEAP | ID + Password | ID + Password | **Low** | Cisco proprietary; **deprecated** |
| EAP-MD5 | ID + Password | None | **Low** | No server auth; not recommended |

**PEAP vs EAP:** PEAP encapsulates EAP within an encrypted and authenticated TLS tunnel. PEAP is more secure than plain EAP because the authentication exchange is protected.

### Wireless Authentication Methods

1. **Open authentication:** Device connects with SSID, no security. Never use for corporate.
2. **Shared key (PSK):** Pre-shared key distributed to all devices. Simple but poor key management.
3. **EAP-based (802.1X):** Separate authentication server validates each user. Supports one- or two-factor authentication.

**Mutual authentication** is ideal: both client and AP verify each other's identity.

### Common Wireless Attacks

- **Evil twin / Rogue AP:** Attacker sets up an AP with the same SSID as a legitimate network. Users connect and their traffic is intercepted.
- **Wardriving:** Driving around scanning for open or poorly secured wireless networks.
- **WEP cracking:** Statistical attack exploiting weak IVs. WEP can be cracked in minutes.
- **KRACK (WPA2):** Key Reinstallation Attack; patched in most implementations.

### Radio Frequency Management

Wireless signals extend beyond walls. AP placement should minimize signal bleed into publicly accessible areas (parking lots). Use directional antennas and reduce transmit power where possible. Segregate guests, employees, and vendors onto separate SSIDs/VLANs.

## Exam Nuance

- Know the evolution: **WEP (broken) → WPA (TKIP) → WPA2 (CCMP/AES) → WPA3 (GCMP)**.
- TKIP uses RC4 with key mixing as a stopgap; it is **not considered secure**.
- EAP-TLS requires **both client and server certificates** — strongest but most complex to deploy.
- LEAP is **deprecated** and Cisco-proprietary — wrong answer on any "most secure" question.
- WPA3 is the answer to "newest/most current" but is flagged as not yet widely deployed.
- 802.1X and EAP are used together; 802.1X is the port-based access control framework; EAP is the authentication protocol within it.

## Cross-links

- [802.11 Standard](../standards/802-11.md)
- [802.1X Standard](../standards/802-1x.md)
- [Network Access Control](./network-access-control.md)
- [Network Attacks](./network-attacks.md)

## Sources

- destination-cissp §4.1.14 (pp. 0582–0593)
