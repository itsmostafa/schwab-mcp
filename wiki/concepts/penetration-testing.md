---
title: "Penetration Testing"
type: concept
domain: 6
tags: [penetration-testing, red-team, blue-team, purple-team, black-box, white-box, grey-box, blind, double-blind]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Penetration Testing

## Definition

A penetration test (pen test) is a controlled, authorized simulation of an attack against a target system, network, or application. Unlike a vulnerability assessment, a pen test **includes an exploitation phase** — testers actively attempt to compromise identified vulnerabilities to confirm they are real (true positives). The key differentiating step is exploitation. See [Assessment vs Audit](assessment-vs-audit.md).

**Approval is mandatory.** Without documented authorization, a pen test is indistinguishable from criminal activity. An NDA is often required because testers may access sensitive data during testing.

## Phases

The source (destination-cissp §6.2.2) describes five phases:

1. **Reconnaissance** — Passive information gathering using publicly available data: DNS queries, WHOIS, LinkedIn, job boards (Indeed), forums. The target cannot detect this activity. Also known as Open Source Intelligence (OSINT).

2. **Enumeration** — Active scanning of IP ranges, open ports (65,536 TCP and UDP ports, 0–65,535), hostnames, and user accounts. The target *can* detect this activity. Port → service mapping (e.g., port 80 = HTTP/web server). Narrows the attack surface to specific systems and services.

3. **Vulnerability Analysis** — Identification of weaknesses in enumerated targets. This is the **fork**: vulnerability assessments stop here (document + report); pen tests continue.

4. **Execution/Exploitation** — Pen tests only. Actively attempt to exploit identified vulnerabilities to confirm they are true positives. Manual skill is the primary driver of quality here.

5. **Document Findings/Reporting** — Applicable to both vulnerability assessments and pen tests. Report contains: techniques tested, tools used, what worked/failed, identified vulnerabilities, mitigation steps, and prioritization of findings. Minimizing false positives in the report is as important as the testing itself.

## Testing Variables

### Perspective (Where the test originates)

| Internal | External |
|---|---|
| Performed from inside the corporate network | Performed from outside (internet-facing) |
| Tests insider threat scenarios and lateral movement | Tests perimeter defenses and defense-in-depth |

### Approach (Who knows about it)

| Blind | Double-Blind |
|---|---|
| Tester has minimal target information | Tester has minimal info AND internal security/IT teams do not know a test is coming |
| Internal IT/security may know a test is scheduled | Only senior management (who commissioned it) knows |
| Tests tester skill | Tests both tester skill and the team's incident detection/response |

### Knowledge (How much information the tester has)

| Black Box (Zero Knowledge) | Gray Box (Partial Knowledge) | White Box (Full Knowledge) |
|---|---|---|
| Tester starts with nothing — same as blind | Tester has some network/system info | Tester has full knowledge: IP ranges, diagrams, password policies |
| Mirrors real-world external attacker | Balanced; simulates insider with limited access | Most thorough; maximizes coverage |

## Red, Blue, and Purple Teams

| Team | Role |
|---|---|
| **Red Team** | Offensive. Simulates attacks using pen testing, social engineering, threat intelligence. |
| **Blue Team** | Defensive. Evaluates security posture, recommends mitigations, monitors, performs incident response, runs security operations. |
| **Purple Team** | Not a distinct team — a collaboration between red and blue. Promotes information sharing and joint exercises. Purely adversarial red-vs-blue dynamics reduce overall organizational security; purple team collaboration avoids this. |

## Additional Testing Concepts

### Test Types (by user intent)

| Type | Description |
|---|---|
| **Positive testing** | Normal usage scenarios; confirms system behaves as expected under correct inputs. |
| **Negative testing** | Introduces normal errors (wrong password, invalid input); confirms system fails gracefully without crashing. |
| **Misuse testing** | Simulates a malicious user/attacker trying to break the system; most adversarially realistic. |

### Boundary Value Analysis and Equivalence Partitioning

Used to make testing efficient rather than exhaustive (destination-cissp §6.2.1):

- **Boundary value analysis** — identifies behavioral boundaries (e.g., password length changes from rejected to accepted at 8 chars and again at 16 chars); focuses testing at the boundaries where bugs are most likely.
- **Equivalence partitioning** — groups inputs exhibiting the same behavior (partitions); tests a representative value from each partition rather than every possible input.

## Cross-links

- [Assessment vs Audit vs Penetration Test](assessment-vs-audit.md)
- [Vulnerability Assessment](vulnerability-assessment.md)
- [Security Testing Types](security-testing-types.md)

## Sources

- destination-cissp §6.2.2 (phases, perspective, approach, knowledge, red/blue/purple teams)
- destination-cissp §6.2.1 (positive/negative/misuse testing, boundary value, equivalence partitioning)
- cissp-exam-outline (domain 6 subtopic list)
