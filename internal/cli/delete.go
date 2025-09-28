package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "Deletes the task ",
	Long:  "Deletes the task by specifying its task number.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		fmt.Printf("%s task deleted", taskId)
	},
}
