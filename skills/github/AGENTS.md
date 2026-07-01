# Agent Instructions — set-gh-token

## Structure

```
github/
├── main.go              # MCP server setup, stdio transport
├── cmd/                 # MCP tool handlers
│   ├── cli.go           # SwapCliToken — swaps gh CLI token
│   └── mcp.go           # SwapMcpToken — swaps MCP token in mcp_config.json
├── pats/                # Shared package
│   └── pats.go          # LoadToken(mode) — reads gh-pats.yaml
├── go.mod
├── go.sum
└── set-gh-token         # compiled binary (gitignored)
```

## Config

PATs file at `~/agents-skills/config/gh-pats.yaml`:
```yaml
work_PAT: "..."
personal_PAT: "..."
```

MCP config at `~/.codeium/windsurf/mcp_config.json`.

## Rebuilding

```bash
cd /Users/farquaad/agents-skills/skills/github && go build -o set-gh-token .
```

## MCP Server

The binary runs as an MCP server over stdio. It exposes two tools:

- `set-gh-mcp-token` — swaps GitHub PAT in mcp_config.json. Args: `mode` (work_mode|personal_mode)
- `set-gh-cli-token` — swaps gh CLI token via `gh auth login --with-token`. Args: `mode` (work_mode|personal_mode)
