---
title: "SDN and NFV"
type: concept
domain: 4
tags: [sdn, nfv, software-defined-networking, control-plane, data-plane, vpc]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# SDN and NFV

Software-Defined Networking (SDN) decouples the control plane from the data plane, enabling network management through software rather than hardware configuration. NFV (Network Functions Virtualization) virtualizes network functions (firewalls, routers, IDS/IPS) as software running on commodity hardware.

## Key Facts

### SDN Architecture

SDN uses three planes:

| Plane | Role | Communication |
|-------|------|--------------|
| **Application plane** | Hosts SDN applications: firewall logic, reporting, network management | → Control plane via **Northbound APIs** |
| **Control plane** | The "brain" — determines routing decisions, maintains routing tables and protocols. One centralized controller for the SDN. | ↔ Data plane via **Southbound APIs** |
| **Data plane** | The "doer" — actually forwards packets based on directions from the control plane. | |

**Northbound APIs:** Application ↔ Control plane communication.
**Southbound APIs:** Control ↔ Data plane communication (controller to physical hardware).

### SDN vs Traditional Networks

| Feature | Traditional Network | SDN |
|---|---|---|
| Control logic | Distributed in each device | Centralized in SDN controller |
| Configuration | Per-device CLI | Software/API-driven; centralized |
| Reconfiguration | Manual, slow | Rapid, programmable |
| Security risk | Per-device exploits | Centralized controller is a single point of failure / high-value target |

### NFV (Network Functions Virtualization)

NFV replaces physical network appliances with software running on standard servers/VMs:
- Virtual firewalls, routers, IDS/IPS, load balancers.
- Reduces hardware costs; enables rapid deployment.
- Enables microsegmentation with per-workload virtual firewalls.

### VPC (Virtual Private Cloud)

A logically isolated portion of a public cloud provider's infrastructure. The isolation is **purely virtual** — no separate physical hardware. Customers can configure their own IP ranges, subnets, routing, and security groups within the VPC.

### Security Implications

- SDN's centralized controller is a **high-value attack target** — compromising it gives control over the entire network.
- Software-based network functions may inherit software vulnerabilities not present in purpose-built hardware.
- However, SDN enables **faster response** to threats — security policy changes can be pushed to the entire network instantly.
- NFV and microsegmentation together enable **zero trust architecture** at scale.

### WAN Technologies (Related)

| Technology | Notes |
|---|---|
| X.25 | Pioneer packet-switching; excellent error correction; high overhead. Legacy. |
| Frame Relay | Focuses on speed over error correction; uses PVCs/SVCs. Legacy. |
| ATM | Connection-oriented; high-speed; fixed 53-byte cells. Legacy. |
| MPLS | Current standard; forwarding labels guarantee customer traffic isolation within provider network; but organization's data should still be encrypted because provider admins could access it. |

## Exam Nuance

- Northbound API = Application ↔ Control; Southbound API = Control ↔ Data. Know the direction.
- SDN centralizes control — **single point of failure** and single point of attack.
- MPLS guarantees traffic isolation within the provider network, but **does not provide confidentiality from the provider itself** — encrypt data over MPLS anyway.
- VPC isolation is **logical, not physical** — relevant to shared responsibility discussions.

## Cross-links

- [Secure Network Design](./secure-network-design.md)
- [Firewalls](./firewalls.md)
- [Network Devices](./network-devices.md)

## Sources

- destination-cissp §4.1.15–4.1.16 (pp. 0593–0602)
