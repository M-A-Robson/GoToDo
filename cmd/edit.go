/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"todo/db"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit todo content",
	Long: `Edit todo by id. For example:
todo edit [id] [new content].
This will replace the content of the specified todo item`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid argument: %s is not an integer\n", args[0])
			return
		}
		todo, err := db.GetTodo(i)
		if err != nil {
			fmt.Printf("Error finding todo: %s\n", err)
		} else {
			fmt.Println("todo", todo.ID, "previous content:", todo.Content)
		}
		new_content := strings.Join(args[1:], " ")
		err = db.EditTodo(i, new_content)
		if err != nil {
			fmt.Printf("Error updating todo: %s\n", err)
		} else {
			fmt.Println("todo", args[0], "updated:", new_content)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
