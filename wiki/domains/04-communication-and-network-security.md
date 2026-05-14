---
title: "Domain 4 — Communication and Network Security"
type: domain
domain: 4
tags: [networking, protocols, osi, tls, vpn, segmentation, wireless]
sources: [cissp-exam-outline, destination-cissp]
updated: 2026-05-13
---

# Domain 4 — Communication and Network Security

Exam weight: **13%**. Covers network architecture, secure protocols, segmentation strategies,
and network components.

## Subtopics (from CISSP Exam Outline)

### 4.1 Apply secure design principles in network architectures

The OSI model's seven layers provide a structured framework for understanding where security controls, protocols, and devices operate; higher layers offer richer functionality at the cost of speed, while lower layers are fast but have minimal intelligence. The TCP/IP model implements the OSI model with four layers (Application, Transport, Internet, Link) and hosts the protocol suite that drives the internet. Security professionals must know which protocol or device lives at each layer, because exam questions routinely ask "at which layer does X operate?" Key design principles span physical topology selection, logical segmentation with VLANs and VRF, microsegmentation for east-west traffic control, and SDN for software-driven policy enforcement.

- OSI and TCP/IP models
- IPv4 and IPv6 (unicast, broadcast, multicast, anycast)
- Secure protocols (IPSec, SSH, SSL/TLS)
- Implications of multilayer protocols
- Converged protocols (iSCSI, VoIP, InfiniBand over Ethernet, Compute Express Link)
- Transport architecture (topology, data/control/management plane, cut-through/store-and-forward)
- Performance metrics (bandwidth, latency, jitter, throughput, signal-to-noise ratio)
- Traffic flows (north-south, east-west)
- Physical segmentation (in-band, out-of-band, air-gapped)
- Logical segmentation (VLANs, VPNs, virtual routing and forwarding, virtual domain)
- Micro-segmentation (network overlays/encapsulation; distributed firewalls, routers, IDS/IPS, zero trust)
- Edge networks (ingress/egress, peering)
- Wireless networks (Bluetooth, Wi-Fi, Zigbee, satellite)
- Cellular/mobile networks (4G, 5G)
- Content distribution networks (CDN)
- Software defined networks (SDN) / Software-Defined WAN / network functions virtualization
- Virtual Private Cloud (VPC)
- Monitoring and management (network observability, traffic flow/shaping, capacity management)

### 4.2 Secure network components

Network components include the transmission media (twisted pair, coaxial, fiber optic, wireless) and the devices that connect them (hubs, switches, routers, firewalls, IDS/IPS). Fiber optic is the most secure wired medium because it is harder to tap than copper and immune to electromagnetic interference. Firewall architectures range from simple packet filtering at Layer 3 to application-proxy firewalls at Layer 7, with the trade-off between intelligence and throughput defining which architecture suits a given security requirement. IDS/IPS devices provide detection and prevention beyond what firewalls offer — an IDS sits passively on a mirror/span port while an IPS is deployed inline to actively block malicious traffic.

- Operation of infrastructure (redundant power, warranty, support)
- Transmission media (physical security, signal propagation quality)
- Network Access Control (NAC) systems (physical and virtual)
- Endpoint security (host-based)

### 4.3 Implement secure communication channels according to design

Secure communication channels protect data in transit over untrusted networks. VPNs (tunneling + encryption) are the primary tool; IPsec operates at Layer 3 and is the preferred VPN for site-to-site connectivity, offering authentication via AH and encryption via ESP in either transport or tunnel mode; TLS VPNs are simpler to establish and encrypt by default at the Transport layer. Remote authentication protocols — RADIUS, TACACS+, and Diameter — complement VPN by verifying the identity of the user, not just the device. VoIP introduces convergence risks such as toll fraud, eavesdropping, and vishing, mitigated by SRTP encryption and strong SIP configuration.

- Voice, video, and collaboration (conferencing, Zoom rooms)
- Remote access (network administrative functions)
- Data communications (backhaul networks, satellite)
- Third-party connectivity (telecom providers, hardware support)

## Key concepts

- [OSI Model](../concepts/osi-model.md) — 7 layers, PDUs, protocols and devices per layer
- [TCP/IP Model](../concepts/tcp-ip-model.md) — 4-layer mapping to OSI, TCP vs UDP
- [IP Addressing](../concepts/ip-addressing.md) — IPv4 classes, CIDR, NAT/PAT, IPv6
- [Network Protocols](../concepts/network-protocols.md) — TCP handshake, UDP, ICMP, ARP, DNS, DHCP; key port numbers
- [Secure Network Design](../concepts/secure-network-design.md) — defense in depth, DMZ, bastion host, microsegmentation
- [Firewalls](../concepts/firewalls.md) — packet filter, stateful, circuit proxy, application proxy; architectures
- [IDS and IPS](../concepts/ids-ips.md) — inline vs passive, signature vs anomaly, false positive/negative
- [VPN](../concepts/vpn.md) — IPsec (AH/ESP, transport/tunnel, IKE), TLS VPN, split tunneling, L2TP/PPTP
- [Wireless Security](../concepts/wireless-security.md) — WEP/WPA/WPA2/WPA3, EAP variants, 802.1X
- [Network Attacks](../concepts/network-attacks.md) — DoS/DDoS, SYN flood, MITM, ARP poisoning, spoofing
- [Network Devices](../concepts/network-devices.md) — hub/switch/router/bridge, VLAN, STP
- [Secure Protocols](../concepts/secure-protocols.md) — SSH vs Telnet, HTTPS, SFTP, SNMPv3, TLS versions
- [VoIP Security](../concepts/voip-security.md) — SIP, H.323, SRTP, vishing, toll fraud
- [Network Access Control](../concepts/network-access-control.md) — NAC, 802.1X, RADIUS, TACACS+
- [SDN and NFV](../concepts/sdn-nfv.md) — control/data/application planes, northbound/southbound APIs

## Sources
- cissp-exam-outline (subtopic list, exam weight)
- destination-cissp §4.1–4.3
