/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Show all todos",
	Long:  `Show all todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := db.GetAllTodos()
		if err != nil {
			fmt.Println("Error fetching todos: ", err)
			return
		}
		completed_tasks := []db.Todo{}
		for _, todo := range todos {
			if todo.Complete {
				completed_tasks = append(completed_tasks, todo)
			}
		}
		if len(completed_tasks) == 0 {
			fmt.Println("No ToDo's marked as complete")
			return
		}
		fmt.Println("ToDo's marked complete")
		for _, todo := range completed_tasks {
			fmt.Println(todo.ID, "-", todo.Content, "completed", todo.Finished)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
