package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks",
	Long:  "Lists the tasks with flags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list of tasks")
	},
}
