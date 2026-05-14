---
title: "Secure Network Design"
type: concept
domain: 4
tags: [dmz, bastion-host, defense-in-depth, segmentation, microsegmentation, nat, proxy]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Secure Network Design

Secure network design applies layered controls to limit blast radius, enforce least privilege at the network level, and create choke points for monitoring.

## Key Facts

**Defense in Depth:** Multiple overlapping security layers — policies/people (outermost), architecture, cabling/switching, OS, firewall configuration (innermost). No single control is sufficient alone.

**Network Segmentation / Partitioning:** Dividing the network into segments using switches, routers, and firewalls. Controls traffic flow between segments and limits lateral movement after a breach.

**DMZ (Demilitarized Zone):** A screened subnet placed between the internet and the internal network. Public-facing services (web, email, DNS, remote access) live here — isolated from the internal network. If a DMZ host is compromised, the attacker still faces a firewall before reaching internal systems.

**Bastion Host:** A hardened host placed in the DMZ, exposed to the internet and prepared for attack (French: *bastion* = fortress). Web servers, mail servers, and jump servers are typical bastion hosts.

**Network Perimeter:** The boundary the organization controls. Best practice: limit to **one ingress/egress choke point** where traffic is inspected. Multiple entry points make monitoring exponentially harder.

**Microsegmentation:** Using virtualization to deploy many small, granular network segments — often a single workload per segment — each with its own firewall rules. Makes lateral movement extremely difficult and allows tighter ACLs per service. Technologies: network overlays/encapsulation, distributed virtual firewalls/routers, IDS/IPS per segment, Zero Trust Architecture.

**NAT / PAT as Security Control:** By hiding internal IP addressing from external observers, NAT/PAT impedes attacker reconnaissance. Not a primary security control but a useful layer of obscurity.

**Proxy:** A device (usually at Layer 7) that acts on behalf of clients. Filters requests, enforces policy, and can block traffic to known malicious destinations. Circuit proxy firewalls (Layer 5) use NAT to hide internal IPs.

**Physical vs Logical Segmentation:**

| Type | Method | Security |
|------|--------|----------|
| In-band management | Same network as user traffic | Less secure |
| Out-of-band management | Separate management network | More secure |
| Air-gapped | Physically isolated — no connection to other networks | Most secure; highest operational cost |
| VLAN | Logical segmentation via 802.1Q tags | Cheaper/flexible; requires proper config |
| VRF | Multiple virtual routing tables on one device | Logical isolation |

**Traffic Flows (data center context):**

- **North-south:** Traffic between clients (internet) and data center servers. Traditionally the focus of perimeter firewalls.
- **East-west:** Traffic between servers within the data center. Increasingly important as attackers use lateral movement after initial compromise. Microsegmentation is the answer.

## Exam Nuance

- A **screened subnet** (two firewalls around a DMZ) is more expensive than a three-legged firewall but offers vendor diversity — a vulnerability in one vendor's firewall is unlikely to exist in another.
- **Bastion host** and **DMZ** are closely related but distinct: the DMZ is the network zone; the bastion host is the hardened device within it.
- Air-gapped networks require **physical presence** to manage — ICS/SCADA systems are the typical use case.
- Microsegmentation is closely linked to **Zero Trust** — treat every segment as untrusted regardless of location.

## Cross-links

- [Firewalls](./firewalls.md)
- [IDS and IPS](./ids-ips.md)
- [Network Devices](./network-devices.md)
- [IP Addressing](./ip-addressing.md)
- [SDN and NFV](./sdn-nfv.md)

## Sources

- destination-cissp §4.2.1 (pp. 0602–0613)
