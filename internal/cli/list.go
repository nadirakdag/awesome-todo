package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks",
	Long:  "Lists all the tasks with their status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list of tasks")
	},
}
