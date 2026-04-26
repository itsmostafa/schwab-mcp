---
tags: [agentic-ai, tool-use, function-calling, llm, interview-prep]
type: concept
last-updated: 2026-04-26
---

# Tool Calling

How LLMs invoke external functions — the mechanism that gives agents real-world capability.

## The Core Idea

An LLM by itself only generates text. Tool calling is the mechanism by which an LLM can **request** the execution of an external function, receive the result, and incorporate it into its reasoning.

Critical distinction: for client-side tools, **the model doesn't execute anything**. It outputs a structured request saying "call this function with these arguments." Your code actually runs the function. Some providers also offer server-side tools, but your application still needs explicit policy, permissions, and audit boundaries around tool access.

## The Protocol (Step by Step)

```
1. Developer defines tools (name + description + JSON Schema/input schema for args)
2. User sends a message
3. LLM reasons, decides a tool is needed
4. LLM outputs a tool_use block: { name, id, input }   ← stop_reason: "tool_use"
5. Your code executes the actual function
6. You append the result as a tool_result block
7. You call the API again with updated conversation
8. LLM reasons with the new data, either calls another tool or returns final answer
9. Repeat until stop_reason == "end_turn"
```

This is a **request-response loop** — the LLM and your runtime take turns.

## Tool Definition Structure

```json
{
  "name": "get_alert_details",
  "description": "Retrieve full alert metadata from the SIEM by alert ID. Use this when you need to understand what triggered an alert before triaging it.",
  "input_schema": {
    "type": "object",
    "properties": {
      "alert_id": {
        "type": "string",
        "description": "The unique SIEM alert identifier"
      },
      "include_raw_logs": {
        "type": "boolean",
        "description": "Whether to include the raw log events that triggered the alert"
      }
    },
    "required": ["alert_id"]
  }
}
```

Three parts that matter:
- **`name`** — must be unambiguous and action-oriented (`get_`, `search_`, `create_`, `close_`)
- **`description`** — this is how the LLM decides *when* to use the tool. Invest here.
- **`input_schema` / `parameters`** — JSON Schema-style argument contract; constrains the model output, but handlers must still validate semantics and permissions

## Anthropic API Mechanics

```python
# 1. Call API with tools
response = client.messages.create(
    model="claude-sonnet-4-6",
    tools=[get_alert_details_tool],
    messages=[{"role": "user", "content": "Triage alert ABC-123"}]
)

# 2. Check if model wants to use a tool
if response.stop_reason == "tool_use":
    tool_use = [b for b in response.content if b.type == "tool_use"][0]
    
    # 3. Execute the actual function
    result = get_alert_details(tool_use.input["alert_id"])
    
    # 4. Append both the model's request and your result to messages
    messages.append({"role": "assistant", "content": response.content})
    messages.append({
        "role": "user",
        "content": [{
            "type": "tool_result",
            "tool_use_id": tool_use.id,   # must match the id from the request
            "content": json.dumps(result)
        }]
    })
    
    # 5. Call API again — model now has the data and can continue
    response = client.messages.create(...)
```

## The Agent Loop (Full Implementation Pattern)

```python
def run_agent(user_message, tools):
    messages = [{"role": "user", "content": user_message}]
    
    while True:
        response = client.messages.create(
            model="claude-sonnet-4-6",
            tools=tools,
            messages=messages
        )
        
        if response.stop_reason == "end_turn":
            # Extract and return final text response
            return next(b.text for b in response.content if b.type == "text")
        
        if response.stop_reason == "tool_use":
            messages.append({"role": "assistant", "content": response.content})
            
            # Handle parallel tool calls (multiple tools in one response)
            tool_results = []
            for block in response.content:
                if block.type == "tool_use":
                    result = execute_tool(block.name, block.input)
                    tool_results.append({
                        "type": "tool_result",
                        "tool_use_id": block.id,
                        "content": json.dumps(result)
                    })
            
            messages.append({"role": "user", "content": tool_results})
```

## Parallel Tool Calls

Modern APIs can return multiple `tool_use` blocks in a single response — the model requests several tools at once. Your code runs them (ideally in parallel), returns all results together, and the model continues.

```
Model: "I need alert details AND the user's history. Let me call both."
→ [get_alert_details(ABC-123), get_user_history(john.doe)]
Your code: runs both concurrently
→ [result_1, result_2]
Model: continues reasoning with both datasets
```

This can be significantly faster than sequential calls for independent lookups, assuming the host runtime actually executes the requested calls concurrently and the downstream systems can handle it.

## SOC-Specific Tool Inventory (Panther Context)

For an alert triage agent, tools would look like:

