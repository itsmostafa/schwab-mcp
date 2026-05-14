---
title: "TLS — Transport Layer Security"
type: standard
domain: 4
tags: [tls, ssl, https, tls13, tls12, drown, handshake, cipher-suite]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# TLS — Transport Layer Security

Transport Layer Security (TLS) is the successor to SSL (Secure Sockets Layer). It provides encrypted, authenticated communication channels between clients and servers. TLS is the protocol behind HTTPS and is the most widely deployed security protocol on the internet.

## Key Facts

### Version History

| Version | Status | Notes |
|---------|--------|-------|
| SSLv2 | **Broken** | DROWN attack — never enable, even for backward compat |
| SSLv3 | **Broken** | POODLE attack |
| TLS 1.0 | **Deprecated** | Removed from most modern browsers/servers |
| TLS 1.1 | **Deprecated** | Removed from most modern browsers/servers |
| TLS 1.2 | Current | Acceptable; widely deployed |
| TLS 1.3 | **Current (2018)** | Preferred; faster handshake; forward secrecy by default; removed weak cipher suites |

### TLS Handshake (TLS 1.2)

1. **Client Hello** — Client sends supported TLS version, cipher suites, random nonce.
2. **Server Hello** — Server selects cipher suite, sends its certificate (contains server public key), signed by a CA.
3. **Key Exchange** — Client verifies cert against trusted CA chain. Client generates symmetric session key, encrypts with server's public key, sends to server. Server decrypts with private key. Both now share the session key.
4. **Secure Session** — All subsequent communication encrypted with the symmetric session key.

**Asymmetric crypto:** Used only for the key exchange phase (steps 2–3).
**Symmetric crypto:** Used for all application data (step 4) — much faster.

### TLS 1.3 Improvements over 1.2

- **Faster handshake:** 1-RTT (round-trip) handshake by default (vs 2-RTT for TLS 1.2).
- **Forward secrecy mandatory:** Uses ephemeral keys (ECDHE) so past sessions cannot be decrypted if private key is later compromised.
- **Removed weak algorithms:** No more RSA key exchange, RC4, DES, 3DES, SHA-1, MD5 in cipher suites.
- **Encrypted certificate:** Server certificate is encrypted in 1.3 (not in 1.2).

### DROWN Attack (SSLv2)

The DROWN attack exploits SSLv2 — even if a server primarily uses TLS 1.2, if it also accepts SSLv2 connections, attackers can use SSLv2's weaknesses to decrypt TLS 1.2 sessions if the same private key is used on both. **Mitigation:** Disable SSLv2 entirely and never share private keys between servers that support SSLv2 and servers that don't.

### Mutual Authentication

By default, TLS authenticates the **server** to the client (client verifies server cert). **Mutual TLS (mTLS)** additionally requires the client to present a certificate, authenticating to the server. Used in zero-trust environments and API security.

## Exam Nuance

- SSL and TLS are the **same protocol family**; TLS is the modern name. The exam may use both terms interchangeably.
- SSLv2 DROWN: enabling it for backward compatibility can compromise TLS 1.2 sessions — **never enable SSLv2**.
- Asymmetric crypto is used to **exchange** the symmetric session key, not to encrypt application data.
- TLS 1.3 is preferred; TLS 1.2 is acceptable. TLS 1.0/1.1 are deprecated.
- TLS VPN vs IPsec VPN: TLS operates at **Layer 4**, IPsec at **Layer 3** — see [VPN](../concepts/vpn.md).

## Cross-links

- [Secure Protocols](../concepts/secure-protocols.md)
- [VPN](../concepts/vpn.md)
- [Network Protocols](../concepts/network-protocols.md)

## Sources

- destination-cissp §4.3.3 (pp. 0650–0656)
