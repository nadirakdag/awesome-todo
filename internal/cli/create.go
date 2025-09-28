package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new task",
	Long:  "Creates a new task with a description to you task list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		fmt.Printf("creating a new task with description: %s", description)
	},
}
