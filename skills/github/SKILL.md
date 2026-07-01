---
name: github
description: you must use this skill before performing an action related to git and github
---

# What is this skill?
This skill provides a local CLI to manage the `<mode>` of the github related tools the user has. It also provides you with the two methods of github related tools you can use. 
The user has two modes of github, `personal` and `work`. This skill ensures you're setting the tokens correctly such that there is no confusion when you are working on a project. Failure to use this skill results in confusion and unneeded token spendage. 

# Step 1: Determine `<mode>`
If the repository that you are in is `tobi-*`, then `<mode>` is `work`
Otherwise, `<mode>` is `personal`.

# Step 2: Set tokens
Run these two commands. Do not explore, list MCP tools, or read source code. Just run them.

```bash
~/agents-skills/skills/github/set-gh mcp <mode>
~/agents-skills/skills/github/set-gh cli <mode>
```

`<mode>` is a positional argument. Not a flag. Examples:
- `~/agents-skills/skills/github/set-gh mcp work`
- `~/agents-skills/skills/github/set-gh cli personal`

If either command fails, inform the user. Do not attempt to debug.

# Step 3: Use github
You must use github mcp for all tasks. If the github mcp is inadequate, you may use `gh` CLI as a backup.

Do not list available MCP tools. You know what they are. Just use them.

# whoami
If `<mode>` is `work`, then your username is `fahadattobi` otherwise it is `LordFarquaadtheCreator`
If `<mode>` is `work`, then your email is `fahad@tobiwealth.com` otherwise it is `fahadfaruqi1@gmail.com`