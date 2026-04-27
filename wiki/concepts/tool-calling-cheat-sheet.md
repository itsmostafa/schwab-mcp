---
tags: [tool-calling, agentic-ai, function-calling, interview-prep]
type: cheat-sheet
last-updated: 2026-04-27
---

# Tool Calling Cheat Sheet

Use this when you need a fast, interview-ready explanation of how LLM tool calls work and how to design them safely.

---

## One-Sentence Definition

Tool calling is how an LLM asks the host application to run a structured external function, then uses the returned result to continue the conversation or agent workflow.

Key distinction:

> The model does not execute client-side tools. It emits a structured request. Your runtime validates, authorizes, runs the function, returns the result, and decides whether the loop continues.

---

## The Loop

Memorize:

```text
Define tools
-> call model with tools + messages
-> model returns tool request
-> validate request
-> authorize action
-> execute handler
-> append tool result
-> call model again
-> repeat until final answer or stop condition
```

Short interview phrasing:

> "The LLM chooses a tool by name and arguments, but my application owns execution, permissions, retries, error handling, and audit logs."

---

## Tool Definition Anatomy

Every tool needs:

| Part | What it does | What good looks like |
|---|---|---|
| `name` | Action the model can request | Clear verb: `get_alert_details`, `search_logs`, `create_ticket` |
| `description` | Teaches the model when to use it | Says what it does, when to use it, and when not to use it |
| `schema` | Shapes the arguments | Required fields, types, enums, descriptions, constraints |
| `handler` | Actual code that runs | Validates, authorizes, executes, returns structured output/errors |
| policy | Execution boundary | RBAC, tenant isolation, rate limits, human approval for risky actions |

Example:

```json
{
  "name": "search_threat_intel",
  "description": "Look up IPs, domains, URLs, or hashes in threat intelligence. Use only for concrete indicators extracted from an alert or log event.",
  "input_schema": {
    "type": "object",
    "properties": {
      "indicator": {
        "type": "string",
        "description": "The IP, domain, URL, or hash to investigate"
      },
      "indicator_type": {
        "type": "string",
        "enum": ["ip", "domain", "url", "hash"]
      }
    },
    "required": ["indicator", "indicator_type"]
  }
}
```

---

## Read vs Write Tools

This is the most important design boundary.

| Tool type | Examples | Policy |
|---|---|---|
| Read | `get_alert_details`, `query_logs`, `get_user_history`, `search_threat_intel` | Safe to retry, still needs auth and tenant isolation |
| Write | `create_ticket`, `close_alert`, `escalate_alert`, `block_ip` | Requires stricter validation, idempotency, audit logs, often human approval |

Good line:

> "I would start with read-only tools, prove quality, then add write tools behind explicit policy gates."

---

## SOC Agent Tool Inventory

For Panther-style alert triage:

```text
get_alert_details(alert_id)
query_logs(query, time_range, tenant_id)
get_user_history(user_id, lookback_window)
get_asset_context(host_id)
search_threat_intel(indicator)
find_similar_alerts(alert_fingerprint)
create_ticket(summary, severity, evidence_ids)
escalate_alert(alert_id, team, reason)
close_alert(alert_id, resolution, evidence_ids)
```

Default policy:

- Let the agent freely request read tools within tenant/RBAC limits.
- Require policy gates for write tools.
- Require human review for high-severity, low-confidence, or containment actions.
- Log every requested tool call, whether allowed or blocked.

---

## Handler Rules

Never trust model arguments blindly.

Handler checklist:

- Validate required fields, types, enum values, IDs, time ranges, and tenant scope.
- Enforce RBAC and tenant isolation server-side.
- Make write operations idempotent where possible.
- Add timeouts, retries, and circuit breakers.
- Return structured errors the agent can recover from.
- Redact secrets and sensitive fields before returning results to the model.
- Attach stable evidence IDs instead of dumping excessive raw data.
- Record audit logs: tool, args, actor, tenant, result status, latency, model/prompt version.

---

## Parallel Tool Calls

Independent lookups can run concurrently:

```text
Alert triage needs:
- alert details
- user history
- asset context
- threat intel
- similar alerts

The model can request several tools in one turn.
The host runtime runs them in parallel and returns all results together.
```

Tradeoff:

- Parallel calls reduce latency.
- They can increase load on downstream services.
- Add rate limits, budgets, and graceful degradation.

---

## Common Failure Modes

| Failure | Why it happens | Fix |
|---|---|---|
| Wrong tool selected | Names/descriptions overlap | Sharpen tool names and descriptions |
| Bad arguments | Weak schema or ambiguous prompt | Add enums, constraints, examples, handler validation |
| Infinite loop | Agent keeps calling tools | Max iterations, progress checks, final-answer fallback |
| Tool overuse | Model verifies everything | Tell it when not to call tools, cache stable reads |
| Unsafe write | Model proposes risky action | External policy gate and human approval |
| Cross-tenant leakage | Retrieval/tool layer trusts prompt | Tenant enforcement in code, not prompt |
| Prompt injection | Tool result/log contains instructions | Treat tool output as data, not instructions |
| Silent failures | Handler hides errors | Return structured error objects |

---

## Security Model

Strong answer:

> "Tool calling is powerful because it gives the model access to systems, but the security boundary must stay in the application. The model can request, but code must authorize."

Guardrails:

- Tool allowlist per agent/task.
- Read/write separation.
- Server-side authorization on every call.
- Tenant isolation at the data layer.
- Human approval for dangerous side effects.
- Prompt-injection resistant formatting for untrusted tool outputs.
- Full audit trail for compliance and debugging.

In SOC context, tool results are often attacker-influenced. A log line can contain malicious instructions, so retrieved content must never override system/developer instructions or tool policy.

---

## Testing Tool Calls

Test at three layers:

1. **Tool handler tests**: validation, auth, idempotency, structured errors.
2. **Agent behavior tests**: given mocked tool results, does the model choose the right next action?
3. **Safety/eval tests**: prompt injection, wrong-tenant access, write-tool approval, high-severity escalation.

Useful fixtures:

- Happy path alert triage
- Missing alert ID
- Threat intel timeout
- Ambiguous indicator
- Malicious log line saying "ignore previous instructions"
- Low-confidence result that must not auto-close
- High-severity result that must escalate

---

## Interview Answer Templates

**"Explain tool calling."**

> "Tool calling lets an LLM request external functions through a schema. The host application executes the function, not the model, which lets us enforce validation, permissions, audit logging, and safety gates."

**"How do you design tools?"**

> "I start with a minimal inventory, separate read from write, give each tool a precise name and description, define a strict input schema, and keep policy enforcement in the handler."

**"How do you prevent unsafe actions?"**

> "I treat model output as a proposal. The runtime validates args, checks RBAC and tenant scope, applies policy, and requires human approval for risky write actions."

**"What would tools look like for alert triage?"**

> "Read tools fetch alert details, related logs, user history, asset context, threat intel, and similar past alerts. Write tools create tickets, escalate, add notes, or close alerts, but only behind policy gates."

**"How do you handle tool errors?"**

> "Return structured errors, let the agent decide whether to retry or degrade gracefully, and cap retries/iterations so it cannot loop forever."

---

## Final Recall

Memorize:

```text
Model requests. Runtime validates. Runtime authorizes. Runtime executes.
Runtime returns structured result. Model continues. Policy stays outside the model.
```

## See Also

- [[concepts/tool-calling]] — full reference
- [[rounds/round-3-cheat-sheet]] — Round 3 AI Integration recall sheet
- [[concepts/agentic-ai]] — agent loop, RAG, feedback loops
