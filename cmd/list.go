/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

// listCmd represents the complete command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all todos",
	Long:  `Show all todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := db.GetAllTodos()
		if err != nil {
			fmt.Println("Error fetching todos: ", err)
			return
		}
		fmt.Println("ToDo List")
		for _, todo := range todos {
			complete := "in progress"
			if todo.Complete {
				complete = "completed " + todo.Finished
			}
			fmt.Println(todo.ID, ")", todo.Content, ":", complete)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
