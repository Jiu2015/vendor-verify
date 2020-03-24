package cmd

import "github.com/spf13/cobra"

var pathCommand workPathCommand

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	pathCommand.Command().SilenceUsage = true
	rootCmd.AddCommand(pathCommand.Command())
}
