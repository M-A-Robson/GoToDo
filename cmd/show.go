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

// addCmd represents the add command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show todo",
	Long: `Show todo by id. For example:
todo show [id].`,
	Args: cobra.ExactArgs(1),
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
			fmt.Println("todo", todo.ID, ":", todo.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
