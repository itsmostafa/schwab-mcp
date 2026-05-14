---
title: "TCP/IP Model"
type: concept
domain: 4
tags: [tcp-ip, networking, tcp, udp, protocols]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# TCP/IP Model

The TCP/IP model is the practical four-layer implementation of the OSI reference model. It is the protocol suite that powers the internet. TCP and UDP are the two workhorses at the Transport layer.

## Key Facts

| TCP/IP Layer | Maps to OSI Layers | Key Protocols |
|---|---|---|
| 4 — Application | 7 (Application) + 6 (Presentation) + 5 (Session) | HTTP/S, FTP, DNS, SSH, SMTP, SNMP |
| 3 — Transport | 4 (Transport) | TCP, UDP, SSL/TLS |
| 2 — Internet | 3 (Network) | IP, ICMP, IPsec |
| 1 — Link | 2 (Data Link) + 1 (Physical) | ARP, L2TP, PPTP, Ethernet |

### TCP vs UDP

| Feature | TCP | UDP |
|---|---|---|
| Reliability | Reliable, ordered | Unreliable, unordered ("send and pray") |
| Connection | Connection-oriented (3-way handshake) | Connectionless |
| Speed | Slower | Faster |
| Use cases | Web, email, file transfer | Streaming, DNS, VoIP |
| PDU name | Segment | Datagram |

### TCP Three-Way Handshake

1. **SYN** — Client sends SYN + random sequence number (e.g., 1000).
2. **SYN-ACK** — Server acknowledges (ACK = 1001) and sends its own SYN (e.g., 2000).
3. **ACK** — Client acknowledges (ACK = 2001). Full-duplex connection established.

Graceful teardown: FIN → ACK → FIN → ACK (two pairs).

### Port Ranges

| Range | Name | Notes |
|---|---|---|
| 0–1023 | Well-known | HTTP, HTTPS, SMTP, etc. |
| 1024–49151 | Registered | IANA-assigned (e.g., Viber UDP 4244) |
| 49152–65535 | Dynamic/Ephemeral | Dynamically assigned by OS for outbound connections |

## Exam Nuance

- OSI is the **model**; TCP/IP is the **implementation**. The exam uses both — know the mapping.
- TCP's random sequence number prevents **session hijacking** — an attacker who cannot predict the sequence number cannot forge packets into the session.
- A **SYN flood** exploits the half-open state between steps 1 and 3, filling the server's connection queue until it crashes; mitigated by SYN proxies, IPS, or SYN cookies.
- SSL/TLS sits at **Layer 4 (Transport)** in the TCP/IP model but is discussed in the OSI context at Layer 5/6.

## Cross-links

- [OSI Model](./osi-model.md)
- [Network Protocols](./network-protocols.md)
- [Network Attacks](./network-attacks.md)
- [VPN](./vpn.md)

## Sources

- destination-cissp §4.1.1, §4.1.7 (pp. 0493–0552)
