# SOC Domain Knowledge

Security Operations Center context for the Panther interview.

## What a SOC Does

A Security Operations Center monitors, detects, analyzes, and responds to cybersecurity threats. Core activities:

1. **Alert triage**: receive alert → determine if real threat vs false positive → prioritize
2. **Investigation**: dig into real threats, gather context, understand blast radius
3. **Response**: contain, eradicate, remediate
4. **Detection engineering**: write and tune detection rules

## The Pain Points Panther Solves

From the job description and SOC domain:
- **Alert fatigue**: analysts receive thousands of alerts/day, most are false positives
- **Manual triage**: repetitive, time-consuming, error-prone when humans do it
- **Context gaps**: alerts lack enrichment — who is this IP? Is this user behavior normal?
- **Scaling headcount**: more data = more alerts = more analysts (expensive, doesn't scale)
- **Detection quality**: hard to write good detection rules; hard to test them

Panther's answer: AI agents that triage at machine speed, learn from analyst decisions, and enable SOC teams to cover 5-10x more data without proportional headcount growth.

## Key Terminology

| Term | Meaning |
|------|---------|
| SIEM | Security Information and Event Management — collects/correlates logs |
| XDR | Extended Detection and Response — cross-layer threat detection (endpoint, network, cloud) |
| SOAR | Security Orchestration, Automation, and Response |
| Alert triage | Classifying alerts as true/false positive and prioritizing response |
| Detection-as-Code | Writing detection rules as code (version controlled, testable) |
| IOC | Indicator of Compromise (IP, hash, domain, etc.) |
| TTPs | Tactics, Techniques, Procedures (MITRE ATT&CK framework) |
| False positive | Alert fired on benign activity |
| True positive | Alert fired on actual malicious activity |
| MTTR | Mean Time to Respond — key metric for SOC efficiency |

## Panther's Product

- **Cloud-native SIEM alternative**: ingestion pipeline + security data lake
- **Detection-as-Code**: Python-based detection rules, version-controlled, tested
- **Panther Query Language (PQL)**: SQL-like queries over security data
- **4 AI SOC agents** (current):
  1. Alert triage agent
  2. Interactive chat
  3. Detection code generation
  4. Text-to-search

## Alert Triage Flow (Manual vs AI)

**Manual flow**:
Alert fires → Analyst reviews → Looks up IOCs → Checks context (past alerts, asset info) → Decides true/false → Documents → Escalates or closes

**AI agent flow**:
Alert fires → Agent fetches enrichment → Searches past similar alerts → Retrieves analyst decisions on similar cases → Classifies risk with evidence → Auto-closes only high-confidence benign/low-risk alerts within policy → Routes risky or inconclusive alerts with context summary

## Why Go + Python Matter Here

- Panther's ingestion pipeline likely uses Go (performance-critical, concurrent)
- Python is the language of ML/AI tooling (langchain, openai SDK, embeddings)
- Detection rules in Panther are Python
- Having both = can work across the stack

## See Also

- [[panther/company]] — Panther product details
- [[rounds/systems-design]] — system design prompts for SOC systems
