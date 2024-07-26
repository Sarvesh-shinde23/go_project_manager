package main

import (
	"log"

	"github.com/Sarvesh-shinde23/go_project_manager/db"
)

func main() {

	sqlStorage, err := db.NewSQLiteStorage("test.db")
	db := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPIServer(":3000", store)
	api.Serve()
}
