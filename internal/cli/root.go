package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Simple Todo App CLI Interface",
	Long:  "Simple Todo App with CLI, REST and gRPC endpoint, this is the CLI interface",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
}
