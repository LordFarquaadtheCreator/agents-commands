# Agent Instructions — set-gh

## Structure

```
github/
├── main.go              # entrypoint, calls cobra Execute
├── cmd/                 # cobra subcommands
│   ├── root.go          # root cobra command
│   ├── cli.go           # cli subcommand — swaps gh CLI token, captures stderr
│   └── mcp.go           # mcp subcommand — swaps MCP token via typed McpConfig struct
├── pats/                # Shared package
│   └── pats.go          # LoadToken(mode) — reads gh-pats.yaml, mode is work|personal
├── go.mod
├── go.sum
└── set-gh               # compiled binary (gitignored)
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
cd /Users/farquaad/agents-skills/skills/github && go build -o set-gh .
```

## CLI

The binary is a cobra CLI with two subcommands:

- `set-gh mcp <mode>` — swaps GitHub PAT in mcp_config.json. `mode` is work|personal
- `set-gh cli <mode>` — swaps gh CLI token via `gh auth login --with-token`. `mode` is work|personal
