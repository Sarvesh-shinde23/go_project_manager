package main

import (
	"log"

	"github.com/Sarvesh-shinde23/go_project_manager/db"
)

func main() {

	connStr := "postgres://postgres:0000@localhost:5432/goprojectmanager?sslmode=disable"
	storage, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPIServer(":3000", store)
	api.Serve()
}
