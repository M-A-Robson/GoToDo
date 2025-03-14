/*
Copyright © 2025 Mark Robson
*/
package main

import (
	"todo/cmd"
	"todo/db"
)

func main() {
	db.InitialiseDatabase()
	defer db.DB.Close()
	cmd.Execute()
}
