---
name: github
description: you must use this skill before performing an action related to git and github
---

If the repository that you are in is `tobi-*`, then you are in `work_mode`
Otherwise, you are in `personal_mode`

# Environment
If you are in `work_mode`, environment is `work_mode`. Otherwise, environment is `personal_mode`

# Github MCP
You **must** use github mcp for all tasks.

Before using the github mcp, you must call the `set-gh-mcp-token` tool from the `set-gh-token` MCP server with the current `<environment>` as the `mode` argument. If the tool call fails, you cannot use github and must inform the user. Try `gh` command as a backup, see below.

# `GH` CLI
If the github mcp is inadequate, then you are allowed to use `gh` - github cli tool. Use this with caution.

The `gh` CLI is a separate tool. Before using `gh`, you must call the `set-gh-cli-token` tool from the `set-gh-token` MCP server with the current `<environment>` as the `mode` argument. If the tool call fails, you cannot use `gh` and must inform the user.

# whoami
If you are in `work_mode`, then your username is `fahadattobi` otherwise it is `LordFarquaadtheCreator`
If you are in `work_mode`, then your email is `fahad@tobiwealth.com` otherwise it is `fahadfaruqi1@gmail.com`