package cli

import (
	"awesome-todo/internal/todo"
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

		taskService := todo.NewService()

		task, err := taskService.Complete(taskId)
		if err != nil {
			fmt.Printf("Got an error while mark the task as completed, id: %s, err: %v", taskId, err)
			return
		}

		printTasks([]todo.Task{task})
	},
}
