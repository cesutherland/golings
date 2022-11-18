package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "golings",
		Short:         "Learn go through interactive exercises",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}

	rootCmd.AddCommand(cmdHint)
	rootCmd.AddCommand(ListCmd("info.toml"))
	rootCmd.AddCommand(RunCmd("info.toml"))
	rootCmd.AddCommand(cmdVerify)

	return rootCmd
}

func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		os.Exit(1)
	}
}
