---
type: audit
last-updated: 2026-04-26
tags: [audit, prep, ai-integration, systems-design]
---

# Prep Plan Method Audit

Scope: `wiki/prep-plan.md`, checked against the official candidate prep guide, role description, referral notes, and round-specific wiki pages.

## Verdict

For a compressed 36-hour window before the AI Integration and Systems Design interviews, the plan is close to the best practical prep strategy.

The core choice is right: stop accumulating more notes and optimize for retrieval under pressure. The official guide says these rounds evaluate reasoning, communication, ambiguity handling, production thinking, security awareness, and depth from real experience. The plan trains those signals directly through out-loud mocks, a canonical SOC alert triage design, tradeoff drills, failure-mode drills, and protected sleep.

It is not the absolute best possible strategy in every timeline. With a week or more, the best plan would add more interviewer-style external mocks, a deeper personal-project retro, and more Panther-specific product research. But given the stated 36-hour schedule, adding much more content would likely make performance worse, not better.

## What Is Strong

- The plan matches the official AI Integration rubric: agent architecture, RAG/retrieval, prompt composition, AI security, production metrics, and SOC fluency.
- The plan matches the official Systems Design rubric: ambiguity handling, clear abstractions, tradeoffs, failure modes, collaboration, and strategic fit.
- The anchor alert triage agent is the right unifying design because it covers both rounds and Panther's actual role focus.
- The study mechanics are correct: active recall, timed mocks, whiteboarding while speaking, interleaving, and two protected sleep windows.
- The "no more deep reading" rule is correct for this stage. The wiki content is already strong; the remaining risk is delivery under pressure.

## Gaps To Fix

1. Add a short "questions to ask Panther" drill.
   The official Systems Design guide explicitly says to come with good questions. The plan has clarifying questions for the design prompt, but not questions that show curiosity about Panther's architecture and decision-making.

2. Add one interviewer-assisted mock if possible.
   Solo mocks are useful, but the rounds evaluate interaction: interruptions, clarification, follow-up pressure, and whether the interviewer can follow the design. One 30-minute mock with another person would be higher value than extra reading.

3. Make the personal system story non-optional and Panther-shaped.
   The plan includes it, but it should be treated as a top-tier asset for both Rounds 3 and 4. It needs concrete scale, metrics, tradeoffs, failure, customer/user impact, and what you would do differently now.

4. Include a "no live AI answers" reminder.
   The job description permits AI for prep and some assignments, but says live conversational and technical interviews are meant to assess the candidate's own thinking. For Rounds 3 and 4, the plan should explicitly say: use AI for preparation only, not to generate live answers.

5. Slightly broaden the Systems Design drill beyond the alert triage anchor.
   The anchor is still right, but the official guide emphasizes general system architecture: processes, storage/databases, networking, APIs, concurrency, latency, and failure handling. The second mock should force one non-triage design to avoid overfitting.

## Recommended Minimal Edits

- In Block A, keep the personal story work exactly where it is and label it the highest-priority task.
- In Block B or C, replace any low-value review with a 15-minute "Panther questions" drill:
  - "Where do you draw the line today between autonomous action and analyst review?"
  - "What signals have mattered most in evaluating SOC agent quality with real customers?"
  - "How do you think about tenant-specific behavior versus shared model/product behavior?"
  - "What constraints from Panther's existing ingestion pipeline shape the agent architecture?"
- In Block C, make the second systems-design mock non-triage: text-to-search, detection generation, or collective intelligence.
- On interview day, add one line: no new AI-generated answers; speak from the internalized scaffolds and real experience.

## Bottom Line

The plan is excellent for the time window. The best version is not "more studying"; it is:

1. Rehearse the triage architecture until it is automatic.
2. Nail the personal system story with metrics and tradeoffs.
3. Do one live mock with interruptions if possible.
4. Prepare 3-4 sharp Panther architecture questions.
5. Sleep.
