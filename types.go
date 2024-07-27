package main

import "time"

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"projectID"`
	AssignedToID int64     `json:"assignedTo"`
	CreatedAt    time.Time `json:"createdAt"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}
type User struct {
	ID        int64     `json:id`
	FirstName string    `json:"firstName"`
	lastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
