package cli

import (
	"awesome-todo/internal/todo"
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [description]",
	Short: "Create a new task",
	Long:  "Creates a new task with the provided description in your task list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		taskService := todo.NewService()

		task, err := taskService.Create(description)
		if err != nil {
			fmt.Printf("Got an error while creating the task, err: %v", err)
			return
		}

		printTasks([]todo.Task{task})
	},
}
