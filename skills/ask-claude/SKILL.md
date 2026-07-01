---
name: ask-claude
description: you must use this skill when you need to plan a feature
---

You must use this skill sparingly, it is expensive. 

First, check if the BrowserOS mcp is active. If not, ask the user to activate it.

Then, go to `https://claude.ai/new` and ask your question. 

The prompt for claude must be as follows:
1) Ask your question with extreme specifics (how to make something, properly structure it, etc.)
2) You must include full relevant context in this prompt, Any information about the application, architecture, design that the website cannot (and should not) infer about the project should be included. 
3) You must include that you need the response in plain markdown.

Then, hit send for the prompt and wait for the response. Wait 60 seconds for the prompt to finish. The response will not be immediate.

# Valid assumptions:
If the BrowserOS MCP is available as listed, it is active.
You must always start by creating a new page to claude, never recycle old pages or consider them. 