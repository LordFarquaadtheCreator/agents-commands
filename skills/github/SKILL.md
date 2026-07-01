---
name: github
description: you must use this skill before performing an action related to git and github
---

If the repository that you are in is `tobi-*`, then `<mode>` is `work`
Otherwise, `<mode>` is `personal`

# Github MCP
You **must** use github mcp for all tasks.

Before using the github mcp, you must call the `set-gh-mcp-token` tool from the `set-gh-token` MCP server with the current `<mode>` as the `mode` argument. If the tool call fails, you cannot use github and must inform the user. Try `gh` command as a backup, see below.

# `GH` CLI
If the github mcp is inadequate, then you are allowed to use `gh` - github cli tool. Use this with caution.

The `gh` CLI is a separate tool. Before using `gh`, you must call the `set-gh-cli-token` tool from the `set-gh-token` MCP server with the current `<mode>` as the `mode` argument. If the tool call fails, you cannot use `gh` and must inform the user.

# whoami
If `<mode>` is `work`, then your username is `fahadattobi` otherwise it is `LordFarquaadtheCreator`
If `<mode>` is `work`, then your email is `fahad@tobiwealth.com` otherwise it is `fahadfaruqi1@gmail.com`