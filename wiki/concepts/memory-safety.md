---
title: "Memory Safety and Buffer Overflow"
type: concept
domain: 8
tags: [buffer-overflow, aslr, bounds-checking, memory-safety, integer-overflow, stack, heap, toctou, memory-reuse]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Memory Safety and Buffer Overflow

## Definition

**Memory safety** refers to the property of software that prevents unintended access to, or modification of, memory. Memory-unsafe code is one of the most common and dangerous classes of vulnerability — it underlies buffer overflows, use-after-free bugs, and many other exploitable conditions.

---

## Buffer Overflow

### What It Is

A **buffer overflow** occurs when data sent to a storage buffer exceeds the buffer's allocated capacity. Buffers are temporary memory storage areas used by applications to handle input, processing, and output. Buffer sizes are typically fixed at design time and do not dynamically expand.

**Consequence:** When overflow occurs, excess data spills into adjacent memory. An attacker can craft the overflow payload to contain executable code, which may then be placed in executable memory regions, allowing the attacker to:
- Execute arbitrary code
- Elevate privileges on the system
- Crash the application (denial of service)

*(source: destination-cissp §8.5.2)*

### Stack vs. Heap Overflow

- **Stack overflow**: Overflow of the call stack (function call frames, local variables). Classic attack vector: overwrite the return address to redirect execution to attacker-controlled code.
- **Heap overflow**: Overflow of heap-allocated memory. Harder to exploit deterministically but still dangerous; can corrupt heap management structures.

### How Attackers Exploit It

1. Attacker identifies an application that accepts input without validating its length.
2. Attacker provides input longer than the buffer.
3. Overflow data (shellcode) overwrites adjacent memory (e.g., the saved return address on the stack).
4. When the function returns, execution jumps to the attacker's code.

---

## Mitigations

| Mitigation | How It Works | Notes |
|---|---|---|
| **ASLR (Address Space Layout Randomization)** | Randomizes the memory addresses where system executables, stack, heap, and libraries are loaded each time the program runs. | Makes it very difficult for an attacker to predict the memory layout needed to craft a reliable exploit. Industry-standard OS feature. |
| **DEP (Data Execution Prevention) / NX bit** | Marks memory regions (stack, heap) as non-executable. Even if attacker injects shellcode, the CPU will refuse to execute it. | Complements ASLR. Can be bypassed by return-oriented programming (ROP) attacks. |
| **Stack canaries** | A known value ("canary") is placed on the stack between local variables and the return address. Before returning, the compiler checks that the canary is intact. | If overflowed, the canary value is corrupted and the program terminates before executing attacker code. |
| **Bounds / parameter checking** | Validate that all inputs to a variable are within acceptable bounds before use (e.g., string length, array index, numeric range). | Preventive; catches overflow conditions before they occur. |
| **Safe programming languages** | Languages like Rust, Go, or Java manage memory automatically and perform bounds checks at runtime. C and C++ are unsafe by default. | Architectural mitigation; not always possible with legacy codebases. |
| **Safe library functions** | Use `strncpy` instead of `strcpy`, `snprintf` instead of `sprintf`, etc. | Prevents specific categories of overflow at the function call level. |
| **Code review and SAST** | Static analysis tools can detect calls to unsafe functions and missing bounds checks. | Process-level mitigation; catches issues before production. |
| **Runtime array/buffer bounds checking** | Language runtimes or compiler flags that check array accesses at runtime. | Adds overhead but catches overflows in test environments. |

> **Exam priority:** ASLR is the top mitigation to know by name. Bounds/parameter checking is the second.

---

## Related Memory Vulnerabilities

| Vulnerability | Description |
|---|---|
| **Integer overflow** | Arithmetic operation produces a result that exceeds the maximum value for the integer type, wrapping around to a small or negative number. Often used to bypass length checks before allocating a buffer. |
| **Use-after-free** | A program continues to use a memory pointer after the referenced memory has been freed/deallocated. Can lead to arbitrary code execution if attacker controls the reallocated memory. |
| **Null pointer dereference** | Code attempts to access memory through a null pointer. Usually causes a crash (denial of service). |
| **Format string vulnerability** | User-supplied input is passed directly to a format string function (e.g., `printf(user_input)` instead of `printf("%s", user_input)`). Can allow arbitrary memory read/write. |
| **Memory/object reuse** | Residual sensitive data from a previous operation remains in memory and can be read by a subsequent process. Mitigation: overwrite memory before reuse. *(destination-cissp §8.5.5)* |

---

## Exam-Relevant Nuance

- Buffer overflows are among the oldest and most prevalent application vulnerabilities; they remain relevant in languages like C/C++.
- ASLR alone does not prevent all buffer overflow attacks — attackers can use info-leak vulnerabilities to defeat ASLR, or use ROP chains to bypass DEP. The combination of ASLR + DEP + stack canaries + bounds checking provides defense-in-depth.
- The fix for a specific buffer overflow in deployed software is typically a **software patch** — ASLR is a systemic mitigation that makes exploitation harder, but the underlying bug still needs to be patched.
- Memory reuse (object reuse) is distinct from buffer overflow — it is about stale data lingering after a free, not about overflow.

---

## Cross-Links

- [Secure Coding Practices](secure-coding-practices.md) — general secure coding; TOCTOU, covert channels
- [Database Security](database-security.md) — SQL injection (input validation failure)
- [Software Testing Types](./security-testing-types.md) — DAST and fuzz testing for finding memory vulnerabilities

## Sources

- destination-cissp §8.5.1 (source-level vulnerabilities table), §8.5.2 (buffer overflow) (pp. 0965–0967)
- cissp-exam-outline (Domain 8.5: Security weaknesses and vulnerabilities at the source-code level)
