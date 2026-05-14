---
title: "Practice Questions — Domain 04: Communication And Network Security"
type: practice
domain: 04
tags: [practice, networking, osi, firewalls, ids, ips, vpn, wireless]
sources: [destination-cissp]
updated: 2026-05-13
---

# Practice Questions — Domain 04: Communication And Network Security

## Questions

### Q1: At which OSI layer does a packet-filtering firewall operate?

**Answer:** Layer 3 — Network.

**Why:** Packet-filtering firewalls make decisions based on IP addresses and port numbers, which exist in the packet header at the Network layer. They cannot inspect the payload. Stateful firewalls operate at Layers 3–4; circuit proxy at Layer 5; application proxy at Layer 7. Higher layer = more intelligence, lower layer = more speed.

*Subtopic: 4.2 — [Firewalls](../concepts/firewalls.md)*

---

### Q2: An organization needs to deploy an IDS. How should it be connected to the network, and why?

**Answer:** The IDS should be connected to a mirror/span/promiscuous port on a switch. This allows it to receive a copy of all traffic transiting the switch without being inline (in the traffic path).

**Why:** An IDS is a passive detection device — it only detects, logs, and alerts; it does not block. Placing it on a mirror/span port means it has no impact on traffic flow. An IPS, by contrast, is placed inline so it can actively drop or modify malicious traffic.

*Subtopic: 4.2 — [IDS and IPS](../concepts/ids-ips.md)*

---

### Q3: Which IPsec sub-protocol provides confidentiality (encryption), and which provides only integrity and authentication?

**Answer:** ESP (Encapsulating Security Payload) provides integrity, authentication, replay protection, AND encryption. AH (Authentication Header) provides integrity, authentication, and replay protection only — it does NOT encrypt the payload.

**Why:** If an exam scenario requires confidentiality, ESP must be used. AH alone is insufficient. Using both AH and ESP in a bidirectional VPN requires 4 Security Associations (one per direction × one per component).

*Subtopic: 4.3 — [VPN](../concepts/vpn.md)*

---

### Q4: What is the difference between IPsec transport mode and tunnel mode?

**Answer:** In **transport mode**, the original IP header is preserved; only the payload is encrypted/authenticated. In **tunnel mode**, the entire original packet (header + payload) is encapsulated inside a new outer IP packet and encrypted.

**Why:** Transport mode is used for host-to-host communication. Tunnel mode is used for site-to-site VPNs because it hides the original IP header, protecting internal network topology from eavesdroppers.

*Subtopic: 4.3 — [VPN](../concepts/vpn.md)*

---

### Q5: An organization is comparing wireless security protocols. WPA uses TKIP; WPA2 uses CCMP. Why is CCMP considered superior?

**Answer:** CCMP is based on AES (Advanced Encryption Standard) with 128-bit keys and provides stronger encryption and integrity than TKIP, which uses RC4 (a stream cipher) with known weaknesses.

**Why:** TKIP was created as a stopgap to fix WEP without requiring hardware replacement. It is no longer considered secure and has been superseded by CCMP/AES. WEP is fully broken due to a weak IV that allowed attackers to statistically recover the key. WPA3 extends this further with GCMP.

*Subtopic: 4.1 — [Wireless Security](../concepts/wireless-security.md)*

---

### Q6: A company's security analyst sees that an attack is occurring but the IDS raised no alert. What type of alert state is this, and why is it considered the worst outcome?

**Answer:** This is a **False Negative** — an attack is occurring, but no alert is generated.

**Why:** False Negatives are the worst outcome because the security team has no visibility into an active attack, meaning no corrective action is taken. False Positives (alerts with no real attack) are bad for efficiency but not dangerous — the team can investigate and dismiss. A False Negative allows attackers to operate undetected.

*Subtopic: 4.2 — [IDS and IPS](../concepts/ids-ips.md)*

---

### Q7: Which remote authentication protocol uses TCP and encrypts the entire packet? How does it differ from RADIUS?

