package cli

import (
	"awesome-todo/internal/todo"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks",
	Long:  "Lists all the tasks with their status",
	Run: func(cmd *cobra.Command, args []string) {

		taskService := todo.NewService()

		taskList, err := taskService.GetList()
		if err != nil {
			fmt.Printf("Got an error while getting the task list, err: %v", err)
		}

		printTasks(taskList)
	},
}

func printTasks(tasks []todo.Task) {
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 10, ' ', 0)

	//header
	fmt.Fprintln(w, "ID\tDescription\tStatus\tCreatedAt")
	fmt.Fprintln(w, "--\t-----------\t------\t---------")

	for _, t := range tasks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", t.ID, t.Description, getTaskStatus(t.Status), t.GetTimeDiff())
	}

	w.Flush()
}

func getTaskStatus(status bool) string {
	if status {
		return "Done"
	} else {
		return "Pending"
	}
}
