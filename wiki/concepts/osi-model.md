---
title: "OSI Model"
type: concept
domain: 4
tags: [osi, networking, layers, protocols, encapsulation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# OSI Model

The Open Systems Interconnection (OSI) model is a seven-layer reference framework that describes how two networked devices communicate. Each layer has specific responsibilities and hands off to the adjacent layer via encapsulation (going down) and decapsulation (going up).

## Key Facts

| Layer | Name | PDU | Key Protocols/Devices |
|-------|------|-----|-----------------------|
| 7 | Application | Data | HTTP/S, FTP, DNS, SSH, SMTP, SNMP; Application-proxy firewalls, gateways |
| 6 | Presentation | Data | XML, JPEG, codecs, encryption/compression |
| 5 | Session | Data | PAP, CHAP, EAP, NetBIOS, RPC; Circuit-proxy firewalls |
| 4 | Transport | Segments / Datagrams | TCP, UDP, SSL/TLS |
| 3 | Network | Packets | IP, ICMP, IPsec, OSPF; Routers, packet-filtering firewalls |
| 2 | Data Link | Frames | ARP, RARP, L2TP, PPTP; Switches, bridges |
| 1 | Physical | Bits | Hubs, repeaters, NICs, cabling |

**Mnemonics (top → bottom):** "All People Seem To Need Data Processing"
**Mnemonics (bottom → top):** "Please Do Not Throw Sausage Pizza Away"

**Encapsulation / Decapsulation:** As data travels down the OSI stack each layer adds a header (and sometimes trailer), forming the PDU for that layer. On the receiving end, each layer strips its header as the data travels back up.

**Firewalls span multiple layers:** Packet-filtering at L3, stateful at L3/L4, circuit-proxy at L5, application-proxy at L7. Higher = more intelligence, lower = more speed.

**The higher the layer, the richer the security control** — and the more processing overhead.

## Exam Nuance

- When a question says "which layer does X operate at?" map the PDU or device to the table above.
- Routers operate at **Layer 3**; switches are **Layer 2** by default (Layer 3 switch is explicitly stated).
- SSL/TLS is listed at **Layer 4 (Transport)** in TCP/IP, but the session is established at **Layer 5** in the OSI model — the source is explicit that authentication protocols (PAP, CHAP, EAP) physically sit at L5 even though they ride PPP (L2).
- Application layer is where **most breaches occur** because it has the most functionality and the most code.
- OSI is a **model**; TCP/IP is the **implementation**.

## Cross-links

- [TCP/IP Model](./tcp-ip-model.md)
- [Firewalls](./firewalls.md)
- [Network Protocols](./network-protocols.md)
- [VPN](./vpn.md)
- [IDS and IPS](./ids-ips.md)

## Sources

- destination-cissp §4.1.1–4.1.10 (pp. 0493–0558)