**Answer:** TACACS+ uses TCP and encrypts the entire packet payload. RADIUS uses UDP and only poorly obfuscates the user's password.

**Why:** TACACS+ was developed by Cisco as an improvement over RADIUS. The full-packet encryption of TACACS+ prevents an attacker from even seeing the structure of the authentication exchange. Diameter is the successor to RADIUS and adds EAP and improved security for modern mobile networks.

*Subtopic: 4.3 — [Network Access Control](../concepts/network-access-control.md)*

---

### Q8: A user reports that when they use split tunneling, they can access both the corporate intranet and the public internet simultaneously. What is the security concern with this setup?

**Answer:** Internet traffic bypasses the corporate VPN tunnel and goes directly from the user's device to the internet, bypassing organizational security controls such as web proxies, URL filtering, DLP, and IDS/IPS.

**Why:** Split tunneling reduces bandwidth consumption (a benefit) but creates a security gap: if the user's device is compromised via the unfiltered internet connection, the attacker may then have a path into the corporate VPN tunnel. Organizations concerned about control should disable split tunneling and route all traffic through the VPN.

*Subtopic: 4.3 — [VPN](../concepts/vpn.md)*

---

### Q9: Which EAP variant provides the strongest authentication and why?

**Answer:** **EAP-TLS** provides the strongest authentication because it requires certificate-based authentication for **both the client and the server** (mutual authentication).

**Why:** PEAP and EAP-TTLS only require the server to present a certificate; the client authenticates with a username/password. LEAP uses password-based authentication for both sides (low security, deprecated). EAP-MD5 provides no server authentication at all. EAP-TLS's certificate requirement for both parties provides the highest assurance.

*Subtopic: 4.1 — [Wireless Security](../concepts/wireless-security.md)*

---

### Q10: A screened subnet firewall architecture is more expensive than a three-legged firewall. What is its key advantage?

**Answer:** A screened subnet uses two separate firewalls from potentially different vendors. If a vulnerability is discovered in one vendor's firewall, the same vulnerability is unlikely to exist in the other vendor's product, providing defense in depth.

**Why:** A three-legged firewall uses a single device with three network interfaces — if that device is compromised, all three network zones are affected. The screened subnet's dual-firewall design creates a true DMZ with two independent enforcement points, significantly raising the bar for an attacker.

*Subtopic: 4.2 — [Firewalls](../concepts/firewalls.md) | [Secure Network Design](../concepts/secure-network-design.md)*

---

### Q11: What is ARP poisoning and why is it easy to execute?

**Answer:** ARP poisoning is an attack where an attacker sends forged ARP replies to associate their MAC address with a legitimate IP address, redirecting traffic intended for the legitimate host to the attacker's machine.

**Why:** ARP has no built-in authentication. Any device can send an ARP reply claiming to own any IP address, and other devices (including switches) will update their ARP tables accordingly. This makes MITM attacks trivially easy on unsegmented networks. Mitigations include dynamic ARP inspection on managed switches, static ARP entries, and network monitoring.

*Subtopic: 4.1 — [Network Attacks](../concepts/network-attacks.md)*

---

### Q12: A security architect needs to ensure that a VoIP deployment is secure against eavesdropping. Which protocol should be used for the media stream and which for session management?

**Answer:** **SRTP (Secure Real-time Transport Protocol)** should be used to encrypt the media stream (voice/video data). **SIP (Session Initiation Protocol)** handles session initiation, maintenance, and teardown — it should be secured using TLS (SIP over TLS).

**Why:** RTP (without the S) carries voice data with no encryption — intercepted packets can be decoded to reconstruct conversations. SRTP adds encryption, authentication, integrity, and replay protection (RFC 3711). SIP governs the control channel; if SIP is unencrypted, attackers can observe or manipulate call setup.

*Subtopic: 4.3 — [VoIP Security](../concepts/voip-security.md)*
