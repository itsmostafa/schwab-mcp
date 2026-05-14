---
title: "Directory Services"
type: concept
domain: 5
tags: [ldap, active-directory, directory, x500, sso, kerberos]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Directory Services

## Definition

Directory services provide a centralized repository for storing, managing, and authenticating identity information (users, groups, devices, policies). They are foundational to SSO and enterprise IAM.

## Key Technologies

### LDAP (Lightweight Directory Access Protocol)

LDAP is the protocol used to query and modify directory services. It is derived from the X.500 standard (simplified / "lightweight" version).

- Operates on a hierarchical, tree-structured database.
- **DN (Distinguished Name):** Uniquely identifies an entry in the directory (e.g., `CN=Alice,OU=Finance,DC=example,DC=com`).
- **OU (Organizational Unit):** A container within the directory hierarchy used to organize entries.
- **X.500 origin:** LDAP was designed as a lighter alternative to the full X.500 DAP (Directory Access Protocol).

### Active Directory (AD)

Microsoft's implementation of a directory service, using LDAP + Kerberos + DNS for full SSO capabilities.

- Required for Kerberos in a Windows environment (destination-cissp §5.2.8).
- Stores user accounts, group memberships, policies, and computer objects.
- Foundational for enterprise SSO within Windows ecosystems.

## Relationship to Access Control

Directory services underpin centralized access control:
- **Centralized administration model** — a single directory stores all identities.
- A compromise of the directory service = compromise of all identities (single point of failure).
- Enables Group Policy (enforcing least privilege, password policies, screen lock, etc.).

## Cross-links

- [Kerberos](kerberos.md) — uses AD for Windows SSO
- [SSO](sso.md)
- [Federation](federation.md) — IDaaS syncs with local directories
- [Identity Lifecycle](identity-lifecycle.md)

## Sources

- destination-cissp §5.2.8 (AD + Kerberos context, p. 705)
- destination-cissp §5.3.1 (IDaaS synced identity from AD)
- cissp-exam-outline §5.2, §5.3
