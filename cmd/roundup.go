/*
Copyright © 2025 Mark Robson https://github.com/M-A-Robson
*/
package cmd

import (
	"fmt"
	"time"
	"todo/db"

	"github.com/spf13/cobra"
)

// roundupCmd represents the roundup command
var roundupCmd = &cobra.Command{
	Use:   "roundup",
	Short: "To see todays todo highlights",
	Long:  `To see todays todo highlights.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get all todos
		todos, err := db.GetAllTodos()
		if err != nil {
			fmt.Println("Error fetching todos: ", err)
			return
		}
		todays_date := time.Now().Format("2006-01-02 15:04:00")[:11]
		fmt.Println("ToDo Roundup for :", todays_date)
		// loop through for those created or completed today
		created_today := []db.Todo{}
		completed_today := []db.Todo{}
		one_day_tasks := 0
		for _, todo := range todos {
			catch := 0
			if todo.Created[:11] == todays_date {
				created_today = append(created_today, todo)
				catch += 1
			}
			if todo.Complete {
				if todo.Finished[:11] == todays_date {
					completed_today = append(completed_today, todo)
					catch += 1
				}
			}
			if catch == 2 {
				one_day_tasks += 1
			}
		}
		// display stats
		fmt.Println(len(created_today), "tasks created")
		fmt.Println(len(completed_today), "tasks completed")
		fmt.Println(one_day_tasks, "tasks created and completed today")
		fmt.Println("Outstanding tasks from today:")
		for _, todo := range created_today {
			if !todo.Complete {
				fmt.Println(todo.ID, "-", todo.Content)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(roundupCmd)
}
