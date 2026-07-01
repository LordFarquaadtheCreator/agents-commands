# MCPs

All MCP servers are defined here. Each MCP must be in its own folder.

## Structure

```
mcps/
├── AGENTS.md              # this file
└── <mcp-name>/
    ├── Dockerfile          # builds and runs the MCP server
    ├── AGENTS.md           # instructions for this specific MCP
    └── mcp-config.json     # copy-pastable MCP config entry for mcp_config.json
```

## Rules

- Each MCP lives in its own directory under `mcps/`
- Every MCP must have a `Dockerfile`, `AGENTS.md`, and a `mcp-config.json` snippet
- `mcp-config.json` contains the copy-pastable entry for `mcpServers` in `~/.codeium/windsurf/mcp_config.json`
- `AGENTS.md` describes what the MCP does, how to build it, and how to run it
