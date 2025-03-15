/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var completed string = "\u2713 "
var in_progress string = "\u2610 "

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
			complete := in_progress
			if todo.Complete {
				complete = completed
			}
			fmt.Println(todo.ID, complete, todo.Content, todo.Finished)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
