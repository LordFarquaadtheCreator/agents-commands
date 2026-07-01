package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"set-gh-token/pats"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var mcpConfigPath = filepath.Join(os.Getenv("HOME"), ".codeium", "windsurf", "mcp_config.json")

type McpConfig struct {
	McpServers map[string]McpServer `json:"mcpServers"`
}

type McpServer struct {
	Disabled      bool              `json:"disabled"`
	ServerUrl     string            `json:"serverUrl,omitempty"`
	Url           string            `json:"url,omitempty"`
	Command       string            `json:"command,omitempty"`
	Args          []string          `json:"args,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	DisabledTools []string          `json:"disabledTools,omitempty"`
}

type McpInput struct {
	Mode string `json:"mode" jsonschema:"required,description=Either work or personal"`
}

type McpOutput struct {
	Message string `json:"message"`
}

func SwapMcpToken(ctx context.Context, ss *mcp.ServerSession, req *mcp.CallToolParamsFor[McpInput]) (*mcp.CallToolResultFor[McpOutput], error) {
	token, err := pats.LoadToken(req.Arguments.Mode)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(mcpConfigPath)
	if err != nil {
		return nil, fmt.Errorf("MCP config not found: %s", mcpConfigPath)
	}

	var config McpConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid JSON in MCP config: %v", err)
	}

	github := config.McpServers["github"]
	github.ServerUrl = "https://api.githubcopilot.com/mcp/"
	github.Headers = map[string]string{}

	github.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	config.McpServers["github"] = github

	output, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(mcpConfigPath, output, 0644); err != nil {
		return nil, fmt.Errorf("cannot write to %s", mcpConfigPath)
	}

	msg := fmt.Sprintf("Successfully updated GitHub token to %s", req.Arguments.Mode)
	return &mcp.CallToolResultFor[McpOutput]{
		Content:           []mcp.Content{&mcp.TextContent{Text: msg}},
		StructuredContent: McpOutput{Message: msg},
	}, nil
}
