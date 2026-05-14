---
title: "Network Devices"
type: concept
domain: 4
tags: [hub, switch, router, bridge, vlan, stp, layer2, layer3]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Network Devices

Network devices connect hosts and route/switch traffic. Knowing which OSI layer each device operates at is essential for CISSP exam questions.

## Key Facts

### Device by OSI Layer

| Device | OSI Layer | Key Behavior | Security Note |
|--------|-----------|-------------|--------------|
| Hub | Layer 1 | Broadcasts all traffic to all ports. Same collision domain. | Highly insecure — everyone sees all traffic |
| Repeater | Layer 1 | Regenerates signal to extend distance. | No security function |
| Concentrator | Layer 1 | Combines multiple signals into one line. | No intelligence |
| Bridge | Layer 2 | Connects two network segments; forwards based on MAC. | No traffic filtering |
| Switch (L2) | Layer 2 | Forwards frames only to the destination port (MAC-based). Separate collision domains per port. | Better than hub; VLAN support adds segmentation |
| Switch (L3) | Layer 3 | Performs routing based on IP addresses; often used with VLANs. | Exam: assume L2 unless explicitly stated L3 |
| Router | Layer 3 | Routes packets between networks based on IP; uses routing protocols (BGP, OSPF, RIP). | Boundary routers act as simple packet-filtering firewalls |
| Packet-filtering firewall | Layer 3 | Filters based on source/destination IP and port headers. | Fast, limited intelligence |

### Hubs vs Switches

**Hub:** Dumb device. Any frame received on one port is flooded to all other ports. All connected devices share one collision domain — every device sees every other device's traffic (major confidentiality risk).

**Switch:** Intelligent. Maintains a CAM table (Content Addressable Memory) mapping MAC addresses to ports. Forwards frames only to the correct destination port. Eliminates unnecessary broadcasts (except for unknown MACs and broadcasts).

### VLANs (Virtual LANs)

VLANs logically segment a physical network without physical rewiring. Implemented via **IEEE 802.1Q** tagging. A switch port is configured as a member of a VLAN; traffic from that port is tagged and only reaches ports in the same VLAN.

**Trunk links:** Carry traffic for multiple VLANs between switches.

**Security benefit:** Isolates broadcast domains; limits lateral movement. A device in VLAN 10 cannot directly communicate with VLAN 20 without a router or Layer 3 switch enforcing the routing.

### STP (Spanning Tree Protocol)

Ethernet networks with redundant links can form loops, causing broadcast storms. STP prevents loops by logically disabling redundant paths. Key points:
- When the active path fails, STP reconverges and activates a redundant path.
- RSTP (Rapid STP) is the modern version.

### Routing Protocols

Routers use routing protocols to dynamically maintain routing tables:
- **BGP (Border Gateway Protocol):** Inter-domain routing; used on the internet.
- **OSPF (Open Shortest Path First):** Interior gateway protocol; includes security features.
- **RIP (Routing Information Protocol):** Older, simpler; less secure.

## Exam Nuance

- **Assume switches are Layer 2** unless the question explicitly says "Layer 3 switch."
- Hub = shared collision domain = everyone sees everything = security risk.
- VLAN segmentation via 802.1Q is a **logical** control, not physical — misconfiguration can break isolation.
- A router used between an internal network and the internet is called a **boundary router** and functions as the simplest packet-filtering firewall.
- Bridges and switches differ: a bridge connects two segments; a switch has many ports and makes per-port forwarding decisions.

## Cross-links

- [OSI Model](./osi-model.md)
- [Secure Network Design](./secure-network-design.md)
- [SDN and NFV](./sdn-nfv.md)
- [Firewalls](./firewalls.md)
- [802.1X Standard](../standards/802-1x.md)

## Sources

- destination-cissp §4.1.2–4.1.5 (pp. 0501–0534)
- destination-cissp §4.1.15 (pp. 0593–0599)
