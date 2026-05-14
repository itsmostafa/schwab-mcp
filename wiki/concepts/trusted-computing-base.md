---
title: "Trusted Computing Base (TCB)"
type: concept
domain: 3
tags: [tcb, reference-monitor, security-kernel, tpm, ring-protection, privilege-levels, process-isolation]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Trusted Computing Base (TCB)

The Trusted Computing Base (TCB) is **the totality of all protection mechanisms within an architecture** — every piece of hardware, firmware, and software that contributes to enforcing the security policy of a system. If any component of the TCB is compromised, the security of the entire system is at risk.

## TCB Components

A TCB can include:
- Processors (CPUs)
- Memory (primary and secondary storage)
- Virtual memory
- Firmware
- Operating systems
- System kernel
- Security policies and procedures
- Access control mechanisms

For an enterprise, the TCB can be massive — encompassing training programs, change management processes, and the full network.

## Subjects and Objects

All security within information systems operates in terms of subjects and objects:

| Term | Definition |
|---|---|
| **Subject** | Active entity — a person, process, or program actively trying to access something |
| **Object** | Passive entity — anything being accessed (file, server, hardware component, process) |

## Reference Monitor Concept (RMC)

The RMC is the **concept** (not an implementation) that a subject accesses an object only through mediation based on a set of rules, with that access being logged and monitored.

RMC features:
- Must mediate **all** access
- Must be **protected from modification**
- Must be **verifiable** as correct
- Must **always be invoked** (cannot be bypassed)

## Security Kernel

The security kernel is the **implementation** of the Reference Monitor Concept. This is a critical distinction:

| Term | Nature |
|---|---|
| Reference Monitor Concept | Theory / concept |
| Security Kernel | Implementation of the RMC |
| System Kernel | Core of the OS — different from the security kernel |

A viable security kernel must have three properties:

| Property | Description |
|---|---|
| **Completeness** | Impossible to bypass the mediation — subjects must always go through it |
| **Isolation** | Mediation rules are tamper-proof; only authorized parties can change them |
| **Verifiability** | Logging, monitoring, and testing confirm the kernel is functioning correctly |

## System Kernel vs. Security Kernel

- **System kernel**: Core of the operating system; has complete control over all system components; runs in privileged/supervisor mode.
- **Security kernel**: Implementation of the RMC; mediates access between subjects and objects.

These are **not the same thing** — a common exam trap.

## Privilege Levels and Ring Protection Model

CPUs operate in **processor states** that restrict operations:

| State | Also called | Description |
|---|---|---|
| **Supervisor state** | Kernel mode | Highest privilege; full access to all CPU capabilities; system kernel runs here |
| **Problem state** | User mode | Lower privilege; limited CPU instruction access; normal application mode |

The **Ring Protection Model** (Figure 3-17 in source) formalizes these levels:

| Ring | Trust Level | Typical Content |
|---|---|---|
| Ring 0 | Highest / most trusted | Firmware, OS kernel, critical system processes |
| Ring 1–2 | Intermediate | OS services, device drivers |
| Ring 3 | Lowest / least trusted | User applications |

Each ring communicates with adjacent rings via system calls. Outer rings can only access inner rings through controlled, trusted interfaces.

## Process Isolation

Process isolation prevents one process from affecting another's memory or resources. Two primary methods:

| Method | Mechanism |
|---|---|
| **Memory Segmentation** | RAM is divided into segments assigned to each application; one application cannot access another's segment |
| **Time-Division Multiplexing** | CPU allocates very small time slices to each process; each process runs in isolation during its time slot |

## Storage Types

| Type | Also called | Characteristics | Examples |
|---|---|---|---|
| **Primary storage** | Volatile memory | Fast, small, temporary (lost on power off) | RAM, cache, CPU registers |
| **Secondary storage** | Non-volatile memory | Slow, large, persistent | HDDs, SSDs, tapes, optical media |

**Virtual memory (paging)**: When RAM fills up, the OS moves inactive data to a "paging file" on secondary storage and restores it when needed. This prevents crashes but can cause latency.

## Firmware

Firmware is low-level software that controls hardware — it boots hardware and brings it online. Modern firmware is updatable (not hard-coded), which introduces vulnerability: firmware can now be attacked and modified.

## Trusted Platform Module (TPM)

See dedicated page: [TPM](tpm.md).

Summary: A TPM is a hardware chip (ISO/IEC 11889) that performs cryptographic operations — key generation and storage — and provides platform integrity verification at boot time. Key concepts:
- **Binding**: Data encrypted to a specific TPM's endorsement key; only that TPM can decrypt it
- **Sealing**: Data encrypted with conditions (e.g., specific system state) that must be met before decryption is allowed

## Exam-Relevant Nuance

- TCB = all protection mechanisms (totality); the security kernel is one component within the TCB.
- Reference Monitor Concept is a concept; Security Kernel is the implementation. Do not confuse.
- Ring 0 is the most trusted ring; Ring 3 is the least trusted.
- The three properties of a security kernel (completeness, isolation, verifiability) are directly tested.
- Process isolation prevents covert channel exploitation between processes sharing memory.

## Cross-links

- [Security Models](security-models.md) — formal models that TCB implements
- [TPM](tpm.md) — hardware root of trust
- [Common Criteria](common-criteria.md) — evaluation of TCB and related systems
- [Virtualization Security](virtualization-security.md) — how hypervisors fit into the TCB

## Sources

- destination-cissp §3.4.1–3.4.10 (pp. 279–297)
- cissp-exam-outline §3.4
