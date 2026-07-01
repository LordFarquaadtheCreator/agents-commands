package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"set-gh-token/pats"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CliInput struct {
	Mode string `json:"mode" jsonschema:"required,description=Either work_mode or personal_mode"`
}

type CliOutput struct {
	Message string `json:"message"`
}

func SwapCliToken(ctx context.Context, ss *mcp.ServerSession, req *mcp.CallToolParamsFor[CliInput]) (*mcp.CallToolResultFor[CliOutput], error) {
	token, err := pats.LoadToken(req.Arguments.Mode)
	if err != nil {
		return nil, err
	}

	ghCmd := exec.Command("gh", "auth", "login", "--with-token")
	ghCmd.Stdin = strings.NewReader(token)
	ghCmd.Stdout = os.Stdout
	ghCmd.Stderr = os.Stderr

	if err := ghCmd.Run(); err != nil {
		return nil, fmt.Errorf("gh auth login failed: %v", err)
	}

	msg := fmt.Sprintf("Successfully updated gh CLI token to %s", req.Arguments.Mode)
	return &mcp.CallToolResultFor[CliOutput]{
		Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		StructuredContent: CliOutput{Message: msg},
	}, nil
}
