---
title: "Secure Protocols"
type: concept
domain: 4
tags: [ssh, https, sftp, snmpv3, tls, ldaps, secure-protocols]
sources: [destination-cissp, cissp-exam-outline]
updated: 2026-05-13
---

# Secure Protocols

For nearly every insecure protocol, a secure alternative exists. CISSP exam questions frequently test whether you can select the correct secure protocol given a scenario.

## Key Facts

### Insecure → Secure Protocol Mapping

| Insecure Protocol | Port | Secure Alternative | Port | Notes |
|-------------------|------|--------------------|------|-------|
| Telnet | 23 | SSH | 22 | SSH uses public-key crypto for remote login |
| FTP | 20/21 | SFTP (SSH FTP) | 22 | SFTP rides inside SSH tunnel |
| HTTP | 80 | HTTPS | 443 | HTTPS = HTTP over TLS |
| SNMP v1/v2 | 161/162 UDP | SNMPv3 | 161/162 UDP | v3 adds auth and encryption |
| LDAP | 389 | LDAPS | 636 | LDAP over SSL/TLS |
| Rlogin/rsh | various | SSH | 22 | Legacy Unix remote tools; replace with SSH |
| TFTP | 69 UDP | SFTP / SCP | 22 | TFTP has no auth; disable it |

### TLS Versions

| Version | Status | Notes |
|---------|--------|-------|
| SSLv2 | **Broken** | DROWN attack; do not support even for backward compatibility |
| SSLv3 | **Broken** | POODLE attack |
| TLS 1.0 | Deprecated | Removed from most browsers |
| TLS 1.1 | Deprecated | Removed from most browsers |
| TLS 1.2 | Current | Still widely used; acceptable if properly configured |
| TLS 1.3 | Current (2018) | Preferred; faster handshake, improved security |

### SSH

- Port 22. Uses public-key cryptography.
- Provides: encrypted remote shell, file transfer (SFTP), port forwarding/tunneling.
- Replaces: Telnet, rsh, rlogin, FTP (when used as SFTP).
- Hardening: Telnet can be tunneled through SSH if a service cannot be replaced.

### HTTPS and the TLS Handshake

1. **Client Hello** — client sends TLS version and supported cipher suites.
2. **Server Hello** — server responds with chosen cipher suite and its certificate (public key).
3. **Key Exchange** — client validates server cert (via CA chain), creates symmetric session key, encrypts it with server's public key, sends to server.
4. **Secure Session** — server decrypts with private key; both sides hold the same session key.

After handshake: URL shows `https://` and padlock. Asymmetric crypto is used only to establish the symmetric session key; all data thereafter is encrypted with the symmetric key (faster).

### DROWN Attack (SSLv2)

Enabling SSLv2 for backward compatibility exposes TLS connections to the DROWN attack, which can compromise the session and expose passwords, credit card numbers, and other sensitive data. **Best defense:** ensure private keys are not shared with servers that accept SSLv2 connections.

## Exam Nuance

- SSH on port **22** serves triple duty: remote login, SFTP, and tunneling other protocols.
- SFTP ≠ FTPS. SFTP uses SSH (port 22); FTPS is FTP over TLS (port 990 or explicit on 21).
- TLS is the **renamed, improved successor to SSL**. The two names are often used interchangeably colloquially, but TLS is correct.
- SSLv2 creates DROWN vulnerability even if the server itself primarily uses TLS 1.2+.
- The TLS handshake uses **asymmetric** crypto to securely exchange the **symmetric** session key.

## Cross-links

- [VPN](./vpn.md)
- [Network Protocols](./network-protocols.md)
- [TLS Standard](../standards/tls-standard.md)
- [Network Access Control](./network-access-control.md)

## Sources

- destination-cissp §4.1.10 (pp. 0556–0558)
- destination-cissp §4.3.3 (pp. 0650–0655)
