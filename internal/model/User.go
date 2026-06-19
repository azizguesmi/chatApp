package model

import "time"

type User struct {
	ID             int       `json:"id"`
	UserName       string    `json:"name"`
	Email          string    `json:"email"`
	PasswordHashed string    `json:"password_hashed"`
	CreatedAt      time.Time `json:"created_at"`
}
