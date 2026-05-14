---
title: "Network Protocols"
type: concept
domain: 4
tags: [tcp, udp, icmp, arp, dns, dhcp, ftp, ssh, http, smtp, ports]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Network Protocols

Network protocols are standardized rules for communication between devices. This page covers the exam-critical protocols, their OSI/TCP-IP layer, and associated security concerns.

## Key Facts

### Critical Port Numbers (memorize these)

| Port | Protocol | Notes |
|------|----------|-------|
| 20 | FTP (data) | Plaintext — insecure |
| 21 | FTP (control) | Plaintext — insecure |
| 22 | SSH / SFTP | Secure shell; also carries SFTP |
| 23 | Telnet | Plaintext remote login — insecure |
| 25 | SMTP | Email sending |
| 53 | DNS | TCP and UDP |
| 69 | TFTP | UDP; highly insecure — disable |
| 80 | HTTP | Plaintext web |
| 110 | POP3 | Email receive |
| 143 | IMAP | Email receive |
| 161/162 | SNMP | UDP; use v3 only |
| 443 | HTTPS | HTTP over TLS |
| 3389 | RDP | Remote Desktop Protocol |

### Protocol Reference

**ARP (Address Resolution Protocol):** Maps IP addresses to MAC addresses. Operates at Layer 2/3 boundary. Has **no built-in authentication** — vulnerability exploited by ARP poisoning. RARP maps MAC → IP.

**ICMP (Internet Control Message Protocol):** Layer 3. Used by `ping` (host reachability) and `traceroute` (path mapping). Frequently filtered at firewalls because attackers use it for reconnaissance.

**DNS (Domain Name System):** Maps hostnames to IPs. Uses port 53 (TCP and UDP). Little built-in security — vulnerable to **DNS cache poisoning**. DNSSEC mitigates this by cryptographically signing DNS records.

**DHCP (Dynamic Host Configuration Protocol):** Automatically assigns IP addresses. Attackers can impersonate a DHCP server (rogue DHCP) to set themselves as the default gateway, enabling traffic interception.

**TCP (Transmission Control Protocol):** Reliable, ordered, connection-oriented. Three-way handshake (SYN→SYN-ACK→ACK). Segments. Slower than UDP.

**UDP (User Datagram Protocol):** Unreliable, connectionless ("send and pray"). Datagrams. Fast — used for DNS, VoIP, streaming.

**SNMP (Simple Network Management Protocol):** Manages network devices. Ports 161/162 UDP. SNMPv1 and v2 are **highly vulnerable** (community strings in plaintext). **Use SNMPv3** only — it adds authentication and encryption.

**HTTP/HTTPS:** HTTP (port 80) is plaintext. HTTPS (port 443) = HTTP over TLS. The HTTPS padlock appears only after TLS handshake completes.

**FTP/SFTP/TFTP:** FTP (20/21) is plaintext. SFTP (port 22, rides SSH) is secure. TFTP (port 69 UDP) is highly insecure — **disable it**.

**SSH (Secure Shell):** Port 22. Uses public-key cryptography. Replaces Telnet for remote login; also secures FTP → SFTP.

**Telnet:** Port 23. Plaintext remote login — always replace with SSH.

## Exam Nuance

- ARP operates at the **Layer 2 / Layer 3 boundary**: it resolves IP addresses (L3) to MAC addresses (L2). The source treats it as an L2 protocol but discusses it in L3 context — this is a common exam trap.
- TFTP uses **UDP port 69** — it has no authentication and should be disabled in all security-conscious environments.
- **SNMPv1 and v2 are insecure** — only v3 provides authentication and encryption.
- DHCP rogue server attacks enable **MITM** by poisoning gateway info distributed to clients.
- DNS poisoning can redirect users to malicious sites even when they type the correct URL; DNSSEC is the countermeasure.

## Cross-links

- [OSI Model](./osi-model.md)
- [TCP/IP Model](./tcp-ip-model.md)
- [Network Attacks](./network-attacks.md)
- [Secure Protocols](./secure-protocols.md)
- [IP Addressing](./ip-addressing.md)

## Sources

- destination-cissp §4.1.7–4.1.10 (pp. 0543–0558)
- destination-cissp §4.1.13 (pp. 0565–0582)
