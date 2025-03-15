package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID       int
	Content  string
	Complete bool
	Created  string
	Finished string
}

var DB *sql.DB

func InitialiseDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatalf("Error opening database: %q\n", err)
	}

	createTableSql := `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	content TEXT,
	complete BOOLEAN default 0,
	created TEXT NOT NULL DEFAULT '',
	finished TEXT NOT NULL DEFAULT '');`

	_, err = DB.Exec(createTableSql)
	if err != nil {
		log.Fatalf("Error creating table: %q\n", err)
	}
}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}
	rows, err := DB.Query("SELECT * FROM todos")
	if err != nil {
		return todos, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Complete,
			&todo.Created,
			&todo.Finished); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodo(id int) (*Todo, error) {
	row := DB.QueryRow(
		"SELECT * FROM todos WHERE id = ?",
		id)
	todo := &Todo{}
	err := row.Scan(
		&todo.ID,
		&todo.Content,
		&todo.Complete,
		&todo.Created,
		&todo.Finished)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func CreateTodo(content string) error {
	current_time := time.Now().Format("01-JAN-2006 15:04:00")
	_, err := DB.Exec(
		"INSERT INTO todos(content, created) VALUES(?,?)",
		content, current_time)
	return err
}

func DeleteTodo(id int) error {
	_, err := DB.Exec(
		"DELETE FROM todos WHERE id = ?",
		id)
	return err
}

func EditTodo(id int, content string) error {
	_, err := DB.Exec(
		"UPDATE todos SET content = ? WHERE id = ?",
		content,
		id)
	return err
}

func SetTodoComplete(id int) error {
	return SetTodoCompleteStatus(id, true)
}

func SetTodoNotComplete(id int) error {
	return SetTodoCompleteStatus(id, false)
}

func SetTodoCompleteStatus(id int, status bool) error {
	var val string = "0"
	var current_time string = ""
	if status {
		val = "1"
		current_time = time.Now().Format("01-JAN-2006 15:04:00")
	}
	_, err := DB.Exec(
		"UPDATE todos SET (complete, finished) = (?, ?) WHERE id = ?",
		val,
		current_time,
		id)
	return err
}
