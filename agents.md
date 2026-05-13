# Agents

This repo contains scripts for managing GitHub tokens across different contexts.

## Scripts

### set-gh-cli-token.py

Swaps GitHub CLI token using `gh auth login --with-token`.

**Usage:**
```bash
python set-gh-cli-token.py <work_mode|personal_mode>
```

**Requirements:**
- `gh` CLI installed and authenticated
- `~/agents-data/config/gh-pats.json` exists with `work_PAT` and `personal_PAT` keys

**Behavior:**
- Reads PAT from config file
- Runs `gh auth login --with-token` to set CLI token
- Prints success message

### set-gh-mcp-token.py

Swaps GitHub PAT in Windsurf MCP config for GitHub server.

**Usage:**
```bash
python set-gh-mcp-token.py <work_mode|personal_mode>
```

**Requirements:**
- `~/.codeium/windsurf/mcp_config.json` exists with github server config
- `~/agents-data/config/gh-pats.json` exists with `work_PAT` and `personal_PAT` keys

**Behavior:**
- Reads PAT from config file
- Updates `mcp_config.json` github server Authorization header
- Prints success message

## Config

### config/gh-pats.json

Contains GitHub PATs for different contexts:
```json
{
  "work_PAT": "ghp_xxx...",
  "personal_PAT": "ghp_yyy..."
}
```

**Note:** This file is gitignored for security.
