---
title: "Network Attacks"
type: concept
domain: 4
tags: [dos, ddos, syn-flood, mitm, arp-poisoning, spoofing, replay, session-hijacking]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Network Attacks

Network attacks follow the same reconnaissance → enumeration → vulnerability analysis → exploitation progression as penetration tests, differing only in intent and respect for rules of engagement.

## Key Facts

### Attack Types: Passive vs Active

| Type | Changes traffic/state? | Target awareness | Example |
|------|----------------------|-----------------|---------|
| **Passive** | No | None | Eavesdropping / sniffing |
| **Active** | Yes | Often alerted | SYN flood, MITM, ARP poisoning |

### DoS / DDoS

- **DoS (Denial of Service):** One machine overwhelms a target until it crashes or becomes unavailable.
- **DDoS (Distributed DoS):** Multiple compromised hosts (botnet) act in unison against the target, amplifying the impact.

### SYN Flood

Attacker sends a flood of SYN packets without completing the three-way handshake (never sends ACK). The server allocates resources for each half-open connection, filling its connection table until it crashes. Mitigation: SYN proxy, SYN cookies, IPS/firewall rate limiting.

**SYN Scan (Stealth scan):** Attacker sends SYN, receives SYN-ACK if port open, then sends RST instead of ACK — never completes the connection. Used for stealthy port discovery (Nmap default).

### IP-Based Attacks

| Attack | Mechanism |
|--------|-----------|
| **Smurf** | Attacker spoofs victim's IP; sends ICMP echo requests to broadcast address → all hosts reply to victim (ICMP flood). Old attack. |
| **Fraggle** | Like Smurf but uses UDP packets to ports 7/19 (CHARGEN). Old attack. |
| **Teardrop** | Sends malformed, overlapping IP fragments. Target cannot reassemble → crash (DoS). |
| **Fragment/Overlapping Fragment** | Overlapping fragments bypass IDS/firewall; reassembly at destination creates attack payload. |

### Man-in-the-Middle (MITM)

Attacker inserts themselves into the communication path between two parties. Can intercept, modify, or inject traffic. Enabled by ARP poisoning, rogue DHCP, evil twin AP. Countered by mutual authentication and TLS.

### Spoofing / Masquerading

Pretending to be another entity (IP, MAC, email address, DNS record). When IP spoofing is used, the attacker **cannot receive reply traffic** — replies go to the spoofed IP's legitimate owner. Useful for DoS but not for interactive attacks.

### ARP Poisoning

ARP has no authentication. An attacker sends gratuitous ARP replies associating their MAC with a legitimate IP, poisoning the victim's (and switch's) ARP table. Traffic intended for the legitimate host is sent to the attacker. Mitigation: dynamic ARP inspection on managed switches, monitoring, DNSSEC for DNS-side equivalent.

### DNS Cache Poisoning

Attacker injects false records into a DNS resolver's cache, redirecting legitimate domain lookups to malicious IPs. DNSSEC cryptographically signs DNS records to prevent this.

### Session Hijacking / Replay

- **Session hijacking:** Attacker takes over an established session (e.g., by stealing session tokens or predicting sequence numbers).
- **Replay attack:** Attacker captures and retransmits valid authentication messages. Countered by nonces, timestamps, and session tokens.

### Attack Phase Model

1. Reconnaissance (passive info gathering)
2. Scanning / Enumeration (active probing for open ports, accounts)
3. Vulnerability Analysis (map weaknesses)
4. Exploitation (execute attack)

Goal: make each phase as expensive and difficult as possible for attackers.

## Exam Nuance

- Smurf and Fraggle are **old attacks** — unlikely in the wild today, but appear on exams.
- IP spoofing does **not** give the attacker a return channel — responses go to the real owner of the spoofed IP.
- SYN scan = stealth/half-open scan; SYN flood = DoS attack. Same mechanism, different purpose.
- ARP has **no authentication** — this is fundamental to understanding why ARP poisoning is so easy.
- Teardrop and fragment attacks target the **IP reassembly** process.

## Cross-links

- [TCP/IP Model](./tcp-ip-model.md)
- [Network Protocols](./network-protocols.md)
- [IDS and IPS](./ids-ips.md)
- [Firewalls](./firewalls.md)
- [Wireless Security](./wireless-security.md)

## Sources

- destination-cissp §4.1.13 (pp. 0565–0582)
