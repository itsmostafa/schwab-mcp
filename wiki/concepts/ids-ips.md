---
title: "IDS and IPS"
type: concept
domain: 4
tags: [ids, ips, intrusion-detection, intrusion-prevention, signature, anomaly, siem]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# IDS and IPS

Intrusion Detection Systems (IDS) and Intrusion Prevention Systems (IPS) provide detection and prevention capabilities that complement firewalls. Firewalls are preventive — IDS/IPS add detection and correction.

## Key Facts

### IDS vs IPS

| Feature | IDS | IPS |
|---------|-----|-----|
| Action | Detects, logs, alerts | Detects, logs, alerts, **and blocks/drops** |
| Placement | **Passive** — connected to mirror/span/promiscuous port | **Inline** — sits directly in the traffic path |
| Impact on traffic | Zero (receives a copy) | Can introduce latency |
| Risk if device fails | None (traffic still flows) | Can block legitimate traffic ("fail-open" vs "fail-closed") |

**Mirror / Span / Promiscuous port:** A switch port configured to receive a copy of all traffic passing through the switch. The IDS connects here so it can inspect all traffic without being inline.

### Network-Based vs Host-Based

| | Network-Based (NIDS/NIPS) | Host-Based (HIDS/HIPS) |
|--|--|--|
| Monitors | Traffic on a network segment | Activity on a single host/server |
| Placement | Sensor at network chokepoints | Agent installed on the endpoint |
| Best for | Detecting network-level attacks | Detecting attacks on critical servers |

Best practice: deploy **both** for layered detection.

### Detection Methods

| Method | How it works | Strength | Weakness |
|--------|-------------|----------|---------|
| **Signature-based** | Compares traffic/activity against known attack signatures (hashes, byte sequences, IP lists) | High accuracy for known threats; low false positives | Cannot detect **unknown (zero-day)** threats |
| **Anomaly-based** | Establishes a baseline of normal behavior; alerts on deviations | Can catch unknown attacks | High false positive rate; computationally expensive |
| **Heuristic** | Uses rules or ML to identify suspicious patterns without exact signatures | Bridges gap between signature and anomaly | May still miss novel attacks |

Anomaly detection methods: stateful matching, statistical anomalies, traffic anomalies, protocol anomalies.

### Alert States (True/False Positive/Negative)

| | Attack Present | No Attack |
|--|--|--|
| **Alert Generated** | True Positive ✓ | False Positive (tune) |
| **No Alert** | False Negative ✗ (worst!) | True Negative ✓ |

**False Negative** is the worst outcome — an attack is ongoing but the security team has no alert. False Positives cause alert fatigue. Tuning aims to minimize both without trading one for the other.

### Sandbox

When IDS/IPS identifies potentially malicious code, a **sandbox** can execute the code in an isolated environment to determine intent without risking production systems. Also used by malware analysts for reverse engineering.

### Ingress vs Egress Monitoring

- **Ingress:** Monitors inbound traffic — prevents malicious traffic from entering.
- **Egress:** Monitors outbound traffic — prevents data exfiltration and detects compromised hosts initiating DDoS or C2 beaconing.

Both directions should be monitored.

### Allow List / Deny List (Whitelist / Blacklist)

- **Allow list (whitelist):** Only listed IPs are permitted; everything else is blocked.
- **Deny list (blacklist):** Listed IPs are blocked; everything else is permitted.

## Exam Nuance

- IDS is **passive** (mirror/span port); IPS is **inline**. This distinction drives almost every IDS vs IPS exam question.
- SIEM integration: IDS sends events to SIEM for correlation; SIEM can trigger firewall rule changes for corrective action.
- False Negative is explicitly called out as the **worst-case scenario** in the source.
- Signature-based cannot detect **zero-day attacks** (no signature exists yet). Anomaly-based can, at the cost of higher false positives.

## Cross-links

- [Firewalls](./firewalls.md)
- [Secure Network Design](./secure-network-design.md)
- [Network Attacks](./network-attacks.md)
- [OSI Model](./osi-model.md)

## Sources

- destination-cissp §4.2.4–4.2.6 (pp. 0621–0637)
