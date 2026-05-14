---
title: "VPN"
type: concept
domain: 4
tags: [vpn, ipsec, tls, tunneling, ike, split-tunneling, pptp, l2tp]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# VPN

A VPN (Virtual Private Network) is **tunneling + encryption**. Tunneling alone (encapsulating one packet inside another) is not a VPN — encryption of the payload is required to make it a VPN.

## Key Facts

### Tunneling Protocols by OSI Layer

| Protocol | OSI Layer | Encryption | Notes |
|----------|-----------|-----------|-------|
| SSH | 7 — Application | Yes | Secures Telnet, FTP, etc. inside encrypted tunnel |
| SOCKS | 5 — Session | Yes | Socket Secure |
| SSL/TLS | 4 — Transport | Yes | Encrypts by default; TLS VPN |
| IPsec | 3 — Network | Yes (via ESP) | Preferred for site-to-site VPN; embedded in IPv6 |
| GRE | Multiple | **No** | Encapsulation only — not secure by itself |
| L2TP | 2 — Data Link | **No** | Requires IPsec for encryption; replaces PPTP |
| PPTP | 2 — Data Link | Limited | Deprecated; uses PAP/CHAP/EAP for auth |
| L2F | 2 — Data Link | No | Layer 2 Forwarding; largely replaced by L2TP |

**VPN = Tunnel + Encryption.** GRE and L2TP alone are just tunnels.

### IPsec (preferred VPN protocol)

IPsec is the preferred method for establishing a VPN. It is natively embedded in IPv6.

**IPsec sub-protocols:**

| Sub-protocol | Provides |
|---|---|
| **AH (Authentication Header)** | Integrity, data-origin authentication, replay protection. **No encryption.** |
| **ESP (Encapsulating Security Payload)** | Everything AH provides **plus payload encryption (confidentiality)**. |

**IPsec Modes:**

| Mode | Original IP Header | What is protected |
|------|-------------------|-------------------|
| **Transport** | Kept (original header used) | Payload only |
| **Tunnel** | Wrapped inside new outer IP header | Original header + payload (full encapsulation) |

Tunnel mode is used for site-to-site VPNs; transport mode is used for host-to-host communication.

**IKE (Internet Key Exchange):** The key management protocol for IPsec. Based on Diffie-Hellman. Generates the same symmetric session key at both ends of the VPN tunnel.

**Security Associations (SAs):** A one-way set of parameters for an IPsec connection. For bidirectional communication using both AH and ESP: **4 SAs total** (2 directions × 2 components).

### TLS VPN vs IPsec VPN

| | TLS VPN | IPsec VPN |
|--|--|--|
| Layer | Transport (4) and above | Network (3) |
| Encryption default | Yes — encrypts by default | No — requires IKE/ESP explicitly |
| Ease of setup | Easier | More complex |
| Attack impact | Compromise of specific systems/apps | Compromise of entire network |
| Key strength | Port-level granularity | IP-level — any host-to-host |

### Split Tunneling

Split tunneling allows the user to access corporate resources via VPN and the internet via their local connection simultaneously — corporate traffic goes through the VPN; non-corporate traffic bypasses it. **Risk:** internet traffic from the user's device bypasses organizational security controls (proxy, filtering, DLP).

### L2TP vs PPTP

- **PPTP:** Deprecated. Uses PAP/CHAP/EAP for auth. Avoid.
- **L2TP:** Layer 2; **no native encryption** — always pair with IPsec (L2TP/IPsec).

## Exam Nuance

- GRE is **not** a VPN — it provides encapsulation but **no encryption**.
- AH ≠ encryption. ESP provides encryption. If confidentiality is needed, **use ESP**.
- IKE is the key exchange protocol for IPsec; it is based on Diffie-Hellman (asymmetric key exchange to establish a shared symmetric key).
- Transport mode: payload encrypted, original IP header visible. Tunnel mode: entire original packet (header+payload) encrypted within new outer packet.
- SA count: AH + ESP, bidirectional = **4 SAs** (one per direction per component).

## Cross-links

- [IPsec](./vpn.md) — details above
- [Secure Protocols](./secure-protocols.md)
- [Network Access Control](./network-access-control.md)
- [OSI Model](./osi-model.md)

## Sources

- destination-cissp §4.3.1–4.3.2 (pp. 0638–0650)
- destination-cissp §4.3.3 (pp. 0650–0656)
