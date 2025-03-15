/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

var inprogCmd = &cobra.Command{
	Use:   "inprog",
	Short: "Show all todos in progress",
	Long:  `Show all todos in progress.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := db.GetAllTodos()
		if err != nil {
			fmt.Println("Error fetching todos: ", err)
		}
		in_progress := []db.Todo{}
		for _, todo := range todos {
			if !todo.Complete {
				in_progress = append(in_progress, todo)
			}
		}
		if len(in_progress) == 0 {
			fmt.Println("No ToDo's in progress")
			return
		}
		fmt.Println("ToDo's not yet marked as complete")
		for _, todo := range in_progress {
			fmt.Println(todo.ID, "-", todo.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(inprogCmd)
}
