package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	Db *sql.DB
}

func InitDB() (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {

		fmt.Println("Error opening database:", err)

		return nil, err
	}
	return &SQLiteStorage{Db: db}, nil
}
