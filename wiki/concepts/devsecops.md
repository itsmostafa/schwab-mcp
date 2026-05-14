---
title: "DevSecOps"
type: concept
domain: 8
tags: [devops, devsecops, secdevops, ci-cd, agile, integrated-product-team, shift-left]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# DevSecOps

## Definition

**DevOps** unifies three previously siloed teams — software development, operations, and quality assurance (QA) — into a single integrated, collaborative process. **DevSecOps** (and preferably **SecDevOps**) extends DevOps by making security an integral part of that process from the beginning, rather than bolting it on at the end.

> *destination-cissp* notes that the **Integrated Product Team (IPT)** is essentially a formal name for the DevOps concept — a team of skilled professionals each bringing specific expertise, collectively committed to delivering a secure and functional product throughout its entire life cycle.

---

## DevOps vs. DevSecOps

| Aspect | Traditional Approach | DevOps | DevSecOps / SecDevOps |
|---|---|---|---|
| **Team structure** | Separate dev, ops, QA teams; hand-offs between them. | Unified dev + ops + QA team from the start. | Unified team + security embedded from inception. |
| **Security timing** | Security is "bolted on" at the end. | Security often still added late; fast iteration conflicts with slow security reviews. | Security is *planned for* and integrated at every step. |
| **Velocity** | Slow; sequential phases. | Fast; continuous iteration. | Fast; automated security testing keeps pace. |
| **Security testing** | Manual, periodic penetration tests. | Traditional security techniques (pen tests) are too slow for rapid iteration. | Automated SAST/DAST in CI/CD pipeline; continuous security. |

**Key problem with missing security in DevOps:** Traditional security techniques (e.g., periodic penetration tests, manual security analysis) are too slow for rapid DevOps iteration. DevSecOps solves this by automating security checks.

---

## Components of a DevSecOps Approach

*(source: destination-cissp §8.1.5)*

1. **Plan for security** — identify security requirements at the start of every sprint/iteration.
2. **Strong engagement** between developers, operations, and security teams.
3. **Engage developers** — security training; make developers security-aware.
4. **Develop using secure techniques and frameworks** — secure coding standards, approved libraries.
5. **Automate security testing** — integrate SAST and DAST into CI/CD pipeline so every code commit is automatically tested.
6. **Use traditional techniques sparingly** — reserve manual penetration testing and deep-dive analysis for periodic reviews, not every release.

---

## Canary Testing and Deployments

**Canary testing** is a deployment strategy where new code or features are released to a small subset of users before a full rollout. This provides an early warning (like the "canary in the coal mine") — if problems are identified with the small group, they can be fixed before affecting all users.

- Limits blast radius of bugs and security issues.
- Allows real-world testing at production scale in a controlled way.
- Compatible with DevSecOps continuous delivery pipelines.

**Smoke testing** is a related technique: quick preliminary testing after a change to identify simple failures in the most important existing functionality ("did the most basic things still work after this change?").

---

## Exam-Relevant Nuance

- **IPT = DevOps** in ISC2 terminology. Recognize both names.
- **DevSecOps vs. SecDevOps**: The source uses both; "SecDevOps" emphasizes security leading the process. Both mean security is integral from the start.
- Agile's Scrum Master manages team efficiency; the combination of Agile + DevSecOps produces the ideal fast, secure development environment.
- Separate teams naturally lead to security being tacked on at the end — the most common failure mode.
- Combining methodologies (e.g., Agile + structured) is fine; DevSecOps applies to any combination.

---

## Cross-Links

- [SDLC](sdlc.md) — full SDLC phases and methodology context
- [CI/CD Security](ci-cd-security.md) — specifics of securing the CI/CD pipeline
- [Secure Coding Practices](secure-coding-practices.md) — what developers in a DevSecOps team apply
- [Software Testing Types](./security-testing-types.md) — SAST/DAST automated in DevSecOps pipeline

## Sources

- destination-cissp §8.1.5–8.1.6 (pp. 0941–0944)
- cissp-exam-outline (Domain 8: Development methodologies, DevOps, DevSecOps)
