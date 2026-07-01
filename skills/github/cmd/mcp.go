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

type McpInput struct {
	Mode string `json:"mode" jsonschema:"required,description=Either work_mode or personal_mode"`
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

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid JSON in MCP config: %v", err)
	}

	mcpServers, ok := config["mcpServers"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("required key missing in config: mcpServers")
	}

	githubEntry, ok := mcpServers["github"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("required key missing in config: mcpServers.github")
	}

	headers, ok := githubEntry["headers"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("required key missing in config: mcpServers.github.headers")
	}

	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)

	output, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(mcpConfigPath, output, 0644); err != nil {
		return nil, fmt.Errorf("cannot write to %s", mcpConfigPath)
	}

	msg := fmt.Sprintf("Successfully updated GitHub token to %s", req.Arguments.Mode)
	return &mcp.CallToolResultFor[McpOutput]{
		Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		StructuredContent: McpOutput{Message: msg},
	}, nil
}
