/*
Copyright © 2025 Mark Robson https://github.com/M-A-Robson
*/
package cmd

import (
	"fmt"
	"strings"

	"todo/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add a new todo. For example:
todo add [task description].`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task_description := strings.Join(args, " ")
		err := db.CreateTodo(task_description)
		if err != nil {
			fmt.Printf("Error creating todo: %s\n", err)
		} else {
			fmt.Println("Added Todo: " + task_description)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
