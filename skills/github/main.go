package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"set-gh-token/cmd"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "set-gh-token", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "set-gh-mcp-token",
		Description: "Swap GitHub MCP token in mcp_config.json based on mode (work_mode or personal_mode)",
	}, cmd.SwapMcpToken)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "set-gh-cli-token",
		Description: "Swap GitHub CLI token using gh auth login --with-token based on mode (work_mode or personal_mode)",
	}, cmd.SwapCliToken)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
