package main

import (
	"log"
)

func main() {

	sqlStorage := SQLiteStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPIServer(":3000", store)
	api.Serve()
}
