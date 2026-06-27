package main

import (
	"os"

	"github.com/spf13/cobra"

	"manage-job/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "manage-job",
	Short: "Track and retrieve job applications",
}

func main() {
	cmd.PatchCmd.Flags().String("matchBy", "", "JSON object to identify the row (required)")
	cmd.PatchCmd.Flags().String("update", "", "JSON object with fields to change (required)")
	cmd.DeleteCmd.Flags().String("matchBy", "", "JSON object to identify the row (required)")

	rootCmd.AddCommand(cmd.GetCmd, cmd.TrackCmd, cmd.PatchCmd, cmd.DeleteCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
