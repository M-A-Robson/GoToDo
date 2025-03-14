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
		} else {
			fmt.Println("ToDo's marked complete")
			for _, todo := range todos {
				if todo.Complete {
					fmt.Println(todo.ID, ")", todo.Content)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
