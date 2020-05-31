package db

// User model
type User struct {
	UserID int    `json:"user_id" example:"1"`
	Name   string `json:"name" example:"test"`
	Email  string `json:"email" example:"test@gmail.com"`
}
