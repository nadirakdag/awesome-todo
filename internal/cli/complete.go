package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "Marks the task completed",
	Long:  "Mark a task as completed by specifying its task number.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		fmt.Printf("%s taks completed \n", taskId)
	},
}