| Tool | Purpose |
|------|---------|
| `get_alert_details(alert_id)` | Fetch what triggered the alert, raw log events |
| `search_threat_intel(indicator)` | Look up IP/domain/hash in threat intel |
| `query_logs(query, time_range)` | Run a detection query against the log store |
| `get_user_history(username, hours)` | Recent activity for a user principal |
| `enrich_ip(ip)` | Geo, ASN, reputation score |
| `get_asset_context(hostname)` | What is this machine? Owner, risk level |
| `create_ticket(title, severity, body)` | Open an incident in the ticketing system |
| `close_alert(alert_id, resolution, reason)` | Mark as true/false positive |
| `escalate(alert_id, team)` | Hand off to a human team |

The last two are **write tools** — they take real-world action. These deserve extra care (guardrails, confirmation, audit logging).

## Design Principles

### Description quality is the most important variable
The description is how the LLM decides *whether* and *when* to call a tool. A vague description leads to wrong tool selection. A good description says: what it does, when to use it, and when *not* to use it.

### Minimal, focused tool sets
More tools generally means a harder selection problem. Start small, name tools distinctly, and group related operations when that makes the decision surface clearer.

### Read vs. write tools
Separate tools that retrieve information from tools that take action. Read tools are safe to retry; write tools can have side effects. Many systems require explicit human confirmation before write tools execute.

### Input validation at the boundary
The model's JSON arguments may be structurally valid but semantically wrong (wrong IDs, out-of-range values). Validate at your handler, don't trust blindly.

### Termination conditions
An agent can loop — call tools indefinitely without making progress. Add a max-iterations guard to every agent loop.

### Observability
Log every tool call and its result. In a SOC context, this is also your audit trail — who (which agent), what (which tool), when, why (the prompt context).

## Common Failure Modes

| Failure | Cause | Fix |
|---------|-------|-----|
| Model picks wrong tool | Ambiguous descriptions | Sharpen descriptions; use distinct names |
| Model hallucinates arguments | Weak schema, no examples | Add examples and constraints to schema |
| Infinite tool loop | No termination condition | Max-iterations guard |
| Tool overuse | Model prefers confirming over reasoning | Instruct model to reason first, call only when needed |
| Stale data | Cached tool results | Don't cache time-sensitive tool calls |
| Silent failures | No error propagation | Return structured errors; model can recover or escalate |

## Talking Points for "Explain Tool Calling"

1. **What it is**: "Tool calling is how we give LLMs controlled access to external data and actions — they request a tool through a structured interface, then incorporate the returned result. Without tools, agents can only reason over the context already provided."

2. **The key insight**: "For client-side tools, the model doesn't execute anything. It requests a function call. Your code runs it. This separation means you can enforce security boundaries, add audit logging, require human confirmation for destructive actions — all at the host layer."

3. **Why the description matters**: "The tool description is effectively a prompt. If it's ambiguous, the model picks the wrong tool. I've seen agents fail not because the code was wrong but because someone wrote a lazy description."

4. **Parallel calls**: "One underappreciated feature is parallel tool calling — the model can ask for multiple independent tools in one shot. For alert triage, threat intel, user history, and asset context can often be fetched concurrently."

5. **Building a system**: "The core loop is: call API → inspect tool-call response → execute allowed tools → append tool results → loop until final response or iteration limit. The complexity is in tool design, error handling, and guardrails around write operations."

## Validation Sources

- OpenAI: function calling is tool calling with function tools defined by JSON Schema, and the application receives arguments to access data or take actions.
- Anthropic: Claude returns `stop_reason: "tool_use"` and `tool_use` blocks for client tools; the application executes them and sends back `tool_result`.
- MCP specification: MCP standardizes how hosts connect LLM applications to resources, prompts, and tools, with explicit user consent and tool-safety requirements.

## Building a System: Key Design Decisions

1. **Tool inventory** — what can the agent do? Start with read-only tools, add write tools carefully
2. **Tool descriptions** — spend disproportionate time here; they determine correctness
3. **Tool handlers** — actual implementations with input validation and error wrapping
4. **The agent loop** — while loop with max-iteration guard, structured result accumulation
5. **Parallel execution** — run independent tool calls concurrently in the handler
6. **Error propagation** — return structured errors so the model can recover or escalate
7. **Guardrails for write tools** — consider human-in-the-loop confirmation for destructive actions
8. **Observability** — log every call with timestamp, inputs, outputs, latency
9. **Testing** — mock tool responses to test agent behavior without hitting real systems

## See Also

- [[concepts/agentic-ai]] — agent loop, planning strategies, memory types
- [[rounds/ai-integration]] — AI integration interview; tool calling is likely to come up here
- [[panther/role]] — Panther's 4 SOC agent capabilities; tool boundaries are relevant to each
