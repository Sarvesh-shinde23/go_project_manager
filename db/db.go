package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	Db *sql.DB
}

func NewSQLstorage() {

}

func InitDB(dbName string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {

		fmt.Println("Error opening database:", err)

		return nil, err
	}
	return &SQLiteStorage{Db: db}, nil
}
