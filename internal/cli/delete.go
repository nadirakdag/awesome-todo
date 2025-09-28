package cli

import (
	"awesome-todo/internal/todo"
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

		taskService := todo.NewService()

		task, err := taskService.Delete(taskId)
		if err != nil {
			fmt.Printf("Got an error while deleting the task, id: %s, err: %v", taskId, err)
			return
		}

		printTasks([]todo.Task{task})
	},
}
