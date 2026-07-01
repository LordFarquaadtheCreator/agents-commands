package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"set-gh/pats"

	"github.com/spf13/cobra"
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

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Swap GitHub MCP token in mcp_config.json",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := args[0]
		token, err := pats.LoadToken(mode)
		if err != nil {
			return err
		}

		data, err := os.ReadFile(mcpConfigPath)
		if err != nil {
			return fmt.Errorf("MCP config not found: %s", mcpConfigPath)
		}

		var config McpConfig
		if err := json.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("invalid JSON in MCP config: %v", err)
		}

		github := config.McpServers["github"]
		github.ServerUrl = "https://api.githubcopilot.com/mcp/"
		github.Headers = map[string]string{}

		github.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
		config.McpServers["github"] = github

		output, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal config: %v", err)
		}

		if err := os.WriteFile(mcpConfigPath, output, 0644); err != nil {
			return fmt.Errorf("cannot write to %s", mcpConfigPath)
		}

		fmt.Printf("Successfully updated GitHub MCP token to %s\n", mode)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(mcpCmd)
}
