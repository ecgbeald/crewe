package main

import (
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectID    int64     `json:"projectID"`
	AssignedToID int64     `json:"assignedTo"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Project struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
