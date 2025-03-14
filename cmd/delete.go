/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"todo/db"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `Deletes a todo. example:
	todo delete 122
	will delete todo with id 122`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid argument: %s is not an integer\n", args[0])
			return
		}
		_, err = db.GetTodo(i)
		if err != nil {
			fmt.Println("Todo", i, "not found")
			return
		}
		err = db.DeleteTodo(i)
		if err != nil {
			fmt.Printf("Error deleting todo: %s\n", err)
			return
		}
		fmt.Println("todo deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
