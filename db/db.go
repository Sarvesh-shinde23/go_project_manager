package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	Db *sql.DB
}

// NewPostgresStorage initializes a new PostgresStorage with a given connection string.
func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return &PostgresStorage{Db: db}, nil
}

// Init initializes the PostgreSQL database by creating necessary tables.
func (s *PostgresStorage) Init() (*sql.DB, error) {
	// initialize the tables
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.Db, nil
}

// createUsersTable creates the users table if it does not exist.
func (s *PostgresStorage) createUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
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
func (s *PostgresStorage) createProjectsTable() error {
	query := `CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`

	_, err := s.Db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating projects table: %w", err)
	}
	return nil
}

// createTasksTable creates the tasks table if it does not exist.
func (s *PostgresStorage) createTasksTable() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
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
