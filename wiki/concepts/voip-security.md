---
title: "VoIP Security"
type: concept
domain: 4
tags: [voip, sip, srtp, h323, vishing, toll-fraud, convergence]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# VoIP Security

Voice over IP (VoIP) enables voice communications over data networks (IP). It is an example of **IP convergence** — native IP networks carrying non-IP traffic. Convergence adds functionality but also introduces new attack surfaces because data networks have no built-in security.

## Key Facts

### Converged Protocols

| Protocol | Purpose |
|----------|---------|
| VoIP | Voice over IP networks |
| FCoE | Fibre Channel over Ethernet — storage traffic over Ethernet |
| iSCSI | SCSI storage commands over IP networks (SANs) |
| InfiniBand | Remote Direct Memory Access (RDMA) — used in ML/HPC |
| Compute Express Link | Connects CPUs to memory/devices at high speed |

### VoIP Protocols

| Protocol | Purpose | Notes |
|----------|---------|-------|
| **SIP (Session Initiation Protocol)** | Initiates, maintains, and terminates voice/video sessions | Can connect PBX to PSTN |
| **H.323** | Older VoIP signaling protocol | Predates SIP; still in use |
| **RTP (Real-time Transport Protocol)** | Carries voice/video data | **No security** |
| **SRTP (Secure RTP)** | Secure version of RTP | Provides encryption, authentication, integrity, replay protection (RFC 3711) |

### VoIP Attacks

| Attack | Description |
|--------|-------------|
| **Vishing** | Voice phishing — attacker calls victim using spoofed caller ID to extract information or induce action. VoIP makes number spoofing trivial. |
| **Toll fraud** | Attacker hijacks a PBX to make unauthorized long-distance/international calls at the victim's expense. |
| **Eavesdropping** | VoIP packets sniffed on the network; calls decoded from RTP streams. Countered by SRTP. |
| **Call hijacking** | Attacker inserts themselves into a VoIP session (similar to MITM). |

**Vishing vs Smishing:**
- **Vishing:** Voice call (phone).
- **Smishing:** SMS text message.

### PBX and PSTN

- **PBX (Private Branch Exchange):** A private internal telephone network within an organization or hotel.
- **PSTN (Public Switched Telephone Network):** The traditional copper-wire telephone network.

## Exam Nuance

- SRTP is the secure version of RTP; **RTP alone has no security**.
- SIP is responsible for **session management** (setup, teardown); SRTP handles **encryption of the media stream**.
- Vishing exploits VoIP's ease of number spoofing — not a technical network attack but a social engineering vector.
- Convergence: adding any non-native protocol to an IP network **always introduces new attack surface** even if the added protocol is itself secure.
- Adding encryption (SRTP) to VoIP adds **latency** — a trade-off to document when justifying security decisions.

## Cross-links

- [Network Protocols](./network-protocols.md)
- [Network Attacks](./network-attacks.md)
- [Secure Protocols](./secure-protocols.md)

## Sources

- destination-cissp §4.1.12 (pp. 0560–0565)
