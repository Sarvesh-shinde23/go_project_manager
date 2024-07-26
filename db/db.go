package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	Db *sql.DB
}

// NewSQLiteStorage initializes a new SQLiteStorage with a given database file name.
func NewSQLiteStorage(dbName string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	fmt.Println("Connected to SQLite!")

	return &SQLiteStorage{Db: db}, nil
}

// Init initializes the SQLite database by creating necessary tables.
func (s *SQLiteStorage) Init() error {
	if err := s.createUsersTable(); err != nil {
		return err
	}

	if err := s.createProjectsTable(); err != nil {
		return err
	}

	if err := s.createTasksTable(); err != nil {
		return err
	}

	return nil
}

// createUsersTable creates the users table if it does not exist.
func (s *SQLiteStorage) createUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

	_, err := s.Db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}
	return nil
}

// createProjectsTable creates the projects table if it does not exist.
func (s *SQLiteStorage) createProjectsTable() error {
	query := `CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`

	_, err := s.Db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating projects table: %w", err)
	}
	return nil
}

// createTasksTable creates the tasks table if it does not exist.
func (s *SQLiteStorage) createTasksTable() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		project_id INTEGER,
		FOREIGN KEY (project_id) REFERENCES projects(id)
	);`

	_, err := s.Db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating tasks table: %w", err)
	}
	return nil
}
