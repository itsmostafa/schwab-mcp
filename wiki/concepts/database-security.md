---
title: "Database Security"
type: concept
domain: 8
tags: [database, dbms, sql, acid, concurrency, polyinstantiation, inference, aggregation, rdbms, primary-key, foreign-key]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Database Security

## Definition

**Database security** encompasses the controls, processes, and technologies that protect database management systems (DBMS) and the data they contain from unauthorized access, disclosure, modification, and destruction. Security must exist within the database itself and within the applications that access it.

---

## DBMS Components

*(source: destination-cissp §8.2.3)*

| Component | Description | Security Concern |
|---|---|---|
| **Hardware** | Dedicated server with RAID storage, redundant power/cooling/network. | Physical access control; hardware tampering; availability. |
| **Software** | Operating system + DBMS application (e.g., Oracle, MySQL, PostgreSQL, SQL Server). | Patch management; application hardening; configuration review. |
| **Language (SQL)** | Structured Query Language — the interface for interacting with data. Dialects: T-SQL, MySQL, PostgreSQL, SQLite. | SQL injection attacks; improper query construction. |
| **Users** | People who access data, either through a UI or directly via SQL queries. | Access control; principle of least privilege; separation of duties. |
| **Data** | The actual records stored in the database — often sensitive (PII, financial, health). | Encryption at rest; data masking; access controls; integrity controls. |

---

## Relational Database Concepts

**RDBMS (Relational Database Management System)** stores data in two-dimensional tables (relations) that can be linked together to drive inference and decision-making.

| Term | Meaning |
|---|---|
| **Tuple** | A single row in a table. |
| **Attribute** | A single column in a table (also called a field). |
| **Field** | The intersection of a row (tuple) and a column (attribute) — a single data value. |
| **Primary key** | One or more columns whose values uniquely identify each row within a table. Every table must have a primary key. |
| **Foreign key** | One or more columns in a table that reference the primary key of another table, creating a relationship between the two tables. |
| **Referential integrity** | Ensures that every foreign key value in one table corresponds to a valid primary key in the referenced table. Enforced by RDBMS automatically. |

---

## ACID Properties

ACID defines how transactions in an RDBMS must behave to maintain data integrity, especially in concurrent environments:

| Letter | Property | Meaning |
|---|---|---|
| **A** | **Atomicity** | All changes in a transaction take effect, or none at all. No partial transactions. |
| **C** | **Consistency** | A transaction must bring the database from one valid state to another — all rules and constraints are satisfied. |
| **I** | **Isolation** | A transaction's intermediate state is invisible to other concurrent transactions until it is committed. |
| **D** | **Durability** | Once a transaction is committed, it is permanent — even system failures cannot undo it. |

---

## Concurrency and Lock Controls

**Concurrency**: The ability for multiple processes to access or modify shared data simultaneously.

**Problem**: If two users try to update the same record simultaneously, data corruption (integrity loss) can occur.

**Solution — Locking**:
- When User A accesses and begins modifying a record, it is **locked**.
- User B can view the record but cannot modify it until User A completes their transaction and the lock is released.
- Prevents lost updates and dirty reads.

---

## Database Attacks

### SQL Injection

The most significant database attack vector — an attacker injects SQL code into an input field that is incorporated into a database query, manipulating the query to return unauthorized data, modify data, or execute operating system commands.

**Types:**
| Type | How It Works |
|---|---|
| **In-band** | Results returned in the same channel as the injection (classic; easy to exploit). Includes error-based and union-based. |
| **Blind** | No direct output; attacker infers information from application behavior (true/false conditions or timing). Includes Boolean-based blind and time-based blind. |
| **Out-of-band** | Results exfiltrated via a different channel (DNS lookup, HTTP request). Requires specific database features. |

**Prevention:**
- **Parameterized queries (prepared statements)** — the query structure is fixed; user input is always treated as data, never as SQL code. This is the primary defense.
- **Stored procedures** (properly implemented) — similar to parameterized queries.
- **Input validation** — validate type, length, and content of all inputs.
- **Principle of least privilege** — database accounts used by applications should have only the permissions they need (no DROP TABLE access for a read-only report).
- **WAF (Web Application Firewall)** — can detect and block common injection patterns in transit.

---

## Inference and Aggregation Attacks

These attacks allow an attacker to derive sensitive information without direct access to it:

| Attack | Description | Mitigation |
|---|---|---|
| **Inference** | Combining lower-classification data points to logically deduce higher-classification information. E.g., inferring a classified project exists from publicly available data. | Polyinstantiation; data labeling; query access controls. |
| **Aggregation** | Combining individually non-sensitive data elements that collectively reveal sensitive information. E.g., name + employer + salary + address can identify a private individual. | Restrict combination queries; row-level security; data minimization. |

---

## Polyinstantiation (Database Context)

Polyinstantiation allows the same named object to exist as multiple independent instances at different security classification levels. Lower-privileged users see the lower-classification version; they cannot infer the existence of higher-classification data.

**Key use case**: Multi-Level Security (MLS) database environments where different users have different clearance levels.

*See also: [Secure Coding Practices](secure-coding-practices.md) for the general concept and military example.*

---

## Data Protection Techniques

| Technique | Purpose |
|---|---|
| **Encryption at rest** | Protects data stored on disk from physical theft or unauthorized OS-level access. |
| **Encryption in transit** | TLS between application and database server prevents eavesdropping. |
| **Data masking** | Replaces sensitive values with realistic but fake data (e.g., in test/dev environments). |
| **Tokenization** | Replaces sensitive data with a non-sensitive token; original data held in secure token vault. |
| **Database activity monitoring** | Logs and alerts on database queries; detects anomalous access patterns. |

---

## Code Obfuscation (Related Concept)

**Code obfuscation** hides or obscures source code to make it difficult to reverse-engineer:

| Type | What It Does |
|---|---|
| **Lexical** | Modifies the appearance of code (removes comments, debugging info, reformats). Easiest but weakest. |
| **Data** | Modifies data structures within the code. |
| **Control flow** | Reorders statements, methods, loops; inserts irrelevant conditional branches. Most effective. |

> **BCM caveat (exam-relevant):** If code is obfuscated, restoring it after a disaster may be impossible without the original. Organizations must maintain a **software vault** with unaltered mission-critical source code as part of their BCM plan.

---

## Exam-Relevant Nuance

- **Column = attribute; Row = tuple.** The exam uses both terminologies interchangeably.
- **Primary key uniquely identifies each row**; foreign key references a primary key in another table.
- **ACID = Atomicity, Consistency, Isolation, Durability**. Each term has a specific meaning.
- **Parameterized queries are the primary defense against SQL injection** — not input filtering alone (filtering can be bypassed).
- **Polyinstantiation prevents inference**, not unauthorized access. Different authorization levels see different data instances for the same named object.
- **Locks prevent write-write conflicts**; they do not prevent read access by other users in most implementations.

---

## Cross-Links

- [Secure Coding Practices](secure-coding-practices.md) — polyinstantiation concept; input validation
- [Memory Safety](memory-safety.md) — memory vulnerabilities
- [OWASP Top 10](owasp-top-10.md) — injection (A03), broken access control (A01) are database-relevant
- [Software Testing Types](./security-testing-types.md) — DAST for detecting SQL injection

## Sources

- destination-cissp §8.2.3 (DBMS, concurrency, locks, ACID) (pp. 0951–0958)
- destination-cissp §8.5.1 (source-level vulnerability table) (p. 0965)
- destination-cissp §8.5.4 (polyinstantiation) (pp. 0970–0971)
- cissp-exam-outline (Domain 8.2, 8.5)
