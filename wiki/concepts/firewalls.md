---
title: "Firewalls"
type: concept
domain: 4
tags: [firewall, acl, packet-filter, stateful, application-proxy, ngfw, waf]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Firewalls

A firewall enforces security rules between two or more networks by filtering traffic. Firewalls are **preventive controls**. They range from simple packet-filtering routers to sophisticated application-layer proxies.

## Key Facts

### Firewall Types by OSI Layer

| Type | OSI Layer | Intelligence | Speed | Filtering basis |
|------|-----------|-------------|-------|-----------------|
| Packet Filtering | Layer 3 | Lowest | Fastest | Source/destination IP and port (header only) |
| Stateful Packet Filtering | Layer 3 & 4 | Medium | Fast | Header + state table (tracks connection context) |
| Circuit Proxy (Gateway) | Layer 5 | Medium | Higher latency | Session-level rules; uses NAT to hide internal IPs |
| Application Proxy | Layer 7 | Highest | Highest latency | Full payload inspection; deep packet inspection |

**Key principle:** Higher OSI layer = more intelligence, more functionality, more overhead and latency.

### Firewall Technologies Summarized

- **Packet Filter:** Uses ACLs. Reads header fields only. No awareness of session state. Fast but easily spoofed.
- **Stateful Inspection:** Maintains a state table of active connections. Can identify unexpected packets (e.g., ACK without a prior SYN). Covers TCP/UDP stateful tracking.
- **Circuit-Level Proxy (SOCKS):** Sits at Layer 5. Validates TCP handshake/sessions rather than packet content. No application-specific logic. Hides internal network via NAT.
- **Application-Level Proxy:** Layer 7. Full packet payload inspection, deep packet inspection, antivirus, content filtering, stateful inspection, IDS signatures. Most capable but slowest.
- **CBAC (Context-Based Access Control):** Firewall software feature that filters TCP/UDP based on application-layer session info; can detect DDoS patterns.
- **WAF (Web Application Firewall):** Application-proxy specialization for HTTP/S traffic.
- **NGFW (Next-Generation Firewall):** Combines stateful inspection, application awareness, IPS, and threat intelligence in a single device.

### Firewall Architectures

| Architecture | Description |
|---|---|
| **Packet Filtering** | Single router between internet and internal network. Simplest. |
| **Dual-Homed Host** | Computer with two NICs; understands all OSI layers; can run any firewall technology. |
| **Screened Host** | Packet-filtering router + bastion host behind it. Two-layer protection. |
| **Screened Subnet** | Two firewalls creating a DMZ between them. Most robust; allows vendor diversity. |
| **Three-Legged Firewall** | Single firewall with 3 interfaces: internet, DMZ, internal network. |

### Implicit Deny

A fundamental ACL principle: **anything not explicitly permitted is denied.** Firewall rules should end with a deny-all rule; relying on implicit deny without logging gives no visibility into blocked traffic.

## Exam Nuance

- A **boundary router** is the simplest firewall — it filters based on packet headers using ACLs.
- Stateful inspection tracks **connectionless** protocols (UDP, RPC) using source/destination IP+port to infer state.
- Circuit proxies (Layer 5) are **not** the same as application proxies (Layer 7) — circuit proxies don't inspect payload; they only validate the session handshake.
- Screened subnet uses **two separate firewalls** — ideally from different vendors (diversity reduces shared vulnerability risk).
- Firewalls are preventive controls; they need IDS/IPS to complete the detect-correct loop.

## Cross-links

- [OSI Model](./osi-model.md)
- [IDS and IPS](./ids-ips.md)
- [Secure Network Design](./secure-network-design.md)
- [Network Attacks](./network-attacks.md)

## Sources

- destination-cissp §4.2.2–4.2.3 (pp. 0614–0621)
