---
title: "Cloud Security Models"
type: concept
domain: 3
tags: [cloud, saas, paas, iaas, caas, faas, shared-responsibility, deployment-models, multitenancy]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Cloud Security Models

Cloud computing enables on-demand access to computing resources (compute, storage, networking, software) over the internet on a pay-as-you-go basis. Security in the cloud requires understanding both the service model (what is shared) and the deployment model (who owns the infrastructure).

## Defining Characteristics of Cloud Computing

| Characteristic | Description |
|---|---|
| **On-demand self-service** | Resources provisioned immediately and automatically without provider intervention |
| **Broad network access** | Accessible from anywhere via any device with internet access |
| **Resource pooling** | Compute, storage, and networking shared among multiple customers (multi-tenancy) |
| **Rapid elasticity/scalability** | Resources scale up or down quickly, often automatically |
| **Measured service** | Usage tracked precisely; customers pay only for what they use |
| **Multi-tenancy** | Multiple customers share the same physical infrastructure — primary security concern in public cloud |

Note: Multi-tenancy does not apply to private cloud.

## Cloud Service Models

| Model | Acronym | What provider manages | What customer manages | Example |
|---|---|---|---|---|
| **Software as a Service** | SaaS | Everything (physical, OS, middleware, app, data storage) | User accounts and access control | Office 365, Gmail |
| **Platform as a Service** | PaaS | Physical, OS, middleware, runtime | Applications and data | AWS Elastic Beanstalk |
| **Infrastructure as a Service** | IaaS | Physical hardware, networking, hypervisor | OS, applications, data, networking config | AWS EC2, Azure VMs |
| **Containers as a Service** | CaaS | Physical, OS, container engine | Containerized applications | AWS ECS |
| **Function as a Service** | FaaS | Everything (serverless) | Code/functions only; pay per invocation | AWS Lambda |

### Shared Responsibility Summary

In every public cloud model, the provider manages physical security, hardware, and hypervisors.

- **SaaS**: Provider manages everything except user account management and access control (shared)
- **PaaS**: Customer manages applications and data
- **IaaS**: Customer has most control — OS, networking config, apps, data

**Critical rule**: The cloud customer is **always accountable** for their data, regardless of service model. Responsibility can be delegated (via SLA); accountability cannot.

### Cloud Roles

| Role | Definition |
|---|---|
| **Cloud customer/consumer** | Purchases and uses cloud services; always accountable |
| **Cloud service provider** | Sells and operates cloud infrastructure |
| **Cloud broker** | Aggregates services from multiple providers; resells to customers (service arbitrage) |
| **Data controller** | Defines how data should be protected (= cloud customer) |
| **Data processor** | Processes data per the controller's rules (= cloud provider) |

## Cloud Deployment Models

| Model | Infrastructure managed by | Infrastructure owned by | Accessible by |
|---|---|---|---|
| **Public** | Third-party provider | Third-party provider | Everyone (untrusted) |
| **Private** | Organization or provider | Organization or provider | Single organization (trusted) |
| **Community** | Organization or provider | Organization or provider | Specific group (e.g., hospitals, government) |
| **Hybrid** | Both | Both | Both trusted and untrusted |

- **GovCloud** (AWS): Largest example of a community cloud — FedRAMP-compliant for US government agencies
- **Private cloud**: Can be on-premises (org owns hardware) or off-premises (provider hosts dedicated environment)
- **Hybrid cloud**: Commonly used to keep high-sensitivity data on-premises/private while using public cloud for lower-sensitivity workloads

## Key Security Considerations

### Data Protection
- Encrypt data **before** migrating to the cloud (encrypt locally, then upload)
- Use strong access controls regardless of service model
- Understand where data physically resides (data residency/sovereignty requirements)

### Identity and Access
- Cloud IAM requires careful design for least privilege and SoD
- **IDaaS** (Identity as a Service): Cloud-based IAM solutions extending traditional AD/LDAP
- **Federated Identity Management (FIM)**: Protocols (SAML, OAuth, SPML) enabling identity portability across organizations
- **Data controller** always owns accountability; **data processor** handles operational access

### Cloud Migration Risks
- **Vendor lock-in**: Once migrated, moving providers is costly — use multiple providers or open standards to mitigate
- **CapEx to OpEx shift**: Capital expenditure becomes operational expense; efficiency benefit but ongoing cost
- **Availability dependency**: Requires reliable high-speed internet; single point of failure if no redundancy

### Forensics in the Cloud
Cloud forensics is more complex than on-premises:
- Physical access to hardware is typically unavailable
- Investigators request VM snapshots and virtual disk images instead
- SaaS: Must rely entirely on CSP for evidence
- IaaS: Can perform forensic analysis on own VMs; may need CSP cooperation for network traffic
- NIST published "Cloud Computing Forensic Science Challenges" (August 2020)

## Exam-Relevant Nuance

- SaaS = provider manages everything; customer manages users and access.
- IaaS = customer manages OS and above; provider manages physical and hypervisor.
- The cloud customer is **always accountable** — this cannot be delegated.
- Multi-tenancy is a characteristic only of public cloud; private cloud has no multi-tenancy.
- FaaS/serverless: No servers to provision — pay only when functions are invoked.
- CaaS enables DevOps agility: containers can be deployed quickly across environments.

## Cross-links

- [Virtualization Security](virtualization-security.md) — hypervisors, VMs, containers in cloud
- [ICS/SCADA Security](ics-scada-security.md) — OT/ICS in cloud contexts
- [IoT Security](iot-security.md) — IoT and cloud-connected devices
- [Secure Design Principles](secure-design-principles.md) — shared responsibility as a design principle

## Sources

- destination-cissp §3.5.9–3.5.14 (pp. 319–341)
- cissp-exam-outline §3.5
