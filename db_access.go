package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func connectToDb() {
	cfg := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "task_db",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error accessing db")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}

func readTasks() ([]Task, error) {
	var tasks []Task

	rows, qErr := db.Query("SELECT id, title, status FROM tasks")
	if qErr != nil {
		return nil, fmt.Errorf("Error reading from DB", qErr)
	}
	defer rows.Close()
	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.TaskId, &task.Title, &task.Status); err != nil {
			return nil, fmt.Errorf("Error reading row", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func readPending() ([]Task, error) {
	var tasks []Task

	rows, qErr := db.Query("SELECT id, title, status FROM tasks where status = 1")
	if qErr != nil {
		return nil, fmt.Errorf("Error reading from DB", qErr)
	}
	defer rows.Close()
	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.TaskId, &task.Title, &task.Status); err != nil {
			return nil, fmt.Errorf("Error reading row", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func addTask(task Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks(title, status) VALUES (?, 1)", task.Title)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func markComplete(taskId string) (int64, error) {
	result, err := db.Exec("UPDATE tasks set status = 2 where id = ?", taskId)
	if err != nil {
		return 0, err
	}

	r, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return r, nil
}

//func main() {
//	connectToDb()
//	tasks, err := readTasks()
//
//	if err != nil {
//		fmt.Println("Error", err)
//	}
//
//	fmt.Println("Taks: ", tasks)
//}
