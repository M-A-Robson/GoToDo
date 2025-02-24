package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID       int
	Content  string
	Complete bool
}

var DB *sql.DB

func initialiseDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatalf("Error opening database: %q\n", err)
	}

	createTableSql := `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	content TEXT
	complete BOOLEAN default 0
	);`

	_, err = DB.Exec(createTableSql)
	if err != nil {
		log.Fatalf("Error creating table: %q\n", err)
	}
}

func getAllTodos() ([]Todo, error) {
	todos := []Todo{}
	rows, err := DB.Query("SELECT id, content, complete FROM todos")
	if err != nil {
		return todos, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.Complete); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func getTodo(id int) (*Todo, error) {
	row := DB.QueryRow(
		"SELECT * FROM todos WHERE id = ?",
		id)
	todo := &Todo{}
	err := row.Scan(&todo.ID, &todo.Content, &todo.Complete)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func createTodo(content string) error {
	_, err := DB.Exec(
		"INSERT INTO todos(content) VALUES(?)",
		content)
	return err
}

func deleteTodo(id int) error {
	_, err := DB.Exec(
		"DELETE FROM todos WHERE id = ?",
		id)
	return err
}

func editTodo(id int, content string) error {
	_, err := DB.Exec(
		"UPDATE todos SET content = VALUES(?) WHERE id = ?",
		content,
		id)
	return err
}

func setTodoComplete(id int) error {
	return setTodoCompleteStatus(id, true)
}

func setTodoNotComplete(id int) error {
	return setTodoCompleteStatus(id, false)
}

func setTodoCompleteStatus(id int, status bool) error {
	var val string = "0"
	if status {
		val = "1"
	}
	_, err := DB.Exec(
		"UPDATE todos SET complete = ? WHERE id = ?",
		val,
		id)
	return err
}
