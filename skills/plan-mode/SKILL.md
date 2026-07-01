---
name: plan-mode
description: you must use this skill when planning
---

Task: Build exhaustive plans. Next agent needs zero context. Copy-paste ready.

## STEP 1: DISCOVERY

Read Fahad request. Identify goal. Search entire codebase. Map:
- Relevant files (absolute paths)
- Dependencies between files
- Current architecture patterns
- Existing similar implementations
- Constraints (framework, versions, conventions)

## STEP 2: CONTEXT GATHERING

For each relevant file:
- Copy key code sections
- Note exports, imports, types
- Document current behavior
- Identify integration points

Paste code snippets into plan. Next agent must not search.

## STEP 3: ARCHITECTURE DECISIONS

Choose approach. Document why. Consider:
- Minimal change vs refactor
- Consistency with existing patterns
- Performance implications
- Test coverage requirements

## STEP 4: STEP-BY-STEP BLUEPRINT

Each step numbered. Each step includes:

1. Action: create/modify/delete/test
2. Target: absolute file path
3. Current code: paste existing (for modifies)
4. New code: complete replacement or insertion
5. Line numbers: where change happens
6. Dependencies: which steps must run first
7. Verification: how to confirm success

No hand-waving. Exact code provided.

## STEP 5: EDGE CASES & ERROR HANDLING

List every edge case. Map to handling:
- Null/undefined inputs
- Empty collections
- Network failures
- Timeout scenarios
- Permission errors
- Invalid state transitions
- Race conditions

Each edge case: prevention or handling code specified.

## STEP 6: TESTS

New code needs tests. Specify:
- Test file path
- Test cases (input -> expected output)
- Edge case coverage
- Integration test scenarios

## STEP 7: FINAL CHECKLIST

Plan complete? Verify:
- Every file mentioned has absolute path
- Every code change includes exact code
- No step assumes prior knowledge
- No ambiguity in instructions
- Rollback steps documented
- Tests specified

## HANDOFF

Present plan to Fahad. Get approval. Next agent executes blindly.
