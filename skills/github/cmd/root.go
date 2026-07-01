package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "set-gh",
	Short: "Manage GitHub tokens for work and personal modes",
}
