package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"set-gh/pats"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Swap gh CLI token via gh auth login",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := args[0]
		token, err := pats.LoadToken(mode)
		if err != nil {
			return err
		}

		ghCmd := exec.Command("gh", "auth", "login", "--with-token")
		ghCmd.Stdin = strings.NewReader(token)
		var stderr bytes.Buffer
		ghCmd.Stderr = &stderr

		if err := ghCmd.Run(); err != nil {
			return fmt.Errorf("gh auth login failed: %v: %s", err, strings.TrimSpace(stderr.String()))
		}
		if stderr.Len() > 0 {
			return fmt.Errorf("gh auth login failed: %s", strings.TrimSpace(stderr.String()))
		}

		fmt.Printf("Successfully updated gh CLI token to %s\n", mode)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(cliCmd)
}
