package model

import (
	"github.com/mayur-tolexo/demo/db"
)

// UserList is the response of list user endpoint
type UserList struct {
	User []UserDetail `json:"user"`
}

// UserDetail detail model
type UserDetail struct {
	db.User
}

//CreateUser is create new user request model
type CreateUser struct {
	Name  string `json:"name" minLength:"4" example:"test"`
	Email string `json:"email" binding:"required" example:"test@gmail.com"`
}
