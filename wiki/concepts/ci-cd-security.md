---
title: "CI/CD Security"
type: concept
domain: 8
tags: [ci-cd, devsecops, supply-chain, sbom, code-signing, pipeline-security, scm, version-control]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# CI/CD Security

## Definition

**CI/CD (Continuous Integration / Continuous Delivery or Deployment)** is a set of practices that automate the steps of integrating code changes, testing them, and releasing them to users. Securing the CI/CD pipeline is a critical component of DevSecOps and software supply chain security.

> **Coverage note:** destination-cissp covers CI/CD at a high level (§8.2.1, §8.2.5). Supply chain attack specifics and SBOM are not detailed in the source; they are included here from cissp-exam-outline scope with that attribution noted.

---

## CI/CD Definitions

*(source: destination-cissp §8.2.1)*

| Stage | Definition |
|---|---|
| **Continuous Integration (CI)** | Automates the process of committing code to a shared repository and running automated tests. Allows code changes to be frequently integrated and immediately tested. |
| **Continuous Delivery** | Extends CI by also automating the release of validated code changes into a staging/pre-production repository. Code is always in a releasable state. |
| **Continuous Deployment** | Extends Continuous Delivery by automatically pushing code to production — without human intervention — provided all tests pass. |

**Key concept:** If any step in the CI/CD pipeline fails (test failure, security scan failure), the code change is returned to the developer for remediation. This is how automated security checks integrate naturally into the pipeline.

---

## Security Controls in the CI/CD Pipeline

| Control | Description |
|---|---|
| **SAST in pipeline** | Static analysis tool runs automatically on every code commit; fails the build if critical issues are detected. Enforces secure coding standards. |
| **DAST in pipeline** | Automated dynamic testing against a deployed test environment as part of the pipeline. |
| **Secrets management** | Credentials, API keys, and certificates must never be hardcoded in source code or pipeline scripts. Use dedicated secrets management tools (HashiCorp Vault, AWS Secrets Manager). |
| **Dependency scanning (SCA)** | Software Composition Analysis tools scan third-party libraries and open source dependencies for known vulnerabilities (CVEs). |
| **Container image scanning** | If building containerized applications, scan base images and final images for vulnerabilities before deployment. |
| **Infrastructure as Code (IaC) scanning** | Scan IaC templates (Terraform, CloudFormation) for misconfigurations before provisioning. |
| **Access controls on repositories** | Principle of least privilege on code repositories; branch protection rules; mandatory code review before merge. |
| **Audit logging** | All pipeline actions (code commits, builds, deployments) logged for review and forensic purposes. |

---

## Software Configuration Management (SCM)

*(source: destination-cissp §8.2.1)*

SCM focuses specifically on managing changes in software as part of overall configuration/change management:
- Establishes and maintains **baselines** (approved versions of code)
- **Revision control** — tracks all changes with who made them and when
- **Build and process management** — consistent, repeatable build procedures
- **Facilitates teamwork** — prevents conflicting changes; supports parallel development

Popular systems: Git, SVN, Mercurial. Most organizations use Git (with GitHub, GitLab, or Bitbucket as hosted platforms).

---

## Code Repositories

*(source: destination-cissp §8.2.1)*

A **code repository** is the authoritative storage location for source code. Modern repositories provide:
- **Versioning and release control** — tag releases; roll back to previous versions
- **Code review** — pull request / merge request workflows; require approvals before merge
- **Bug tracking and issue management**
- **Document management**
- **Patch distribution**

Security concerns for repositories:
- Access control (who can read, write, merge, delete branches or tags)
- Protection of the default branch (prevent direct pushes; require reviews)
- Signed commits (verify code authorship)
- Detection of secrets accidentally committed

---

## Supply Chain Security

> **Coverage gap:** destination-cissp does not detail software supply chain attacks. Content below reflects cissp-exam-outline scope.

The **software supply chain** includes all components, libraries, tools, and processes used to build and deliver software. Attackers increasingly target the supply chain rather than the final application.

**Notable example:** SolarWinds (2020) — attackers compromised SolarWinds' build system, inserting malicious code into a digitally signed software update that was distributed to ~18,000 organizations.

| Supply Chain Attack Vector | Mitigation |
|---|---|
| Compromised build system | Isolate and harden CI/CD infrastructure; monitor build processes; reproducible builds. |
| Malicious dependency (typosquatting, dependency confusion) | Dependency scanning (SCA); pin dependency versions with hashes; use private registries. |
| Compromised package registry | Verify package integrity (hashes); use SBOM; prefer well-maintained packages. |
| Backdoored updates | Code signing of all released artifacts; verify signatures before installation. |
| Compromised developer credentials | MFA for all repository access; audit logs; secrets rotation. |

---

## Code Signing

**Code signing** is the practice of using a digital certificate to sign software artifacts (executables, packages, container images, scripts). Recipients can verify:
1. The artifact was produced by the claimed publisher (authenticity).
2. The artifact has not been modified since signing (integrity).

Part of the supply chain security chain of trust. Necessary but not sufficient — if the signing infrastructure itself is compromised (as in SolarWinds), signed code can still be malicious.

---

## SBOM (Software Bill of Materials)

An **SBOM** is a formal inventory of all components in a software product — libraries, frameworks, dependencies, and their versions. Enables:
- Rapid identification of affected products when a CVE is published for a component
- Supply chain risk management
- Compliance verification (e.g., license auditing)

Formats: SPDX, CycloneDX. Increasingly required in government and critical infrastructure contracts.

---

## Exam-Relevant Nuance

- **CI/CD pipeline = automated build, test, deploy cycle**. SAST/DAST integrated here is the DevSecOps "shift left" in action.
- **Secrets in code** is a common real-world failure. Pipeline secrets management is a key control.
- **SCM** (Software Configuration Management) ensures reproducibility and accountability for all code changes.
- **Code signing** provides integrity assurance for software distribution — without it, distributed software cannot be verified.
- **Supply chain attacks** target the toolchain, not just the application — an important CISSP exam concept.

---

## Cross-Links

- [DevSecOps](devsecops.md) — DevSecOps philosophy; CI/CD as the automation backbone
- [SDLC](sdlc.md) — CI/CD as a methodology within SDLC
- [Software Acquisition Security](software-acquisition-security.md) — SBOM; third-party dependency risk
- [Change Management Software](change-management-software.md) — version control and change governance
- [Software Testing Types](./security-testing-types.md) — SAST/DAST in pipeline context

## Sources

- destination-cissp §8.2.1 (CI/CD, SCM, code repositories) (pp. 0947–0948)
- destination-cissp §8.2.5 (development ecosystems: CI/CD, SOAR, SCM) (p. 0959)
- cissp-exam-outline (Domain 8.2: CI/CD, code repositories, software configuration management)
- Supply chain attack content: cissp-exam-outline scope (coverage gap in destination-cissp noted)
