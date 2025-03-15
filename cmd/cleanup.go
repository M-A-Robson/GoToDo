/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Deletes all completed todos and updates todo ids",
	Long:  `Deletes all completed todos and updates todo ids`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.DeleteCompletedTodos()
		if err != nil {
			fmt.Println("Error deleting completed todos: ", err)
			fmt.Println("ToDo cleanup exited")
			return
		}

		todos, err := db.GetAllTodos()
		if err != nil {
			fmt.Println("Error fetching todos: ", err)
			fmt.Println("ToDo cleanup exited")
			return
		}

		for i, todo := range todos {
			if todo.ID != i+1 {
				err = db.EditTodoId(todo.ID, i+1)
			}
			if err != nil {
				fmt.Println("Error editing todo", todo.ID, ":", err)
				fmt.Println("ToDo cleanup exited")
				return
			}
		}

		err = db.SetAutoIncrementCounter(len(todos))
		if err != nil {
			fmt.Println("Error resetting counter:", err)
			return
		}

		if err == nil {
			fmt.Println("ToDo cleanup completed")
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
