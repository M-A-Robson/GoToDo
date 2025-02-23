/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo as done",
	Long: `Mark a todo as done by id. For example:
todo complete 16.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid argument: %s is not an integer\n", args[0])
			return nil
		}
		if !indexExists(i) {
			fmt.Printf("Invalid argument: %s does not exist\n", args[0])
			return nil
		} else {
			// todo add business logic here to mark todo[id] as complete
			fmt.Printf("Todo %s marked as complete\n", args[0])
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func indexExists(i int) bool {
	// todo check that index is in our table
	// todo move to app wide context as many calls can use this check
	return true
}
