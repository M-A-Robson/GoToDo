/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// roundupCmd represents the roundup command
var roundupCmd = &cobra.Command{
	Use:   "roundup",
	Short: "To see todays todo highlights",
	Long:  `To see todays todo highlights.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("roundup called")
		// get all todos
		// loop through for those created or completed today
		// display stats: # created, # completed, list completed and outstanding
	},
}

func init() {
	rootCmd.AddCommand(roundupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// roundupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// roundupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
