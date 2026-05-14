---
title: Public Key Infrastructure (PKI)
type: concept
domain: 3
tags: [PKI, certificate-authority, X509, CRL, OCSP, digital-certificate, certificate-pinning, trust]
sources: [destination-cissp]
updated: 2026-05-13
---

# Public Key Infrastructure (PKI)

## Definition

PKI is the complete suite of technology, processes, and policies that enables secure distribution of public keys and verification of key owners' identities. Without PKI, you can still encrypt data but cannot verify the identity of the other party.

## Why PKI Is Necessary

Sending a raw public key over the network creates a risk: a MITM could substitute their own public key. PKI solves this by having a **trusted third party (CA)** cryptographically vouch for the binding between an entity and their public key via a **digital certificate**.

## Digital Certificate

A digital certificate binds an **entity's identity** to their **public key**. The CA signs the certificate with its private key — anyone with the CA's public key (pre-installed in browsers and OSes) can verify the certificate's authenticity.

### X.509 Standard

All CAs produce certificates conforming to the **X.509** standard. Key fields:
- Certificate version
- Serial number
- Issuing CA name
- Validity period (not before / not after)
- Subject distinguished name
- Subject's public key
- Signature algorithm identifier
- CA's digital signature

## PKI Components

| Component | Role |
|---|---|
| **Certificate Authority (CA)** | Root of trust; signs certificates with its private key |
| **Registration Authority (RA)** | Performs identity proofing on behalf of the CA |
| **Intermediate / Issuing CA** | Issues certificates on behalf of the root CA; root CA is kept offline |
| **Certificate Database** | Stores all issued certificates and the revocation list |
| **Certificate Store** | On the user's machine; stores the user's private key and certificates |
| **Validation Authority (VA)** | Responds to CRL/OCSP queries |

## CA Hierarchy (Chain of Trust)

```
Root CA (self-signed; kept OFFLINE — root of trust)
    └── Intermediate CA (subordinate CA; signs sub-CAs or issuing CAs)
            └── Issuing CA (signs entity-level certificates)
                    └── Entity Certificate (Alice, example.com, etc.)
```

- Root CA's private key is the foundation of the entire system. If compromised, the entire PKI collapses.
- **Best practice**: keep Root CA **offline**; use Intermediate CAs for day-to-day certificate issuance.
- If an intermediate CA is compromised, only the certificates it issued need to be revoked — far less damage than a compromised root.

## Certificate Lifecycle

| Phase | Description |
|---|---|
| **Enrollment** | Entity generates key pair; submits CSR (Certificate Signing Request) containing public key and identity info |
| **Issuance** | RA performs identity proofing; intermediate/issuing CA signs the certificate per X.509 |
| **Validation** | Browser/system queries CA to confirm certificate is valid (not expired, not revoked) |
| **Revocation** | Certificate invalidated if private key compromised or enrollment error |
| **Renewal** | Certificates expire (typically 12-month cycles); renewal confirms original CSR info |

## Revocation Confirmation Methods

| Method | Mechanism | Efficiency |
|---|---|---|
| **CRL** (Certificate Revocation List) | Client downloads the entire list of revoked certificate serial numbers from CA | Slow and bandwidth-heavy |
| **OCSP** (Online Certificate Status Protocol) | Client queries CA for status of a specific certificate; CA responds yes/no | Fast and efficient |

> Exam: OCSP is **newer and better** than CRL; CRL is the old, bulk-download method.

## Certificate Pinning

With certificate pinning, when a certificate from a server is trusted, **no new certificate request is made on subsequent visits**. Two implementation methods:
1. **Hardcoded in the application**: the certificate (or its hash) is embedded at build time.
2. **Browser pinning**: on first visit, browser pins the received certificate and uses it directly on future visits.

Prevents MITM certificate substitution attacks; eliminates key distribution risk for pinned connections.

## Cross-Certification and Bridge CA

- **Cross-certification**: two organizations' Root CAs mutually trust each other, enabling their certificate holders to trust each other.
- **Bridge CA**: a central CA that cross-certifies multiple organizations' CAs, creating a network of trust without requiring every pair to cross-certify.

## Exam Traps

- Root CA self-signs its own certificate; everything else traces back to it.
- Root CA should be **kept offline** to protect the root of trust.
- CRL = big list download; OCSP = one-certificate query.
- Certificate pinning removes the need for certificate exchange on subsequent visits.
- Without PKI, identity cannot be verified — you can still encrypt, but you cannot be sure who you're encrypting for.
- The CA encrypts (signs) the certificate with the **CA's private key**; anyone with the **CA's public key** can verify it.

## Cross-References

- [Digital Signatures](digital-signatures.md) — CA uses digital signatures to vouch for certificates
- [Asymmetric Crypto](asymmetric-crypto.md) — key pair mechanics
- [Key Management](key-management.md) — key storage (HSM, TPM)

## Sources

- destination-cissp §3.6.10–3.6.11 (pages 0425–0440)
