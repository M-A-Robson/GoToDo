/*
Copyright © 2025 Mark Robson https://github.com/M-A-Robson
*/
package cmd

import (
	"fmt"
	"strconv"

	"todo/db"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo as done",
	Long: `Mark a todo as done by id. For example:
todo complete 16.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid argument: %s is not an integer\n", args[0])
			return
		}
		err = db.SetTodoComplete(i)
		if err != nil {
			fmt.Printf("Error marking todo complete: %s\n", err)
		} else {
			fmt.Printf("Todo %s marked as complete\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
