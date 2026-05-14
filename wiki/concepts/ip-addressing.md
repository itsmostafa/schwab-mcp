---
title: "IP Addressing"
type: concept
domain: 4
tags: [ip, ipv4, ipv6, nat, pat, subnetting, cidr]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# IP Addressing

IP addressing provides the logical addressing scheme that allows packets to be routed across networks. IPv4 uses 32-bit addresses; IPv6 expands this to 128 bits to solve address exhaustion.

## Key Facts

### IPv4

- **Format:** Four decimal octets separated by dots (e.g., `192.168.1.1`). Each octet = 8 bits. Valid range 0–255.
- **Address space:** 2^32 = ~4.3 billion addresses.
- **.0 suffix** = network address; **.255 suffix** = broadcast address (neither assignable to hosts).

**IP Classes (subnetting):**

| Class | Range | Host bits | Max hosts |
|-------|-------|-----------|-----------|
| A | 1–126.x.x.x | 24 | 16,777,214 |
| B | 128–191.x.x.x | 16 | 65,534 |
| C | 192–223.x.x.x | 8 | 254 |
| D | 224–239.x.x.x | — | Multicast |
| E | 240–255.x.x.x | — | Reserved |

**Private IPv4 ranges (RFC 1918 — non-routable):**

| From | To |
|------|----|
| 10.0.0.0 | 10.255.255.255 |
| 172.16.0.0 | 172.31.255.255 |
| 192.168.0.0 | 192.168.255.255 |

Private addresses are not routable on the internet. Two organizations can use the same private range internally.

### IPv6

- **Format:** 128 bits, written as eight groups of four hex digits separated by colons (e.g., `0000:0000:0000:0000:0000:ffff:0a00:0001`).
- **Address space:** 2^128 — effectively inexhaustible.
- **IPsec support:** Native to IPv6 (optional in IPv4).
- **Backward compatible:** Organizations can run IPv4 and IPv6 simultaneously during transition.

### NAT and PAT

**NAT (Network Address Translation):** Translates private (internal, non-routable) IP addresses to a public (routable) IP address when traffic leaves the network, and reverses the mapping for inbound replies. Security benefit: hides internal IP addressing from external reconnaissance.

**PAT (Port Address Translation):** An extension of NAT that additionally translates source port numbers, allowing many internal hosts to share one public IP by using unique port numbers per connection. PAT is how a home router serves multiple devices through a single ISP address.

## Exam Nuance

- Know the max host counts for Class A (16,777,214), Class B (65,534), Class C (254) — these appear on the exam.
- Private addresses are **non-routable** — a key security feature; they cannot be the direct target of an internet-based attack.
- NAT provides **obscurity** (not full security) by hiding internal topology.
- IPv6 mandates IPsec support; IPv4 treats it as optional.

## Cross-links

- [Network Protocols](./network-protocols.md)
- [Secure Network Design](./secure-network-design.md)
- [Network Devices](./network-devices.md)

## Sources

- destination-cissp §4.1.5–4.1.6 (pp. 0530–0543)
